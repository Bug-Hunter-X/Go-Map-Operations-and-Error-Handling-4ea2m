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

    // Accessing map elements
    fmt.Println(m["a"]) // Output: 10
    fmt.Println(m["d"]) // Output: 0 (zero value)

    // Deleting an element
    delete(m, "b")

    // Checking for element existence
    value, ok := m["c"]
    if ok {
        fmt.Println("c exists", value) // Output: c exists 30
    }
    value, ok = m["b"]
    if ok {
        fmt.Println("b exists", value) //This line won't execute
    } else {
        fmt.Println("b does not exist") // Output: b does not exist
    }

    //Iterating over map
    for key, value := range m {
        fmt.Println(key, value)
    }

}
```