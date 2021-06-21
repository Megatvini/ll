package main

import (
	"flag"
	"fmt"
	"github.com/llir/ll/ast"
	"github.com/llir/ll/selector"
	"os"
	"path/filepath"
)

func printTreeAst(node *ast.Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}

	literal := ""
	// is leaf
	if len(node.Children(selector.Any)) == 0 {
		literal = node.Text()
	}

	fmt.Printf("%s,%d,%d,%s", node.Type().String(), node.Offset(), node.Endoffset(), literal)
	fmt.Printf("\n")
	for _, n := range node.Children(selector.Any) {
		printTreeAst(n, depth+1)
	}
}

func main() {
	llFilePathPointer := flag.String("ll-file-path", "example.ll", "Path to the ll input file")
	flag.Parse()

	llFilePath := *llFilePathPointer
	contentBytes, err := os.ReadFile(llFilePath)
	if err != nil {
		panic(err)
	}

	content := string(contentBytes)

	dir, _ := filepath.Split(llFilePath)

	tree, err := ast.Parse(dir, content)
	if err != nil {
		panic(err)
	}
	printTreeAst(tree.Root(), 0)
}
