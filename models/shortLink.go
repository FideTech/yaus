package models

//ShortLink holds the data for the shortened urls
type ShortLink struct {
	Key string `yaml:"key" json:"key"`
	URL string `yaml:"url" json:"url"`
}
