package main

import (
	"fmt"
	"github.com/dminGod/Dwara/config"
)

func main() {

	fmt.Println("Shri ganeshai namaha!")

	config.GetConfig()

	config.ShowConfig()

}




