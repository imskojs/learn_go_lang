* To run all goroutines in sequence use `atomic`, and `sync`.

    * `atomic` functions to synchronize access to *integers* and *pointers*
        ```go
        var (
            counter int64
            wg sync.WaitGroup
        )
        func incCounter(){
            defer wg.Done()
            for i := 0; i < 10; i++ {
                // only one goroutine will add 1 to counter at a time.
                // Automatically synchronized against the variable that's referenced.
                atomic.AddInt64(&counter, 1)
                // give a logical process to next goroutine.
                // Q:: i is saved and next time this goroutine comes back into que it runs from where it left off?
                runtime.Gosched()
            }
        }

        wg.Add(2)
        go incCounter()
        go incCounter()
        wg.Wait()

        fmt.Println(counter)
        ```
    * Safe read with `atomic.LoadInt64`, and write with `atomic.StoreInt64`.
        * If `LoadInt64` is called at the same time as `StoreInt64` to a same variable, it will be synchronized.
        
    * `sync.Mutex` is used to create a code block which only one goroutine at a time can execute.
        ```go
        var mutex sync.Mutex
        var wg sync.WaitGroup
        func incCounter() {
            for i := 0; i < 10; i++ {
                // Allow only one goroutine to run at a time
                mutex.Lock()
                // DO SOME WORK
                // Q::Gosched does not have any effect?
                runtime.Gosched()
                mutex.Unlock()
            }
        }
        ```
        Can think of code between `mutex.Lock()` and `mutex.Unlock()` to be a critical section that only one goroutine can enter at a time.
        
* Atomic functions and mutexes are old school. `Channels` are a better tool.

    * built-in types, named types, struct types and reference types can be shared through a channel
    
    * Channel has to be made with `make` built-in function `channel := make(chan int)`
    
    * There are buffered and unbuffered channel
        ```go
        unBufferedChannel := make(chan int)
        bufferedChannel := make(chan int, 10)
        ```
    
* Unbuffered Channel
    * No capacity to hold any value before it's received.
    * requires both a sending and receiving goroutine to be ready at the same instance.
    * If two go routines aren't ready at the same instance, the channels makes send or receive operations wait.
    * because of this wait behavior, synchronization is inherent in the interaction.

* Buffered Channel
    * Can hold more than one values before they're received.
    * Receive blocks when there are no value in the channel to receive.
    * Send blocks when there is no available buffer to place the value being sent.
    * Each goroutine is receiving work distributed from the channel. i.e. if chan <- 1, chan <-2 is called, depending on
    the number of goroutines it will 1, and 2 will be distributed evenly, no duplicate receives.
    * Q::But empty receive is duplicated to all goroutines? so that it knows when to return/exit?
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    