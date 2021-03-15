package core

import (
	"log"

	"github.com/FideTech/yaus/storage"
	"github.com/FideTech/yaus/storage/boltdb"
)

var (
	_dataStorage storage.Store
)

func Start() error {
	log.Println("Starting the core")

	store, err := boltdb.New()
	if err != nil {
		return err
	}

	_dataStorage = store

	return nil
}

func Shutdown() error {
	log.Println("Shutting the store down")

	if err := boltdb.Close(); err != nil {
		return err
	}

	_dataStorage = nil

	return nil
}
