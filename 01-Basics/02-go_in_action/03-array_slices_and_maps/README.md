* Array;
    ```go
    // For array, `...` will identify the length of the array based on passed in number.
    array := [...]int{1, 2, 3, 4}

    // Can initialize part of arrays
    array = [5]int{1: 10, 2: 20}
  
    // array of points, change the memory being shared
    array = [5]*int{0: new(int), 1: new(int)}
    *array[0] = 10
    ```
    
* Slice:
    ```go
    // nil slice has a pointer value of nil, i.e. no underlying array.
    var slice []int
  
    // empty slice has a pointer to zero-element array that allocates no storage
    slice = make([]int, 0)
    slice = []int{}
  
    // slice of 100 length and capacity.
    slice = []string{99: ""}
    
    // changing a value in slice is the same as changing value in underlying array
    slice[1] = 10 // is the same as;
    *slice[1] = 10 // which has the same effect of changing value in array.
    underlyingArray[1] = 10
  
    // If underlying array has a enough capacity, slice's `append` changes underlying array's value
    slice = []int{0, 1, 2, 3}
    newSlice := slice[:2]
    newSlice = append(newSlice, 20) // changes index 2's value which is shared by `slice`
    slice[2] == 20 // slice[2] == 20, and newSlice[2] == 20, and underlyingArray[2] == 20
  
    // If underlying array has not enough capacity new array is created that is unrelated to base array.
    slice = []int{0, 1}
    newSlice = append(slice, 2) // newSlice has an underlying array that is not related to original slice.
    // newSlice has an underlying array that is twice the capacity of slice's underlying array.
    // Array capacity doubles when elements < 1000. Increase by 25% when elements > 1000.
  
    // When slice is made it has a default capacity that is the same as the array.
    // Capacity of slice determines whether a new underlying array is created or not.
    //  So when creating new slice, Go has a way of setting capacity, to control how underlying array is manipulated.
    slice = []int{0, 1, 2}
    slice = slice[0:1:2] // slice length 1, capacity 2.
    
    ```
    
    










