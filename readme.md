# Features of Go

- Statically typed
- Fast run time
- Supports concurrency through goroutines and channels
- Does not support classes, objects and inheritance

# Go quickstart
- Type `go mod init tutorial` to initialize the go.mod file
- Run the code `go run .\main.go`

# Syntax
### A Go file consist of the following parts:

  - Packages declaration
  - Import packages
  - Functions
  - Statements and expressions
### Packages
- Every Go programs is made up of packages
- Programs start running in package `main`

### Variables
#### Datatype
  - `int` `int8` `int16` `int32` `int64`
  - `uint` `uint8` `uint16` `uint32` `uint64`
  - `string`
  - `complex64` `complex128`
  - `bool` stores values with two states: true or false
  - `float32` `float64` stores floating point numbers, with decimals

#### Declare a variable
  - `var name = "Hiệp"` (global)
  - `name := "Hiệp"` (function)
  - With the above two methods, the compiler can automatically infer the datatype for the variable
#### Naming conventions
  - Use camelCase (not snake_case): `userName (true)`
  - Use PascalCase for exported identifiers
      - If a variable/function/type starts with a capital letter, its exported (public)
    
            var UserName string // Exported - can be used outside package
    
      - If it starts with a lowercase letter, it's unexported (private to the package)
      
            var userName string // Unexported - only within package

### Arrays
#### Declare an Array
- With `var` keyword
        
        var my_number = [4]int{0,1,2,3} // length us defined
        
        or

        var mu_number = [...]int{0,1,2,3}  // length is infered

- With the `:=` sign

        my_number := [4]int{0,1,2,3} or my_number := [...]int{0,1,2,3}

### Slices
Slices are similar to arrays, but are more powerful and flexible

#### Syntax
        myslice := []int{}
#### Create a slice from array
        var myarray = [length]datatyped{}
        myslice := myarray[start:end]
- Slice `[start:end]` gets elements from index `start` to index `end - 1`
#### Create a slice with the make() function
        slice := make([]datatyped, len, cap)
#### Append elements to a slice
        slice = append(slice, element1, element2,...)
- Go automatically increases capacity according to the "double" rule when it needs to expand a slice.

### Conditions
#### Syntax
        if <ccondition> { // statements }

### Switch
#### Syntax
- Single case

        switch expression {
            case x:
                // code block
            ...
            default:
                // code block
        }

  - Multi-case

          switch expression {
              case x,y:
              // code block if expression is evaluated to x or y
              case v,w:
              // code block if expression is evaluated to v or w
              case z:
              ...
              default:
              // code block if expression is not found in any cases
          }

### Loops
#### Syntax
            for <init>; <evaluated-for-each>; <counter>{}
#### The range keyword
- The `range` keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

            for index, value := range array|slice|map {}

### Function
#### Syntax
- Create a function

        func funcName() {}
- Call a func
        
        funcName()
- Function return 

        func FunctionName(param1 type, param2 type){
            return output
        }

### Struct
#### Declare
        type struct_name struct {
            member1 datatype
            member2 datatype ...
        }

### Maps
- Maps are used to store data values in key:value pairs
- Each element in a map is a key:value pair
- The default value of a map is nil
- Maps hold references to an underlying hash table

#### Create maps (using var and :=)
        var a = map[keyType]ValueType{key1:value1, key2:value2,...}
        b := map[keyType]ValueType{key1:value1, key2:value2,...}
#### Create maps (using make() function)
        var a = make(map[KeyType]ValueType)
        b := make(map[KeyType]ValueType)
#### Update and Add Map Elements
        map_name[key] = value
#### Remove element
        delete(map_name, key)
#### Check for specific elements 
        val, ok := map_name
#### Iterate over maps
        for k, v := range a {
            fmt.Printf("%v : %v, ", k, v)
        }
#### 