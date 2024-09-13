# RemoveAtIndex

RemoveAtIndex is a Go package that provides a simple and efficient way to remove an element at a specified index from a slice of any type.

## Installation

To install RemoveAtIndex, use `go get`:

```bash
go get github.com/mike-keough/removeatindex
```
**Useage:
```go

package main

import (
    "fmt"
    "github.com/mike-keough/removeatindex"
)

//using Slice interface for anytype
var newSlice = Slice{1,2,"foo","bar",5}

func main() {
    RemoveAtIndex(newSlice, 2) //output [1, 2, "bar", 5]
}

```

**API: 
```go 
type Slice []interface{} 
```
Slice is a slice of empty interfaces, allowing it to hold elements of any type.

**Function:
```go
func RemoveAtIndex(s Slice, index int) Slice
```
RemoveAtIndex removes the element at the specified index from the slice and returns the modified slice. If the index is out of bounds, it returns the original slice unchanged.

**Features:
*Remove elements from slices of any type
*Efficient implementation using Go's built-in append function
*Preserves the order of remaining elements
*Handles out-of-bounds indices gracefully

**Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

**License:
This project is licensed under the MIT License - see the LICENSE file for details.