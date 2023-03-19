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

// Return int16 array from comma separated integers from string
// something like this "12,14,-56,1044"
// @stringArray: (string) line to be examined
// @return_0: ([]int16) Array with numbers
// @return_1: (bool) Atatus of operation
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

// Generated an array of integers of requested size
// @length: (int) Length (size) of the generating array
// @return: ([]int)
func GenerateIntArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(999)
	}
	return arr
}
