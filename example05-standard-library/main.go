package main

import (
	"fmt"
	"log"
)

// OrderError represents a general error related to orders
type OrderError struct {
	OrderID string
	Message string
}

// Error is a method that returns the error message
func (e OrderError) Error() string {
	return fmt.Sprintf("Order %s: %s", e.OrderID, e.Message)
}

// PaymentDeclinedError represents a payment failure
type PaymentDeclinedError struct {
	OrderID string
	Reason  string
}

// Error is a method that returns the error message
func (e PaymentDeclinedError) Error() string {
	return fmt.Sprintf("Payment declined for order %s: %s", e.OrderID, e.Reason)
}

func LogError(err error) {
	log.Fatal(fmt.Errorf("received error: %w", err))
}

func main() {
	LogError(OrderError{OrderID: "123", Message: "Order not found"})
	LogError(PaymentDeclinedError{OrderID: "123", Reason: "Insufficient funds"})
}
