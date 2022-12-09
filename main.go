package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	c "github.com/gnongs/goc-relayer-evidence/config"
	t "github.com/gnongs/goc-relayer-evidence/types"
)

var (
	chains          c.Config
	evidence        []t.Evidence
	accountTotalTxs []int
)

func init() {
	chains = c.LoadConfigFile("./config.json")
}

func main() {
	var wg sync.WaitGroup

	for idx, c := range chains {
		wg.Add(1)
		evidence = append(evidence, t.Evidence{ChainName: c.Chain})
		accountTotalTxs = append(accountTotalTxs, 0)
		evidence[idx].ChainName = c.Chain

		fmt.Println("Search all packets on", c.Chain, "....")
		go func(x int, chain, addr, rpc string) {
			defer wg.Done()
			accountTotalTxs[x] = getTotalTxs(rpc, addr)
			evidence[x].TotalHeights = getHeight(rpc, addr, accountTotalTxs[x])
			fmt.Println(chain, "packets checked")
		}(idx, c.Chain, c.Address, c.RPCEndpoint)
	}

	wg.Wait()

	for idx, e := range evidence {
		fmt.Println(e.ChainName, "'s height is verified ....")
		fmt.Println("Verifying height of", e.ChainName)
		for _, b := range e.TotalHeights {
			for _, t := range isVerified(chains[idx].RPCEndpoint, chains[idx].Address, b).Txs {
				var trigger bool
				for _, m := range t.Body.Messages {
					if m.Type == "/ibc.core.client.v1.MsgUpdateClient" {
						trigger = true
						if m.Signer == chains[idx].Address {
							e.VerifiedHeights = append(e.VerifiedHeights, b)
						}
					}
				}
				if trigger {
					break
				}
			}
		}
		fmt.Println(e.ChainName, "'s total_count:", len(e.VerifiedHeights), "of", len(e.TotalHeights))
		fmt.Println(e.VerifiedHeights)
	}
}

func sendRequest(rpc, route string, v url.Values) []byte {
	url := rpc + route

	req, _ := http.NewRequest("GET", url, nil)
	req.URL.RawQuery = v.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	bArr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	return bArr
}

func getTotalTxs(rpc, address string) int {
	var txs t.RetrunTxs

	parameters := url.Values{}
	parameters.Add("events", "message.action='/ibc.core.client.v1.MsgUpdateClient'")
	parameters.Add("events", "message.sender='"+address+"'")
	parameters.Add("events", "recv_packet.packet_dst_channel='channel-0'")
	parameters.Add("events", "recv_packet.packet_dst_port='consumer'")
	parameters.Add("limit", "100")

	json.Unmarshal(sendRequest(rpc, "/cosmos/tx/v1beta1/txs", parameters), &txs)

	totalTxs, _ := strconv.Atoi(txs.Pagination.Total)

	return totalTxs
}

func getHeight(rpc, address string, tPackets int) []string {
	var heights []string
	var txs t.RetrunTxs

	for i := 0; i <= (tPackets / 100); i++ {
		parameters := url.Values{}
		parameters.Add("events", "message.action='/ibc.core.client.v1.MsgUpdateClient'")
		parameters.Add("events", "message.sender='"+address+"'")
		parameters.Add("events", "recv_packet.packet_dst_channel='channel-0'")
		parameters.Add("events", "recv_packet.packet_dst_port='consumer'")
		parameters.Add("limit", "100")
		parameters.Add("pagination.offset", strconv.Itoa(i*100+1))

		json.Unmarshal(sendRequest(rpc, "/cosmos/tx/v1beta1/txs", parameters), &txs)

		for _, txr := range txs.TxResponses {
			for _, m := range txr.Tx.Body.Messages {
				if m.TypeURL != "/ibc.core.client.v1.MsgUpdateClient" {
					continue
				}
				if m.Header.SignedHeader.Header.ValidatorsHash != m.Header.SignedHeader.Header.NextValidatorsHash {
					heights = append(heights, txr.Height)
				}
			}
		}
	}

	return heights
}

func isVerified(rpc, address, height string) t.ReturnTxsBlock {
	query := url.Values{}
	route := "/cosmos/tx/v1beta1/txs/block/" + height

	var blockData t.ReturnTxsBlock

	json.Unmarshal(sendRequest(rpc, route, query), &blockData)

	return blockData
}
