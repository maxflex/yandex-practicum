package main_custom

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Для входных данных
0 0 1 0 1 1 1 0 0 0
на каждом шаге делаем -1, если 0; и +1, если 1. В итоге получаем:
-1 -2 -1 -2 -1 0 1 0 -1 -2
Получается, шаги со значением 0 – это шаги, когда счёт сравнялся
(можно делать это in-place, чтобы не расходовать доп. память, т.к.
a[i] мы уже проанализировали и теперь мы можем его перезаписать)


Далее, алгоритм: чтобы найти максимальную последовательность для шага i
мы как бы "обнуляем" всё что было до шага i (кладём это в переменную add)
и начинаем смотреть с конца массива в поисках 0 (с учетом сдвига add)
Самый правый 0 – и есть самая длинная последовательность для шага i
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	arr := make([]int, 0, len(words))
	for _, word := range words {
		val, _ := strconv.Atoi(word)
		arr = append(arr, val)
	}
	var max, counter int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			counter--
		} else {
			counter++
		}
		if counter == 0 {
			max = i
		}
		arr[i] = counter
	}
	for i := 1; i < len(arr); i++ {
		add := arr[i-1] * -1
		// если i-j стал меньше max, то последовательности длиннее уже не будет
		for j := len(arr) - 1; j > i && j-i > max; j-- {
			if arr[j]+add == 0 {
				if i > max {
					max = i
				}
				break
			}
		}
	}
	fmt.Print(max + 1)
}
