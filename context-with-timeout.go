package main

import (
	"context"
	"fmt"
	"time"
)

func FunctionB(ctx context.Context) bool {
	// Simulate some condition
	select {
	case <-ctx.Done():
		return false
	case <-time.After(6 * time.Second):
		return true
	}
}

func FunctionA() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wait for FunctionB to return true using the context
	fmt.Println("Waiting for FunctionB...")
	if FunctionB(ctx) {
		fmt.Println("FunctionB returned true")
	} else {
		fmt.Println("FunctionB did not return true within the timeout")
	}
}

func main() {
	FunctionA()
}
