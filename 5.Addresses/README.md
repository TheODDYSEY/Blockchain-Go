# Building Blockchain in Go. Part 5: Addresses

Welcome to the Blockchain Project! This project implements a simplified version of Bitcoin's transaction mechanism, including transaction signing and verification using ECDSA.

## Features

- **Wallet Creation**: Generate new wallet addresses.
- **Transaction Handling**: Create, sign, and verify transactions.
- **Blockchain Operations**: Add transactions to blocks and mine new blocks.
- **Balance Retrieval**: Check wallet balances.
- **Security**: Transactions are signed to ensure authenticity and prevent unauthorized spending.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Creating a Wallet](#creating-a-wallet)
  - [Creating a Blockchain](#creating-a-blockchain)
  - [Sending Transactions](#sending-transactions)
  - [Checking Balances](#checking-balances)
- [Implementation Details](#implementation-details)
  - [Transaction Signing](#transaction-signing)
  - [Transaction Verification](#transaction-verification)
- [Testing](#testing)
- [Conclusion](#conclusion)
- [References](#references)

## Usage

### Creating a Wallet
Generate a new wallet address:

```
blockchain_go createwallet
```

### Creating a Blockchain
Create a new blockchain and mine the genesis block:

```
blockchain_go createblockchain -address <your-wallet-address>
```

### Sending Transactions
Send coins from one address to another:

```
blockchain_go send -from <sender-address> -to <recipient-address> -amount <amount>
```

### Checking Balances
Check the balance of a specific wallet:

```
blockchain_go getbalance -address <wallet-address>
```

## Implementation Details

### Transaction Signing
Transactions are signed using ECDSA to ensure that only the owner of the private key can authorize spending. The `Sign` method creates a trimmed copy of the transaction, hashes it, and signs the hash.

```go
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
  if tx.IsCoinbase() {
    return
  }

  txCopy := tx.TrimmedCopy()

  for inID, vin := range txCopy.Vin {
    prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
    txCopy.Vin[inID].Signature = nil
    txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
    txCopy.ID = txCopy.Hash()
    txCopy.Vin[inID].PubKey = nil

    r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
    signature := append(r.Bytes(), s.Bytes()...)

    tx.Vin[inID].Signature = signature
  }
}
```

### Transaction Verification
The `Verify` method checks the signature of each transaction input against the public key to ensure it was authorized by the owner.

```go
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
  txCopy := tx.TrimmedCopy()
  curve := elliptic.P256()

  for inID, vin := range tx.Vin {
    prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
    txCopy.Vin[inID].Signature = nil
    txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
    txCopy.ID = txCopy.Hash()
    txCopy.Vin[inID].PubKey = nil

    r := big.Int{}
    s := big.Int{}
    sigLen := len(vin.Signature)
    r.SetBytes(vin.Signature[:(sigLen / 2)])
    s.SetBytes(vin.Signature[(sigLen / 2):])

    x := big.Int{}
    y := big.Int{}
    keyLen := len(vin.PubKey)
    x.SetBytes(vin.PubKey[:(keyLen / 2)])
    y.SetBytes(vin.PubKey[(keyLen / 2):])

    rawPubKey := ecdsa.PublicKey{curve, &x, &y}
    if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
      return false
    }
  }

  return true
}
```

## Testing
To ensure everything works as expected:
- Create wallets and a blockchain.
- Send transactions and verify balances.
- Comment out the signing step to verify that unsigned transactions are rejected.

## Conclusion
This project demonstrates key concepts in blockchain, including transaction signing and verification, using a simplified Bitcoin-like protocol.

## References
- [Public-key cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography)
- [Digital signatures](https://en.wikipedia.org/wiki/Digital_signature)
- [Elliptic curve cryptography](https://en.wikipedia.org/wiki/Elliptic-curve_cryptography)
- [ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
- [Bitcoin addresses](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses)
- [Base58](https://en.bitcoin.it/wiki/Base58Check_encoding)
- [Elliptic curve introduction](https://en.wikipedia.org/wiki/Elliptic_curve)
