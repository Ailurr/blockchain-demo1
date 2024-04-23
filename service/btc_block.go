package service

import (
	"bytes"
	"demo1/model"
	"encoding/json"
	"fmt"
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

func NewBTCClient() *BTCClient {
	return &BTCClient{
		client: &http.Client{},
	}
}
func newRequest(c *BTCClient, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", "https://go.getblock.io/da7c6bec0e4b456c951f790cf5b84c1b", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	respBytes, _ := io.ReadAll(res.Body)
	return respBytes, nil
}
func (c *BTCClient) getBlock(hex string) (model.Block, error) {
	arg := BTCRequest{
		Method: "getblock",
		Params: []interface{}{hex},
	}
	body, err := json.Marshal(arg)
	if err != nil {
		return model.Block{}, fmt.Errorf("json marshal err: %w", err)
	}
	respBytes, err := newRequest(c, body)
	if err != nil {
		return model.Block{}, fmt.Errorf("newRequest err: %w", err)
	}
	var res model.BlockRes
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return model.Block{}, fmt.Errorf("json unmarshal err: %w", err)
	}
	return res.Result, nil
}
func (c *BTCClient) getTransaction(hex string) (model.Trans, error) {
	arg := BTCRequest{
		Method: "getrawtransaction",
		Params: []interface{}{hex, true},
	}
	body, err := json.Marshal(arg)
	if err != nil {
		return model.Trans{}, fmt.Errorf("json marshal err: %w", err)
	}
	respBytes, err := newRequest(c, body)
	if err != nil {
		return model.Trans{}, fmt.Errorf("newRequest err: %w", err)
	}
	var res model.TransRes
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return model.Trans{}, fmt.Errorf("json unmarshal err: %w", err)
	}
	return res.Result, nil
}

func parseTransactionInfo(t *model.Trans) (model.ParsedTrans, error) {
	totalIn := 0.0
	totalOut := 0.0
	in := make([]model.UTXO, 0)
	out := make([]model.UTXO, 0)
	for _, s := range t.Vin {
		transaction, err := btcClient.getTransaction(s.Txid)
		if err != nil {
			return model.ParsedTrans{}, fmt.Errorf("get transaction err %w", err)
		}
		val := transaction.Vout[s.Vout].Value
		totalIn += val
		in = append(in, model.UTXO{
			Address: transaction.Vout[s.Vout].ScriptPubKey.Address,
			Value:   val,
		})
	}

	for _, s := range t.Vout {
		totalOut += s.Value
		out = append(out, model.UTXO{
			Address: s.ScriptPubKey.Address,
			Value:   s.Value,
		})
	}
	fee := totalIn - totalOut
	res := model.ParsedTrans{
		TxId: t.Txid,
		In:   in,
		Out:  out,
		Fee:  fee,
	}
	return res, nil
}

func GetBlockTransInfo(hash string) ([]model.ParsedTrans, error) {
	block, err := btcClient.getBlock(hash)
	if err != nil {
		return []model.ParsedTrans{}, fmt.Errorf("getBlock err %w", err)
	}
	result := make([]model.ParsedTrans, 0)

	for i, tx := range block.Tx {
		if i >= 3 {
			break
		}
		transaction, err := btcClient.getTransaction(tx)
		if err != nil {
			return []model.ParsedTrans{}, fmt.Errorf("getTransaction err %w", err)
		}
		if transaction.Vin[0].Coinbase != "" {
			continue
		}
		parsedTrans, err := parseTransactionInfo(&transaction)
		if err != nil {
			return []model.ParsedTrans{}, fmt.Errorf("parseTransactionInfo err %w", err)
		}
		result = append(result, parsedTrans)
	}

	return result, nil
}
