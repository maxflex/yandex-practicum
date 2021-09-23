/**
ID посылки: 53382508

-- ПРИНЦИП РАБОТЫ --
Для решения используется структура данных стек.
Сканируем по словам.
– Если наткнулись на операнд, кладём его на вершину стека
– Если наткнулись на знак операции – достаём из стека два последних операнда
  и применяем арифметику соответствующую знаку операции

В итоге финальный результат будет находиться на вершине стека

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Амортизированная сложность добавления/удаления из стека – O(1)
Сложность алгоритма линейная: O(n), где n – кол-во операций и чисел на входе

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В стеке хранятся только операнды, значит, алгоритм расходует O(m) памяти,
где m – кол-во операндов на входе.

*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var s Stack
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		switch word {
		case "+", "-", "*", "/":
			b, _ := s.Pop()
			a, _ := s.Pop()
			switch word {
			case "+":
				s.Push(a + b)
			case "-":
				s.Push(a - b)
			case "*":
				s.Push(a * b)
			case "/":
				s.Push(int(math.Floor(float64(a) / float64(b))))
			}
		default:
			operand, _ := strconv.Atoi(word)
			s.Push(operand)
		}
	}
	result, _ := s.Pop()
	fmt.Print(result)
}
