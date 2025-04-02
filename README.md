# 🔗 BlockchainGo: Distributed Ledger Implementation

<div align="center">

![BlockchainGo](https://img.shields.io/badge/BlockchainGo-v1.0-brightgreen?style=for-the-badge&logo=bitcoin&logoColor=white)
[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![BoltDB](https://img.shields.io/badge/BoltDB-Database-orange?style=for-the-badge&logo=databricks&logoColor=white)](https://github.com/boltdb/bolt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](https://opensource.org/licenses/MIT)

</div>

<p align="center">
  <b>A professional Bitcoin-like blockchain implementation in Go, built from the ground up.</b>
</p>

<p align="center">
  <a href="#-overview">Overview</a> •
  <a href="#-implementation-stages">Implementation Stages</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-key-features">Key Features</a> •
  <a href="#-technical-concepts">Technical Concepts</a> •
  <a href="#-contributing">Contributing</a> •
  <a href="#-resources">Resources</a>
</p>

---

## 🌐 Overview

> **BlockchainGo** is a complete educational implementation of a Bitcoin-like blockchain, showcasing fundamental concepts of distributed ledger technology through a progressive, step-by-step approach.

This project demonstrates the core principles underlying blockchain technology, from basic data structures to complex distributed consensus mechanisms. Each implementation stage builds upon previous concepts, providing a comprehensive understanding of how cryptocurrencies work.

<div align="center">
  <img src="https://komodoplatform.com/en/academy/content/images/2024/07/BlockchainF1.jpg" alt="Blockchain Architecture" width="650">
</div>

## 📚 Implementation Stages

The repository is organized into seven stages, each representing a significant advancement in blockchain technology:

### 🧱 1. Basic Prototype

<details>
<summary><b>Foundation of blockchain architecture</b></summary>

**Core Concepts:**
- 📊 Blockchain data structure fundamentals
- 🔄 Basic block creation and chaining
- 🔒 SHA-256 cryptographic hashing

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
</details>

### ⛏️ 2. Proof-of-Work

<details>
<summary><b>Mining and computational consensus</b></summary>

**Core Concepts:**
- 💻 Mining with computational puzzles
- 🧮 Hash-based proof of work similar to Bitcoin
- 📈 Difficulty adjustment

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
</details>

### 💾 3. Persistence & CLI

<details>
<summary><b>Database storage and command interface</b></summary>

**Core Concepts:**
- 🗄️ Database storage for blockchain persistence
- 🖥️ Command-line interface for blockchain interaction
- 🔍 Block retrieval and iteration

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
```
</details>

### 💰 4. Transactions (Part 1)

<details>
<summary><b>UTXO transaction model implementation</b></summary>

**Core Concepts:**
- 💸 Bitcoin-style UTXO transaction model
- 📥 Transaction inputs and outputs
- 🏆 Coinbase transactions (block rewards)

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

# Send coins
./blockchain_tx send -from Ivan -to Pedro -amount 6
```
</details>

### 🔐 5. Addresses

<details>
<summary><b>Cryptographic wallets and signatures</b></summary>

**Core Concepts:**
- 🗝️ Cryptographic wallet implementation
- 📝 ECDSA-based key pairs
- 🏷️ Base58Check encoded addresses
- ✍️ Transaction signing and verification

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
```
</details>

### ⚡ 6. Transactions (Part 2)

<details>
<summary><b>UTXO optimization and mining rewards</b></summary>

**Core Concepts:**
- 🚀 UTXO set optimization
- 💎 Mining rewards integration
- ⚡ Efficient transaction validation

**Implementation Details:**
- `UTXOSet` struct for caching unspent outputs
- Reindexing mechanism for UTXO set maintenance
- Coinbase transaction implementation for mining rewards
- UTXO set updates when new blocks are mined

**Commands:**
```bash
cd 6.Transactions-PT2
go build -o blockchain_tx2

# Send coins with mining reward
./blockchain_tx2 send -from WALLET_ADDRESS -to RECIPIENT -amount 6 -mine
```
</details>

### 🌐 7. Network

<details>
<summary><b>P2P networking and distributed consensus</b></summary>

**Core Concepts:**
- 📡 Peer-to-peer network communication
- 📢 Block and transaction broadcasting
- 🔄 Distributed blockchain synchronization
- 🤝 Multi-node mining and consensus

**Implementation Details:**
- Multiple node identification with NODE_ID
- TCP server implementation for node communication
- Message protocols for version, block, and transaction exchange
- Blockchain synchronization between nodes

**Commands:**
```bash
cd 7.Network
go build -o blockchain_network

# Start a node with mining capabilities
export NODE_ID=3000
./blockchain_network startnode -miner YOUR_WALLET_ADDRESS
```
</details>

## 🚀 Installation

### Prerequisites

- 🔧 Go 1.16+
- 🗄️ BoltDB

### Getting Started

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

## 🔍 Usage

<div align="center">
  <img src="https://www.researchgate.net/profile/Mrityunjay-Ghosh/publication/355223429/figure/fig3/AS:1079123128659968@1634294411034/Flow-of-a-transaction-in-Blockchain-technology-39.ppm" alt="Blockchain Flow" width="700">
</div>


### 👛 Wallet Management

```bash
# Generate a new wallet address
./blockchain_go createwallet
# Output: Your new address: 13Uu7B1vDP4ViXqHFsWtbraM3EfQ3UkWXt

# List all addresses in your wallet
./blockchain_go listaddresses
```

### ⛓️ Blockchain Operations

```bash
# Create a new blockchain (first-time setup)
./blockchain_go createblockchain -address YOUR_WALLET_ADDRESS

# Print the entire blockchain
./blockchain_go printchain
```

### 💸 Transaction Management

```bash
# Send coins from one address to another
./blockchain_go send -from SENDER_ADDRESS -to RECIPIENT_ADDRESS -amount 10

# Check wallet balance
./blockchain_go getbalance -address WALLET_ADDRESS
```

### 🌐 Network Operations

```bash
# Set your node ID
export NODE_ID=3000

# Start a node (with optional mining)
./blockchain_go startnode -miner MINER_ADDRESS
```

## 🛠️ Key Features

- 🔗 **Immutable Ledger** - Cryptographically secure blockchain with tamper-proof blocks
- ⚒️ **Mining Engine** - Complete Proof-of-Work implementation with adjustable difficulty
- 👛 **Wallet System** - Secure ECDSA key generation with Base58 address encoding
- 💰 **UTXO Architecture** - Bitcoin-style transaction model with unspent transaction outputs
- 🔏 **Cryptographic Security** - Transaction signing and verification with ECDSA
- 📡 **P2P Network** - Full node-to-node communication with blockchain synchronization
- 🗄️ **Efficient Storage** - Optimized persistence with BoltDB and UTXO caching

## 📘 Technical Concepts

<table>
  <tr>
    <td align="center"><img src="https://img.icons8.com/color/48/000000/blockchain-technology.png"/><br/>Blockchain Structure</td>
  <td align="center"><img src="https://img.icons8.com/ios-filled/50/000000/lock.png"/><br/>Cryptography</td>
    <td align="center"><img src="https://img.icons8.com/color/48/000000/database-restore.png"/><br/>Consensus</td>
    <td align="center"><img src="https://img.icons8.com/color/48/000000/network.png"/><br/>P2P Networking</td>
  </tr>
  <tr>
    <td>
      <ul>
        <li>Block structure with headers and data</li>
        <li>Merkle trees for transaction validation</li>
        <li>Linked hash-based architecture</li>
      </ul>
    </td>
    <td>
      <ul>
        <li>SHA-256 hashing algorithms</li>
        <li>ECDSA for digital signatures</li>
        <li>Public/private key infrastructure</li>
      </ul>
    </td>
    <td>
      <ul>
        <li>Proof-of-Work mining</li>
        <li>Difficulty adjustment mechanism</li>
        <li>Chain selection logic</li>
      </ul>
    </td>
    <td>
      <ul>
        <li>Node discovery protocol</li>
        <li>Block propagation</li>
        <li>Transaction broadcasting</li>
      </ul>
    </td>
  </tr>
</table>

## 👥 Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

1. 🍴 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add some amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🎁 Open a Pull Request

## 📚 Resources

- 📄 [Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
- 📕 [Mastering Bitcoin](https://github.com/bitcoinbook/bitcoinbook)
- 📘 [Go Documentation](https://golang.org/doc/)
- 📙 [BoltDB](https://github.com/boltdb/bolt)
- 📺 [Blockchain Tutorials](https://www.youtube.com/results?search_query=blockchain+golang+tutorial)

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

<div align="center">

⭐ **Found this project useful? Consider giving it a star!** ⭐

</div>

