# Introducing Go

* `Goroutines` are like threads
  * are functions that run concurrently.
  * Go runtime automatically schedule the execution of goroutines against a set of configured logical processors.
  * Each logical processor is bound to a single OS thread.
  * goroutines have minimal overhead, can spawn thousands of goroutines.

* `Channels` are data structure that let you send typed messages between goroutines sync.
  * Enforce the pattern that only one goroutine should modify the data at any time.
  * Eliminates a need for locks and synchronization mechanisms.

* `Types` are composed of smaller types.

* Go `interfaces` are used to express the behavior of a type.
  * Go interface typically represents just a single action.
  
* Install from https://golang.org/doc/install

* Compile with `go build`, and run ./hello_world
  
