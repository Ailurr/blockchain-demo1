package service

import (
	"bytes"
	"demo1/model"
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
func (c *BTCClient) getBlock(hex string) (model.Block, error) {
	arg := BTCRequest{
		Method: "getblock",
		Params: []interface{}{hex},
	}
	body, _ := sonic.Marshal(arg)
	respBytes := newRequest(c, body)
	var res model.BlockRes
	sonic.Unmarshal([]byte(string(respBytes)), &res)
	return res.Result, nil
}
func (c *BTCClient) getTransaction(hex string) (model.Trans, error) {
	arg := BTCRequest{
		Method: "getrawtransaction",
		Params: []interface{}{hex, true},
	}
	body, _ := sonic.Marshal(arg)
	respBytes := newRequest(c, body)
	var res model.TransRes
	sonic.Unmarshal([]byte(string(respBytes)), &res)
	return res.Result, nil
}

func parseTransactionInfo(t *model.Trans) model.ParsedTrans {
	totalIn := 0.0
	totalOut := 0.0
	in := make([]model.UTXO, 0)
	out := make([]model.UTXO, 0)
	for _, s := range t.Vin {
		transaction, _ := btcClient.getTransaction(s.Txid)
		val := transaction.Vout[s.Vout].Value
		totalIn += val
		in = append(in, model.UTXO{
			Address: transaction.Vout[s.Vout].ScriptPubKey.Address,
			Value:   val,
		})
		//fmt.Printf("Address:%s | amount：%.8f BTC\n", transaction.Vout[s.Vout].ScriptPubKey.Address, val)
	}

	for _, s := range t.Vout {
		totalOut += s.Value
		out = append(out, model.UTXO{
			Address: s.ScriptPubKey.Address,
			Value:   s.Value,
		})
		//fmt.Printf("Adress:%s | amount：%.8f BTC\n", s.ScriptPubKey.Address, s.Value)
	}
	fee := totalIn - totalOut
	res := model.ParsedTrans{
		TxId: t.Txid,
		In:   in,
		Out:  out,
		Fee:  fee,
	}
	return res
}

func GetBlockTransInfo(hash string) ([]model.ParsedTrans, error) {
	block, err := btcClient.getBlock(hash)
	if err != nil {
		return []model.ParsedTrans{}, err
	}
	result := make([]model.ParsedTrans, 0)

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
