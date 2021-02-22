package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {
	fmt.Println("Testing...")
	t.Errorf("Simulate an error!")
}