package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

//<TODO>  write virustotal.go
//<TODO>

func filterAndPrint(slice []string) {
	//filter outputWordlist and print it
	for i := 0; i < len(slice); i++ {
		if contains(slice[i+1:], slice[i]) {
			slice = remove(slice, i)
			i--
		}
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%s\n", slice[i])
	}

}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func usage() {
	fmt.Printf("Usage for finding subdomains:")
	fmt.Printf("cat domains.txt | domainRecon")
	fmt.Printf("Usage for crawling urls")
	fmt.Printf("cat urls.txt | domainRecon")
}

func main() {

	var threads int
	flag.IntVar(&threads, "threads", 10, "number of threads")

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "verbose option")

	var c bool
	flag.BoolVar(&c, "c", false, "crawl urls")

	var d int
	flag.IntVar(&d, "d", 2, "depth to crawl")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	if c {
		for scanner.Scan() {
			crawl(scanner.Text(), d, threads)
		}

		if scanner.Err() != nil {
			usage()
			os.Exit(3)
		}

	} else {
		for scanner.Scan() {
			crtshDomains, err := GetCrtshSubs(scanner.Text())
			if err != nil {
				fmt.Printf("crtsh failed")
			}

			certSpotterDomains, err := GetCertSpotterSubs(scanner.Text())
			if err != nil {
				fmt.Printf("crtsh failed")
			}

			// virusTotalDomains, err := GetVirusTotalSubs(scanner.Text())
			// if err != nil {
			// 	fmt.Printf("crtsh failed")
			// }

			// need to do this in a better way
			allDomains := append(crtshDomains, certSpotterDomains...)
			// allDomains = append(allDomains, virusTotalDomains...)

			filterAndPrint(allDomains)
		}

		if scanner.Err() != nil {
			usage()
			os.Exit(3)
		}
	}
}
