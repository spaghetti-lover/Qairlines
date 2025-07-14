package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Bài viết số 1", "bai-viet-so-1"},
		{"  Clean --- title  ", "clean-title"},
		{"Hello World", "hello-world"},
		{"Go is great!!!", "go-is-great"},
		{"Tiêu đề với ký tự đặc biệt #@!", "tieu-de-voi-ky-tu-dac-biet"},
		{"Chữ có     nhiều     khoảng trắng", "chu-co-nhieu-khoang-trang"},
		{"", ""},
		{"---", ""},
	}

	for _, test := range tests {
		result := Slugify(test.input)
		assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
	}
}
