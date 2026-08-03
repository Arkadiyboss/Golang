package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestNoLeaks(t *testing.T) {
	defer goleak.VerifyNone(t)
	main()
}
