package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint, resume chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d: Checkpoint reached\n", id)
	checkpoint <- struct{}{}
	<-resume
	fmt.Printf("Worker %d: Resuming\n", id)
}

func main() {
	var numWorkers int
	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&numWorkers)
	checkpoint := make(chan struct{})
	resume := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}

	for i := 1; i <= numWorkers; i++ {
		<-checkpoint
	}

	fmt.Println("All workers reached the checkpoint")

	close(resume)
	wg.Wait()

	fmt.Println("All workers completed their work")
}

// Output
// Enter the number of workers: 5
// Worker 5: Starting
// Worker 2: Starting
// Worker 3: Starting
// Worker 4: Starting
// Worker 1: Starting
// Worker 1: Checkpoint reached
// Worker 2: Checkpoint reached
// Worker 3: Checkpoint reached
// Worker 4: Checkpoint reached
// Worker 5: Checkpoint reached
// All workers reached the checkpoint
// Worker 1: Resuming
// Worker 4: Resuming
// Worker 5: Resuming
// Worker 2: Resuming
// Worker 3: Resuming
// All workers completed their work

package main
import(
    "fmt"
    "sync"
    "time"
)
func main(){
    var num int
    var wg sync.WaitGroup
    start:=make(chan struct{})
    fmt.Print("Enter the number of workers: ")
	fmt.Scan(&num)
    for i:=1;i<=num;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            <-start
            fmt.Printf("Worker %d is starting to work\n",id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d has reached checkpoint\n",id)
            time.Sleep(time.Second)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d resuming the work\n",id)
            time.Sleep(time.Second)
        }(i)
    }
    fmt.Println("Starting workers.....")
    close(start)
    wg.Wait()
    fmt.Println("All workers completed their work")
}
