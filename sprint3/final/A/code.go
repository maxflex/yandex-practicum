/**
ID успешной посылки: 54356469

Вычислительная сложность: 2*log(n) => O(log(n))
log(n) на поиск индекса правильного начала массива и
log(n) на поиск элемента методом бинарного поиска

Пространственная сложность: O(1)
дополнительная память не требуется ни на каком из этапов
*/
package main

/**
Сначала находим сдвиг (правильный индекс начала массива)
далее делаем модификацию бинарного поиска с учётом сдвига
*/
func brokenSearch(arr []int, k int) int {
	l := 0
	r := len(arr) - 1
	offset := getOffset(arr, l, r)
	mid := (l + r) / 2
	offsetMid := (mid + offset) % len(arr)
	for arr[offsetMid] != k {
		if k > arr[offsetMid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
		if l > r {
			return -1
		}
		mid = (l + r) / 2
		offsetMid = (mid + offset) % len(arr)
	}
	return offsetMid
}

/**
Найти сдвиг (правильный индекс начала массива)

Вычислительная сложность: log(n)
Алгоритм: divide & conquare

Смотрим в центр массива, и если следующий за ним элемент меньше, значит,
следующий элемент и есть правильный индекс начала, т.к. в отсортированном
по возрастанию массиве a[i+1] всегда должен быть больше a[i]

Повторяем операцию для левой половины массива, если первый элемент в нём
больше центрального, т.к. в отсортированном по возрастанию массиве a[0]
всегда должен быть больше a[n/2]. Иначе повторяем для правой половины
*/
func getOffset(arr []int, l, r int) int {
	for {
		mid := (l + r) / 2
		next := (mid + 1) % len(arr)
		if arr[next] < arr[mid] {
			return next
		}
		if arr[l] > arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
		if l > r {
			return -1
		}
	}
}