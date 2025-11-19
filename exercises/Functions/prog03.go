package main

import "fmt"

/*
--- PROBLEM: THE RUNNING AVERAGE GENERATOR ---

Objective:
Create a factory function called RunningAverage that returns a closure. 
This closure must maintain the running sum and count of all numbers it has received.

Requirements:
1. Factory Function (RunningAverage):
   - Input: None.
   - Output: Returns a single closure function with the signature: func(newValue float64) float64

2. The Returned Closure:
   - Must capture (remember) two variables from the outer scope:
     a. sum (float64, initialized to 0.0).
     b. count (int, initialized to 0).
   - Logic: When the closure is called with a newValue:
     i. Update the captured 'sum' and 'count' variables.
     ii. Calculate the average (sum / count).
     iii. Return the calculated average as a float64.
*/
// Define your RunningAverage function here:
func RunningAverage() func(newValue float64) float64 {
    count:=0
	var sum float64 = 0.0

	return func (newValue float64) float64 {
		count++
		sum+=newValue
		return sum/float64(count)
	}
}


func main() {
    // 1. Initialize two independent average calculators.

	avgCalcA:= RunningAverage()
    
    // 2. Call Counter A with data points.
    fmt.Println("--- Calculator A ---")
    
    avg := avgCalcA(10.0) // Expected: 10.0 / 1 = 10.0
    fmt.Printf("New Value: 10.0 | Current Average: %.2f\n", avg)
    
    avg = avgCalcA(20.0) // Expected: (10.0 + 20.0) / 2 = 15.0
    fmt.Printf("New Value: 20.0 | Current Average: %.2f\n", avg)

    avg = avgCalcA(30.0) // Expected: (10.0 + 20.0 + 30.0) / 3 = 20.0
    fmt.Printf("New Value: 30.0 | Current Average: %.2f\n", avg)
    
    // 3. (Optional) Initialize and use a second instance (avgCalcB) to confirm state isolation.

	avgcalcB:=RunningAverage()

	fmt.Println("\n--- Calculator B ---")
	avg1:=avgcalcB(8.0)
	fmt.Printf("New Value: 8.0 | Current Average: %.2f\n", avg1)
	avg1=avgcalcB(12.0) // Expected: (8.0 + 12.0) / 2 = 10.0
	fmt.Printf("New Value: 12.0 | Current Average: %.2f\n", avg1)
}