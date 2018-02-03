package blockchain

import (
	"fmt"
	"log"
)

func (cli *CLI) CreateBlockchain(address string) error {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc, err := CreateBlockchain(address)
	if err != nil {
		return err
	}
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
	return nil
}
