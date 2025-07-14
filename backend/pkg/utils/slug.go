package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func removeDiacritics(str string) string {
	// Thay thế ký tự đ/Đ trước vì nó không thuộc nhóm Mn
	str = strings.ReplaceAll(str, "đ", "d")
	str = strings.ReplaceAll(str, "Đ", "D")

	// Loại dấu bằng cách chuẩn hóa Unicode
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
