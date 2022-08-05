package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []int = []int{1, 40, 30, 31, 20}
	Sort(numbers)
	fmt.Println(numbers)
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root.String())
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var values []string
	var searchTree func(t *tree)
	searchTree = func(t *tree) {
		if t.left != nil {
			searchTree(t.left)
		}
		if t.right != nil {
			searchTree(t.right)
		}
		values = append(values, strconv.Itoa(t.value))
	}
	searchTree(t)
	result := strings.Join(values, ",")
	return result
}
