# Bitcoin Implementation in Go

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![BoltDB](https://img.shields.io/badge/BoltDB-Database-orange?style=flat-square)](https://github.com/boltdb/bolt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![ECDSA](https://img.shields.io/badge/ECDSA-Cryptography-blue?style=flat-square)](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
[![P2P](https://img.shields.io/badge/P2P-Network-green?style=flat-square)](https://en.wikipedia.org/wiki/Peer-to-peer)

A step-by-step implementation of a Bitcoin-like blockchain in Go, from basic prototype to a fully functional P2P network. This educational project demonstrates core blockchain concepts through progressive implementation stages.

## Project Stages

Each directory in this repository represents a specific stage in building a blockchain from scratch:

### 1. Basic Prototype

**Core Concepts:**
- Blockchain data structure fundamentals
- Basic block creation and chaining
- SHA-256 cryptographic hashing

**Implementation Details:**
- `Block` struct with timestamp, data, and hash fields
- `Blockchain` container to manage the sequence of blocks
- Genesis block generation
- Simple block validation

**Commands:**
```bash
cd 1.Basic-Prototype
go run main.go
```
Running this will create a blockchain with a genesis block and add two sample blocks, displaying each block's details.

### 2. Proof-of-Work

**Core Concepts:**
- Mining with computational puzzles
- Hash-based proof of work similar to Bitcoin
- Difficulty adjustment

**Implementation Details:**
- `ProofOfWork` struct tied to each block
- Difficulty target with leading zeros
- Nonce discovery through computational work
- Hash validation

**Commands:**
```bash
cd 2.Proof-of-Work
go run main.go
```
This demonstrates mining blocks with increasing difficulty and validates the mined blocks.

### 3. Persistence & CLI

**Core Concepts:**
- Database storage for blockchain persistence
- Command-line interface for blockchain interaction
- Block retrieval and iteration

**Implementation Details:**
- BoltDB integration for blockchain storage
- CLI commands for adding blocks and viewing the chain
- Database buckets for organizing blockchain data

**Commands:**
```bash
cd 3.Persistence-and-CLI
go build -o blockchain_cli

# Print the blockchain (creates a new one if it doesn't exist)
./blockchain_cli printchain

# Add a new block
./blockchain_cli addblock -data "Send 1 BTC to Ivan"

# View the chain again
./blockchain_cli printchain
```

### 4. Transactions (Part 1)

**Core Concepts:**
- Bitcoin-style UTXO transaction model
- Transaction inputs and outputs
- Coinbase transactions (block rewards)

**Implementation Details:**
- `Transaction` struct with inputs and outputs
- `TXInput` and `TXOutput` for transaction components
- Balance calculation using transaction history
- UTXO identification and tracking

**Commands:**
```bash
cd 4.Transactions-PT1
go build -o blockchain_tx

# Create a blockchain with the genesis block
./blockchain_tx createblockchain -address Ivan

# Check balance
./blockchain_tx getbalance -address Ivan

# Send coins
./blockchain_tx send -from Ivan -to Pedro -amount 6

# Check balances after transaction
./blockchain_tx getbalance -address Ivan
./blockchain_tx getbalance -address Pedro
```

### 5. Addresses

**Core Concepts:**
- Cryptographic wallet implementation
- ECDSA-based key pairs
- Base58Check encoded addresses
- Transaction signing and verification

**Implementation Details:**
- `Wallet` struct with private/public key pair
- Address generation with RIPEMD160 hashing
- Digital signatures for transaction authorization
- Signature verification for transaction validity

**Commands:**
```bash
cd 5.Addresses
go build -o blockchain_wallet

# Create a new wallet
./blockchain_wallet createwallet
# Output: Your new address: 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt

# Create a blockchain with your wallet address
./blockchain_wallet createblockchain -address 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt

# Check your balance (should show mining reward)
./blockchain_wallet getbalance -address 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt

# Send coins to another address
./blockchain_wallet send -from 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt -to 15pUhCbtrGh3JUx5iHnXjfpyHyTgawvG5h -amount 6
```

### 6. Transactions (Part 2)

**Core Concepts:**
- UTXO set optimization
- Mining rewards integration
- Efficient transaction validation

**Implementation Details:**
- `UTXOSet` struct for caching unspent outputs
- Reindexing mechanism for UTXO set maintenance
- Coinbase transaction implementation for mining rewards
- UTXO set updates when new blocks are mined

**Commands:**
```bash
cd 6.Transactions-PT2
go build -o blockchain_tx2

# Create a wallet and blockchain
./blockchain_tx2 createwallet
# Output: Your new address: 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1

./blockchain_tx2 createblockchain -address 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1

# Create another wallet
./blockchain_tx2 createwallet
# Output: Your new address: 12DkLzLQ4B3gnQt62EPRJGZ38n3zF4Hzt5

# Send coins with mining reward (note the -mine flag)
./blockchain_tx2 send -from 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1 -to 12DkLzLQ4B3gnQt62EPRJGZ38n3zF4Hzt5 -amount 6 -mine

# Check balances
./blockchain_tx2 getbalance -address 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1
./blockchain_tx2 getbalance -address 12DkLzLQ4B3gnQt62EPRJGZ38n3zF4Hzt5
```

### 7. Network

**Core Concepts:**
- Peer-to-peer network communication
- Block and transaction broadcasting
- Distributed blockchain synchronization
- Multi-node mining and consensus

**Implementation Details:**
- Multiple node identification with NODE_ID
- TCP server implementation for node communication
- Message protocols for version, block, and transaction exchange
- Blockchain synchronization between nodes

**Commands:**
```bash
cd 7.Network
go build -o blockchain_network

# For each node, set a different NODE_ID and run in separate terminals

# Terminal 1: Create central node (3000)
export NODE_ID=3000
./blockchain_network createwallet
# Save this wallet address for blockchain creation
./blockchain_network createblockchain -address YOUR_WALLET_ADDRESS_3000
cp blockchain_3000.db blockchain_genesis.db
./blockchain_network startnode

# Terminal 2: Create wallet node (3001)
export NODE_ID=3001
./blockchain_network createwallet
# Save this wallet address
./blockchain_network startnode

# Terminal 3: Create miner node (3002)
export NODE_ID=3002
./blockchain_network createwallet
# Save this wallet address for mining
cp blockchain_genesis.db blockchain_3002.db
./blockchain_network startnode -miner YOUR_WALLET_ADDRESS_3002

# Send transaction from node 3000 (back in Terminal 1)
./blockchain_network send -from YOUR_WALLET_ADDRESS_3000 -to YOUR_WALLET_ADDRESS_3001 -amount 10

# Check balances on different nodes
./blockchain_network getbalance -address YOUR_WALLET_ADDRESS_3001
```

## Installation & Quick Start

### Requirements
- Go 1.16+
- BoltDB

### Installation
```bash
# Clone repository
git clone https://github.com/yourusername/blockchain-go.git
cd blockchain-go

# Install dependencies
go mod init blockchain
go get github.com/boltdb/bolt
go get golang.org/x/crypto/ripemd160

# Build any stage
cd <stage-directory>
go build -o blockchain_go *.go
```

### Basic Usage Workflow

1. **Create wallet**
   ```bash
   ./blockchain_go createwallet
   # Save the generated address
   ```

2. **Initialize blockchain**
   ```bash
   ./blockchain_go createblockchain -address YOUR_WALLET_ADDRESS
   ```

3. **Send coins**
   ```bash
   ./blockchain_go send -from SENDER -to RECIPIENT -amount 10
   ```

4. **Check balance**
   ```bash
   ./blockchain_go getbalance -address ADDRESS
   ```

5. **Run a node (with optional miner)**
   ```bash
   export NODE_ID=3000
   ./blockchain_go startnode -miner MINER_ADDRESS
   ```

## Key Features

- ⛓️ **Blockchain Core** - Immutable chain of cryptographically linked blocks
- ⚒️ **Mining** - Bitcoin-inspired Proof-of-Work consensus with adjustable difficulty
- 🔐 **Wallets** - ECDSA key pairs with Base58 address encoding and secure key storage
- 💸 **Transactions** - Secure UTXO model with digital signatures and transaction verification
- 📡 **P2P Network** - Node communication, block propagation, and distributed consensus
- 💾 **Persistence** - Efficient blockchain storage with BoltDB and UTXO set optimization

## Learning Value

This project demonstrates core blockchain concepts including:

- **Cryptography**: SHA-256 hashing, ECDSA signatures, public/private key infrastructure
- **Consensus**: Proof-of-Work mining with difficulty adjustment
- **Distributed Systems**: P2P networking, blockchain synchronization
- **Database Design**: Efficient blockchain and UTXO storage
- **Digital Assets**: Transaction models, wallet management, and balance tracking

Each implementation stage builds upon the previous one, providing a comprehensive understanding of blockchain technology from the ground up.

## References

- [Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
- [Mastering Bitcoin](https://github.com/bitcoinbook/bitcoinbook)
- [Go Documentation](https://golang.org/doc/)
- [BoltDB](https://github.com/boltdb/bolt)

## License

MIT License