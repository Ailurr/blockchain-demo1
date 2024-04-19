package point

func Point2(privatekey string, to string, gas uint64) {
	//fmt.Println("-------------------------Point 2-----------------------------")
	//
	//privateKey, err := crypto.HexToECDSA(privatekey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal(err)
	//}
	//
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//value := big.NewInt(0)
	//// in units
	//gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//toAddress := common.HexToAddress(to)
	//var data []byte
	//tx := types.NewTx(&types.LegacyTx{
	//	Nonce:    nonce,
	//	GasPrice: gasPrice,
	//	Gas:      gas,
	//	To:       &toAddress,
	//	Value:    value,
	//	Data:     data,
	//})
	//
	//chainID, err := ethClient.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = ethClient.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
