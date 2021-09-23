package main

/**
Хотя в задании написано, что Value — целое число
храним элементы Дека строками, чтобы не конвертировать
туда-сюда данные со входа.
Можем себе позволить, т.к. никаких арифметических действий
с элементами Дека не выполняется
*/
type Deque struct {
	items       []string
	front, back int16
}

func NewDeque(size uint16) Deque {
	return Deque{
		items: make([]string, size),
		front: -1,
		back:  -1,
	}
}

func (d *Deque) PushBack(x string) bool {
	if d.IsFull() {
		return false
	}
	if d.IsEmpty() {
		d.front, d.back = 0, 0
	} else {
		d.front = d.next(d.front)
	}
	d.items[d.front] = x
	return true
}

func (d *Deque) PushFront(x string) bool {
	if d.IsFull() {
		return false
	}
	if d.IsEmpty() {
		d.front, d.back = 0, 0
	} else {
		d.back = d.prev(d.back)
	}
	d.items[d.back] = x
	return true
}

func (d *Deque) PopFront() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}
	popped := d.items[d.back]
	if d.back == d.front {
		d.front, d.back = -1, -1
	} else {
		d.back = d.next(d.back)
	}
	return popped, true
}

func (d *Deque) PopBack() (string, bool) {
	if d.IsEmpty() {
		return "", false
	}
	popped := d.items[d.front]
	if d.back == d.front {
		d.front, d.back = -1, -1
	} else {
		d.front = d.prev(d.front)
	}
	return popped, true
}

func (d *Deque) IsFull() bool {
	return d.next(d.front) == d.back
}

func (d *Deque) IsEmpty() bool {
	return d.front == -1
}

func (d *Deque) next(x int16) int16 {
	return (x + 1) % int16(cap(d.items))
}

func (d *Deque) prev(x int16) int16 {
	if x == 0 {
		return int16(cap(d.items)) - 1
	}
	return x - 1
}
