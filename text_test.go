package lgtm

import (
	"testing"
)

func TestPaddingText_String(t *testing.T) {
	tests := []struct {
		name string
		p    PaddingText
		want string
	}{
		{
			name: "正常 文字の間に半角スペースをpaddingする",
			p:    PaddingText("LGTM"),
			want: "L G T M",
		},
		{
			name: "正常 文字の間に半角スペースをpaddingする(元の文字に半角スペースあり)",
			p:    PaddingText("Looks Good To Me"),
			want: "L o o k s   G o o d   T o   M e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("PaddingText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaddingText_HasJP(t *testing.T) {
	tests := []struct {
		name string
		p    PaddingText
		want bool
	}{
		{
			name: "英数字のみ",
			p:    PaddingText("LGTM"),
			want: false,
		},
		{
			name: "日本語交じり",
			p:    PaddingText("ABCあいう"),
			want: true,
		},
		{
			name: "日本語のみ",
			p:    PaddingText("あいう"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HasJP(); got != tt.want {
				t.Errorf("PaddingText.HasJP() = %v, want %v", got, tt.want)
			}
		})
	}
}
