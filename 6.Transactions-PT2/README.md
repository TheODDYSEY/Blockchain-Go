# Transactions (Part 6)

## Introduction
This part of the blockchain implementation in Go focuses on enhancing transactions, particularly integrating mining rewards and optimizing transaction lookup using the UTXO set. 

## Previous Parts:
1. Basic Prototype
2. Proof-of-Work
3. Persistence and CLI
4. Transactions 1
5. Addresses

## Key Enhancements in This Part:
- Implementing mining rewards using a **coinbase transaction**.
- Introducing the **UTXO Set** to improve efficiency in finding unspent transactions.
- Optimizing transaction lookups to avoid scanning the entire blockchain for validation.

## Reward Implementation
The mining reward is given through a **coinbase transaction** that is added to the beginning of each block’s transactions.

### Updated `send` Command:
```go
func (cli *CLI) send(from, to string, amount int) {
    bc := NewBlockchain()
    UTXOSet := UTXOSet{bc}
    defer bc.db.Close()

    tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
    cbTx := NewCoinbaseTX(from, "")
    txs := []*Transaction{cbTx, tx}

    newBlock := bc.MineBlock(txs)
    fmt.Println("Success!")
}
```
Here, the miner receives the reward for creating a new block.

## Introducing the UTXO Set
The **UTXO Set** (Unspent Transaction Output Set) improves efficiency by keeping a cache of unspent outputs, allowing transactions to be validated quickly without scanning the entire blockchain.

### `UTXOSet` Struct:
```go
type UTXOSet struct {
    Blockchain *Blockchain
}
```
This struct is tied to the blockchain but stored separately to optimize lookups.

### UTXO Set Reindexing:
```go
func (u UTXOSet) Reindex() {
    db := u.Blockchain.db
    bucketName := []byte(utxoBucket)

    err := db.Update(func(tx *bolt.Tx) error {
        err := tx.DeleteBucket(bucketName)
        _, err = tx.CreateBucket(bucketName)
    })

    UTXO := u.Blockchain.FindUTXO()
    err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket(bucketName)
        for txID, outs := range UTXO {
            key, err := hex.DecodeString(txID)
            err = b.Put(key, outs.Serialize())
        }
    })
}
```
This method creates the UTXO set by scanning all blocks and storing only unspent outputs.

## Updating the UTXO Set
Each time a new block is mined, the UTXO set must be updated:
```go
func (u UTXOSet) Update(block *Block) {
    db := u.Blockchain.db

    err := db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(utxoBucket))

        for _, tx := range block.Transactions {
            if tx.IsCoinbase() == false {
                for _, vin := range tx.Vin {
                    updatedOuts := TXOutputs{}
                    outsBytes := b.Get(vin.Txid)
                    outs := DeserializeOutputs(outsBytes)

                    for outIdx, out := range outs.Outputs {
                        if outIdx != vin.Vout {
                            updatedOuts.Outputs = append(updatedOuts.Outputs, out)
                        }
                    }

                    if len(updatedOuts.Outputs) == 0 {
                        err := b.Delete(vin.Txid)
                    } else {
                        err := b.Put(vin.Txid, updatedOuts.Serialize())
                    }
                }
            }

            newOutputs := TXOutputs{}
            for _, out := range tx.Vout {
                newOutputs.Outputs = append(newOutputs.Outputs, out)
            }

            err := b.Put(tx.ID, newOutputs.Serialize())
        }
    })
}
```
This method ensures that:
- Spent outputs are removed.
- New outputs from recent transactions are added.

## Running the Commands
### Create a Blockchain:
```sh
$ blockchain_go createblockchain -address 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1
```

### Send Coins:
```sh
$ blockchain_go send -from 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1 -to 12DkLzLQ4B3gnQt62EPRJGZ38n3zF4Hzt5 -amount 6
```

### Get Balance:
```sh
$ blockchain_go getbalance -address 1JnMDSqVoHi4TEFXNw5wJ8skPsPf4LHkQ1
```

## Conclusion
In this part, we have:
- Implemented **mining rewards** via coinbase transactions.
- Introduced the **UTXO Set** for better transaction validation.
- Improved transaction efficiency by reducing blockchain scans.
