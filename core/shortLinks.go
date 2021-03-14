package core

import (
	"github.com/teris-io/shortid"

	"github.com/FideTech/yaus/models"
)

func GetAllDynamicShortLinks() ([]models.ShortLink, error) {
	return _dataStorage.ShortLinks().GetAll()
}

func GetDynamicShortLinkByKey(key string) (models.ShortLink, error) {
	return _dataStorage.ShortLinks().GetByKey(key)
}

func CreateDynamicShortLink(shortLink *models.ShortLink) error {
	if shortLink.Key == "" {
		key, err := shortid.Generate()
		if err != nil {
			return err
		}

		shortLink.Key = key
	}

	return _dataStorage.ShortLinks().Create(shortLink)
}
