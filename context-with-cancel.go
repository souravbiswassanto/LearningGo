package main

import (
	"context"
	"fmt"
	"time"
)

func PerformTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task was canceled")
		return
	default:
		// Simulate some task
		//time.Sleep(2 * time.Second)
		fmt.Println("Task completed")
		go ant(ctx)
		go PerformTask(ctx)
	}
}
func ant(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task was canceled ant")
		return
	default:
		// Simulate some task
		//time.Sleep(2 * time.Second)
		fmt.Println("Task completed in ant")
		go ant(ctx)
		go PerformTask(ctx)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a task
	go PerformTask(ctx)

	// Cancel the task after a delay
	//time.Sleep(3 * time.Second)
	time.Sleep(60 * time.Microsecond)
	cancel()
	fmt.Println("Main program finished")
	time.Sleep(1 * time.Second)

}
/* 
Task completed
Task completed
Task completed
Task completed
Task completed
Main program finished
Task was canceled ant
Task was canceled ant
Task completed in ant
Task was canceled
Task was canceled
Task completed in ant
Task was canceled
Task was canceled ant
Task was canceled ant
Task was canceled ant

*/
