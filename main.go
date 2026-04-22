// main.go
package main

import (
	"context"
	"os"

	"github.com/dottmp/asd-uptime/components"
)

func main() {
	component := components.Greet("John", 25)
	component.Render(context.Background(), os.Stdout)
}
