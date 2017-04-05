package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/movio/itrgen"
)

func main() {
	pkg := flag.String("package", "main", "the package that this Itr will be in. Defaults to 'main'")
	name := flag.String("type", "", "the main type that we want to map from. eg, string or MyStruct")
	mapTo := flag.String("map-to", "", "comma separated list of types that can be mapped to")
	flag.Parse()
	if *name == "" {
		fmt.Println("You must specify -type=...")
		os.Exit(2)
	}

	maps := strings.Split(strings.Trim(*mapTo, "\""), ",")

	t := itrgen.NewType(
		strings.Trim(*pkg, "\""),
		strings.Trim(*name, "\""),
		maps,
	)

	if err := itrgen.Generate(t); err != nil {
		fmt.Printf("Error in generation: %v\n", err)
		os.Exit(2)
	}
}
