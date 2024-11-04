package main

import (
	"fmt"
	"sync"

	"github.com/DarRo9/test-task/internal/pool"
)

func main() {
	wp := workerpool.NewPool(1)
	wp.RemoveWorker()              
	fmt.Println(wp.RemoveWorker()) 

	fmt.Println(wp.AddWork("test")) 

	wp.AddWorker() 
	wp.AddWorker() 
	wp.AddWorker()
	wp.AddWorker()  

	strings := []string{"Меркурий",
    "Венера",
    "Земля",
    "Марс",
    "Юпитер",
    "Сатурн",
    "Уран",
    "Нептун",
	}

	wg := sync.WaitGroup{}
	for i, s := range strings {
		if i == 1 {
			wp.AddWorker() 
		} else if i == 5 {
			wp.RemoveWorker() 
		}

		wg.Add(1)
		go func() {
			wp.AddWork(s)
			wg.Done()
		}()
	}

	wp.RemoveWorker() 

	wg.Wait()

	wp.Stop()
}
