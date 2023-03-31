package usecons

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Prints a looking good header
// @header: (string) Header text to print
// @return: none
func Header(header *string) string {
	len := utf8.RuneCountInString(*header)
	var dashes []rune

	for i := 0; i < len; i++ {
		dashes = append(dashes, '-')
	}

	return fmt.Sprintf("\n%s\n%s\n", *header, string(dashes))
}

// Converts int16 array from comma-separated integers string to an array
// in this manner: "12,14,-56,1044" -> [12, 14, -56, 1044]
// @stringArray: (string) line to be examined
// @return_0: ([]int16) Array with numbers
// @return_1: (bool) Status of result
func CommaSeparatedIntToArray(stringArray string) ([]int16, bool) {
	ok := true
	var numbers []int16
	_numbers := strings.Split(stringArray, ",")

	if len(_numbers) < 2 {
		return numbers, false
	}

	for i := 0; i < len(_numbers); i++ {
		// _number, err := strconv.Atoi(_numbers[i])
		_number, err := strconv.ParseInt(_numbers[i], 10, 16)

		if err != nil {
			ok = false
		} else {
			numbers = append(numbers, int16(_number))
		}
	}
	return numbers, ok
}

// Generates an array of integers with the requested length (size)
// @length: (int) Length (size) of the generating array
// @return: ([]int) Requested size array
func GenerateIntArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(999)
	}
	return arr
}

// Generates a square matrix of integers with the requested level
// @level: (int) Length (size) of the matrix
// @return: ([]int) Requested size array
func GenerateIntSquareMatrix(level int) [][]int {
	m := [][]int{}
	for i := 0; i < level; i++ {
		m = append(m, GenerateIntArray(level))
	}
	return m
}

// Generates a matrix of integers with the requested levels
// @h: (int) Height (number of rows) of the matrix
// @w: (int) Width (number of columns) of the matrix
// @return: ([]int) Requested size array
func GenerateIntMatrix(h int, w int) [][]int {
	m := [][]int{}
	for i := 0; i < h; i++ {
		m = append(m, GenerateIntArray(w))
	}
	return m
}
