package object

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFont_FontFace(t *testing.T) {
	type args struct {
		size float64
	}
	tests := []struct {
		name string
		f    Font
		args args
	}{
		{
			name: "NotoSansMonoフォントが埋め込まれているかテスト",
			f:    NotoSansMono,
			args: args{
				size: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.f.FontFace(tt.args.size)
			assert.NoError(t, err)
		})
	}
}
