package awesomeProject2

import "strings"

type Filter interface {
	filter(language Language) Language
}

type Language struct {
	content string
}
type ObsceneLanguageFilter struct {

}
func (ObsceneLanguageFilter ObsceneLanguageFilter) filter(language Language) Language{
	if strings.Contains(language.content,"fuck") {
		strs:=strings.Split(language.content,"fuck");
		language.content="";
		for _,i :=range strs{
			language.content=language.content+i;
		}
	}
	return language;
}
