package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	t "github.com/gnongs/goc-relayer-evidence/types"
)

const (
	rpc     = "https://api-c.hero.nodestake.top/"
	address = "cosmos1f9k0728rcngqj7p2eeqk27v8dw5sw5ccydx77j"
)

var (
	tPackets int
)

func main() {
	evidenceHeight := getHeight(rpc, address)

	fmt.Println("TOTAL_COUNT:", len(evidenceHeight), "of", tPackets)
	fmt.Println(evidenceHeight)
}

func sendQuery(v url.Values) t.ReturnData {
	u := rpc + "/cosmos/tx/v1beta1/txs"
	var d t.ReturnData

	client := &http.Client{}

	req, _ := http.NewRequest("GET", u, nil)
	v.Set("events", "message.action='/ibc.core.client.v1.MsgUpdateClient'")
	v.Set("events", "message.sender='"+address+"'")
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

func totalPackets() int {
	parameters := url.Values{}
	totalPackets, _ := strconv.Atoi(sendQuery(parameters).Pagination.Total)

	return totalPackets
}

func getHeight(rpc, address string) []string {
	var heights []string

	tPackets = totalPackets()

	for i := 0; i <= (tPackets / 100); i++ {
		var value = url.Values{}
		value.Set("pagination.offset", strconv.Itoa(i*100+1))
		resp := sendQuery(value)

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
