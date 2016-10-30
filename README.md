# don

Don (do on) is a library to remove clutter from repetitive `if-fmt` or `if-log` calls.

Don functions match the corresponding fmt or log function names. The only addtion is that the first parameter of every function consists of a boolean value, and that the function is only executed if it is true.

Version: 0.1  
Installation: `go get github.com/thibran/don`


Examples
-------

```go
import (
    "github.com/thibran/don"
    "os"
    )

func main() {
    // if true, exit with os.Exit(1) – like log.Fatalln
    don.Fatalln(len(os.Args) < 2, "No download URL specified!")
    // print (using fmt.Println internally) only if the condition is true
    don.Println(os.Getenv("foobar") != "", "env is set")
}
```


The don package contains additional X+Else functions. If the condition is true, the first slice value is taken, otherwise the second.

```go
import "github.com/thibran/don"

func main() {
    number := 2
    don.PrintfElse(number == 1, "downloading %v file%s\n",
        []interface{}{"one", ""},
        []interface{}{number, "s"},
    ) // prints to stdout, using fmt.Printf: downloading 2 files
}
```

```go
import (
    "github.com/thibran/don"
    "flag"
)

func main() {
    showVersion := flag.Bool("version", false, "myapp version")
    flag.Parse()
    // exit the app with os.Exit(0), if the version flag has been set
    don.ExitlnSuccessful(*showVersion, "myapp 0.7")
}
```
