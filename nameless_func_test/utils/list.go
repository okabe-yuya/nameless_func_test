package utils

// リストを受け取り、値の登場回数をカウント
func CountAggregeter(lst []string) map[string]int {
	// 集計用のmap
	aggregater := make(map[string]int, 0)
	for _, val := range lst {
		// 既出であれば+1
		if _, ok := aggregater[val]; ok {
			aggregater[val] += 1
		} else {
			// 初登場であれば1を格納
			aggregater[val] = 1
		}
	}
	return aggregater
}
