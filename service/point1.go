package service

import (
	"bytes"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
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
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	NTx               int      `json:"nTx"`
	Previousblockhash string   `json:"previousblockhash"`
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
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence    int64    `json:"sequence"`
		Coinbase    string   `json:"coinbase"`
		Txinwitness []string `json:"txinwitness"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm     string `json:"asm"`
			Desc    string `json:"desc"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
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
func (c *BTCClient) getBlock(hex string) (Block, error) {
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
func (c *BTCClient) getTransaction(hex string) (Trans, error) {
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

type UTXO struct {
	Address string  `json:"address,omitempty"`
	Value   float64 `json:"value,omitempty"`
}
type ParsedTrans struct {
	TxId string  `json:"txid"`
	In   []UTXO  `json:"in"`
	Out  []UTXO  `json:"out"`
	Fee  float64 `json:"fee"`
}

func parseTransactionInfo(t *Trans) ParsedTrans {
	totalIn := 0.0
	totalOut := 0.0
	in := make([]UTXO, 0)
	out := make([]UTXO, 0)
	for _, s := range t.Vin {
		transaction, _ := btcClient.getTransaction(s.Txid)
		val := transaction.Vout[s.Vout].Value
		totalIn += val
		in = append(in, UTXO{
			Address: transaction.Vout[s.Vout].ScriptPubKey.Address,
			Value:   val,
		})
		//fmt.Printf("Address:%s | amount：%.8f BTC\n", transaction.Vout[s.Vout].ScriptPubKey.Address, val)
	}

	for _, s := range t.Vout {
		totalOut += s.Value
		out = append(out, UTXO{
			Address: s.ScriptPubKey.Address,
			Value:   s.Value,
		})
		//fmt.Printf("Adress:%s | amount：%.8f BTC\n", s.ScriptPubKey.Address, s.Value)
	}
	fee := totalIn - totalOut
	res := ParsedTrans{
		TxId: t.Txid,
		In:   in,
		Out:  out,
		Fee:  fee,
	}
	return res
}

func Point1(hash string) ([]ParsedTrans, error) {

	block, err := btcClient.getBlock(hash)
	if err != nil {
		return []ParsedTrans{}, err
	}
	result := make([]ParsedTrans, 0)
	for i, tx := range block.Tx {
		if i >= 3 {
			break
		}
		transaction, _ := btcClient.getTransaction(tx)
		if transaction.Vin[0].Coinbase != "" {
			continue
		}
		parsedTrans := parseTransactionInfo(&transaction)
		result = append(result, parsedTrans)
	}

	return result, nil
}
