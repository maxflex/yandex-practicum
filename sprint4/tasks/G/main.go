package main

/**
https://stackoverflow.com/questions/26650293
https://ideone.com/d8cZTH

По сути я сам почти додумался до решения (см. main-custom.go)
Делаем +1/-1 на каждом шаге, и если выходим в 0 – занчит на этом шаге счёт сравнялся
Но смотрим не только в 0, а в любые одинаковые цифры. Например, от -1 до следующего -1 – это место,
где сравнялся счёт. Или от -2 до следующего крайнего -2 (с правой стороны).

Сложно было понять вот это:
first := map[int]int{0: 0}
Это означает лишь, что на нулевом шаге у нас счёт равный
Если взять входные данные [0 1], то без этого уточнения не сработает (додумаешься почему)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	first := map[int]int{0: 0}
	var counter, max int
	for i, v := range arr {
		i++
		if v == "0" {
			counter--
		} else {
			counter++
		}
		if j, ok := first[counter]; ok {
			if i-j > max {
				max = i - j
			}
		} else {
			first[counter] = i
		}
	}
	fmt.Print(max)
}
