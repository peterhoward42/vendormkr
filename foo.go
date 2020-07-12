package main

// For each dependency you want to vendor, import the package, and then
// add something from the package to the *deps* slice. (So that it 
// will compile without an unused-import error).

// The thing you choose to add is semantically arbitrary, but it has to
// evaluate to an expression to satisfy the Null interface type defined for
// *dep's* slice members. Functions are perhaps the easiest choice.

// See README.md for how to then vendorise these dependencies.


import (
	"github.com/Shopify/sarama"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	deps := []interface{}{
		sarama.NewConfig,
		govalidator.SetFieldsRequiredByDefault,
		gin.Default,
		godotenv.Load,
		envconfig.Process,
	}

	// Prevent "deps declared and not used" compilation error.
	_ = deps
}

