package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

// Pair is what we use as the mutation struct. x = from, y = to.
type Pair struct {
    x int
    y int
}

// Checks to see if the error is nil.
func check(err error) {
    if err != nil {
        panic(err)
    }
}

// Helper that converts a string to an int.
func toInt(s string) int {
    x, err := strconv.Atoi(s)
    check(err)
    return x
}

// Checks to see if an item is in a slice.
func contains(slice []string, item string) bool {
    // Create a map of strings as the keys and the values as an empty struct.
    set := make(map[string]struct{}, len(slice))
    // Go through and set each key to an element in the slice.
    for _, s := range slice {
        set[s] = struct{}{}
    }
    // Finally search for the item in the map. If it is not there
    // it will return an error.
    _, ok := set[item]
    return ok
}

func swapPosition(ops []Pair, list []string) []string {
    // Create result array.
    result := make([]string, len(list))
    // Give default value of "" to array.
    for i, _ := range result {
        result[i] = "" 
    }

    // For each operation (x to y) move x from list to y in result.
    for _, op := range ops {
        result[op.y-1] = list[op.x-1]
    }

    // For each element in list, if the value is not already in result, put 
    // the value at the first "" it finds.
    for _, x := range list {
        if !contains(result, x) {
            found := false
            counter := 0
            for !found {
                if result[counter] == "" {
                    result[counter] = x
                    found = true
                }
                counter++
            }
        }
    }
    // Return the result.
    return result
}

// Helper method to print the array.
func printArray(slice [] string) {
    for i, x := range slice {
        if i != len(slice)-1 {
            fmt.Printf("%s ", x)
        } else {
            fmt.Printf("%s", x)
        }
    }
    fmt.Println()
}

func main() {
    // Reads from stdin.
    reader := bufio.NewReader(os.Stdin)
 
    // Gets the number of data sets from the input.
    datasetsStr, err := reader.ReadString('\n')
    check(err) 
    // Strips the new line off of the string.
    datasetsStr = strings.TrimRight(datasetsStr, "\r\n") 
    // Convert the number of data sets to a string.
    datasets := toInt(datasetsStr)

    // For all data sets.
    for i := 0; i < datasets; i++ {
        // Mutation information is on one line. We split it up here.
        startStr, err := reader.ReadString('\n')
        check(err)
        startStr = strings.TrimRight(startStr, "\r\n")
        
        startLine := strings.Split(startStr, " ")
        for j, string := range startLine {
            startLine[j] = strings.TrimSpace(string)
        }
        // Finally get the items and the operations off that one line.
        items := toInt(startLine[0])
        operations := toInt(startLine[1]) 

        // This is the actual list we will be looking at.
        listStr, err := reader.ReadString('\n')
        check(err)
        listStr = strings.TrimRight(listStr, "\r\n")

        // Splits the list up by space.
        list := make([]string, items)
        list = strings.Split(listStr, " ")
        
        // Create a list of pairs to store our operations.
        ops := []Pair{} 
        // Set all the operations we are going to do.
        for k := 0; k < operations; k++ {
            opStr, err := reader.ReadString('\n')
            check(err)
            opStr = strings.TrimRight(opStr, "\r\n")

            opsArray := strings.Split(opStr, " ")
            x := toInt(opsArray[0])
            y := toInt(opsArray[1])
            pair := Pair{x: x, y:y}
            ops = append(ops, pair)
        }
        // Apply the swaps.
        result := swapPosition(ops, list)
        // Print the results.
        printArray(result)
    }
}
