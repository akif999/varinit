package a

import "fmt"

func a() {
	var hoge string = "hoge" // want "identifier is meaningless"

	fuga(hoge) // want "identifier is meaningless" "identifier is meaningless"
}

func fuga(s string) { // want "identifier is meaningless"
	fmt.Println(s)
}
