// Package lexer tokenizes the lines into nodes that has limited hierarchy.
package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type constructor interface {
	check(string) bool
	construct(string) Line
}

var precedence = []constructor{
	&Activate{},
	&Alt{},
	&And{},
	&Box{},
	&Break{},
	&Create{},
	&Critical{},
	&Deactivate{},
	&Destroy{},
	&Else{},
	&End{},
	&LifelineDecl{},
	&Loop{},
	&Message{},
	&Note{},
	&Option{},
	&Parallel{},
	&WideNote{},
}

func fromLine(line string) Line {
	for _, constructor := range precedence {
		if ok := constructor.check(line); !ok {
			continue
		}
		if lm := constructor.construct(line); lm != nil {
			return lm
		}
	}
	return nil
}

func FromReader(src io.Reader) (*Diagram, error) {
	diagram := &Diagram{
		Lines: []Line{},
		Opts:  DiagramOpts{},
	}

	scanner := bufio.NewScanner(src)
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "%%") {
			continue
		}

		switch line {
		case "sequenceDiagram":
			continue

		case "autoNumber":
			diagram.Opts.AutoNumber = true

		default:
			if l := fromLine(line); l != nil {
				diagram.Lines = append(diagram.Lines, l)
			}
			fmt.Printf("skipping line #%d: %s\n", i, line)
		}
	}

	return diagram, nil
}
