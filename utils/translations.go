package utils

var translations []Translation

func GetTranslations() *[]Translation {
	return &translations
}

func GetTranslation(name string) Translation {
	var ret Translation
	for _, x := range translations {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}
func GetTranslationsByLanguage(language string) []Translation {

	var ret []Translation
	for _, x := range translations {
		if x.Language == language {
			ret = append(ret, x)
		}
	}
	return ret
}

func SetBaseValueTranslations() {

	//EVENTS

}
