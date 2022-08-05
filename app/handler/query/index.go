package query

import "net/url"

// 数字のクエリを取得し、文字列形式で返す
func GetQueryNumInStr(v url.Values, s string, d string) string {
	n := v.Get(s)
	if n == "" {
		// クエリに何も指定していない場合はデフォルト値を代入
		n = d
	}

	return n
}
