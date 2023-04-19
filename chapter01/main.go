package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	s := "https://google.com/search?q=rush&safe=active"

	u, err := url.Parse(s)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

	// Expected result: https://google.com/search?q=golang&safe=active
}
