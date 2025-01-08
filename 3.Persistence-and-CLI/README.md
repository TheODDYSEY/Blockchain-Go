# Building Blockchain in Go. Part 3: Persistence and CLI

This project is a simple implementation of a blockchain in Go, featuring basic blockchain functionalities such as adding blocks, proof-of-work, and persistence using BoltDB.

## Features

- **Blockchain**: A chain of blocks, each containing data, a hash, and a reference to the previous block's hash.
- **Proof-of-Work**: A mechanism to ensure that adding new blocks requires computational work.
- **Persistence**: Blocks are stored in a BoltDB database for persistence.

## Prerequisites

- Go 1.16 or later
- BoltDB

## Installation

### Initialize the Go module:
```
go mod init blockchain
```
### Install dependencies:
```
go get github.com/boltdb/bolt
```
### Build the project:
```
go build -o blockchain_go main.go
```
## Usage

### Add a Block

To add a new block to the blockchain, use the `addblock` command followed by the data for the block:
```
./blockchain_go addblock -data "Send 1 BTC to Ivan"
```
### Print the Blockchain

To print all the blocks in the blockchain, use the `printchain` command:
```
./blockchain_go printchain
```
### Example Commands

#### Print the blockchain (this will create a new blockchain if none exists)
```
./blockchain_go printchain
```
#### Add a new block with data "Send 1 BTC to Ivan"
```
./blockchain_go addblock -data "Send 1 BTC to Ivan"
```
#### Add another block with data "Pay 0.31337 BTC for a coffee"
```
./blockchain_go addblock -data "Pay 0.31337 BTC for a coffee"
```

#### Print the blockchain again to see the added block
```
./blockchain_go printchain
```

## How It Works

### Blockchain

- **Block**: Contains a timestamp, data, the hash of the previous block, its own hash, and a nonce.
- **Blockchain**: A sequence of blocks, starting with a genesis block.

### Proof-of-Work

- **ProofOfWork**: Ensures that adding a new block requires computational work by finding a hash that meets a certain target.

### Persistence

- **BoltDB**: Used to store blocks persistently. Each block is stored with its hash as the key.

## Viewing the Database

To view the contents of the BoltDB database, you can use tools like Bolt CLI or Bolt Browser, or write a simple Go script to read and print the database contents.

## Viewing the Database with Bolt CLI
You can use the Bolt CLI tool to inspect the contents of your BoltDB database.

### Install Bolt CLI

First, install the Bolt CLI tool:
```
go install github.com/boltdb/bolt/cmd/bolt@latest
```


### View the Database

Navigate to the directory containing your `blockchain.db` file and use the following commands:

#### Get Basic Info
```
bolt info blockchain.db
```
#### Get Database Stats
```
bolt stats blockchain.db
```
#### List Pages
```
bolt pages blockchain.db
```
## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.