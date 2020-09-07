package main

// For each dependency you want to vendor, import the package, and then
// add something from the package to the *deps* slice. (So that it 
// will compile without an unused-import error).

// The thing you choose to add is semantically arbitrary, but it has to
// evaluate to an expression to satisfy the Null interface type defined for
// *dep's* slice members. Functions are perhaps the easiest choice.

// See README.md for how to then vendorise these dependencies.


import (
	"github.com/stretchr/testify/assert"
	"github.com/antonmedv/expr"
	yaml "gopkg.in/yaml.v2"

	"google.golang.org/grpc"
)

func main() {
	deps := []interface{}{
		assert.Equal,
		expr.Env,
		yaml.Marshal,
		grpc.SupportPackageIsVersion4,
	}

	// Prevent "deps declared and not used" compilation error.
	_ = deps
}

