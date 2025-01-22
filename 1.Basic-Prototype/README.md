# Basic Blockchain Implementation in Go

## Overview
This project implements a basic blockchain prototype in Go. It demonstrates the fundamental concepts of blockchain technology, including blocks, chain formation, and basic data storage.

## Features
- Basic blockchain data structure
- Block creation and mining
- SHA256 hashing
- Genesis block initialization
- Chain validation
- Data storage in blocks

## Technical Details

### Block Structure
Each block in the blockchain contains:
- Timestamp: Unix timestamp of block creation
- Data: The actual data stored in the block
- PrevBlockHash: Hash of the previous block
- Hash: Current block's hash
