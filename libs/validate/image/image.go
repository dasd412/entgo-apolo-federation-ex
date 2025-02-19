package image

import (
	"errors"
	"regexp"
)

var ErrInvalidImageUrlFormat = errors.New("invalid image url format")

const UrlPattern = `^https?://(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(?:/[^/]+)*/.+\.(jpg|jpeg|png|gif|bmp|webp)(?:\?.*)?$`

func ValidateImageURL(imageURL string) error {
	urlRegex := regexp.MustCompile(UrlPattern)
	match := urlRegex.MatchString(imageURL)

	if match {
		return nil
	} else {
		return ErrInvalidImageUrlFormat
	}
}
