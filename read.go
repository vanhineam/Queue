package main

import (
    "fmt"
    "bufio"
    "os"
)

// Checks to see if the error is nil.
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    // Reads from stdin.
    reader := bufio.NewReader(os.Stdin)

    // Read until end of file
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Print(line)
    }
 
}
