package storage

import (
	"errors"

	"github.com/FideTech/yaus/models"
)

var (
	ErrNotFound         = errors.New("item not found")
	ErrBucketNotFound   = errors.New("bucket not found")
	ErrKeyAlreadyExists = errors.New("key already exists")
)

type Store interface {
	ShortLinks() ShortLinkStore
}

type ShortLinkStore interface {
	GetAll() ([]models.ShortLink, error)
	GetByKey(key string) (models.ShortLink, error)

	Create(shortLink *models.ShortLink) error

	EnsureBucketExists() error
}
