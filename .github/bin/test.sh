#!/bin/bash

# tests.csv ファイルを読み込む
while IFS=',' read -ra row; do
    # 1列目の標準入力を実行し、2列目の標準出力と比較する
    if [ "$(npx ts-node problems/1.two-sum/main.ts "${row[0]}")" == "${row[1]}" ]; then
        echo "Success: ${row[0]} -> ${row[1]}"
    else
        echo "Failure: ${row[0]} -> ${row[1]}"
    fi
done < problems/1.two-sum/tests.csv