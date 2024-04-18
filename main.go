package main

import "demo1/point"

type gasPrice struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func main() {
	point.Point1()

	//point.Point2("private_key","to_adress",10000000)

	var blockHigh int64 = 5671745
	point.Point3(blockHigh)

}
