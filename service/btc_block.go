package service

import (
	"bytes"
	"demo1/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
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

func parseTransactionInfo(t model.Trans) (model.ParsedTrans, error) {
	totalInNum := 0.0
	totalOutNuM := 0.0
	in := make([]model.UTXO, len(t.Vin))
	out := make([]model.UTXO, len(t.Vout))
	var lock sync.Mutex
	asyctl := make(chan struct{}, 4)
	for i := 0; i < 4; i++ {
		asyctl <- struct{}{}
	}

	wg := sync.WaitGroup{}
	fmt.Printf("%s vin len %d\n", t.Txid, len(t.Vin))
	for i, utxoIn := range t.Vin {
		wg.Add(1)
		<-asyctl
		uIn := utxoIn
		i := i
		go func() {
			transaction, err := btcClient.getTransaction(uIn.Txid)
			if err != nil {
				fmt.Printf("getTransaction err: %s\n", err.Error())
			}
			val := transaction.Vout[uIn.Vout].Value
			lock.Lock()
			totalInNum += val
			lock.Unlock()
			in[i] = model.UTXO{
				Address: transaction.Vout[uIn.Vout].ScriptPubKey.Address,
				Value:   val,
			}
			asyctl <- struct{}{}
			wg.Done()
		}()
	}
	for i, s := range t.Vout {
		totalOutNuM += s.Value
		out[i] = model.UTXO{
			Address: s.ScriptPubKey.Address,
			Value:   s.Value,
		}
	}
	wg.Wait()
	fee := totalInNum - totalOutNuM
	res := model.ParsedTrans{
		TxId: t.Txid,
		In:   in,
		Out:  out,
		Fee:  fee,
	}

	return res, nil
}

func GetBlockTransInfo(hash string) ([]model.ParsedTrans, error) {

	num := 24
	block, err := btcClient.getBlock(hash)
	if err != nil {
		return []model.ParsedTrans{}, fmt.Errorf("getBlock err %w", err)
	}
	res := make([]model.ParsedTrans, num)

	asyctl := make(chan struct{}, 12)
	for i := 0; i < 12; i++ {
		asyctl <- struct{}{}
	}
	wg := sync.WaitGroup{}
	fmt.Printf("block.Tx len: %d, only use %d for example\n", len(block.Tx), num)
	for i, tx := range block.Tx {
		if i >= num {
			break
		}
		tx := tx
		i := i
		<-asyctl
		wg.Add(1)
		go func() {
			transaction, err := btcClient.getTransaction(tx)
			if err != nil {
				fmt.Printf("getTransaction err %s", err.Error())
			}
			if transaction.Vin[0].Coinbase != "" {
				return
			}
			parsedTrans, err := parseTransactionInfo(transaction)
			if err != nil {
				fmt.Printf("parseTransactionInfo err %s.", err.Error())
			}
			res[i] = parsedTrans
			asyctl <- struct{}{}
			wg.Done()
		}()
	}

	//TODO 必须额外加一个wg.Done()才能结束
	wg.Done()
	wg.Wait()
	return res, nil
}
