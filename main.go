/*
Copyright © 2024 Michael Putera Wardana <michaelputeraw@krobot.my.id>
*/
package main

import (
	"github.com/api-monolith-template/cmd"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	cmd.Execute()
}
