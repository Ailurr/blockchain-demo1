package main

import (
	"demo1/point"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//old:=os.Stdout
	os.Stdout = file

	hash := "0000000000000000000285f2233538754a9d409e939c5800ba8f88862fdc55e5"
	point.Point1(hash)

	//point.Point2(private_key,to_address,10000000)

	//os.Stdout=old
	//
	//var blockHigh int64 = 5729013
	//point.Point3(blockHigh)

}
