* The package name is the same as the last element of the import path.
    * package `math/rand` has files that begin with `package rand`

* Factored import statement
    * `import ("fmt" "math")`
    * same as `import "fmt"` `import "math"`

* An identifier is exported if it begins with a capital letter

* Multiple results
    ```go
    func swap(x, y string) (string, string) {
        return x, y
    }
    ```

* If an initializer is present, the type can be omitted.
    * `var i = 1 // int type`

* `:=`, short assignment statement, can be used inside a function.

* Outside a function, every statement begins with a keyword.
    * Variable declared without initializer are given zero value of the type.
    * variable declaration can be factored
        ```go
        var (
            ToBe bool = false
            MaxInt uint64 = 1<<64 - 1
        )
        ```
    
* Type conversion by `T(v)`
    * `var i = 42`, `var f float64 = float64(i)`
    
* Constants can only be character, string, boolean, numeric values.
    * cannot be declared using `:=`, short assignment statement.

* `for` loop consists of optional init statement, optional condition expression, and optional post statement
    * `for i := 0; i < 10; i++ {  }` 
    * `for ; i < 10; { }` simplifies to `for i < 10 { }` which is the same as `while`.
    * `for { }` infinite loop
    
    
* `if` statement
    * `if x < 0 { }`
    * `if v := math.Pow(1, 2); v < 10 { }` can start with a short assignment statement
    
* `switch` statement. No fall through in Go unlike JS.
    * can start with short assignment statement.
      ```go
      switch os := runtime.GOOS; os {
      case "one":
	      fmt.Println("case one")
      case "two":
	      fmt.Println("case two")      
      default:
	      fmt.Println("default case")
      }

      ```
    * can have expression in each case
        ```go
        switch time.Saturday {
        case today + 0:
  	       fmt.Println("Today.")
        case f(1):
  	       fmt.Println("Whatever")
        }
        default:
  	       fmt.Println("default case")
        ```
     * Shorter if then else statement with `switch true`
        ```go
        switch { // same as switch true
        case t.Hour() < 12: // if
 	        fmt.Println("")
        case t.Hour() < 15: // else if
 	        fmt.Println("")
        default:           // else
            fmt.Println("")
        }
        ```
        
* `defer` runs function after return.
    * arguments are evaluated immediately but function call is deferred.
    * Stacking or calling multiple defer has `last-in-first-out` order.
   
