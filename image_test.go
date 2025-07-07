package lgtm

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedImage_Image(t *testing.T) {
	// Gopher画像が埋め込まれているかテスト
	_, err := GopherPng.Image()
	assert.NoError(t, err)
}
