// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://www.rust-lang.org/",
			"https://www.rust-lang.org/learn",
			"https://play.rust-lang.org/",
			"https://doc.rust-lang.org/stable/book/",
			"https://www.rust-lang.org/",
			"https://www.rust-lang.org/learn",
			"https://play.rust-lang.org/",
			"https://doc.rust-lang.org/stable/book/",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}
