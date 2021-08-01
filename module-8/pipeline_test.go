package pipeline

import (
	"fmt"
	"runtime"
	"testing"
)

var urls = []string{
	"https://golang.org/ref/spec",
	"https://golang.org/doc/effective_go.html",
	"https://dave.cheney.net/2016/08/20/solid-go-design",
	"https://dave.cheney.net/2019/06/10/constant-time",
	"https://dave.cheney.net/2019/05/07/prefer-table-driven-tests",
	"https://dave.cheney.net/2019/01/29/you-shouldnt-name-your-variables-after-their-types-for-the-same-reason-you-wouldnt-name-your-pets-dog-or-cat",
	"https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors",
	"https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common",
	"https://dave.cheney.net/2018/12/28/the-office-coffee-model-of-concurrent-garbage-collection",
	"https://dave.cheney.net/2018/09/03/maybe-adding-generics-to-go-is-about-syntax-after-all",
	"https://dave.cheney.net/2018/07/16/using-go-modules-with-travis-ci",
	"https://dave.cheney.net/2018/07/14/taking-go-modules-for-a-spin",
	"https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics",
	"https://dave.cheney.net/2018/07/12/slices-from-the-ground-up",
	"https://dave.cheney.net/2018/01/08/gos-hidden-pragmas",
}

func TestSearchWord(t *testing.T) {
	input := make(chan string, len(urls))

	for _, url := range urls {
		input <- url
	}
	close(input)

	count, err := SearchWord(input, "Go")
	if err != nil {
		t.Error(err)
		return
	}

	if count != 518 {
		t.Fail()
		t.Fatalf("unexpected result got %d, expected %d\n", count, 541)
	}
}

func BenchmarkSearchWord15(b *testing.B) {
	fmt.Println(runtime.NumCPU())
	for i := 0; i < b.N; i++ {

		input := make(chan string, len(urls))
		for _, url := range urls {
			input <- url
		}
		close(input)

		count, err := SearchWord(input, "Go")
		if err != nil {
			b.Error(err)
			return
		}

		if count != 518 {
			b.Fail()
			b.Fatalf("unexpected result got %d, expected %d\n", count, 5)
		}
	}
}
