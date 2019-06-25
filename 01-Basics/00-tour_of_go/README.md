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
    

* `*T`: "a pointer to T value type", pointer type.
    * zero value is `nil`.
    * eg. `var p *int`: p is a pointer to int value type.

* All values have or can have a pointer. `&` operator *generates* a pointer.
    * dereferencing or indirecting with `*`, note `*` is used to denote pointer type as well.
        ```go
        i := 42
        p = &i // & generates a pointer to 42
        fmt.Println(*p) // read variable i value through the pointer type `p`
        *p = 21 // set variable i value through the pointer `p`
        ```
   
* `struct`, which is like data class, consists of fields, which is like a property of data class.
    * `struct` is a *type* as well as *class-like* data structure.
        ```go
        type Vertex struct { // Note::define data structure with `type` keyword
  	       X int  // Q::Capital for public?
           Y int
        }
        vertex := Vertex{1, 2}
        fmt.Println(vertex.X)
        v.X = 4
        ```
        
    * field values are optional, can be created using named fields
        ```go
        vertex2 := Vertex{} // vertex2.X == 0, vertex2.Y == 0
        vertex3 := Vertex{X: 1} // vertex3.X == 1, vertex3.Y == 0
        ```
    * As with any structure defined with `type` keyword, it's an allocated literal value, `&` operator will generate a pointer.
        ```go
        // Where as literal value like 42 cannot have pointer generated;
        i := &42 // CompileError:  cannot take the address of 42
        p := &Vertex{1, 2} // has pointer type `*Vertex`
        ```
        
    * When accessing field through struct pointer implicit dereferencing happens.
        ```go
        v := Vertex{1, 2}
        p := &v
        (*p).X = 4 // normally we would need to explicitly dereference.
        p.X = 4 // compiler supports implicit dereferencing for dot accessors.
        ```

* Array: a built in type.
    * type `[n]T` is an array of n values of type `T`. Cannot resize.
        ```go
        var a [10]int // empty array of fixed length
        a[0] = 1
        
        strings := [4]int{"a", "b", "c", "d"}
        fmt.Println(strings[0])
        ```
* Slice: a built in dynamic references to arrays.

    * type `[]T` is an slice of type `T`. Can resize.

    * It's a reference to literal value(array) (i.e. reference value). (Note::pointer value, literal value, reference value)
        ```go
        nums := [3]int{1, 2, 3}
        xs := nums[0:2] // [1, 2]
        ys := nums[1:3] // [2, 3]
        xs[1]  = 20
        fmt.Println(ys) // [20, 3]
        ```
        
    * Slice has both a `length` and a `capacity`.
        * `length`: number of elements it contains. `len(someSlice)`
        * `capacity`: number of elements in the underlying array. `cap(someSlice)`
        * Note::both of above are dynamic measure. eg) `length` is not fixed like an array.
    
    * Slice always has an underlying array. And don't store data.
        * Slice can be created from array, with `a[begInclusive:endExclusive]`;
            ```go
            strs := [4]string{"a", "b", "c", "d"}
            sliced := strs[0:2] // ["a", "b"]
            x := str[:2] // same as str[0:2]
            y := str[1:] // same as str[1:str.length]
            z := str[:]  // same as str[0: str.length]
            ```
            
            * When created from array a slice cannot go over the `capacity`.
            
        * Slice can be created from type. (Q:: what do you call this? constructor?)
            ```go
            lists := []bool{true, true, false}
            dataList := []struct { 
  	          i int 
              b bool 
            }{
              {2, true},
              {3, false}
            }
            ```
         
        * Slice from make (TODO)
    
    

