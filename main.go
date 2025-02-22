/*
Copyright © 2024 Michael Putera Wardana <michaelputeraw@gmail.com>
*/
package main

import (
	"github.com/api-monolith-template/cmd"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	cmd.Execute()
}
