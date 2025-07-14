package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func removeDiacritics(str string) string {
	// Replace Vietnamese-specific characters đ/Đ first, as they are not in the Mn (Mark, Nonspacing) Unicode group.
	// Note: This is language-specific. If supporting other languages with unique characters, extend or refactor this logic.
	str = strings.ReplaceAll(str, "đ", "d")
	str = strings.ReplaceAll(str, "Đ", "D")

	// Remove diacritics by normalizing Unicode and filtering out nonspacing marks.
	t := norm.NFD.String(str)
	result := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func Slugify(title string) string {
	title = strings.ToLower(title)
	title = removeDiacritics(title)

	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	title = re.ReplaceAllString(title, "")

	title = strings.ReplaceAll(title, " ", "-")

	re = regexp.MustCompile("-+")
	title = re.ReplaceAllString(title, "-")

	title = strings.Trim(title, "-")
	return title
}
