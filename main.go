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
	chains   c.Config
	evidence []t.Evidence
	tPackets []int
)

func init() {
	chains = c.LoadConfigFile("./config.json")
}

func main() {
	var wg sync.WaitGroup

	for idx, c := range chains {
		wg.Add(1)
		evidence = append(evidence, t.Evidence{ChainName: c.Chain})
		tPackets = append(tPackets, 0)
		evidence[idx].ChainName = c.Chain

		fmt.Println("Search Tx on", c.Chain, "....")
		go func(x int, chain, addr, rpc string) {
			defer wg.Done()
			tPackets[x] = totalPackets(rpc, addr)
			evidence[x].Heights = getHeight(rpc, addr, tPackets[x])
			fmt.Println(chain, "packets checked")
		}(idx, c.Chain, c.Address, c.RPCEndpoint)
	}

	wg.Wait()

	for idx, e := range evidence {
		fmt.Println(e.ChainName, "'s total_count:", len(e.Heights), "of", tPackets[idx])
		fmt.Println(e.Heights)
	}
}

func sendQuery(rpc, address string, v url.Values) t.ReturnData {
	u := rpc + "/cosmos/tx/v1beta1/txs"
	var d t.ReturnData

	client := &http.Client{}

	req, _ := http.NewRequest("GET", u, nil)
	v.Set("events", "message.action='/ibc.core.client.v1.MsgUpdateClient'")
	v.Set("events", "message.sender='"+address+"'")
	v.Set("limit", "100")
	req.URL.RawQuery = v.Encode()

	resp, err := client.Do(req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return d
}

func totalPackets(rpc, address string) int {
	parameters := url.Values{}
	totalPackets, _ := strconv.Atoi(sendQuery(rpc, address, parameters).Pagination.Total)

	return totalPackets
}

func getHeight(rpc, address string, tPackets int) []string {
	var heights []string

	for i := 0; i <= (tPackets / 100); i++ {
		var value = url.Values{}
		value.Set("pagination.offset", strconv.Itoa(i*100+1))
		resp := sendQuery(rpc, address, value)

		for _, txr := range resp.TxResponses {
			for _, m := range txr.Tx.Body.Messages {
				if m.TypeURL == "/ibc.core.client.v1.MsgUpdateClient" {
					if m.Header.SignedHeader.Header.ValidatorsHash != m.Header.SignedHeader.Header.NextValidatorsHash {
						heights = append(heights, txr.Height)
					}
				}
			}
		}
	}

	return heights
}
