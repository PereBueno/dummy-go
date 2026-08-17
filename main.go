package main

// Import CloudBees Unify SDK
import (
	"fmt"

	"github.com/rollout/rox-go/v6/server"
)

// Create Rox flags in the Flags container class
type Flags struct {
	EnableTutorial server.RoxFlag
}

var flags = &Flags{
	// Define the feature flags
	EnableTutorial: server.NewRoxFlag(false),
}

var rox *server.Rox

func main() {

	options := server.NewRoxOptions(server.RoxOptionsBuilder{})

	rox := server.NewRox()

	// Register the flags container with CloudBees Unify.
	rox.RegisterWithEmptyNamespace(flags)

	// Setup the feature management environment key
	<-rox.Setup("fc6699f7-ab42-4ed0-aadc-0654a006b10b", options)

	// Boolean flag example
	fmt.Printf("EnableTutorial's value is %t\n", flags.EnableTutorial.IsEnabled(nil))
}
