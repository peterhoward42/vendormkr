package main

// For each dependency you want to vendor, import the package, and then
// add function from the package to the *deps* slice. (So that it will compile
// without an unused import error).

// See README.md for how to vendorise the dependencies.


import (
    "github.com/kelseyhightower/envconfig"
)

func main() {
	deps := []interface{}{
		envconfig.Process,
	}

	// Prevent "deps declared and not used".
	_ = deps
}

