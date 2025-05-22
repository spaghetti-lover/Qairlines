package utils

import (
	"regexp"
	"strings"
)

// Slugify converts a string into a URL-friendly slug.
func Slugify(title string) string {
	// Chuyển đổi thành chữ thường
	title = strings.ToLower(title)

	// Xóa ký tự không phải chữ cái, số hoặc khoảng trắng
	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	title = re.ReplaceAllString(title, "")

	// Thay thế khoảng trắng bằng dấu gạch ngang
	title = strings.ReplaceAll(title, " ", "-")

	// Xóa các dấu gạch ngang liên tiếp
	re = regexp.MustCompile("-+")
	title = re.ReplaceAllString(title, "-")

	// Xóa khoảng trắng ở đầu và cuối chuỗi
	title = strings.TrimSpace(title)

	return title
}
