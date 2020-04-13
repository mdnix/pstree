package main

import (
	"log"
	"fmt"
	"pstree/pkg/pstree"
)

func main() {
	tree := make(pstree.ProcessTree)

	if err := tree.Populate(); err != nil {
		log.Fatalf("failed getting proceses: %v", err)
		return
	}

	fmt.Print(tree)
}
