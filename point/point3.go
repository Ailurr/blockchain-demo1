package point

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"math/big"
)

func Point3(high int64) {
	fmt.Println("-------------------------Point 3-----------------------------")
	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(high))
	if err != nil {
		return
	}
	//trans hash: 0x16a8ac8af36f0227e6bc17fbf44fe88e72412505076c9ba64943206b457a5445
	receipt, err := ethClient.TransactionReceipt(context.Background(), block.Transactions()[8].Hash())
	b, _ := sonic.Marshal(receipt.Logs)
	formatPrint(b)
	//fmt.Printf("%s | %s | %s | %s \n", "token", "from", "to", "value")
	//for _, tx := range block.Transactions() {
	//
	//
	//	chainID, err := ethClient.NetworkID(context.Background())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("%s | %s | %s | %s \n", "", from.Hex(), tx.To().Hex(), tx.Value())
	//}
}
