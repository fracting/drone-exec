package runner

import "github.com/fracting/drone-exec/parser"

func Load(tree *parser.Tree) *Build {
	return &Build{tree: tree}
}
