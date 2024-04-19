package main

import (
	"demo1/point"
	"math/big"
)

func main() {

	//file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	////old:=os.Stdout
	//os.Stdout = file
	//
	//hash := "0000000000000000000285f2233538754a9d409e939c5800ba8f88862fdc55e5"
	//point.Point1(hash)

	private_key := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	to_adress := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	value := big.NewInt(0)
	point.Point2(private_key, to_adress, value)
	//os.Stdout=old
	//
	//var blockHigh int64 = 5729013
	//point.Point3(blockHigh)

}
