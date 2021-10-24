/**
ID успешной посылки: 55344447

Вычислительная сложность: O(n)
Операции вставки, удаления и поиска по ключу выполняются в среднем за О(1)
Значит, сложность алгоритма O(n), где n - кол-во операций на входе

Пространственная сложность: O(1)
По условию "число хранимых в таблице ключей не превосходит 10^5"
Можно воспользоваться этой информацией и сразу инициировать таблицу размером m=10^5
В ограничение памяти мы более чем укладываемся

Разрешение коллизий: метод цепочек
На практике коллизий не возникнет, т.к. по условию "в таблице хранятся уникальные ключи.",
а значит, при m=10^5 каждый попадёт в свою уникальную корзину

* в задаче выбрано m=100003 - ближайшее простое больше 10^5, хотя это необязательно,
т.к. коллизий не возникнет. Чтобы потестировать разрешение коллизий, надо выбрать меньшее m
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HashMap struct {
	bucket []*ListNode
}

type HashEntry struct {
	key, val int
}

type ListNode struct {
	data HashEntry
	next *ListNode
}

func NewHashMap(size int) HashMap {
	return HashMap{make([]*ListNode, size)}
}

func (hm *HashMap) put(key, val int) {
	index := hm.getIndex(key)
	entry := &ListNode{data: HashEntry{key, val}}
	if hm.bucket[index] == nil {
		hm.bucket[index] = entry
	} else {
		// разрешение коллизии методом цепочек
		node := hm.bucket[index]
		for {
			if node.data.key == key {
				// перезаписываем значение по ключу
				node.data.val = val
				return
			}
			if node.next == nil {
				// записываем новое значение
				node.next = entry
				return
			}
			node = node.next
		}
	}
}

func (hm HashMap) get(key int) int {
	index := hm.getIndex(key)
	node := hm.bucket[index]
	for node != nil {
		if node.data.key == key {
			return node.data.val
		}
		node = node.next
	}
	return -1
}

func (hm *HashMap) delete(key int) int {
	index := hm.getIndex(key)
	var prev *ListNode
	node := hm.bucket[index]
	for node != nil {
		if node.data.key == key {
			if prev == nil {
				// удаление из головы списка
				hm.bucket[index] = node.next
			} else {
				// удаление промежуточного узла списка
				prev.next = node.next
			}
			return node.data.val
		}
		prev = node
		node = node.next
	}
	return -1
}

func (hm HashMap) getIndex(key int) int {
	return hm.hash(key) % len(hm.bucket)
}

func (hm HashMap) hash(key int) int {
	return key
}

func main() {
	var n, key, val int
	fmt.Scan(&n)
	hm := NewHashMap(100003)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		scanner.Scan()
		args := strings.Fields(scanner.Text())
		result := -2
		if len(args) > 1 {
			key, _ = strconv.Atoi(args[1])
		}
		if len(args) > 2 {
			val, _ = strconv.Atoi(args[2])
		}
		switch args[0] {
		case "get":
			result = hm.get(key)
		case "put":
			hm.put(key, val)
		case "delete":
			result = hm.delete(key)
		}
		switch result {
		case -2:
		case -1:
			fmt.Fprintln(writer, "None")
		default:
			fmt.Fprintln(writer, result)
		}
	}
	writer.Flush()
}
