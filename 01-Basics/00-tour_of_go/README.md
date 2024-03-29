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
         
        * Slice from make. It's like Array(n).
            ```go
            a := make([]int, 5) // Slice of length and cap 5 with underlying array that has 5 zeros.
            b := make([]int, 0, 5) // Slice of length 0 and capacity 5
            ```
    
    * `append` is like push.
    	* If baking array is too small, the returned slice of `append` will point to the newly allocated array.
            ```go
            var s []int
            s = append(s, 1, 2, 3) // s == [1, 2, 3]
            ```

* Map maps keys to values.
    * As with any type, variable declaration without initialization get initialized with zero value of type. Map's zero value is `nil`. `nil` map cannot be used at all as no keys and values can be added to `nil` map. Hence map is often used with `make`.
        ```go
        type Vertex struct {
          Lat, Long float64
        }
        m := make(map[string]Vertex)
        m["SomeKey"] = Vertex{2.2, 3.3}
        ```
	
    * Map literals
        ```go
        m := map[string]Vertex{
          "One": Vertex{2.3, 4.5},
          "Two": Vertex{1.1, 2.2},
        }
        // If value of a map is a type name, in this case `Vertex`, type can be omitted inside the map literal
        m := map[string]Vertex{
          "One": {2.3, 4.5},
          "Two": {1.1, 2.2},
        }
        ```
    
    * Normally map is used by mutating.
        ```go
        m["One"] = Vertex{1, 2}
        val := m["One"]
        delete(m, "One")
        val, exists = m["One"] // exist is a boolean
        ```

* Function values
    * Higher order functions.
        ```go
        func outer(someFunc func(float64, float64) float64) float64 {
          return someFunc(3, 4)
        }
        ```
	
    * Closures. Interestingly if a inner function can access and assign to the outer varaible, gophers say that an inner function is "bound" to the variables.
        ```go
        func outer() func(int) int {
          sum := 0
          return func(x int) int {  // this inner function is bound to sum variable.
            sum += x
            return sum
          }
        }	
	    ```

* `range` form of the `for` loop iterages over a *slice* or *map*
    * `range` for slice returns index, and a copy of the element at that index
        ```go
        list := []int{1, 2, 3}
        for i, v := range list {
          fmt.Println("v is a copy of list[i]")
        }
        ```
	
    * Skip index or value by assigning to _.
        ```go
        for i, _ := range list
        for _, v := range list
        for i := range list // if only first value (index) is needed second variable can be omitted.
        ```

* No class in Go. Type is used like a class. Function with a receiver is a method
    
    * Method is a function with a *receiver* argument.
        ```go
        //   (receiver)
        func (v Vertex) Abs() float64 {
          return math.Sqrt(v.X * v.X)
        }
        // non-struct type can be used as a receiver as well
        type MyFloat float64
        func (val MyFloat) Abs() float64 {
          return float64(val)
        }
        ```
    
    * Only types in the same package can be used as a receiver.
        
        * Even built in type such as `int` cannot be directly used as a receiver as it is not in the same package.
            ```go
            func (num int) Abs() float64 { // Error as int type is not defined in the same package.
              return float64(num)
            }
            ```
    
    * Receiver can be a pointer type
    
    	* Methods with pointer receivers can modify the value to which the receiver points. This also means if a method as a value receivers, it cannot modify and will be copied by value just like function args are copied by value.
            ```go
            // Here v is passed by a reference. (passed by address value)
            func (v *Vertex) Scale(f float64) {
              v.X = v.X * f
            }
            // Here v is passed by a value. (similar to cloned Object which is not connected to original Object at all)
            func (v Vertex) Scale2(f float64) {
              v.X = v.X * f
            }
            v.Scale(10) // results in original v mutating.
            v.Scale2(10) // original v is not affected at all.
            ```
	
* Different behaviors of pointer indirection(dereferencing) between a receiver and function argument.
	
	* Function parameter of pointer type *must* be given a pointer (an address value)
		```go
		vertex := Vertex{3, 4}
		func SomeFunc(v *Vertex, f float64) {
		  v.X = v.X * f
		}
		val := SomeFunc(vertex) // Compile Error.
		val := SomeFunc(&vertex)
		
		```
	
	* But method with pointer receiver can take both a value and a pointer.
		```go
		vertex := Vertex{1,2}
		func (v *Vertex) Scale(f float64) {
		  v.X = v.X * f
		}
		vertex.Scale(5) // OK. For convenience compiler converts it to (&vertex).Scale(5)
		pVertex := &vertex
		pVertex.Scale(5) // OK
		```
	
	* Reverse is true, methods with value receivers can take both a value and a pointer. Note since it is passed by value a value is copied hence still no mutation is done on the dereferenced address.
		```go
		func (v Vertex) Scale(f float64) {
		  v.X = v.X * f
		}
		pVertex := &Vertex{1, 2}
		pVertex.Scale(5) // Ok. For convenience compiler converts it to (*pVertex).Scale(5)
		vertex := *pVertex
		vertex.Scale(5) // OK.
		```
	
	* In general, all methods have either all pointer receiver or all value receiver. And not a mixture.
		

* Interface type has a set of method signatures (method types). Implicit interface rather than explicit interface (no `implements` keyword)
    * When defining methods in typed language we usually make an interface and `implement` it.
    In Go, we don't implement interface, but if we want to assign a type to a variable of specific interface, a Type must have all methods defined in an interface.
    In this regards, interface is a explicit contract that focuses on a behavior of a variable of specific type.
        ```go
        type Abser interface {
          Abs() float64
        }
        type MyFloat float64

        func (f MyFloat) Abs() float64 {
          return float64(f)
        }

        var a Abser
        a = MyFloat(12) // MyFloat has an Abs method so it is considered to implements Abser interface hence is assignable to `a` variable which has Abser interface type here.
        ```
    * For interface methods defined on pointer receiver, in order to assign to a variable of interface type, a value has to be a pointer value.
        ```go
        type Abser interface {
          Abs() float64
        }
        func (v *Vertex) Abs() float64 {
            return float64(v.X)
        }
        var a Abser
        a = Vertex{1, 1} // Error. As no Abs method is defined in Vertex value type.
        a = &Vertex{1, 1} // OK. as Abs is defined in *Vertex pointer type.
        ```
    
    * Interface value has a tuple of a value and a concrete type. Internally something like `(3, int)` where `3` is value `int` is type
        
        * If the value of a internal tuple is `nil`, method is called with `nil` receiver. So if a `nil` handling error exists in method there will be no nil pointer error.
            ```go
            type I interface {
                SomeMethod() float64
            }
            func (v *Vertex) SomeMethod() float64 {
              if v == nil {
                // Handling nil pointer
              }
              return float64(v.X)
            }
            var i I
            var v *Vertex // nil pointer Vertex
            i = v // nil
            i.SomeMethod() // even though i is nil SomeMethod still works, no null point error as long as there is a null receiver handler in a method.
            ```
            Hence, an interface value that holds a nil concrete value (a value in internal tuple) is itself non-nil
            
        * `nil` interface value, unlike above, holds neither concrete value nor concrete type. This leads to nil pointer error.
        
        * type `any` = `interface{}`
            ```go
            var i interface{}
            i = 42
            i = "Hello"
            ```
            
* Type Assertion, `i.(T)`, to convert or check any `interface{}`
    * second return value is used to check whether type assertion is successful or not
        ```go
        var i interface{} = Vertex{1, 2}
        t, ok := i.(Vertex) // t == Vertex{1, 2}, ok == true

        i = "some string"
        t, ok = i.(Vertex) // t == "", a zero value of string, ok == false
        ```
    
    * if second variable (`ok`, above) is not provided, and the interface does not hold a asserted type `(Vertex)` of `i.(Vertex)` above,
    it will throw error (cause panic).
        ```go
        var i interface{}
        t := i.(Vertex) // causes panic as `i` does not hold asserted type, and second variable is not defined to catch `error`
        ```

* `type` keyword can be used in switch statement with same syntax as Type Assertion.

    * Each switch case is compared against the `type`.
        ```go
        switch v := i.(type) { // note `type` is a keyword, and no second variable needed
            case Vertex:
                // v is of type Vertex
            case float64:
                // v is of type float64
            default:
                // v is of type i
        }
        ```

* `Stringer` built-in interface, `type Stringer interface { String() string }` is like `toString()`

    * It is used as a method that is called when any type value has to be expressed as a string value.
        ```go
        type Person struct {
            Name string
            Age int
        }
        func (p *Person) String() string {
            return string(p.Age)
        }
        person := &Person{"Dave", 33}
        fmt.Println(person) // `fmt.Println` uses String() method to convert `person` to a string value
        // prints "33"
        ```
        
* `error` built-in interface, `type error interface { Error() string}`

    * some packages look for the `error` interface, for example `fmt` uses `Error()` method when passed in type is an `error` type
        ```go
        type MyError struct {
            When time.Time
            What string
        }

        func (e *MyError) Error() string { // Since MyError implements Error method it is now compatible with `error` interface
            return fmt.Println(e.What)
        }
        var err error
        err = &MyError{time.Now(), "it's an error"}

        fmt.Println(err) // since err is error type, Println will call Error() method to get an error message.
        ```

    
* *Channel*, `chan`, is a medium where data is sent and received between go routines.

    * Channels can be explained with analogy of RxJS even though it's not related at all.
    
        * In RxJS values are emitted one by one. eg)
            ```typescript
            // from function sends each element in an array, subscriber receives each element one by one
            from([1, 2, 3]).subscribe(x => console.log(x)) // logs 1, then 2, then 3
            ```
          In Go, channels send and receive manually one by one. eg)
            ```go
            ch := make(chan int, 3) // need to specify how many will be sent
            // have to explicitly send each individually
            ch <- 1 // <- means send to channel `ch`
            ch <- 2
            ch <- 3
            // have to explicitly receive each item
            var x int
            x = <-ch
            fmt.Println(x) // logs 1
            x = <-ch
            fmt.Println(x) // logs 2
            x = <-ch
            fmt.Println(x) // logs 3
            ```
          Surely, this is not applicable with a large number of sends and receives. So there is a better way using `range` keyword with `channel` type data.
          Above can be written like below.
            ```go
            ch := make(chan int, 3)
            for x := 1; x <= 3; x++ {
  	           ch <- x
            }
            close(ch) // need to manually close. It's like complete event in RxJS.
  	        // Loops through each sent integer as though it's an array. 
            //   Similar to RxJS's `map` which has same interface for an Array and Observable, 
            //   `range` has the same interface for a `slice`, and a `channel`.
            for x := range ch { // ch must be `close`d for `for` loop to terminate, otherwise it will loop infinitely
  	           fmt.Println(x)
            }
            ```
          Event though above example is in Sync setting, same logic applies to Async setting with `go`routines, just like RxJS treats Sync and Async the same.
        
         
* *Goroutine*, `go`, is a lightweight thread. Can spawn tens of thousands without a problem.

    * `go f(x, y, z)`, function `f`, and arguments `x`, `y`, `z` is evaluated in current thread (main thread) but function body (logic) is called in new independent thread.
    
* *Select* is like a async compatible switch. 

    * It can be used to block(wait) for one of the channel specified in `case` statement.
        ```go
        c := make(chan int)
        quit := make(chan int)
        go func() {
          for i := 0; i < 10; i++ {
              fmt.Println(<-c)
          }
          quit <- 0
        }()
        x := 0
        y := 1
        for {
  	      select {
          case c <- x:  // send to channel c. this can be run when goroutine listens to c, i.e. <-c exists in declaration.
    	    fmt.Println(x)
          case <-quit: // when receives from quit channel. can receive value by `val := <-quit`.
    	    fmt.Println("quit")
            return
          }
        }
        ```
        Note that `select` statement blocks(waits) until at least one of the `case` can be run.
        
    * Or it can be non-blocking(passes-through) with `default` case.
        ```go
        tick := time.Tick(100 * time.Millisecond) // tick is a channel that sends every 100 ms, like time(100, 100) in rxjs
        boom := time.After(500 * time.Millisecond) // boom is a channel
        for {
            select {
            case <-tick:
                fmt.Println("tick.")
            case <-boom:
                fmt.Println("BOOM!")
                return
            default:
                fmt.Println("...")
                time.Sleep(50 * time.Millisecond)
            }
        }
        ```

    *  Note for both cases, compiler evaluates whether channel has both send and receive defined at declaration time.

* *Mutex*, mutual exclusion, is another way for managing shared state. `sync.Mutex` allows only one goroutine to change state at a time.

    * Block of code to run in mutual exclusion (sync) by wrapping code in `Lock` and `Unlock` method of `sync.Mutex`, Q::internally this must be a que of goroutines.
        ```go
        type SafeCounter struct {
            v int
            mux sync.Mutex
        }

        func (c *SafeCounter) Inc() {
            c.mux.Lock()
            c.v++
            c.mux.Unlock()
        }

        c := SafeCounter{0}
        for i := 0; i < 1000; i++ {
            go c.Inc()
        }
        time.Sleep(time.Second) // sync code run before goroutine so need a way for it to run on next tick (naively sleep here)
        fmt.Println(c.v)
        ```