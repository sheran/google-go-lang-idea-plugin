package main

// simple variable
var /*begin*/va/*end.Unused global*/ int = 3

var ExportedVariable int = 10

// variable with parenthesis
var (
    /*begin*/vb/*end.Unused global*/ int = 2
    vc float = 3
    /*begin*/v/*end.Unused global*/ = 3
)

func main() {
    a, _, /*begin*/k/*end.Unused variable*/ := 5, 6, 7
    println(vc)
}
