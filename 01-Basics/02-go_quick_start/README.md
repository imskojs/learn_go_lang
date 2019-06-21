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
  


