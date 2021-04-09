package tailcat_test

import (
	"embed"
	"fmt"
	"io/fs"
	
	"github.com/parrogo/tailcat"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use tailcat.Func()
func ExampleFunc() {
	fmt.Println(tailcat.Func())
	// Output: 42
}
