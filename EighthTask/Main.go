package EighthTask

import (
	"fmt"
	"golang.org/x/tour/tree"
	"reflect"
	"sort"
	"time"
)

// Упражнение: равнозначные двоичные деревья

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	tree1Array := make([]int, 10)
	tree2Array := make([]int, 10)
	t1Chan := make(chan int)
	go Walk(t1, t1Chan)
	t2Chan := make(chan int)
	go Walk(t2, t2Chan)
	for i := 0; i < 10; i++ {
		num := <- t1Chan
		num2 := <- t2Chan
		tree1Array[i] = num
		tree2Array[i] = num2
	}
	sort.Ints(tree1Array)
	sort.Ints(tree2Array)
	return reflect.DeepEqual(tree1Array, tree2Array)
}

// Упражнение: поисковый робот

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchResult struct {
	body string
	urls[]string
	err error
}

type void struct{}
var member void

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Implementing a function without changing the signature (almost)
func Crawl(url string, depth int, fetcher Fetcher, fetchedUrls... map[string]void) {
	if depth <= 0 {
		return
	}
	var localFetchedUrls map[string]void
	if len(fetchedUrls) != 0 {
		localFetchedUrls = fetchedUrls[0]
	} else {
		localFetchedUrls = make(map[string]void)
	}
	_, urlExists := localFetchedUrls[url]
	if urlExists {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	localFetchedUrls[url] = member
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, localFetchedUrls)
	}
	time.Sleep(time.Millisecond)
	return
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func Main() {
	// Упражнение: равнозначные двоичные деревья
	myTree := tree.New(1)
	myTree2 := tree.New(1)
	fmt.Println(Same(myTree, myTree2))

	// Упражнение: поисковый робот
	Crawl("http://golang.org/", 4, fetcher)
}

