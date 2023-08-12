package main

import "fmt"

type Calculator struct {
    value float64
}

func NewCalculator(initialValue float64) *Calculator {
    return &Calculator{value: initialValue}
}

func (c *Calculator) Add(n float64) *Calculator {
    c.value += n
    return c
}

func (c *Calculator) Multiply(n float64) *Calculator {
    c.value *= n
    return c
}

func (c *Calculator) GetValue() float64 {
    return c.value
}

func main() {
    calc := NewCalculator(10.0)
    result := calc.Add(5.0).Multiply(2.0).GetValue()
    fmt.Println("Result:", result) // Output: Result: 30
}
