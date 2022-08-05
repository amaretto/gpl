package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	sort.Sort(sort.Reverse(Items(result.Items)))

	t := time.Now()
	monthAgo := t.AddDate(0, -1, 0)
	yearAgo := t.AddDate(-1, 0, 0)
	tmp := 0
	for _, item := range result.Items {

		if item.CreatedAt.Unix() > monthAgo.Unix() {
			if tmp != 1 {
				fmt.Println("---Created less than a month ago---")
				tmp = 1
			}
		} else if item.CreatedAt.Unix() > yearAgo.Unix() {
			if tmp != 2 {
				fmt.Println("---Created less than a year ago---")
				tmp = 2
			}
		} else {
			if tmp != 3 {
				fmt.Println("---Created over a year ago---")
				tmp = 3
			}
		}

		fmt.Printf("%v #%-5d %9.9s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
	}
}
