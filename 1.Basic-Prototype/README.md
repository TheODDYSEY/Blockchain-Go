# Basic Prototype

This project provides a simple implementation of a blockchain in Go, demonstrating the basics of data structures, cryptographic hashing, and decentralization concepts. 

## Overview

A blockchain is a distributed and secure digital ledger that chains blocks of data together. Each block contains a timestamp, some data, a hash of the previous block, and its own hash, ensuring data integrity and securing it against tampering.

### Key Components

- **Block**: Represents a single block in the blockchain, containing fields for storing important information about each block.
  - `Timestamp`: The time at which the block was created.
  - `Data`: The actual information stored in the block (e.g., transaction details).
  - `PrevBlockHash`: The hash of the previous block in the chain.
  - `Hash`: The unique hash of the current block computed based on its content.

- **Blockchain**: Represents the entire chain of blocks.
  - `blocks`: A slice of pointers to `Block` structs, representing all the blocks in the blockchain.

### Functions

- **NewBlock**: Creates a new block using the provided data and the hash of the previous block, automatically setting its hash with `SetHash`.

- **AddBlock**: Adds a new block to the blockchain by creating it from the provided data.

- **NewGenesisBlock**: Creates the very first block of the blockchain, known as the genesis block.

- **NewBlockchain**: Initializes a new blockchain with the genesis block.

### Usage

To run the blockchain example:

1. Clone this repository or copy the provided code into a `.go` file (e.g., `main.go`).
2. Open your terminal and navigate to the directory containing the Go file.
3. Run the Go program with the command:
```
go run main.go
```

4. The program will create a new blockchain, add two blocks to it, and print the details of each block in the blockchain.   

Output
The output will display the previous block's hash, the data contained in each block, and its own hash, for example:
```
Prev. hash: 0000000000000000000000000000000000000000000000000000000000000000
Data: Genesis Block
Hash: 6c2c84ffb4da0c47e5c9ace5507c9d062a0214c39e5938280080b9c7a3a5976a

Prev. hash: 6c2c84ffb4da0c47e5c9ace5507c9d062a0214c39e5938280080b9c7a3a5976a
Data: Send 1 BTC to Ivan
Hash: 68bfb9be140aece9c3f8c477f2dadf7931f87fa6751fc7ed08f72e9b8576a8a1

Prev. hash: 68bfb9be140aece9c3f8c477f2dadf7931f87fa6751fc7ed08f72e9b8576a8a1
Data: Send 2 more BTC to Ivan
Hash: 43bc7aef2da54d2f44bbdb5c6af57ad3ab0d99fefbf22e48c5cd8b8c06e5c4eb

```

Conclusion
This simple Go implementation provides a foundation for understanding how blockchains function. 