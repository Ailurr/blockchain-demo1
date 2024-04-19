package main

import "demo1/point"

func main() {
	//file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	////old:=os.Stdout
	//os.Stdout = file
	//
	//hash := "0000000000000000000285f2233538754a9d409e939c5800ba8f88862fdc55e5"
	//point.Point1(hash)

	//private_key := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//to_adress := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//value := big.NewInt(0)
	//point.Point2(private_key, to_adress, value)
	//https://sepolia.etherscan.io/tx/0x65685d6efbb79ebb64a3dcdb2fe9548ba1269b9e643021691013c1cc568e15fe

	//os.Stdout=old

	var blockHigh int64 = 5731717
	point.Point3(blockHigh)

}
