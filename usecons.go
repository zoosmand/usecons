package usecons

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Prints header of task.
// @header: (string) line to be printed as header
// @return: none
func Header(header *string) string {
	len := utf8.RuneCountInString(*header)
	var dashes []rune

	for i := 0; i < len; i++ {
		dashes = append(dashes, '*')
	}

	return fmt.Sprintf("\n%s\n%s\n", *header, string(dashes))
}

// Checks input like "12,14,-56,1044"
// @numbersLine: (string) line to be examined
// @return_0: ([]int16) array with numbers
// @return_1: (bool) status of operation
func CheckUserInputNumbersViaComma(numbersLine string) ([]int16, bool) {
	ok := true
	var numbers []int16
	_numbers := strings.Split(numbersLine, ",")

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
