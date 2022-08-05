package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/amaretto/study-go/golang_study2/ch05/ex13/links"
)

type work struct {
	link  string
	depth int
}

var tokens = make(chan struct{}, 20)

func crawl(target work) []work {
	fmt.Println(target.link)
	tokens <- struct{}{}
	list, err := links.Extract(target.link)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	var works []work
	for _, targetLink := range list {
		works = append(works, work{targetLink, target.depth + 1})
	}
	return works
}

func main() {
	depth := flag.Int("depth", 3, "crawl depth")
	flag.Parse()

	worklist := make(chan []work)
	unseenLinks := make(chan work)
	//finish := make(chan struct{}, 20)

	var initWorks []work
	for _, arg := range flag.Args() {
		initWorks = append(initWorks, work{arg, 0})
	}

	go func() { worklist <- initWorks }()

	for i := 0; i < 20; i++ {
		go func() {
			for target := range unseenLinks {
				if target.depth > *depth {
					//					finish <- struct{}{}
					continue
				}
				foundLinks := crawl(target)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	//	go func() {
	//		var finishCount int
	//		for range finish {
	//			finishCount++
	//			if finishCount == 20 {
	//				break
	//			}
	//		}
	//		close(unseenLinks)
	//		close(worklist)
	//	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, target := range list {
			if !seen[target.link] {
				seen[target.link] = true
				unseenLinks <- work{target.link, target.depth + 1}
			}
		}
	}
}
