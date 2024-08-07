# GoCryptoTrader package cache

<img src="https://github.com/aaabigfish/gocryptotrader/blob/master/web/src/assets/page-logo.png?raw=true" width="350px" height="350px" hspace="70">


[![Build Status](https://github.com/aaabigfish/gocryptotrader/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/aaabigfish/gocryptotrader/actions/workflows/tests.yml)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/aaabigfish/gocryptotrader/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/aaabigfish/gocryptotrader?status.svg)](https://godoc.org/github.com/aaabigfish/gocryptotrader/portfolio)
[![Coverage Status](http://codecov.io/github/thrasher-corp/gocryptotrader/coverage.svg?branch=master)](http://codecov.io/github/thrasher-corp/gocryptotrader?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/aaabigfish/gocryptotrader)](https://goreportcard.com/report/github.com/aaabigfish/gocryptotrader)


This cache package is part of the GoCryptoTrader codebase.

## This is still in active development

You can track ideas, planned features and what's in progress on this Trello board: [https://trello.com/b/ZAhMhpOy/gocryptotrader](https://trello.com/b/ZAhMhpOy/gocryptotrader).

Join our slack to discuss all things related to GoCryptoTrader! [GoCryptoTrader Slack](https://join.slack.com/t/gocryptotrader/shared_invite/enQtNTQ5NDAxMjA2Mjc5LTc5ZDE1ZTNiOGM3ZGMyMmY1NTAxYWZhODE0MWM5N2JlZDk1NDU0YTViYzk4NTk3OTRiMDQzNGQ1YTc4YmRlMTk)

## Current Features for cache package

+ Basic LRU cache system with both goroutine safe (via mutex locking) and non-goroutine safe options

## How to use

##### Basic Usage:

```go
package main

import ("github.com/aaabigfish/gocryptotrader/common/cache")

func main() {
	lruCache := cache.New(5)
	lruCache.Add("hello", "world")
	c := lruCache.Contains("hello")
	if !c {
		fmt.Println("expected cache to contain \"hello\" key")
	}

	v := lruCache.Get("hello")
	if v == nil {
		fmt.Println("expected cache to contain \"hello\" key")
	}
	fmt.Println(v)
}
```
## Contribution

Please feel free to submit any pull requests or suggest any desired features to be added.

When submitting a PR, please abide by our coding guidelines:

+ Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
+ Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
+ Code must adhere to our [coding style](https://github.com/aaabigfish/gocryptotrader/blob/master/doc/coding_style.md).
+ Pull requests need to be based on and opened against the `master` branch.

## Donations

<img src="https://github.com/aaabigfish/gocryptotrader/blob/master/web/src/assets/donate.png?raw=true" hspace="70">

If this framework helped you in any way, or you would like to support the developers working on it, please donate Bitcoin to:

***bc1qk0jareu4jytc0cfrhr5wgshsq8282awpavfahc***

