package main

import (
	"algos/datastructures"
	"fmt"
)

func main() {
	l := datastructures.LinkedList{}
	l.InsertAt(7, 1)
	val, ok := l.Get(0)
	fmt.Println(val, ok)
	l.Print()
	println(l.Length())

	l.Append(0)
	l.Prepend(2)

	l.RemoveAt(1)
	l.Print()
	l.Append(5)
	l.Append(7)
	l.Append(9)

	res, _ := l.Get(2)
	if res != 9 {
		t.Errorf("Expected 9, got %d", res)
	}

	res, _ = l.RemoveAt(1)
	if res != 7 {
		t.Errorf("Expected 7, got %d", res)
	}

	if l.Length() != 2 {
		t.Errorf("Expected 2, got %d", l.Length())
	}

	l.Append(11)
	res, _ = l.RemoveAt(1)
	if res != 9 {
		t.Errorf("Expected 9, got %d", res)
	}

	_, ok := l.Remove(9)
	if ok != false {
		t.Errorf("Expected false when removing 9")
	}

	res, _ = l.RemoveAt(0)
	if res != 5 {
		t.Errorf("Expected 5, got %d", res)
	}

	res, _ = l.RemoveAt(0)
	if res != 11 {
		t.Errorf("Expected 11, got %d", res)
	}

	if l.Length() != 0 {
		t.Errorf("Expected 0, got %d", l.Length())
	}

	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)

	res, _ = l.Get(2)
	if res != 5 {
		t.Errorf("Expected 5, got %d", res)
	}

	res, _ = l.Get(0)
	if res != 9 {
		t.Errorf("Expected 9, got %d", res)
	}

	res, _ = l.Remove(9)
	if res != 9 {
		t.Errorf("Expected 9, got %d", res)
	}

	if l.Length() != 2 {
		t.Errorf("Expected 2, got %d", l.Length())
	}

	res, _ = l.Get(0)
	if res != 7 {
		t.Errorf("Expected 7, got %d", res)
	}

}
