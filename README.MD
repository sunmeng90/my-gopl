# Go programming language


## Basic Syntax
```go
s := "
```
**short variable declaration may be used only  within a function, not for package-level variables.**  
**`new()` function can be used to create a unamed variable of type T and intializes the value to zero value of T and return the variable address**  
**variable escape: go compiler may choose to allocate variables on heap or on the stack no matter how the variable is declared(by var or by new function)**
```go
    var global *int
    func f(){
        var x int
        x = 1
        global = &x //x will be allocated on heap, because it is still reachable by global after function f returned, so we can say x escapes from function f
    }
```

## Data types

a `rune` value in go is like a `char` value in other language: 


**Return a pointer from function**

If we return a local var pointer the go compiler will allocate the variable on the heap

```go
package main

import (  
    "fmt"
)

func hello() *int {  
    i := 5
    return &i
}
func main() {  
    d := hello()
    fmt.Println("Value of d", *d)
}
```
In line no. 9 of the program above, we return the address of the local variable i from the function hello. The behavior of this code is undefined in programming languages such as C and C++ as the variable  i goes out of scope once the function hello returns. But in the case of Go, the compiler does a escape analysis and allocates i on the heap as the address escapes the local scope. Hence this program will work and it will print, 
```
Value of d 5  
```

**Do not pass a pointer to an array as a argument to a function. Use slice instead.**
https://golangbot.com/pointers/

**we can assign value to named return value directly in function body**
```go
func (r *stringReader) Read(p []byte) (n int, err error) {
    n = copy(p, r.s)//n is not a new declared var, but return var that are assigned directly
    r.s = r.s[n:]
    if len(r.s) == 0 {
        err = io.EOF //assign to err
    }
    return
}
```

# Struct and interface

* struct are mainly for data structure definition
* interface meant to behavior
**you can directly define methods for struct by bind function to struct type
but it's recommend to use interface for the sake of decouple???** 
<https://www.cnblogs.com/zhangdewang/p/10612089.html>
