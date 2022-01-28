package main

import (
	"fmt"
	"strings"
	int "github.com/conorburke/learn_go_with_tests/integers"
)

const prefix = "Hello,"
const spanish = "spanish"

func main() {
	fmt.Println("Hi from Main")
	fmt.Println(int.Add(2,3))
	fmt.Println(string("chaos"))
}

func SayHello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if strings.ToLower(language) == spanish {
		return "Hola, " + name
	}
	return fmt.Sprintf("%s %s", prefix, name)
}


// package main

// import (
//     "fmt"
//     "io"
//     "log"
//     "net/http"
// )

// func Greet(writer io.Writer, name string) {
//     fmt.Fprintf(writer, "Hello, %s", name)
// }

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
//     Greet(w, "world")
// }

// func main() {
//     log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }