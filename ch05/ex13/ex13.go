package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/amaretto/study-go/golang_study2/ch05/ex13/links"
)

func main() {
	err := breadthFirst(crawl, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

var dirPath string = "copy/"

func crawl(u string) ([]string, error) {
	var newList []string
	fmt.Println(u)
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	parsedRoot, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("the root url cannot parse:%s", u)
	}
	// create root file
	dstPath := dirPath + parsedRoot.Hostname()
	err = createTargetDir(dstPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(dstPath + parsedRoot.Path)
	downloadFile(u, dstPath+parsedRoot.Path+"/"+parsedRoot.Hostname()+".html")
	if err != nil {
		return nil, err
	}

	for _, link := range list {
		parsedURL, err := url.Parse(link)
		if err != nil {
			return nil, fmt.Errorf("the link cannot parse:%s", link)
		}
		if strings.HasSuffix(parsedURL.Hostname(), parsedRoot.Hostname()) {
			newList = append(newList, link)
			dstPath := dirPath + parsedURL.Hostname()
			createTargetDir(dstPath)
			dstFileName := dstPath + parsedURL.Path
			downloadFile(link, dstFileName)
		}
	}
	return newList, nil
}

func createTargetDir(dstPath string) error {
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		err := os.Mkdir(dstPath, 0777)
		if err != nil {
			return fmt.Errorf("cannot create directory:%v", err)
		}
	}
	return nil
}

func downloadFile(link, dstFileName string) error {
	resp, err := http.Get(link)
	if err != nil {
		return fmt.Errorf("cannot get file from the url: %v", err)
	}
	defer resp.Body.Close()

	f, err := os.Create(dstFileName)
	if err != nil {
		return fmt.Errorf("cannot create file from the path: %s, err:%v", dstFileName, err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("cannot copy to local file:%v", err)
	}
	return nil
}

func breadthFirst(f func(item string) ([]string, error), worklist []string) error {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				tmp, err := f(item)
				if err != nil {
					return err
				}
				worklist = append(worklist, tmp...)
			}
		}
	}
	return nil
}
