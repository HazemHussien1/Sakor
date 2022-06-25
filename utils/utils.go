package utils

import "fmt"

//filter outputWordlist and print it
func FilterAndPrint(slice []string) {
  for i := 0; i < len(slice); i++ {
	if Contains(slice[i+1:], slice[i]) {
	  slice = Remove(slice, i)
	  i--
	}
  }
  for i := 0; i < len(slice); i++ {
	fmt.Printf("%s\n", slice[i])
  }
}

func Remove(slice []string, s int) []string {
  return append(slice[:s], slice[s+1:]...)
}

func Contains(s []string, e string) bool {
  for _, a := range s {
	if a == e {
	  return true
	}
  }
  return false
}

func Usage() {
  fmt.Printf("Usage for finding subdomains:")
  fmt.Printf("cat domains.txt | ./Sakor")
  fmt.Printf("Usage for crawling urls:")
  fmt.Printf("cat urls.txt | ./Sakor")
}
