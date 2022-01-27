package pkg

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/go-playground/validator/v10"
	"jim/internal/repository"
	"regexp"
)

func MobileValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, value)
	if result {
		return true
	} else {
		return false
	}

}

func ImageValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	imageFormatSlice := []interface{}{repository.ImageFormatOfJpg, repository.ImageFormatOfJpeg, repository.ImageFormatOfGif, repository.ImageFormatOfPng, repository.ImageFormatOfBmp}
	imageFormatSet := mapset.NewSetFromSlice(imageFormatSlice)
	return imageFormatSet.Contains(value)

}
