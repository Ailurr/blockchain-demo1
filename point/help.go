package point

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethClient *ethclient.Client

var btcClient *BTCClient

func init() {
	ethClient, _ = ethclient.Dial("https://sepolia.infura.io/v3/d4a09685d62a40738b42e40880995927")
	btcClient = NewBTCClient()
}
func formatPrint(s []byte) {
	var out bytes.Buffer
	json.Indent(&out, s, "", "  ")
	fmt.Println(out.String())
}
