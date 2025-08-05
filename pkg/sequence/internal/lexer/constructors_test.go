package lexer

import (
	"maps"
	"slices"
	"testing"

	"github.com/ufukty/diagramer/pkg/sequence/lexer/tokens"
)

func TestActivate(t *testing.T) {
	tcs := map[string]*Activate{
		"activate":   {Lifeline: ""},
		"activate a": {Lifeline: "a"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Activate{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Activate{}).construct(input).(*Activate)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestAlt(t *testing.T) {
	tcs := map[string]*Alt{
		"alt":             {},
		"alt description": {Description: "description"},
		"alt abc 123 +%&": {Description: "abc 123 +%&"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Alt{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Alt{}).construct(input).(*Alt)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	tcs := map[string]*And{
		"and":             {},
		"and action-name": {Action: "action-name"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(And{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (And{}).construct(input).(*And)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestBox(t *testing.T) {
	tcs := map[string]*Box{
		"box #000 title":      {Color: "#000", Title: "title"},
		"box #000":            {Color: "#000"},
		"box #0000 title":     {Color: "#0000", Title: "title"},
		"box #000000 title":   {Color: "#000000", Title: "title"},
		"box #00000000 title": {Color: "#00000000", Title: "title"},
		"box t#000":           {Title: "t#000"},
		"box title #000":      {Title: "title #000"},
		"box title":           {Title: "title"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Box{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Box{}).construct(input).(*Box)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestBreak(t *testing.T) {
	tcs := map[string]*Break{
		"break":             {},
		"break description": {Description: "description"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Break{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Break{}).construct(input).(*Break)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tcs := map[string]*Create{
		"create actor a as Alice":       {LifelineDecl: LifelineDecl{Type: "actor", Name: "a", Alias: "Alice"}},
		"create actor a":                {LifelineDecl: LifelineDecl{Type: "actor", Name: "a", Alias: ""}},
		"create participant a as Alice": {LifelineDecl: LifelineDecl{Type: "participant", Name: "a", Alias: "Alice"}},
		"create participant a":          {LifelineDecl: LifelineDecl{Type: "participant", Name: "a", Alias: ""}},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Create{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Create{}).construct(input).(*Create)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestCritical(t *testing.T) {
	tcs := map[string]*Critical{
		"critical":             {},
		"critical description": {Description: "description"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Critical{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Critical{}).construct(input).(*Critical)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestDeactivate(t *testing.T) {
	tcs := map[string]*Deactivate{
		"deactivate a": {Lifeline: "a"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Deactivate{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Deactivate{}).construct(input).(*Deactivate)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tcs := map[string]*Destroy{
		"destroy a": {Name: "a"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Destroy{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Destroy{}).construct(input).(*Destroy)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestElse(t *testing.T) {
	tcs := map[string]*Else{
		"else":             {},
		"else description": {Description: "description"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Else{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Else{}).construct(input).(*Else)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestEnd(t *testing.T) {
	tcs := map[string]*End{
		"end": {},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(End{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (End{}).construct(input).(*End)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestLifelineDecl(t *testing.T) {
	tcs := map[string]*LifelineDecl{
		"actor a as Alice":       {Type: "actor", Name: "a", Alias: "Alice"},
		"actor a":                {Type: "actor", Name: "a", Alias: ""},
		"participant a as Alice": {Type: "participant", Name: "a", Alias: "Alice"},
		"participant a":          {Type: "participant", Name: "a", Alias: ""},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(LifelineDecl{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (LifelineDecl{}).construct(input).(*LifelineDecl)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestLoop(t *testing.T) {
	tcs := map[string]*Loop{
		"loop":             {},
		"loop description": {Description: "description"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Loop{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Loop{}).construct(input).(*Loop)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	tcs := map[string]*Message{
		"a->>b-: ACK": {From: "a", To: "b", Content: "ACK", Activation: tokens.Deactivate},
		"a->>b-":      {From: "a", To: "b", Content: "", Activation: tokens.Deactivate},
		"a->>b: ACK":  {From: "a", To: "b", Content: "ACK"},
		"a->>b":       {From: "a", To: "b", Content: ""},
		"a->>b+: ACK": {From: "a", To: "b", Content: "ACK", Activation: tokens.Activate},
		"a->>b+":      {From: "a", To: "b", Content: "", Activation: tokens.Activate},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Message{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Message{}).construct(input).(*Message)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestNote(t *testing.T) {
	tcs := map[string]*Note{
		"note left of a: lorem ipsum":  {Lifeline: "a", Pos: tokens.LeftOf, Content: "lorem ipsum"},
		"note left of a:":              {Lifeline: "a", Pos: tokens.LeftOf, Content: ""},
		"note left of a":               {Lifeline: "a", Pos: tokens.LeftOf, Content: ""},
		"note over a: lorem ipsum":     {Lifeline: "a", Pos: tokens.Over, Content: "lorem ipsum"},
		"note over a:":                 {Lifeline: "a", Pos: tokens.Over, Content: ""},
		"note over a":                  {Lifeline: "a", Pos: tokens.Over, Content: ""},
		"note right of a: lorem ipsum": {Lifeline: "a", Pos: tokens.RightOf, Content: "lorem ipsum"},
		"note right of a:":             {Lifeline: "a", Pos: tokens.RightOf, Content: ""},
		"note right of a":              {Lifeline: "a", Pos: tokens.RightOf, Content: ""},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Note{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Note{}).construct(input).(*Note)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestOption(t *testing.T) {
	tcs := map[string]*Option{
		"opt":             {},
		"opt description": {Description: "description"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Option{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Option{}).construct(input).(*Option)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestParallel(t *testing.T) {
	tcs := map[string]*Parallel{
		"par":        {},
		"par action": {Action: "action"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Parallel{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Parallel{}).construct(input).(*Parallel)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}

func TestWideNote(t *testing.T) {
	tcs := map[string]*WideNote{
		"note over a, b: lorem ipsum": {From: "a", To: "b", Content: "lorem ipsum"},
		"note over a, b:":             {From: "a", To: "b"},
		"note over a, b":              {From: "a", To: "b"},
		"note over a,b: lorem ipsum":  {From: "a", To: "b", Content: "lorem ipsum"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(WideNote{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (WideNote{}).construct(input).(*WideNote)
			if *got != *expected {
				t.Errorf("assertion:\nexpected %#v\ngot      %#v", expected, got)
			}
		})
	}
}
