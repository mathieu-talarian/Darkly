package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var url = `http://192.168.105.128/.hidden`

func flags() (url string, err error) {
	flag.StringVar(&url, "u", "", "Darkly url")
	flag.Parse()
	return
}

func main() {
	URL, err := flags()
	if err != nil {
		log.Fatal(err)
	}
	if URL != "" {
		getURL(URL)
	} else {
		getURL(url)
	}

}

func getURL(url string) {
	r, err := regexp.Compile(`(?m)<a href="(.+)">`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range r.FindAllSubmatch(body, -1) {
		// match[1]
		ns1 := fmt.Sprintf("%s/%s", url, string(match[1]))
		// fmt.Println(ns, ns1)
		if string(match[1]) == "../" {
			continue
		} else if string(match[1]) == "README" {
			getREADME(fmt.Sprintf("%s/%s", url, string(match[1])))
		} else {
			getURL(ns1)
		}
	}
}

func getREADME(url string) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(string(body), "voisin") {
		if !strings.Contains(string(body), "craquer") {
			if !strings.Contains(string(body), "aide") {
				if !strings.Contains(string(body), "Non") {
					fmt.Println(url, "found:", string(body))
				}
			}
		}
	}
}
