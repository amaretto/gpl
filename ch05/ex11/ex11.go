package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compiler": {
		"data structures",
		"formal languages",
		"computesr organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"linear algebra":        {"calculus"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		log.Fatalf("Topological sort is failed :%v\n", err)
	}
	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	var stack []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if exists(stack, item) {
				return fmt.Errorf("the data containes circulating reference  item:%s", item)
			}
			if !seen[item] {
				seen[item] = true
				stack = append(stack, item)
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				stack = stack[:len(stack)-1]
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return order, err
	}
	return order, nil
}

func exists(stack []string, target string) bool {
	for _, s := range stack {
		if s == target {
			return true
		}
	}
	return false
}
