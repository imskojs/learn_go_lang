* All code files in a same folder must use the same package name.
  * Common practice to name the package after the folder.
  
* Package defines a unit of compiled code.

* To run side effect by importing go uses `_`. 
If `_` keyword is used before package name it will call `init` method of importing package.
If `_` not used Go compiler will not import unused packages
eg)
```go
package main
import (
	_ "github.com/matchers"
)
```

* Also *ALL* `init` functions of various imported and used packages will be called before main function.

* The compiler will always look for the packages from `GOROOT` and `GOPATH` environment variables

* Variable located outside the scope of any function => package-level variable

* Variable Declaration example
```go
var matchers = make(map[string]Matcher)
```
Reads, map of `Matcher` type values with a key of `string` type

* Package scoped. Even if variables or functions (identifiers) are in different files, as long as it is in the same package
it can be accessed without import in different files.

* Exported identifiers start with a capital letter. Need to import if in different package.

* Unexported identifiers start with a lowercase letter, and can't be directly accessed by code in other packages.

* Unexported identifiers is like a private variable which can be used in other packages through closure.

* Both exported and unexported identifiers are available in different files of the same package.

* `make` is a way to declare and define types to a variable. Like `const value: Object` in TypesScript.
It merely allocates type and not the value or reference. Only takes `map`, `slice`, `channel` types as first arg.

* All variables are initialized to their zero value.
  * number -> 0
  * string -> ""
  * Boolean -> false
  * Pointers -> nil
  
* Function can have multiple return values.
  * Hence it is common to return Ok value and Err value.

* `:=` is used to both declare and initialize variable at the same time.
  * `type` of a variable initialized this way is determined by function's return type.

* Use `var` when declaring variables that will be initialized to their zero value.

* Use `:=` when making a function call.

* Channels implement a queue of typed values that are used to communicate data between goroutines.

* Once the `main` function returns, all goroutines still running will also be terminated.

* Looking up key in a map returns two values. 1. value, 2. boolean indicating existence.

* When a key doesn't exist the map will return the zero value of the type of value.

* `go` is used with function call to run a function concurrently. Hence `anonymous function` is
often used to make a piece of code concurrent. eg) `go func(param1 Sometype, param2 SomeTYpe){}(x, y)`

* Pointer Type is expressed with `*`. eg) `*Feed` read as, type that is an address of a value of type Feed.

* The value of a pointer variable is the address to the memory. Passing pointer variables between
functions is still considered a pass by value, where value is the address to the memory.

* `for channel := range someChannel`, looping only happens when result is written to the channel.
This implies loop is not based on index but rather when channel is written. Loop finishes when channel closes.

* Go compiler can deduce the type from the value on the rhs of the assignment operator.

* `defer` is used to schedule a function to be called right after a function returns or terminates unexpectedly.
  * Helps readability and reduce bug.
  
* Another way to declare variable without using make;
```go
var feeds []*Feed
```
Above declares nil slice that contains pointers to Feed type values.

* `*` is used for pointer to Type. `&` is used to access actual address.

* 뭔가... File 자체가 Class 인 느낌이지...

* `type WHATEVER struct` is used to define shape.

* `type WHATEVER interface` is used to define behavior. By the methods that are declared within the interface type.

* suffix `er` for an interface that has only one method. eg)
```go
type Matcher interface {
	Search()
}
```

* The name of the interface should relate to its general behavior.

* Compiler will reference or dereference the value depending on whether we are using a value or pointer to make a method call.

* Receiver is a way to declare a method for a type. eg)

```go
type defaultMatcher struct {}
func (m defaultMatcher) Search(param1 SomeType) ([]*Result, err) {
	return nil, nil
}
```

above is similar to

```typescript
class defaultMatcher {
    public Search(param1: SomeType): null {
        return null
    }
}
```

* Methods declared with pointer receivers can only be called by interface type values that contain pointers.

* Methods declared with value receivers can be called by interface type values that contain both values and pointers.

* If function argument specifies the interface type, values as well as pointers that implement the interface will be accepted.

* SHOULD GO THOUGH TOUR OF GO FIRST

































