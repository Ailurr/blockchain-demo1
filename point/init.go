package point

import "github.com/ethereum/go-ethereum/ethclient"

var client *ethclient.Client

func init() {
	client, _ = ethclient.Dial("https://sepolia.infura.io/v3/d4a09685d62a40738b42e40880995927")
}
