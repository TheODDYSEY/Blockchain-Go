# Blockchain Go

This project is a simple implementation of a blockchain in Go, featuring basic blockchain functionalities such as creating a blockchain, adding blocks, performing transactions, and checking balances. It also includes a network component to simulate multiple nodes interacting with each other.

## Features

- **Blockchain**: A chain of blocks, each containing transactions, a hash, and a reference to the previous block's hash.
- **Proof-of-Work**: A mechanism to ensure that adding new blocks requires computational work.
- **Persistence**: Blocks are stored in a BoltDB database for persistence.
- **Transactions**: Ability to send and receive coins between addresses.
- **Balance Checking**: Ability to check the balance of any address.
- **Networking**: Simulate multiple nodes interacting with each other.

## Prerequisites

- Go 1.16 or later
- BoltDB

## Installation

### Initialize the Go module:

```bash
go mod init blockchain

Install dependencies:

go get github.com/boltdb/bolt
go get github.com/stretchr/testify/assert
go get golang.org/x/crypto/ripemd160

Build the project:
go build -o blockchain_go *.go

Usage
Set Up Node 3000
Open Terminal for Node 3000: Set the NODE_ID environment variable to 3000.

export NODE_ID=3000

Generate a Wallet Address: Use the createwallet command to generate a new wallet address.

./blockchain_go createwallet

This will output a new wallet address. Let's assume the generated address is 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7.

Create a Blockchain with the Valid Address: Use the generated wallet address to create the blockchain.

./blockchain_go createblockchain -address 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7

Save the Genesis Block: Copy the blockchain database to a genesis block file.
cp blockchain_3000.db blockchain_genesis.db

Set Up Node 3001
Open Terminal for Node 3001: Set the NODE_ID environment variable to 3001.
export NODE_ID=3001

Generate Wallet Addresses for Node 3001: Generate some wallet addresses.

./blockchain_go createwallet
./blockchain_go createwallet
./blockchain_go createwallet

Let's assume the generated addresses are:

WALLET_1: 14DhJuzJuwnQTakCk9bbXGeURcW7iVxnJM
WALLET_2: 1KfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
WALLET_3: 1PfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
Send Coins to Wallet Addresses: Switch back to Terminal 1 (Node 3000) and send some coins from 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7 to the wallet addresses
.

./blockchain_go send -from 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7 -to 14DhJuzJuwnQTakCk9bbXGeURcW7iVxnJM -amount 10 -mine
./blockchain_go send -from 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7 -to 1KfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR -amount 10 -mine

Start the Node: Start the node.

./blockchain_go startnode

Set Up Node 3001 with Genesis Block
Copy the Genesis Block: Copy the genesis block file to the blockchain database for node 3001.
sudo cp blockchain_genesis.db blockchain_3001.db

Run the Node: Start the node.

./blockchain_go getbalance -address 14DhJuzJuwnQTakCk9bbXGeURcW7iVxnJM
./blockchain_go getbalance -address 1KfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
./blockchain_go getbalance -address 19Ephd29bSdbC4gzizu58WrbCjBaCdUVc7

Set Up Node 3002
Open Terminal for Node 3002: Set the NODE_ID environment variable to 3002.
export NODE_ID=3002

Generate a Wallet for Node 3002: Generate a wallet for the miner node.

./blockchain_go createwallet

Let's assume the generated address is 1RfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR.

Initialize the Blockchain: Copy the genesis block file to the blockchain database for node 3002.

sudo cp blockchain_genesis.db blockchain_3002.db

Start the Node: Start the node with the miner wallet.

./blockchain_go startnode -miner 1RfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR

Perform Transactions
Send Coins from WALLET_1 and WALLET_2: Send some coins from WALLET_1 and WALLET_2 to other addresses.

./blockchain_go send -from 14DhJuzJuwnQTakCk9bbXGeURcW7iVxnJM -to 1PfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR -amount 1
./blockchain_go send -from 1KfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR -to 1QfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR -amount 1

Mine the New Block: Switch to the miner node (Terminal 4, Node 3002) and see it mining a new block.

Download the Newly Mined Block: Switch to the wallet node (Terminal 3, Node 3001) and start it to download the newly mined block.

./blockchain_go startnode

Check Balances: Stop the node and check the balances.

./blockchain_go getbalance -address 14DhJuzJuwnQTakCk9bbXGeURcW7iVxnJM
./blockchain_go getbalance -address 1KfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
./blockchain_go getbalance -address 1PfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
./blockchain_go getbalance -address 1QfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR
./blockchain_go getbalance -address 1RfVdDvvQ977oVCpUqz7zAPUEiXKrX5avR


Conclusion
By following these steps, you should be able to create a blockchain using valid wallet addresses and perform the scenario successfully. This project demonstrates the basic functionalities of a blockchain, including creating blocks, performing transactions, and interacting with multiple nodes in a network.


This README file provides an overview of the project, installation instructions, usage examples, and detailed commands to set up and interact with multiple nodes in your blockchain network.
This README file provides an overview of the project, installation instructions, usage examples, and detailed commands to set up and interact with multiple nodes in your blockchain network.