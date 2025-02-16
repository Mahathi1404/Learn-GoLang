/* closures example 2 */

package main

import "fmt"

/**
 * 1. The Factory Function
 * Goal: Takes a tax rate and returns a closure (the specialized calculator).
 */
func createTaxCalculator(rate float64) func(price float64) float64 {
	// 🎯 YOUR CODE HERE:
	// a) Capture the 'rate' variable in the closure.
	// b) Return the inner function that calculates price * (1 + captured_rate).
	cur_rate := rate
	return func(price float64) float64 {
		return price * (1 + cur_rate)
	}
}

func main() {
	// --- Demonstration of Independent Closures ---
	
	// 2. Create the first calculator for Region A (5% tax)
	// The closure should capture 0.05
	regionA_Calc := createTaxCalculator(0.05) 

	// 3. Create the second calculator for Region B (10% tax)
	// The closure should capture 0.10
	regionB_Calc := createTaxCalculator(0.10) 

	// 4. Test the independence
	
	itemPrice := 100.0 // Base price is $100.00
	
	// Region A: Should calculate $100 * (1 + 0.05) = $105.00
	priceA := regionA_Calc(itemPrice) 
	
	// Region B: Should calculate $100 * (1 + 0.10) = $110.00
	priceB := regionB_Calc(itemPrice) 
	
	// 5. Print results to confirm isolation
	fmt.Printf("Base Price: $%.2f\n\n", itemPrice)
	fmt.Printf("Region A (5%%): Final Price = $%.2f\n", priceA)
	fmt.Printf("Region B (10%%): Final Price = $%.2f\n", priceB)
}