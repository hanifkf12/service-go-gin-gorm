package libraries

import "strings"

func SplitToken(header string) string {
	tokens := strings.Split(header, " ")
	token := tokens[1]
	return token
}
