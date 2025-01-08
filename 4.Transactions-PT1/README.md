# Building Blockchain in Go. Part 4: Transactions 1

This project is a simple implementation of a blockchain in Go, featuring basic blockchain functionalities such as creating a blockchain, adding blocks, performing transactions, and checking balances.

## Features

- **Blockchain**: A chain of blocks, each containing transactions, a hash, and a reference to the previous block's hash.
- **Proof-of-Work**: A mechanism to ensure that adding new blocks requires computational work.
- **Persistence**: Blocks are stored in a BoltDB database for persistence.
- **Transactions**: Ability to send and receive coins between addresses.
- **Balance Checking**: Ability to check the balance of any address.

## Prerequisites

- Go 1.16 or later
- BoltDB

## Installation

Initialize the Go module:
```
go mod init blockchain
```
Install dependencies:
```
go get github.com/boltdb/bolt
```

Build the project:
```
go build -o blockchain_go main.go
```

Create a Blockchain
To create a new blockchain and send the genesis block reward to a specified address:
```
blockchain_go createblockchain -address Ivan
```

Check Balance
To check the balance of a specific address:
```
blockchain_go getbalance -address Ivan

Balance of 'Ivan': 10
```

Send Coins
To send coins from one address to another:
```
blockchain_go send -from Ivan -to Pedro -amount 6

00000001b56d60f86f72ab2a59fadb197d767b97d4873732be505e0a65cc1e37

Success!
```



Check balance of Ivan
```
blockchain_go getbalance -address Ivan

Balance of 'Ivan': 4
```
Check balance of Pedro
```
blockchain_go getbalance -address Pedro

Balance of 'Pedro': 6
```

Send 2 coins from Pedro to Helen
```
blockchain_go send -from Pedro -to Helen -amount 2

00000099938725eb2c7730844b3cd40209d46bce2c2af9d87c2b7611fe9d5bdf
Success!
```

Send 2 coins from Ivan to Helen
```
blockchain_go send -from Ivan -to Helen -amount 2

000000a2edf94334b1d94f98d22d7e4c973261660397dc7340464f7959a7a9aa
Success!
```

Send 3 coins from Helen to Rachel
```
blockchain_go send -from Helen -to Rachel -amount 3

000000c58136cffa669e767b8f881d16e2ede3974d71df43058baaf8c069f1a0
 Success!
```


Check balance of Ivan
```
blockchain_go getbalance -address Ivan

Balance of 'Ivan': 2
```

Check balance of Pedro
```
blockchain_go getbalance -address Pedro

Balance of 'Pedro': 4
```

Check balance of Helen
```
blockchain_go getbalance -address Helen

Balance of 'Helen': 1
```

Check balance of Rachel
```
blockchain_go getbalance -address Rachel

Balance of 'Rachel': 3
```

### Attempt to send 5 coins from Pedro to Ivan (should fail due to insufficient funds)
```
blockchain_go send -from Pedro -to Ivan -amount 5

Output: panic: ERROR: Not enough funds
```

Check balance of Pedro
```
blockchain_go getbalance -address Pedro
Output: Balance of 'Pedro': 4
```

Check balance of Ivan
```
blockchain_go getbalance -address Ivan
Output: Balance of 'Ivan': 2
```
## How It Works

### Blockchain

- **Block**: Contains a timestamp, transactions, the hash of the previous block, its own hash, and a nonce.
- **Blockchain**: A sequence of blocks, starting with a genesis block.

### Proof-of-Work

- **ProofOfWork**: Ensures that adding a new block requires computational work by finding a hash that meets a certain target.

### Persistence

- **BoltDB**: Used to store blocks persistently. Each block is stored with its hash as the key.

### Transactions

- **Transaction**: Represents a transfer of coins from one address to another.
- **UTXO (Unspent Transaction Output)**: Represents unspent outputs from previous transactions that can be used as inputs in new transactions.