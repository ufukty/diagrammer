package lexer

import (
	"regexp"
	"strings"

	"github.com/ufukty/diagramer/pkg/sequence/lexer/tokens"
)

func startsWith(line, word string) bool {
	return strings.Split(line, " ")[0] == word
}

func (Activate) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "activate")
}

func (Alt) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "alt")
}

func (And) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "and")
}

func (Box) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "box")
}

func (Break) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "break")
}

func (Create) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "create")
}

func (Critical) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "critical")
}

func (Deactivate) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "deactivate")
}

func (Destroy) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "destroy")
}

func (Else) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "else")
}

func (End) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "end")
}

func (LifelineDecl) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "participant") || startsWith(strings.TrimSpace(line), "actor")
}

func (Loop) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "loop")
}

func (Message) check(line string) bool {
	return strings.Contains(line, "->>")
}

var wideNote = &WideNote{}

func (Note) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "note") && !wideNote.check(line)
}

func (Option) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "opt")
}

func (Parallel) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "par")
}

func (WideNote) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "note") && strings.Contains(strings.Split(line, ":")[0], ",")
}

var (
	regexActivate   = regexp.MustCompile(`activate\s+(\w+)`)
	regexAlt        = regexp.MustCompile(`alt\s+(.+)`)
	regexAnd        = regexp.MustCompile(`and\s+(.+)`)
	regexBox        = regexp.MustCompile(`box(?:\s+(#[0-9A-Fa-f]{3,8}))?(?:\s+([^#].*))?`)
	regexBreak      = regexp.MustCompile(`break\s+(.+)`)
	regexCritical   = regexp.MustCompile(`critical\s+(.+)`)
	regexDeactivate = regexp.MustCompile(`deactivate\s+(\w+)`)
	regexDestroy    = regexp.MustCompile(`destroy\s+(\w+)`)
	regexElse       = regexp.MustCompile(`else\s+(.+)`)
	regexLifeline   = regexp.MustCompile(`(participant|actor)\s+(\w+)(?:\s+as\s+(.+))?`)
	regexLoop       = regexp.MustCompile(`loop\s+(.+)`)
	regexMessage    = regexp.MustCompile(`(\w+)\s*(->>)\s*(\w+)([-+]{1})?(?::\s*(.+))?`)
	regexNote       = regexp.MustCompile(`note(?:\s+(left of|right of|over)(?:\s+(\w+)(?::(?:\s+(.*))?)?)?)?`)
	regexOption     = regexp.MustCompile(`opt\s+(.+)`)
	regexParallel   = regexp.MustCompile(`par\s+(.+)`)
	regexWideNote   = regexp.MustCompile(`note(?:\s+over(?:\s+(\w+)(?:\s*,\s*(\w+)(?::\s+(.*))?)?)?)?`)
)

func (Activate) construct(line string) Line {
	a := &Activate{}
	if ms := regexActivate.FindStringSubmatch(line); len(ms) > 1 {
		a.Lifeline = ms[1]
	}
	return a
}

func (Alt) construct(line string) Line {
	a := &Alt{}
	if ms := regexAlt.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

func (And) construct(line string) Line {
	a := &And{}
	if ms := regexAnd.FindStringSubmatch(line); len(ms) > 1 {
		a.Action = ms[1]
	}
	return a
}

func (Box) construct(line string) Line {
	a := &Box{}
	if ms := regexBox.FindStringSubmatch(line); len(ms) > 2 {
		a.Color = ms[1]
		a.Title = ms[2]
	}
	return a
}

func (Break) construct(line string) Line {
	a := &Break{}
	if ms := regexBreak.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

var lifelineDecl = &LifelineDecl{}

func (Create) construct(line string) Line {
	a := &Create{}
	if ld := lifelineDecl.construct(strings.TrimSpace(strings.TrimPrefix(line, "create"))); ld != nil {
		if ld, ok := ld.(*LifelineDecl); ok {
			a.LifelineDecl = *ld
		}
	}
	return a
}

func (Critical) construct(line string) Line {
	a := &Critical{}
	if ms := regexCritical.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

func (Deactivate) construct(line string) Line {
	a := &Deactivate{}
	if ms := regexDeactivate.FindStringSubmatch(line); len(ms) > 1 {
		a.Lifeline = ms[1]
	}
	return a
}

func (Destroy) construct(line string) Line {
	a := &Destroy{}
	if ms := regexDestroy.FindStringSubmatch(line); len(ms) > 1 {
		a.Name = ms[1]
	}
	return a
}

func (Else) construct(line string) Line {
	a := &Else{}
	if ms := regexElse.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

func (End) construct(line string) Line {
	return &End{}
}

func (LifelineDecl) construct(line string) Line {
	match := regexLifeline.FindStringSubmatch(line)
	if len(match) < 3 {
		return nil
	}
	p := &LifelineDecl{
		Type:  match[1],
		Alias: "",
		Name:  match[2],
	}
	if len(match) > 3 {
		p.Alias = match[3]
	}
	return p
}

func (Loop) construct(line string) Line {
	a := &Loop{}
	if ms := regexLoop.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

func (Message) construct(line string) Line {
	m := &Message{}
	if match := regexMessage.FindStringSubmatch(line); len(match) > 5 {
		m.From = match[1]
		m.To = match[3]
		m.Content = match[5]
		switch match[4] {
		case "+":
			m.Activation = tokens.Activate
		case "-":
			m.Activation = tokens.Deactivate
		}
	}
	return m
}

func (Note) construct(line string) Line {
	a := &Note{}
	if ms := regexNote.FindStringSubmatch(line); len(ms) > 3 {
		a.Pos = tokens.NotePos(ms[1])
		a.Lifeline = ms[2]
		a.Content = ms[3]
	}
	return a
}

func (Option) construct(line string) Line {
	a := &Option{}
	if ms := regexOption.FindStringSubmatch(line); len(ms) > 1 {
		a.Description = ms[1]
	}
	return a
}

func (Parallel) construct(line string) Line {
	a := &Parallel{}
	if ms := regexParallel.FindStringSubmatch(line); len(ms) > 1 {
		a.Action = ms[1]
	}
	return a
}

func (WideNote) construct(line string) Line {
	a := &WideNote{}
	if ms := regexWideNote.FindStringSubmatch(line); len(ms) > 3 {
		a.From = ms[1]
		a.To = ms[2]
		a.Content = ms[3]
	}
	return a
}
