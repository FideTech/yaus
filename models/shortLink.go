package models

import (
	"time"
)

//ShortLink holds the data for the shortened urls
type ShortLink struct {
	Key string `yaml:"key" json:"key"`
	URL string `yaml:"url" json:"url"`

	Redirects int `yaml:"-" json:"redirects"`

	CreatedAt time.Time `yaml:"-" json:"createdAt"`
	UpdatedAt time.Time `yaml:"-" json:"updatedAt"`
}
