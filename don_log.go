package don

import "log"

// Fatal is like log.Fatal, but is only executed if ok is true.
func Fatal(ok bool, v...interface{}) {
    if ok { log.Fatal(v...) }
}

// Fatalf is like log.Fatalf, but is only executed if ok is true.
func Fatalf(ok bool, format string, v...interface{}) {
    if ok { log.Fatalf(format, v...)  }
}

// Fatalln is like log.Fatalln, but is only executed if ok is true.
func Fatalln(ok bool, v...interface{}) {
    if ok { log.Fatalln(v...) }
}

// Panic is like log.Panic, but is only executed if ok is true.
func Panic(ok bool, v...interface{}) {
    if ok { log.Panic(v...) }
}

// Panicf is like log.Panicf, but is only executed if ok is true.
func Panicf(ok bool, format string, v...interface{}) {
    if ok { log.Panicf(format, v...) }
}

// Panicln is like log.Panicln, but is only executed if ok is true.
func Panicln(ok bool, v...interface{}) {
    if ok { log.Panicln(v...) }
}
