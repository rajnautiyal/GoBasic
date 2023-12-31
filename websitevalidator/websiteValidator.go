package websitevalidator

import (
	"fmt"
	"net/http"
)

func WebValidator() {
	links := []string{
		"https://facebook.com/",
		"https://google.com/",
		"https://stackoverflow.com/",
		"https://golang.org",
	}

	ch := make(chan string)
	for _, link := range links {
		go func(link string) {
			checkLinkUsingChannel(link, ch)
		}(link)
	}
	for l := range ch {
		go checkLinkUsingChannel(l, ch)
	}
}
func checkLinkUsingChannel(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("the link is not working:", link, err)
		ch <- link
		return
	}

	fmt.Println(link, " is up !")
	ch <- link
}

/*func checkLinkUsingChannel(link1 string, ch chan string) (bool, error) {
	_, err := http.Get(link1)
	if err != nil {
		fmt.Println("the link is not working:", link1, err)
		ch <- "might be we are doing "
		return false, err
	}
	ch <- fmt.Sprintf("the link is working fine: %s", link1)
	return true, nil
}*/
