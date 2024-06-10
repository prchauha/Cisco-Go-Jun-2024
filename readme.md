# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Software Requirement
- Go Tools (https://go.dev/dl)
    - verification
        ```shell
        go version
        ```
- Visual Studio Code (https://code.visualstudio.com)

## Repository
- https://github.com/tkmagesh/cisco-go-jun-2024

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (15 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 3:00 PM (15 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoints
- Code & Discuss
- Floor is open for Q&A at all times during the class

## Why Go?
- Concurrency
    - Managed Concurrency using "Goroutines"
    - Very cheap (~4KB)
    - Very efficient
    - Support is built in  the "language" itself
        - go "keyword", channel "data type", channel "operator" ( <- ), range & select-case "constructs"
    - API support through standard library
        - "sync" package
        - "sync/atomic" package
- Lightweight
    - Compiled to machine code
    - Execution Speed (equivalent to c++)
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public/private/protected)
    - No reference types (everything is a value)
    - No class (only structs)
    - No inheritance (only composition)
    - No exceptions (only errors)
    - No "try-catch-finally"
    - No implicity type conversion
    - No pointer arithmatic

## Standard Library
    - https://pkg.go.dev/std

## "go" Tool
Create a build
```shell
go build [file_name_1.go] [file_name_2.go] ...
go build -o [binary_name] [file_name_1.go] [file_name_2.go] ...
```
Build and execute
```shell
go run [file_name.go]
```
## Data Types
- string
- bool
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- alias
    - byte (alias for unnsigned int)
    - rune (alias for unicode code point)

## variable
### Declaration
    - Using the "var" keyword
        - memory allocated
        - memory initialized with the "zero" value of the data type (by default)
        - can be used in package & function scope
    - Using ":="
        - memory allocated
        - memory initialized with the given value
        - cannot use in the package scope
