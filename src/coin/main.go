package main

import (
	"core"
	"fmt"
)
func main(){
	bc := core.NewBlockchain()	//初始化区块链，创建第一个区块(创世纪区块)

	bc.AddBlock("send 1 BTC to bysu")
	bc.AddBlock("send 2 BTC to bysu")

	for _,block := range bc.Blocks{
		fmt.Printf("Prev.hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Println()
	}
}
