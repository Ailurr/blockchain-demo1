package point

import (
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)
import (
	"context"
	"fmt"
	"math/big"
)

func Point3(high int64) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(high))
	if err != nil {
		return
	}
	fmt.Println()
	fmt.Printf("%s | %s | %s | %s \n", "token", "from", "to", "value")
	for _, tx := range block.Transactions() {

		//TODO  get the token of transaction

		if tx.Type() != types.LegacyTxType && tx.Type() != types.AccessListTxType {
			continue
		}

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s | %s | %s | %s \n", "", from.Hex(), tx.To().Hex(), tx.Value())
	}
}
