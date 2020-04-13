package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	isSilence := len(remark) == 0
	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := hasLetters && strings.ToUpper(remark) == remark

	if isSilence {
		return "Fine. Be that way!"
	}

	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if isYelling {
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
