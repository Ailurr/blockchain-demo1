package model

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
