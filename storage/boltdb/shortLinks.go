package boltdb

import (
	"encoding/json"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/FideTech/yaus/models"
	"github.com/FideTech/yaus/storage"
)

//ShortLinks is the store for short links
type ShortLinksStore struct {
	store  *boltStore
	bucket []byte
}

var shortLinksStorageInstance *ShortLinksStore

func (b *boltStore) ShortLinks() storage.ShortLinkStore {
	if shortLinksStorageInstance == nil {
		shortLinksStorageInstance = &ShortLinksStore{
			store:  b,
			bucket: []byte("shortLinks"),
		}
	}

	return shortLinksStorageInstance
}

func (sl *ShortLinksStore) EnsureBucketExists() error {
	return sl.store.DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(sl.bucket); err != nil {
			return err
		}

		return nil
	})
}

func (sl *ShortLinksStore) GetAll() ([]models.ShortLink, error) {
	results := []models.ShortLink{}

	viewErr := sl.store.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(sl.bucket)

		if b == nil {
			return storage.ErrBucketNotFound
		}

		return b.ForEach(func(key []byte, val []byte) error {
			var item models.ShortLink
			if err := json.Unmarshal(val, &item); err != nil {
				return err
			}

			results = append(results, item)

			return nil
		})
	})

	return results, viewErr
}

func (sl *ShortLinksStore) GetByKey(key string) (models.ShortLink, error) {
	result := models.ShortLink{}

	viewErr := sl.store.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(sl.bucket)

		if b == nil {
			return storage.ErrBucketNotFound
		}

		val := b.Get([]byte(key))
		if val == nil || len(val) == 0 {
			return storage.ErrNotFound
		}

		return json.Unmarshal(val, &result)
	})

	return result, viewErr
}

func (sl *ShortLinksStore) Create(shortLink *models.ShortLink) error {
	return sl.store.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(sl.bucket)

		if b == nil {
			return storage.ErrBucketNotFound
		}

		key := []byte(shortLink.Key)

		if result := b.Get(key); result != nil || len(result) > 0 {
			return storage.ErrKeyAlreadyExists
		}

		shortLink.CreatedAt = time.Now()

		val, err := json.Marshal(shortLink)
		if err != nil {
			return err
		}

		return b.Put(key, val)
	})
}

func (sl *ShortLinksStore) Update(shortLink *models.ShortLink) error {
	return sl.store.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(sl.bucket)

		if b == nil {
			return storage.ErrBucketNotFound
		}

		key := []byte(shortLink.Key)

		if result := b.Get(key); result == nil || len(result) == 0 {
			return storage.ErrNotFound
		}

		shortLink.UpdatedAt = time.Now()

		val, err := json.Marshal(shortLink)
		if err != nil {
			return err
		}

		return b.Put(key, val)
	})
}
