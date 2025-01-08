package main

import (
    "fmt"
    "log"

    "github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

func main() {
    db, err := bolt.Open(dbFile, 0600, nil)
    if err != nil {
        log.Panic(err)
    }
    defer db.Close()

    err = db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        if b == nil {
            return fmt.Errorf("Bucket %q not found!", blocksBucket)
        }

        err := b.ForEach(func(k, v []byte) error {
            fmt.Printf("Key: %x\n", k)
            fmt.Printf("Value: %x\n", v)
            fmt.Println()
            return nil
        })
        return err
    })

    if err != nil {
        log.Panic(err)
    }
}