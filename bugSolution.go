```go
package main

import "fmt"

func main() {

    //Creating a map
    m := make(map[string]int)

    //Inserting values into map
    m["a"] = 10
    m["b"] = 20
    m["c"] = 30

    // Accessing map elements and handling potential errors
    value, ok := m["a"]
    if ok {
        fmt.Println("a exists", value)
    } else {
        fmt.Println("a does not exist")
    }

    value, ok = m["d"]
    if ok {
        fmt.Println("d exists", value)
    } else {
        fmt.Println("d does not exist")
    }

    // Deleting an element
    delete(m, "b")

    //Checking for element existence
    value, ok = m["c"]
    if ok {
        fmt.Println("c exists", value)
    } else {
        fmt.Println("c does not exist")
    }

    //Iterating over map
    for key, value := range m {
        fmt.Println(key, value)
    }

}
```