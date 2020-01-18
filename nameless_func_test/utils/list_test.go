package utils

import (
	"errors"
	"testing"
)

func TestCountAggregeter(t *testing.T) {
	// 何度も同じ処理を行うため無名関数化
	debuger := func(lst []string, expectMap map[string]int) error {
		// 集計処理を実行
		res := CountAggregeter(lst)

		// 期待する値を用いてloop処理
		for key, val := range expectMap {
			// 値が集計結果に含まれているかどうか
			if cnt, ok := res[key]; ok {
				// カウント数が一致しないのであればerror
				if val != cnt {
					return errors.New("[Error] カウント数が一致しません")
				}
			} else {
				// そもそも結果に含まれていなければerror
				return errors.New("[Error] 想定値か集計値が間違っています")
			}
		}
		return nil
	}

	// シンプルなケース
	if err := debuger([]string{"A", "B", "C", "D"}, map[string]int{"A": 1, "B": 1, "C": 1, "D": 1}); err != nil {
		t.Error(err)
	}

	// 重複した値がうまくカウントされているか
	if err := debuger([]string{"A", "B", "A", "C"}, map[string]int{"A": 2, "B": 1, "C": 1}); err != nil {
		t.Error(err)
	}

	// リストが空の場合
	if err := debuger([]string{}, map[string]int{}); err != nil {
		t.Error(err)
	}

	// 複数の値が重複する場合
	if err := debuger(
		[]string{"A", "B", "A", "C", "D", "C", "E", "F", "D", "F", "G"},
		map[string]int{"A": 2, "B": 1, "C": 2, "D": 2, "E": 1, "F": 2, "G": 1},
	); err != nil {
		t.Error(err)
	}

	// 重複の発生が3回以上
	if err := debuger([]string{"A", "A", "A"}, map[string]int{"A": 3}); err != nil {
		t.Error(err)
	}
}
