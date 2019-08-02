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

### 3.6.1 Constant generator `iota` (enumeration)

`iota` will start with a 0 or with a expression and repeat for each enum entry

```go
// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota //start with 0 and increment by 1 for each entry
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


type Flags uint
// start with 1 << 0 and shift left by one bit for each entry
const (
	FlagUp           Flags = 1 << iota // interface is up
	FlagBroadcast                      // interface supports broadcast access capability
	FlagLoopback                       // interface is a loopback interface
	FlagPointToPoint                   // interface belongs to a point-to-point link
	FlagMulticast                      // interface supports multicast access capability
)
```

# 5 Functions

In go, all variables are passed by value 

```go
// parameter are parsed by value
func test_string(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n", s, &s)
}

func main() {
	s := "a"
	fmt.Printf("outer: %v, %v\n", s, &s)
	test_string(s)
	fmt.Printf("outer: %v, %v\n", s, &s)
}
// output:
// outer: a, 0xc0000101e0
// inner: a, 0xc000010200
// inner: b, 0xc000010200
// outer: a, 0xc0000101e0
```

```go
func test_map(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}

func main() {
    // 运算符:=是make()的语法糖
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	} // A pointer to the header for a Go map.

	fmt.Printf("outer: %v, %p\n", m, m)
	test_map(m)
	fmt.Printf("outer: %v, %p\n", m, m)
}
// output:
// outer: map[a:1 b:2 c:3], 0xc0000762d0
// inner: map[a:1 b:2 c:3], 0xc0000762d0
// inner: map[a:11 b:2 c:3], 0xc0000762d0
// outer: map[a:11 b:2 c:3], 0xc0000762d0
```

# Struct and interface

* struct are mainly for data structure definition
* interface meant to behavior
**you can directly define methods for struct by bind function to struct type
but it's recommend to use interface for the sake of decouple???** 
<https://www.cnblogs.com/zhangdewang/p/10612089.html>
