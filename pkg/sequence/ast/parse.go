package ast

import (
	"fmt"
	"io"
	"strings"

	"github.com/ufukty/diagramer/pkg/sequence/ast/internal/lexer"
)

func Parse(src io.Reader) (*Diagram, error) {
	l, err := lexer.FromReader(src)
	if err != nil {
		return nil, fmt.Errorf("lexer: %w", err)
	}

	diagram := &Diagram{
		Stmts: []Stmt{},
		Opts:  DiagramOpts{},
	}

	errs := []string{}
	lls := map[string]*LifelineDecl{} // name => node
	stack := []ScopeDefining{}
	for _, stmt := range l.Lines {
		latest := stack[len(stack)-1]

		switch stmt := stmt.(type) {

		case *lexer.Activate:
			panic("not implemented")

		case *lexer.Alt:
			panic("not implemented")

		case *lexer.And:
			panic("not implemented")

		case *lexer.Box:
			panic("not implemented")

		case *lexer.Break:
			panic("not implemented")

		case *lexer.Create:
			panic("not implemented")

		case *lexer.Critical:
			panic("not implemented")

		case *lexer.Deactivate:
			panic("not implemented")

		case *lexer.Destroy:
			panic("not implemented")

		case *lexer.Else:
			panic("not implemented")

		case *lexer.End:
			panic("not implemented")

		case *lexer.LifelineDecl:
			ll := &LifelineDecl{
				Type:  stmt.Type,
				Alias: stmt.Alias,
				Name:  stmt.Name,
			}
			lls[ll.Name] = ll
			latest.AppendStmt(ll)

		case *lexer.Loop:
			panic("not implemented")

		case *lexer.Message:
			latest.AppendStmt(&Message{
				Activation: stmt.Activation,
				Content:    stmt.Content,
				From:       stmt.From,
				To:         stmt.To,
			})

		case *lexer.Note:
			panic("not implemented")

		case *lexer.Option:
			panic("not implemented")

		case *lexer.Parallel:
			panic("not implemented")

		case *lexer.WideNote:
			panic("not implemented")

		default:
			fmt.Printf("skipping #v\n")
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("found %d errors:\n%s", len(errs), strings.Join(errs, "\n"))
	}

	return diagram, nil
}
