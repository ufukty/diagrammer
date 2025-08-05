package tokens

type Activation string

const (
	Activate   Activation = "activate"
	Deactivate Activation = "deactivate"
)

type NotePos string

const (
	Over    NotePos = "over"
	LeftOf  NotePos = "left of"
	RightOf NotePos = "right of"
)
