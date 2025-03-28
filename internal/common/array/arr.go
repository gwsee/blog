package array

import (
	"fmt"
	"strings"
)

type IArrayType interface {
	~int | ~uint | ~int32 | ~int64 | ~uint64 | ~uint32 | ~string | ~float32 | ~float64 | ~int8 | ~uint8 | ~int16 | ~uint16
}

func In[T IArrayType](item T, arr []T) bool {
	for _, one := range arr {
		if one == item {
			return true
		}
	}
	return false
}

func Diff[T IArrayType](a []T, b []T) []T {
	var diffArray []T
	temp := map[T]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}
	return diffArray
}

func Intersect[T IArrayType](a []T, b []T) []T {
	var inter []T
	mp := make(map[T]bool)
	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}
	return inter
}

func Join[T IArrayType](arr []T) string {
	if len(arr) == 0 {
		return ""
	}
	var strBuilder strings.Builder
	for _, v := range arr {
		strBuilder.WriteString(fmt.Sprintf("%v,", v))
	}
	return strings.TrimRight(strBuilder.String(), ",")
}

func Sum[T IArrayType](arr []T) T {
	var num T
	for _, v := range arr {
		num += v
	}
	return num
}

func Unique[T IArrayType](arr []T) []T {
	var newArr []T
	temp := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := temp[v]; ok {
			continue
		}
		temp[v] = struct{}{}
		newArr = append(newArr, v)
	}
	return newArr
}
