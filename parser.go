// generated by Textmapper; DO NOT EDIT

package ll

import (
	"fmt"
)

// Parser is a table-driven LALR parser for llvm.
type Parser struct {
	listener Listener

	next symbol
}

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int16
}

func (p *Parser) Init(l Listener) {
	p.listener = l
}

const (
	startStackSize = 256
	noToken        = int32(UNAVAILABLE)
	eoiToken       = int32(EOI)
	debugSyntax    = false
)

func (p *Parser) Parse(lexer *Lexer) error {
	return p.parse(0, 2354, lexer)
}

func (p *Parser) parse(start, end int16, lexer *Lexer) error {
	state := start

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			action = lalr(action, p.next.symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				entry.sym.offset, entry.sym.endoffset = p.next.offset, p.next.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule(rule, &entry, rhs, lexer); err != nil {
				return err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			state = gotoState(state, p.next.symbol)
			stack = append(stack, stackEntry{
				sym:   p.next,
				state: state,
			})
			if debugSyntax {
				fmt.Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
			}
			if state != -1 && p.next.symbol != eoiToken {
				p.next.symbol = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	if state != end {
		if p.next.symbol == noToken {
			p.fetchNext(lexer, stack)
		}
		err := SyntaxError{
			Line:      lexer.Line(),
			Offset:    p.next.offset,
			Endoffset: p.next.endoffset,
		}
		return err
	}

	return nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int16, symbol int32) int16 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
	case INVALID_TOKEN:
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

func (p *Parser) applyRule(rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}
