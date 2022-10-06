package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	Transactions string
	Nonce        string
	Previoushash string
	Currenthash  string
}

func (b *Block) Calculatehash(stringtoHash string) {
	hash := sha256.Sum256([]byte(stringtoHash))
	b.Currenthash = string(hash[:])
}

type List struct {
	Chain    []Block
	LastHash string
}

func (ls *List) NewBlock(transactions string, nonce string, previoushash string) *Block {
	new_blk := new(Block)
	new_blk.Transactions = transactions
	new_blk.Nonce = nonce
	new_blk.Previoushash = previoushash
	new_blk.Calculatehash(transactions + nonce + previoushash)
	ls.Chain = append(ls.Chain, *new_blk)
	ls.LastHash = new_blk.Currenthash
	return new_blk
}

func (ls *List) Listblock() {
	for i := range ls.Chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Printf("Transaction: %s\n", ls.Chain[i].Transactions)
		fmt.Printf("Transaction: %s\n", ls.Chain[i].Nonce)
		fmt.Printf("previous bloch hash: %x\n", ls.Chain[i].Previoushash)
		fmt.Printf("current block hash: %x\n\n\n", ls.Chain[i].Currenthash)
	}

	if ls.Verify() {
		fmt.Printf("Blockchain is not changed.\n\n\n\n\n")
	} else {
		fmt.Printf("Blockchain is changed.\n\n")
	}
}

func (ls *List) Changeblock() {
	ls.Chain[0].Transactions = "Tihamk sent 300 coins to Talha Pasha"
}

func (ls *List) Verify() bool {
	for i := range ls.Chain {
		if i == 0 {
			ls.Chain[i].Calculatehash(ls.Chain[i].Transactions + ls.Chain[i].Nonce + ls.Chain[i].Previoushash)
		} else {
			ls.Chain[i].Calculatehash(ls.Chain[i].Transactions + ls.Chain[i].Nonce + ls.Chain[i-1].Currenthash)
		}
	}

	if ls.LastHash == ls.Chain[len(ls.Chain)-1].Currenthash {
		return true
	} else {
		return false
	}
}

func main() {
	blockChain := new(List)
	blockChain.NewBlock("Tihami sent 1 coin to Talha Pasha", "12345", "")
	blockChain.NewBlock("Ali sent 200 coins to Abdullah", "6789", blockChain.Chain[len(blockChain.Chain)-1].Currenthash)
	blockChain.NewBlock("Zaka sent 2000000 coins to zaka", "8978789879", blockChain.Chain[len(blockChain.Chain)-1].Currenthash)

	blockChain.Listblock()
	blockChain.Changeblock()
	blockChain.Listblock()

}
