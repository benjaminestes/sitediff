package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/benjaminestes/sitediff/sitemap"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tsitediff get <url>")
	fmt.Println("\tsitediff diff <old> <new>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	if os.Args[1] == "get" && len(os.Args) == 3 {
		urls, err := sitemap.FetchAll(os.Args[2])
		if err != nil {
			log.Fatalf("failed to fetch: %v", err)
		}

		for _, u := range urls {
			fmt.Println(u)
		}

		return
	}

	if os.Args[1] == "diff" && len(os.Args) == 4 {
		oldfile, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatalf("error opening old: %v", err)
		}

		oldlines := make(map[string]bool)

		oldscan := bufio.NewScanner(oldfile)

		for oldscan.Scan() {
			oldlines[oldscan.Text()] = true
		}

		newfile, err := os.Open(os.Args[3])
		if err != nil {
			log.Fatalf("error opening old: %v", err)
		}

		newscan := bufio.NewScanner(newfile)

		for newscan.Scan() {
			if oldlines[newscan.Text()] == false {
				fmt.Println(newscan.Text())
			}
		}

		return
	}

	usage()
	return
}
