package Errors

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//  1. вставьте определение типа для []error
//  2. определите метод Error для вашего типа, который будет выводить
//     все ошибки слайса
//  3. реализуйте функцию MyCheck
//
// ...
type StrErrors []error

func (errs StrErrors) Error() string {
	var out string
	for _, err := range errs {
		out += err.Error() + ";"
	}
	return strings.TrimRight(out, `;`)
}

func MyCheck(str string) error {
	var errs StrErrors
	ws_count := 0
	has_number := false

	if len(str) >= 20 {
		errs = append(errs, errors.New("line is too long"))
	}

	for _, el := range str {
		if el == ' ' {
			ws_count++
		} else if el >= '0' && el <= '9' && !has_number {
			has_number = true
		}
	}

	if has_number {
		errs = append(errs, errors.New("found numbers"))
	}
	if ws_count < 2 {
		errs = append(errs, errors.New("no two spaces"))
	}

	return errs
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
