// © 2018 Imhotep Software LLC. All rights reserved.
package main

import (
	"flag"
	"fmt"

	"github.com/imhotepio/letsgo_labs/dictionary"
)

func main() {
	var (
		dic = flag.String("dic", "default", "Specifies a dictionary name. Defaults is default")
		dir = flag.String("dir", "./assets", "Specifies dictionaries load directory. Defaults is ./assets")
	)

	flag.Parse()

	if wl, err := dictionary.Load(*dir, *dic); err != nil {
		panic(err)
	} else {
		fmt.Printf("The word of the day is `%s\n", dictionary.Pick(wl))
	}
}
