# Go Programming Tutorials

This repository contains a collection of Go programming tutorials covering various fundamental concepts and advanced features of the Go programming language.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Tutorial Topics](#tutorial-topics)
  - [Basic Syntax and Functions](#basic-syntax-and-functions)
  - [Arrays and Slices](#arrays-and-slices)
  - [Maps](#maps)
  - [Structs and Interfaces](#structs-and-interfaces)
  - [Pointers](#pointers)
  - [Goroutines and Concurrency](#goroutines-and-concurrency)
- [Running the Examples](#running-the-examples)
- [Contributing](#contributing)

## Introduction

This repository serves as a learning resource for Go programming language. Each tutorial focuses on specific Go concepts with practical examples to help you understand and apply these concepts in real-world scenarios.

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.16 or later recommended)
- Basic understanding of programming concepts

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/go-tutorials.git
   cd go-tutorials
   ```

2. Run any example by navigating to its directory and using the `go run` command:
   ```bash
   cd cmd/tutorial/Arrays
   go run array_1.go
   ```

## Tutorial Topics

### Basic Syntax and Functions

Located in `cmd/tutorial/main.go`, this tutorial covers:

- Basic Go syntax
- Function declarations and calls
- Error handling with multiple return values
- Control flow with if-else statements and switch cases

Example:

```go
func intDivision(numerator int, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("Cannot divide by zero")
    }
    result := numerator / denominator
    remainder := numerator % denominator
    return result, remainder, nil
}
```

### Arrays and Slices

Located in `cmd/tutorial/Arrays/array_1.go`, this tutorial covers:

- Array declaration and initialization
- Array indexing and slicing
- Slice operations (append, make)
- Memory layout of arrays
- Differences between arrays and slices

Example:

```go
// Array declaration
intArr := [...]int32{44, 2, 3, 4, 5}

// Slice declaration and operations
var intSlice []int32 = []int32{4, 5, 6}
intSlice = append(intSlice, 7)
```

### Maps

Located in `cmd/tutorial/Map/map1.go`, this tutorial covers:

- Map declaration and initialization
- Adding, accessing, and deleting map elements
- Checking if a key exists in a map
- Iterating over maps
- Various loop constructs in Go

Example:

```go
var myMap2 = map[string]uint8{"Adam": 12, "Sarah": 23, "John": 34}
delete(myMap2, "John")

// Check if key exists
var age, ok = myMap2["John"]
if ok {
    fmt.Println("John is in the map")
} else {
    fmt.Println("John is not in the map")
}

// Iterate over map
for name, age := range myMap2 {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### Structs and Interfaces

Located in `cmd/tutorial/Strucs_and_interfaces/main.go`, this tutorial covers:

- Struct definition and initialization
- Methods on structs
- Interface definition and implementation
- Polymorphism through interfaces

Example:

```go
type gasEngine struct {
    mpg uint16
    gals uint16
}

type engine interface{
    milesLeft() uint16
}

func (e gasEngine) milesLeft() uint16 {
    return e.gals * e.mpg
}

// Generic function that works with any engine type
func canMakeIt(e engine, miles uint16) bool {
    remaining := e.milesLeft()
    // Implementation...
}
```

### Pointers

Located in `cmd/tutorial/pointer/pointer1.go`, this tutorial covers:

- Pointer declaration and usage
- Memory addresses and dereferencing
- Passing by reference vs. passing by value
- Memory efficiency with pointers

Example:

```go
func square(thing2 *[5]float64) [5]float64 {
    fmt.Printf("\nThe memory location that thing2 points to is: %p", thing2)
    for i := range thing2 {
        thing2[i] = thing2[i] * thing2[i]
    }
    return *thing2
}
```

### Goroutines and Concurrency

Located in `cmd/tutorial/GoRoutines/routing.go`, this tutorial covers:

- Goroutines for concurrent execution
- Synchronization with WaitGroups
- Mutex for safe concurrent access to shared resources
- Simulating database calls with random delays
- Measuring execution time of concurrent operations

Example:

```go
func main() {
    t0 := time.Now();
    for i := 0; i < len(dbData); i++ {
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
    fmt.Printf("\n Total execution time: %s", time.Since(t0))
}

func save(data string) {
    m.Lock()
    result = append(result, data)
    m.Unlock()
}
```

## Running the Examples

Each tutorial can be run independently. Navigate to the specific directory and use the `go run` command:

```bash
# Run the goroutines example
cd cmd/tutorial/GoRoutines
go run routing.go

# Run the pointers example
cd cmd/tutorial/pointer
go run pointer1.go
```

## Contributing

Contributions are welcome! If you'd like to add more examples or improve existing ones:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request
