package core

import (
	"github.com/FideTech/yaus/models"
)

func GetAllDynamicShortLinks() ([]models.ShortLink, error) {
	return _dataStorage.ShortLinks().GetAll()
}

func GetDynamicShortLinkByKey(key string) (models.ShortLink, error) {
	return _dataStorage.ShortLinks().GetByKey(key)
}

func CreateDynamicShortLink(shortLink *models.ShortLink) error {
	return _dataStorage.ShortLinks().Create(shortLink)
}
