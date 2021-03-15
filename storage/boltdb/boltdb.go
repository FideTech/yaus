package boltdb

import (
	"time"

	"github.com/FideTech/yaus/config"
	"github.com/FideTech/yaus/storage"
	bolt "go.etcd.io/bbolt"
)

type boltStore struct {
	*bolt.DB
}

var boltStoreInstance *boltStore

func New() (storage.Store, error) {
	db, err := bolt.Open(config.Config.System.Database.Path, 0600, &bolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}

	boltStoreInstance = &boltStore{db}

	if err := ensureBucketsExist(); err != nil {
		return nil, err
	}

	return boltStoreInstance, nil
}

func Close() error {
	if boltStoreInstance == nil {
		return nil
	}

	return boltStoreInstance.Close()
}

func ensureBucketsExist() error {
	return boltStoreInstance.ShortLinks().EnsureBucketExists()
}
