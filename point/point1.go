package point

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

type BTCClient struct {
	client *http.Client
}
type BTCRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	Merkleroot        string   `json:"merkleroot"`
	Time              int      `json:"time"`
	Mediantime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        int      `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	NTx               int      `json:"nTx"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
	Strippedsize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Tx                []string `json:"tx"`
}
type BlockRes struct {
	Result Block  `json:"result"`
	Error  error  `json:"error"`
	Id     string `json:"id"`
}
type Trans struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Coinbase string `json:"coinbase"`
		Sequence int64  `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm  string `json:"asm"`
			Desc string `json:"desc"`
			Hex  string `json:"hex"`
			Type string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}
type TransRes struct {
	Result Trans  `json:"result"`
	Error  error  `json:"error"`
	Id     string `json:"id"`
}

func NewBTCClient() *BTCClient {
	return &BTCClient{
		client: &http.Client{},
	}
}
func newRequest(c *BTCClient, body []byte) []byte {
	req, _ := http.NewRequest("POST", "https://go.getblock.io/da7c6bec0e4b456c951f790cf5b84c1b", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, _ := c.client.Do(req)
	defer res.Body.Close()
	respBytes, _ := io.ReadAll(res.Body)
	return respBytes
}
func (c *BTCClient) GetBlock(hex string) (Block, error) {
	arg := BTCRequest{
		Method: "getblock",
		Params: []interface{}{hex},
	}
	body, _ := sonic.Marshal(arg)
	respBytes := newRequest(c, body)
	var res BlockRes
	sonic.Unmarshal([]byte(string(respBytes)), &res)
	return res.Result, nil
}
func (c *BTCClient) GetTransaction(hex string) (Trans, error) {
	arg := BTCRequest{
		Method: "getrawtransaction",
		Params: []interface{}{hex, true},
	}
	body, _ := sonic.Marshal(arg)
	respBytes := newRequest(c, body)
	var res TransRes
	sonic.Unmarshal([]byte(string(respBytes)), &res)
	return res.Result, nil
}
func Point1() {
	newbtc := NewBTCClient()
	block, err := newbtc.GetBlock("00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09")
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", block)
	fmt.Println()
	transaction, err := newbtc.GetTransaction(block.Tx[0])
	fmt.Printf("%+v\n", transaction)
}
