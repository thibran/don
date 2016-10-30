package don

import (
    "fmt"
    "os"
)

// ExitSuccessful exits with return-code 0 if ok is true and prints v to stdout.
func ExitSuccessful(ok bool, v...interface{}) {
    if ok {
        fmt.Print(v...)
		os.Exit(0)
    }
}

// ExitfSuccessful exits with return-code 0 if ok is true and prints v to stdout.
func ExitfSuccessful(ok bool, format string, v...interface{}) {
    if ok {
        fmt.Printf(format, v...)
		os.Exit(0)
    }
}

// ExitlnSuccessful exits with return-code 0 if ok is true and prints v to stdout.
func ExitlnSuccessful(ok bool, v...interface{}) {
    if ok {
        fmt.Println(v...)
		os.Exit(0)
    }
}
