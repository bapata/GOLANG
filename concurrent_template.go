package main
import (
  "fmt"
  "time"
  "sync"
)

func long_running() {
  time.Sleep(time.Second*2)
  fmt.Println("Done long_running")
}

func main() {
  fmt.Println("Start main")
  wg := &sync.WaitGroup{}
  for ii:=0;ii<10;ii++ {
    wg.Add(1)
    go func() {
      long_running()
      wg.Done()
    }()
  }
  wg.Wait()
}
