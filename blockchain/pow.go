package blockchain

import (
	"BeegoDemo/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const DIFFFICULTY = 12

/**
 *工作量证明结构体
 */
type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明算法对应的哪个区块
	Block Block
}

/**
 * 实例化一个pow算法实例
 */
func NewPoW(block Block)ProofOfWork{
	target := big.NewInt(1)//初始值
	target.Lsh(target,255-DIFFFICULTY)//左移
	pow := ProofOfWork{
		Target: target,
		Block: block,
	}
	return pow
}
/**
 * pow算法：寻找符合条件的nonce值
 */
func (p ProofOfWork) Run() ([]byte,int64){
	var nonce int64
	//var bigBlock *big.Int//声明
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block := p.Block
		heightBytes,_ :=util.IntToBytes(block.Height)
		timeBytes,_ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes,_ := util.IntToBytes(nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})
		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash = sha256Hash.Sum(nil)
		bigBlock = bigBlock.SetBytes(block256Hash)
		if p.Target.Cmp(bigBlock) == 1{
			break
		}
		nonce++
	}
	return block256Hash, nonce
}
