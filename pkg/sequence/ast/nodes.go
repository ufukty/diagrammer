package ast

import "github.com/ufukty/diagramer/pkg/sequence/lexer/tokens"

type (
	LifelineDecl struct {
		Type  string
		Alias string
		Name  string
	}

	Create struct {
		LifelineDecl
	}

	Destroy struct {
		Name string
	}

	Box struct {
		Color string
		Title string
	}

	Activate struct {
		Lifeline string
	}

	Deactivate struct {
		Lifeline string
	}
)

type (
	Message struct {
		From, To   string
		Content    string
		Activation tokens.Activation
	}

	Note struct {
		Lifeline string
		Pos      tokens.NotePos
		Content  string
	}

	WideNote struct {
		From, To string
		Content  string
	}
)

type Case struct {
	Annotation string
	Stmts      []Stmt
}

type (
	Break    Case
	Loop     Case
	Optional Case
)

type (
	Alternative []Case
	Critical    []Case
	Parallel    []Case
)

type (
	DiagramOpts struct {
		AutoNumber bool
	}

	Diagram struct {
		Stmts []Stmt
		Opts  DiagramOpts
	}
)
