package common

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"slices"
	"sort"
	"strconv"
)

func GetRandomNumBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetSortedIntSliceFromMapValues(mapOfValues map[int]int) []int {
	size := len(mapOfValues)
	arr := make([]int, 0, size)
	for _, v := range mapOfValues {
		arr = append(arr, v)
	}
	slices.Sort(arr)
	return arr
}

func GetIntSliceFromMapValues(mapOfValues map[[32]byte][]int) [][]int {
	size := len(mapOfValues)
	arr := make([][]int, 0, size)
	for _, v := range mapOfValues {
		arr = append(arr, v)
	}
	return arr
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func ComputeHashKeyForList(list []int,
	delim string) [32]byte {
	var buffer bytes.Buffer
	sort.Ints(list)
	for i, _ := range list {
		buffer.WriteString(
			strconv.Itoa(list[i]))
		buffer.WriteString(delim)
	}
	return (sha256.Sum256(buffer.Bytes()))
}
