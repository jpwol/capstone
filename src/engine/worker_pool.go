package engine

import (
	// "fmt"
	// "sync"
	// "slices"
)

// NOTE: moved to worker.go
//
// type Worker struct {
//      id int
//      work []int
// }


//func InitializePool(width, height, maxIter, workers int) []int {
//        var wg sync.WaitGroup
//        out := make([]int, width*height)
//        jobs := make(chan int, height)
//       for w:=0; w<workers; w++{
//                wg.Add(1)
//
//                go func(){
//
//                        defer wg.Done()
//                        for i := range jobs{
//
//                                Row(out, i, width, height, maxIter)
//                        }
//                }()
//        }
//       for py := 0; py<height; py++{
//                jobs <- py
//        }
//        close(jobs)
//        wg.Wait()
//        return(out)
//        panic("Bad Pool, No Return")
//}
