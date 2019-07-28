package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block keeps block headers
type Block struct {
	Timestamp int64	//区块创建时间
	Data []byte	//区块包含的数据
	PrevBlockHash []byte	//前一个区块的哈希值
	Hash []byte	//区块自身的哈希值，用于校验区块数据有效
}

//NewBlock creates and returns Block
func NewBlock(data string,prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}

//SetHash calculates and sets block hash
func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
