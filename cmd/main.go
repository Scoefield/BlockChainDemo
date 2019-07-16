package main

import "demochain/core"

func main()  {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to jacky")
	bc.SendData("Send 2 EOS to Mike")
	bc.Print()
}