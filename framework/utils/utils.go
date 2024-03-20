package utils

import "strings"

func GetDomain(url string) string {
	res := strings.Split(url, "//")
	return res[1]
}

func GenerateFileName(url string) string {
	arr := strings.Split(url, "//")
	arr = arr[1:]
	res := strings.ReplaceAll(strings.Join(arr, "."), "/", ".")
	if res[len(res)-1] == '.' {
		res = res[:len(res)-1]
	}
	return res
}
