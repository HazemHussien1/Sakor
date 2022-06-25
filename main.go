package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)


func main() {

  var threads int
  flag.IntVar(&threads, "t", 10, "number of threads")
  flag.IntVar(&threads, "threads", 10, "number of threads")

  var verbose bool
  flag.BoolVar(&verbose, "v", false, "verbose option")
  flag.BoolVar(&verbose, "verbose", false, "verbose option")

  var c bool
  flag.BoolVar(&c, "c", false, "crawl urls")
  flag.BoolVar(&c, "crawl", false, "crawl urls")

  var d int
  flag.IntVar(&d, "d", 2, "depth to crawl")
  flag.IntVar(&d, "depth", 2, "depth to crawl")

  flag.Parse()

  scanner := bufio.NewScanner(os.Stdin)
  if c {
	for scanner.Scan() {
	  Crawl(scanner.Text(), d, threads)
	}

	if scanner.Err() != nil {
	  usage()
	  os.Exit(3)
	}

  } else {
	for scanner.Scan() {
	  crtshDomains, err := subEnum.GetCrtshSubs(scanner.Text())
	  if err != nil {
		fmt.Printf("crtsh failed")
	  }

	  certSpotterDomains, err := subEnum.GetCertSpotterSubs(scanner.Text())
	  if err != nil {
		fmt.Printf("crtsh failed")
	  }

	  virusTotalDomains, err := subEnum.GetVirusTotalSubs(scanner.Text())
	  if err != nil {
	  	fmt.Printf("crtsh failed")
	  }

	  // need to do this in a better way
	  allDomains := append(crtshDomains, certSpotterDomains...)
	  allDomains = append(allDomains, virusTotalDomains...)

	  utils.FilterAndPrint(allDomains)
	}

	if scanner.Err() != nil {
	  utils.Usage()
	  os.Exit(3)
	}
  }
}
