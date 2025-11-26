#!/bin/bash

# テスト画像からLGTM画像を生成するスクリプト

set -e

# スクリプトのディレクトリに移動
cd "$(dirname "$0")"

# testdata/imagesの全画像に対してLGTM画像を生成
echo "Generating LGTM images with concentration lines..."
for img in testdata/images/*.jpg; do
    if [ -f "$img" ]; then
        output_file="testdata/results/result_$(basename "$img")"
        echo "Processing: $img -> $output_file"
        go run ./cmd/lgtm/main.go -i "$img" -o "$output_file" --conc
    fi
done

echo "Done! Generated $(ls testdata/results/result_*.jpg | wc -l) images."
