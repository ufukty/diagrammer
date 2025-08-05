package lexer

import "github.com/ufukty/diagramer/pkg/sequence/lexer/tokens"

// MARK: Lifelines and rel.

type LifelineDecl struct {
	Type  string
	Alias string
	Name  string
}

type Create struct {
	LifelineDecl
}

type Destroy struct {
	Name string
}

// Box expects [End]
type Box struct {
	Color string
	Title string
}

type Activate struct {
	Lifeline string
}

type Deactivate struct {
	Lifeline string
}

// MARK: Messages and rel.

type Message struct {
	From, To   string
	Content    string
	Activation tokens.Activation
}

type Note struct {
	Lifeline string
	Pos      tokens.NotePos
	Content  string
}

type WideNote struct {
	From, To string
	Content  string
}

// MARK: Single block

// Break expects [End]
type Break struct {
	Description string
}

// Loop expects [End]
type Loop struct {
	Description string
}

// MARK: Multiple blocks

// Critical accepts [Option], expects [End]
type Critical struct {
	Description string
}

// Option can be used inside [Critical]
type Option struct {
	Description string
}

// End is used after [Alt], [Box], [Break], [Critical], [Loop], [Parallel]
type End struct{}

type Parallel struct {
	Action string
}

// And is used after [Parallel]
type And struct {
	Action string
}

type Alt struct {
	Description string
}

// Else is used after [Alt], [Parallel]
type Else struct {
	Description string
}

// MARK: Diagram

type DiagramOpts struct {
	AutoNumber bool
}

type Line interface {
	_line()
}

func (*Activate) _line()     {}
func (*Alt) _line()          {}
func (*And) _line()          {}
func (*Box) _line()          {}
func (*Break) _line()        {}
func (*Create) _line()       {}
func (*Critical) _line()     {}
func (*Deactivate) _line()   {}
func (*Destroy) _line()      {}
func (*Else) _line()         {}
func (*End) _line()          {}
func (*LifelineDecl) _line() {}
func (*Loop) _line()         {}
func (*Message) _line()      {}
func (*Note) _line()         {}
func (*Option) _line()       {}
func (*Parallel) _line()     {}
func (*WideNote) _line()     {}

type Diagram struct {
	Opts  DiagramOpts
	Lines []Line
}
