package don

import (
    "fmt"
    "io"
)

type thisOrThatFn func([]interface{}) (int, error)

func thisOrThat(ok bool, a []interface{}, b []interface{}, f thisOrThatFn) (int, error) {
    if ok {
        return f(a)
    }
    return f(b)
}

// PrintElse is like fmt.Print. If ok, a is passed to the function, otherwise b.
func PrintElse(ok bool, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Print(x...)
    })
}

// PrintfElse is like fmt.Printf. If ok, a is passed to the function, otherwise b.
func PrintfElse(ok bool, format string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Printf(format, x...)
    })
}

// PrintlnElse is like fmt.Println. If ok, a is passed to the function, otherwise b.
func PrintlnElse(ok bool, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Println(x...)
    })
}

// FprintElse is like fmt.Fprint. If ok, a is passed to the function, otherwise b.
func FprintElse(ok bool, w io.Writer, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fprint(w, x...)
    })
}

// FprintfElse is like fmt.Fprintf. If ok, a is passed to the function, otherwise b.
func FprintfElse(ok bool, w io.Writer, format string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fprintf(w, format, x...)
    })
}

// FprintlnElse is like fmt.Fprintln. If ok, a is passed to the function, otherwise b.
func FprintlnElse(ok bool, w io.Writer, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fprintln(w, x...)
    })
}

// FscanElse is like fmt.Fscan. If ok, a is passed to the function, otherwise b.
func FscanElse(ok bool, r io.Reader, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fscan(r, x...)
    })
}

// FscanfElse is like fmt.Fscanf. If ok, a is passed to the function, otherwise b.
func FscanfElse(ok bool, r io.Reader, format string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fscanf(r, format, x...)
    })
}

// FscanlnElse is like fmt.Fscanln. If ok, a is passed to the function, otherwise b.
func FscanlnElse(ok bool, r io.Reader, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Fscanln(r, x...)
    })
}

// ScanElse is like fmt.Scan. If ok, a is passed to the function, otherwise b.
func ScanElse(ok bool, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Scan(x...)
    })
}

// ScanfElse is like fmt.Scanf. If ok, a is passed to the function, otherwise b.
func ScanfElse(ok bool, format string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Scanf(format, x...)
    })
}

// ScanlnElse is like fmt.Scanln. If ok, a is passed to the function, otherwise b.
func ScanlnElse(ok bool, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Scanln(x...)
    })
}

// SscanElse is like fmt.Sscan. If ok, a is passed to the function, otherwise b.
func SscanElse(ok bool, str string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Sscan(str, x...)
    })
}

// SscanfElse is like fmt.Sscanf. If ok, a is passed to the function, otherwise b.
func SscanfElse(ok bool, str string, format string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Sscanf(str, format, x...)
    })
}

// SscanlnElse is like fmt.Sscanln. If ok, a is passed to the function, otherwise b.
func SscanlnElse(ok bool, str string, a []interface{}, b []interface{}) (int, error) {
    return thisOrThat(ok, a, b, func(x []interface{}) (int, error) {
        return fmt.Sscanln(str, x...)
    })
}

// SprintElse is like fmt.Sprint. If ok, a is passed to the function, otherwise b.
func SprintElse(ok bool, a []interface{}, b []interface{}) string {
    if ok {
        return fmt.Sprint(a...)
    }
    return fmt.Sprint(b...)
}

// SprintfElse is like fmt.Sprintf. If ok, a is passed to the function, otherwise b.
func SprintfElse(ok bool, format string, a []interface{}, b []interface{}) string {
    if ok {
        return fmt.Sprintf(format, a...)
    }
    return fmt.Sprintf(format, b...)
}

// SprintlnElse is like fmt.Sprintln. If ok, a is passed to the function, otherwise b.
func SprintlnElse(ok bool, a []interface{}, b []interface{}) string {
    if ok {
        return fmt.Sprintln(a...)
    }
    return fmt.Sprintln(b...)
}

// ErrorfElse is like fmt.Errorf. If ok, a is passed to the function, otherwise b.
func ErrorfElse(ok bool, format string, a []interface{}, b []interface{}) error {
    if ok {
        return fmt.Errorf(format, a...)
    }
    return fmt.Errorf(format, b...)
}
