# Proof-of-Work Implementation

## Introduction

This project implements a basic blockchain in Go, focusing on the Proof-of-Work (PoW) mechanism inspired by Bitcoin. The PoW concept ensures that adding new blocks to the blockchain is computationally intensive, thereby securing the network against tampering.

## Proof-of-Work Mechanism

The Proof-of-Work mechanism implemented here involves:
- Generating a block with a specific difficulty target.
- Finding a nonce (counter) such that the SHA-256 hash of the block's data meets the difficulty target.
- This process ensures that creating a new block (mining) requires significant computational effort.

## Implementation Details

### Components
- **Block**: Represents a unit of data in the blockchain.
- **ProofOfWork**: Manages the PoW logic for a block.
- **Blockchain**: Manages a sequence of blocks using linked references.

### Workflow
1. **NewBlock**: Creates a new block with provided data and previous block hash.
2. **NewProofOfWork**: Initializes a PoW instance for a given block.
3. **Run**: Executes the mining process by incrementing the nonce until the hash meets the target difficulty.
4. **Validate**: Checks if a mined block's PoW is valid.

## Usage

To run the example:
```shell
go run main.go
```

This will demonstrate mining blocks with PoW and validating their hashes.

## Next Steps

This implementation lays the groundwork for a blockchain with:
- Basic block structure and hashing.
- PoW for block creation.
- Validation of PoW for mined blocks.

Future enhancements could include:
- Persistent blockchain storage.
- Transactions, wallets, and addresses.
- Consensus mechanisms (e.g., Proof-of-Stake).

## Resources

- [Blockchain Hashing Algorithm](https://en.bitcoin.it/wiki/Block_hashing_algorithm)
- [Proof of Work](https://en.bitcoin.it/wiki/Proof_of_work)
