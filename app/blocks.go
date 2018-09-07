package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Car represents car data
type Car struct {
	LicensePlate string
	Owner        string
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Car       Car
	Hash      string
	PrevHash  string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// CarInput takes incoming JSON payload for writing heart rate
type CarInput struct {
	LicensePlate string `json:"license_plate"`
	Owner        string `json:"owner"`
}

// HackInput takes incoming JSON payload for writing heart rate
type HackInput struct {
	Index int    `json:"index"`
	Hash  string `json:"hash"`
	Owner string `json:"owner"`
}

// ReplaceChain make sure the chain we're checking is longer than the current blockchain
func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

// IsBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// CalculateHash SHA256 hasing
func CalculateHash(block Block) string {
	carJSON, err := json.Marshal(block.Car)
	if err != nil {
		panic(err)
	}

	record := string(block.Index) + block.Timestamp + string(carJSON) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock create a new block using previous block's hash
func GenerateBlock(lp string, o string) (Block, error) {

	car := Car{
		LicensePlate: lp,
		Owner:        o,
	}

	oldBlock := Blockchain[len(Blockchain)-1]

	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Car = car
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

// HackBlock edit some fields in some block
func HackBlock(index int, hash string, owner string) *Block {

	block := &Blockchain[index]

	block.Car.Owner = owner
	block.Hash = hash

	return block
}

func IsChainValid(bc []Block) bool {
	for i := 1; i < len(bc); i++ {
		currentBlock := bc[i]
		previousBlock := bc[i-1]
		if currentBlock.Hash != CalculateHash(currentBlock) {
			return false
		}
		if currentBlock.PrevHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
