package main

import (
	"github.com/gagandeepahuja09/event-delivery/internal"
)

func main() {
	internal.InitProviders()
	internal.ProcessEvents()
	internal.InitRouter()
}
