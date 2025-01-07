package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "fmt"
    "log"
    "math"
    "math/big"
    "strconv"
    "time"
)

// Block represents a block in the blockchain
type Block struct {
    Timestamp     int64
    Data          []byte
    PrevBlockHash []byte
    Hash          []byte
    Nonce         int
}

// Blockchain represents a blockchain
type Blockchain struct {
    blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := NewBlock(data, prevBlock.Hash)
    bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock creates a new genesis block
func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewBlock creates and returns a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
    pow := NewProofOfWork(block)
    nonce, hash := pow.Run()

    block.Hash = hash[:]
    block.Nonce = nonce

    return block
}

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
    block  *Block
    target *big.Int
}

var (
    maxNonce = math.MaxInt64
)

const targetBits = 16 // Reduce the difficulty

// IntToHex converts an int64 to a hex string
func IntToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }

    return buff.Bytes()
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-targetBits))

    pow := &ProofOfWork{b, target}

    return pow
}

// prepareData prepares data for hashing
func (pow *ProofOfWork) prepareData(nonce int) []byte {
    data := bytes.Join(
        [][]byte{
            pow.block.PrevBlockHash,
            pow.block.Data,
            IntToHex(pow.block.Timestamp),
            IntToHex(int64(targetBits)),
            IntToHex(int64(nonce)),
        },
        []byte{},
    )

    return data
}

// Run performs the proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
    var hashInt big.Int
    var hash [32]byte
    nonce := 0

    fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
    start := time.Now()
    for nonce < maxNonce {
        data := pow.prepareData(nonce)
        hash = sha256.Sum256(data)
        fmt.Printf("\r%x", hash)
        hashInt.SetBytes(hash[:])

        if hashInt.Cmp(pow.target) == -1 {
            break
        } else {
            nonce++
        }
    }
    duration := time.Since(start)
    fmt.Printf("\n\nMining completed in %s\n", duration)
    fmt.Printf("Nonce: %d\n", nonce)
    fmt.Printf("Hash: %x\n", hash[:])

    return nonce, hash[:]
}

// Validate validates the proof-of-work
func (pow *ProofOfWork) Validate() bool {
    var hashInt big.Int

    data := pow.prepareData(pow.block.Nonce)
    hash := sha256.Sum256(data)
    hashInt.SetBytes(hash[:])

    isValid := hashInt.Cmp(pow.target) == -1

    return isValid
}

// main is the entry point for the application
func main() {
    bc := NewBlockchain()

    bc.AddBlock("Send 1 BTC to Ivan")
    bc.AddBlock("Send 2 more BTC to Ivan")

    for _, block := range bc.blocks {
        fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }
}