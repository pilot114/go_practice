package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// пример кастомного типа ошибок
type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}

// В Go, как правило, не использует разные типы ошибок
// Встроенные (типа io.EOF и io.ErrNoProgress) являются просто ссылками,
// позвозяет делать таке простые проверки: if err == io.EOF {...}
func main() {
	// функция восстановления при аварии
	// её можно для нажености использовать в библиотечных функциях,
	// превращая аварии в ошибки
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()

	args := os.Args[1:]
	// область видимосте err - условный блок
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}

	// аварии прекращают работу программы, если нет обработчика
	// в panic передается что угодно, но лучше всего - ошибки
	panic(errors.New("Чета случилось"))

	// авария в сопрограмме приведет к остановке всей программы,
	// так как сопрограмма имеет свой стек (в котором ищется recover),
	// что игнорирует обработку аварии в программе.
	// ПОЭТОМУ panic также следует обрабатывать внутри сопрограммы.
	// пример враппера-обработчика аварий: github.com/Masterminds/cookoo/safely
}

// Concat : объединяет строки, разделяя их пробелами.
// Она возвращает пустую строку и ошибку, если не получила ни одной строки.
func Concat(parts ...string) (string, error) {
	// хороший подход: возвращать валидный результат даже в случае ошибки
	if len(parts) == 0 {
		// fmt.Errorf может заменять errors.New, если требуется форматирование
		return "", errors.New("No strings supplied")
	}
	return strings.Join(parts, " "), nil
}
