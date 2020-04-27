package main

import (
	"fmt"
)

// Fetcher interface
type Fetcher interface {
	// Fetch 返回 URL 的 body 内容, 并且将在这个页面上找到的 URL 放到一个 slice 中
	Fetch(url string)(body string, urls []string, err error)
}

var lockx = make(chan int, 1)

// SafeRun function
func SafeRun(f func()){
	<-lockx
	f()
	lockx <- 1
}

var visited = make(map[string]bool)

// Crawl 使用 fetcher 从某个 URL 开始地柜的爬取页面， 直到
func Crawl(url string, depth int, fetcher Fetcher, quit chan int) {
	// TODO: 并行的爬取 URL
	// TODO: 不重复抓取页面
	// 下面并没有实现上面的两种情况。

	// 延迟调用匿名函数
	defer func(){
		quit <- 1
	}()

	// 如果是访问过的 URL 就直接返回
	if depth <= 0 || visited[url] {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	// 通过 lock 锁把 URL 添加到 visited 中
	SafeRun(func(){
		visited[url] = true
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	subQuit := make(chan int, len(urls))

	for _, u := range urls {
		Crawl(u, depth - 1, fetcher, subQuit)
	}

	for i := 0; i < len(urls); i++ {
		<-subQuit
	}

	return 
}

func main() {
	lockx <- 1
	quit := make(chan int, 1)
	Crawl(url, 4, fetcher, quit)
	fmt.Println(<-quit)
}

// fakeFetcher 是返回若干结果的 Fetcher
type fakeFetcher map[string]*fakeResult


type fakeResult struct {
	body string
	urls []string
}


func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)

}


var url = "https://docs.studygolang.com/"

var fetcher = fakeFetcher{
	url: &fakeResult{
		"The Go Programming Language",
		[]string{
			url + "pkg/",
			url + "cmd/",
		},
	},
	url+"pkg/": &fakeResult{
		"Packages",
		[]string{
			url,
			url + "cmd/",
			url + "pkg/fmt/",
			url + "pkg/os/",
		},
	},
	url + "pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			url,
			url + "pkg/",
		},
	},
	url + "pkg/os/": &fakeResult{
		"Package os",
		[]string{
			url,
			url + "pkg/",
		},
	},
}







