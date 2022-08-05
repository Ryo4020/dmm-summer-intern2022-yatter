package query

import "net/url"

// 数字のクエリを取得し、文字列形式で返す
func GetQueryNumInStr(v url.Values, s string, d string) string {
	n := v.Get(s)
	if n == "" {
		n = d
	}

	return n
}
