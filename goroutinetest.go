package main

const(
    NumGo = 100000
    )

func main() {
    var chans [NumGo]chan int
    for i := 0; i<NumGo; i++ {
        c := get_empty_goroutine()
        chans[i] = c
    }
    for i := 0; i<NumGo; i++ {
        <-chans[i]
    }
}

func get_empty_goroutine() chan int {
    c := make(chan int)
    go func() { c <- 1 }()
    return c
}

func some() {
    
}


// variables 