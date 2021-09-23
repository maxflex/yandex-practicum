/**
ID посылки: 53389963

-- ПРИНЦИП РАБОТЫ --
Для реализации двухсторонней очереди используется массив
на кольцевом буфере. Есть 2 указателя front и back, которые
соответствуют индексам начала и конца массива.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Сложность всех операций в Деке – O(1)
Значит, вычислительная сложность алгоритма линейная:
O(n), где n – кол-во команд на входе

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Придется хранить лишь элементы Дека, значит,
алгоритм расходует O(m) памяти, где m – размер Дека

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var size uint16
	fmt.Scan(&size) // пропускаем n, будем читать до конца ввода
	fmt.Scan(&size) // m – размер дека
	d := NewDeque(size)
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "push_front":
			if ok := d.PushFront(line[1]); !ok {
				fmt.Fprintln(writer, "error")
			}
		case "push_back":
			if ok := d.PushBack(line[1]); !ok {
				fmt.Fprintln(writer, "error")
			}
		case "pop_front":
			if popped, ok := d.PopFront(); ok {
				fmt.Fprintln(writer, popped)
			} else {
				fmt.Fprintln(writer, "error")
			}
		case "pop_back":
			if popped, ok := d.PopBack(); ok {
				fmt.Fprintln(writer, popped)
			} else {
				fmt.Fprintln(writer, "error")
			}
		}
	}
	writer.Flush()
}
