#Corrected Code

this is the code that is given to correct.

package main
import (
"fmt"
"sync"
)
func main() {
var wg sync.WaitGroup
ch := make(chan func(), 10)
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg *sync.WaitGroup) {
   for {
        x := <-ch
             x()
        }
    }(&wg)
}
     for i := 0; i < 5; i++ {
        ch <- func() {
    fmt.Println("Work -", i)
    }
}
    wg.Wait()
}


this code is to create the goroutines that execute the function in the channel.
the major issue was there is no closing for the  channel.if we dont close the channel goroutines will wait indefinitely.
