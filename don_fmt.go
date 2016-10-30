package don

import (
    "fmt"
    "io"
)

type fmtFn func() (int, error)

func formatIt(ok bool, f fmtFn) (int, error) {
    if ok {
        return f()
    }
    return 0, nil
}

// Print is like fmt.Print, but is only executed if ok is true.
func Print(ok bool, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Print(a...)
    })
}

// Printf is like fmt.Printf, but is only executed if ok is true.
func Printf(ok bool, format string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Printf(format, a...)
    })
}

// Println is like fmt.Println, but is only executed if ok is true.
func Println(ok bool, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Println(a...)
    })
}

// Fprint is like fmt.Fprint, but is only executed if ok is true.
func Fprint(ok bool, w io.Writer, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fprint(w, a...)
    })
}

// Fprintf is like fmt.Fprintf, but is only executed if ok is true.
func Fprintf(ok bool, w io.Writer, format string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fprintf(w, format, a...)
    })
}

// Fprintln is like fmt.Fprintln, but is only executed if ok is true.
func Fprintln(ok bool, w io.Writer, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fprintln(w, a...)
    })
}

// Fscan is like fmt.Fscan, but is only executed if ok is true.
func Fscan(ok bool, r io.Reader, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fscan(r, a...)
    })
}

// Fscanf is like fmt.Fscanf, but is only executed if ok is true.
func Fscanf(ok bool, r io.Reader, format string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fscanf(r, format, a...)
    })
}

// Fscanln is like fmt.Fscanln, but is only executed if ok is true.
func Fscanln(ok bool, r io.Reader, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Fscanln(r, a...)
    })
}

// Scan is like fmt.Scan, but is only executed if ok is true.
func Scan(ok bool, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Scan(a...)
    })
}

// Scanf is like fmt.Scanf, but is only executed if ok is true.
func Scanf(ok bool, format string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Scanf(format, a...)
    })
}

// Scanln is like fmt.Scanln, but is only executed if ok is true.
func Scanln(ok bool, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Scanln(a...)
    })
}

// Sscan is like fmt.Sscan, but is only executed if ok is true.
func Sscan(ok bool, str string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Sscan(str, a...)
    })
}

// Sscanf is like fmt.Sscanf, but is only executed if ok is true.
func Sscanf(ok bool, str string, format string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Sscanf(str, format, a...)
    })
}

// Sscanln is like fmt.Sscanln, but is only executed if ok is true.
func Sscanln(ok bool, str string, a ...interface{}) (int, error) {
    return formatIt(ok, func() (int, error){
        return fmt.Sscanln(str, a...)
    })
}

// Errorf is like fmt.Errorf, but is only executed if ok is true.
func Errorf(ok bool, format string, a ...interface{}) error {
    if ok {
        return fmt.Errorf(format, a...)
    }
    return nil
}
