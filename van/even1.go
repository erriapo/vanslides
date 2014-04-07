func even(done <-chan struct{}) <- chan int {
    out := make(chan int)
    go func() {
       defer close(out)
       var i int
       for {
           select {
           case <- done: return      // HL
           case out <-i: i = next(i)
           }
       }
    }()
    return out
}
