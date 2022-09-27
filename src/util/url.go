package util

import (
	"net/url"
	"regexp"
	"strings"
)

var matchPlaceHolders = [...]string{"$0", "$1", "$2", "$3", "$4", "$5", "$6", "$7", "$8", "$9"}

func ReplaceUrl(reMatch *regexp.Regexp, find, replace string) (*url.URL, error) {
	matches := reMatch.FindStringSubmatch(find)
	if len(matches) > len(matchPlaceHolders) {
		matches = matches[:len(matchPlaceHolders)]
	}

	replacerParam := make([]string, 0, len(matches)*2)
	for i := range matches {
		replacerParam = append(replacerParam, matchPlaceHolders[i], matches[i])
	}
	replacer := strings.NewReplacer(replacerParam...)
	target := replacer.Replace(replace)

	if len(target) == 0 {
		target = "/"
	}

	return url.Parse(target)
}
