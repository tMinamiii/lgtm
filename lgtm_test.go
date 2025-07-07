package lgtm

import (
	"testing"
)

func TestTextDrawer_WhiteText(t *testing.T) {
	type args struct {
		mainText   string
		subText    string
		textColor  TextColor
		inputPath  string
		outputPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "images/test_extreme_tall.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_extreme_tall.jpg",
				outputPath: "testdata/results/result_test_extreme_tall.jpg",
			},
		},
		{
			name: "images/test_extreme_wide.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_extreme_wide.jpg",
				outputPath: "testdata/results/result_test_extreme_wide.jpg",
			},
		},
		{
			name: "images/test_huge_square.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_huge_square.jpg",
				outputPath: "testdata/results/result_test_huge_square.jpg",
			},
		},
		{
			name: "images/test_large_landscape.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_large_landscape.jpg",
				outputPath: "testdata/results/result_test_large_landscape.jpg",
			},
		},
		{
			name: "images/test_large_portrait.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_large_portrait.jpg",
				outputPath: "testdata/results/result_test_large_portrait.jpg",
			},
		},
		{
			name: "images/test_micro.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_micro.jpg",
				outputPath: "testdata/results/result_test_micro.jpg",
			},
		},
		{
			name: "images/test_rect_300x200.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_rect_300x200.jpg",
				outputPath: "testdata/results/result_test_rect_300x200.jpg",
			},
		},
		{
			name: "images/test_rect_400x300.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_rect_400x300.jpg",
				outputPath: "testdata/results/result_test_rect_400x300.jpg",
			},
		},
		{
			name: "images/test_rect_500x300.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_rect_500x300.jpg",
				outputPath: "testdata/results/result_test_rect_500x300.jpg",
			},
		},
		{
			name: "images/test_rect_600x400.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_rect_600x400.jpg",
				outputPath: "testdata/results/result_test_rect_600x400.jpg",
			},
		},
		{
			name: "images/test_ribbon_h.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_ribbon_h.jpg",
				outputPath: "testdata/results/result_test_ribbon_h.jpg",
			},
		},
		{
			name: "images/test_ribbon_v.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_ribbon_v.jpg",
				outputPath: "testdata/results/result_test_ribbon_v.jpg",
			},
		},
		{
			name: "images/test_small.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_small.jpg",
				outputPath: "testdata/results/result_test_small.jpg",
			},
		},
		{
			name: "images/test_square_200.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_square_200.jpg",
				outputPath: "testdata/results/result_test_square_200.jpg",
			},
		},
		{
			name: "images/test_square_300.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_square_300.jpg",
				outputPath: "testdata/results/result_test_square_300.jpg",
			},
		},
		{
			name: "images/test_square_500.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_square_500.jpg",
				outputPath: "testdata/results/result_test_square_500.jpg",
			},
		},
		{
			name: "images/test_square_800.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_square_800.jpg",
				outputPath: "testdata/results/result_test_square_800.jpg",
			},
		},
		{
			name: "images/test_tall_2to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_tall_2to1.jpg",
				outputPath: "testdata/results/result_test_tall_2to1.jpg",
			},
		},
		{
			name: "images/test_tall_3to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_tall_3to1.jpg",
				outputPath: "testdata/results/result_test_tall_3to1.jpg",
			},
		},
		{
			name: "images/test_tall_4to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_tall_4to1.jpg",
				outputPath: "testdata/results/result_test_tall_4to1.jpg",
			},
		},
		{
			name: "images/test_tall_5to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_tall_5to1.jpg",
				outputPath: "testdata/results/result_test_tall_5to1.jpg",
			},
		},
		{
			name: "images/test_thin_h.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_thin_h.jpg",
				outputPath: "testdata/results/result_test_thin_h.jpg",
			},
		},
		{
			name: "images/test_thin_v.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_thin_v.jpg",
				outputPath: "testdata/results/result_test_thin_v.jpg",
			},
		},
		{
			name: "images/test_tiny.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_tiny.jpg",
				outputPath: "testdata/results/result_test_tiny.jpg",
			},
		},
		{
			name: "images/test_ultra_tall.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_ultra_tall.jpg",
				outputPath: "testdata/results/result_test_ultra_tall.jpg",
			},
		},
		{
			name: "images/test_ultra_wide.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_ultra_wide.jpg",
				outputPath: "testdata/results/result_test_ultra_wide.jpg",
			},
		},
		{
			name: "images/test_wide_2to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_wide_2to1.jpg",
				outputPath: "testdata/results/result_test_wide_2to1.jpg",
			},
		},
		{
			name: "images/test_wide_3to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_wide_3to1.jpg",
				outputPath: "testdata/results/result_test_wide_3to1.jpg",
			},
		},
		{
			name: "images/test_wide_4to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_wide_4to1.jpg",
				outputPath: "testdata/results/result_test_wide_4to1.jpg",
			},
		},
		{
			name: "images/test_wide_5to1.jpg",
			args: args{
				mainText:   "LGTM",
				subText:    "Looks Good To Me",
				textColor:  TextColorWhite,
				inputPath:  "testdata/images/test_wide_5to1.jpg",
				outputPath: "testdata/results/result_test_wide_5to1.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main := NewMainText(tt.args.mainText, tt.args.textColor)
			sub := NewSubText(tt.args.subText, tt.args.textColor)

			d := NewTextDrawer(main, sub, tt.args.inputPath, tt.args.outputPath)
			if err := d.Draw(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
