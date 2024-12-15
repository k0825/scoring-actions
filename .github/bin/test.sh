#!/bin/bash

# tests.csv ファイルを読み込む
while IFS=',' read -r input output; do
    # 1列目の標準入力を実行し、2列目の標準出力と比較する
    if [ "$(npx ts-node ./problems/1.two-sum/main.ts "${input}")" = "${output}" ]; then
        echo "Success: ${input} -> ${output}"
    else
        echo "Failure: ${input} -> ${output}"
    fi
done < ./problems/1.two-sum/test.csv