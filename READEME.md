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
```

### Install dependencies:

```bash
go get github.com/boltdb/bolt
go get github.com/stretchr/testify/assert
go get golang.org/x/crypto/ripemd160
```

### Build the project:

```bash
go build -o blockchain_go *.go
```

## Usage

### Set Up Node 3000

1. **Open Terminal for Node 3000**: Set the `NODE_ID` environment variable to 3000.
   
   ```bash
   export NODE_ID=3000
   ```

2. **Generate a Wallet Address**: Use the `createwallet` command to generate a new wallet address.

   ```bash
   ./blockchain_go createwallet
   ```

3. **Create a Blockchain**: Use the generated wallet address to create the blockchain.

   ```bash
   ./blockchain_go createblockchain -address <your_wallet_address>
   ```

4. **Save the Genesis Block**: Copy the blockchain database to a genesis block file.

   ```bash
   cp blockchain_3000.db blockchain_genesis.db
   ```

### Set Up Node 3001

1. **Open Terminal for Node 3001**: Set the `NODE_ID` environment variable to 3001.

   ```bash
   export NODE_ID=3001
   ```

2. **Generate Wallet Addresses**: Generate wallet addresses for Node 3001.

   ```bash
   ./blockchain_go createwallet
   ./blockchain_go createwallet
   ./blockchain_go createwallet
   ```

3. **Send Coins**: Use Terminal 1 (Node 3000) to send coins to the wallet addresses.

   ```bash
   ./blockchain_go send -from <your_wallet_address> -to <wallet_1_address> -amount 10 -mine
   ```

4. **Start the Node**: Start the node.

   ```bash
   ./blockchain_go startnode
   ```

### Set Up Node 3002

1. **Open Terminal for Node 3002**: Set the `NODE_ID` environment variable to 3002.

   ```bash
   export NODE_ID=3002
   ```

2. **Generate a Wallet for Node 3002**:

   ```bash
   ./blockchain_go createwallet
   ```

3. **Copy the Genesis Block**: Copy the genesis block file to the blockchain database for Node 3002.

   ```bash
   sudo cp blockchain_genesis.db blockchain_3002.db
   ```

4. **Start the Node**: Start the node with the miner wallet.

   ```bash
   ./blockchain_go startnode -miner <miner_wallet_address>
   ```

### Perform Transactions

1. **Send Coins**: Send coins from one wallet to another.

   ```bash
   ./blockchain_go send -from <wallet_1_address> -to <wallet_2_address> -amount 1
   ```

2. **Mine the New Block**: Mine a new block using the miner node.

3. **Check Balances**: Check the balances of the addresses.

   ```bash
   ./blockchain_go getbalance -address <wallet_address>
   ```

## Conclusion

By following these steps, you can create a blockchain using valid wallet addresses and perform transactions across multiple nodes. This project demonstrates essential blockchain functionalities, including creating blocks, handling transactions, and managing a network of nodes.
