package parser

// EntryType encodes a valid gopher menu type.
type EntryType string

type Entry struct {
	Type    EntryType
	Content string
}

const (
	// INFO is an info-line
	INFO = "INFO"

	// INVALID denotes an invalid menu entry
	INVALID = "INVALID"
)

var types = map[byte]EntryType{
	'i': INFO,
}

// Parser encapsulates the internal state while reading a gopher menu.
type Parser struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

// New returns a reference to a Parser, fully initialized to the input.
func New(input string) *Parser {
	return &Parser{
		input:    input,
		position: 0,
	}
}

// Line parses a line from a gopher menu and returns the type according to RFC1436
func (p *Parser) Line() *Entry {
	if len(p.input) <= 0 {
		return nil
	}

	t, ok := types[p.input[0]]
	if !ok {
		return &Entry{
			Type:    INVALID,
			Content: "",
		}
	}

	p.input = p.input[1:]

	content := p.nextSegment()

	return &Entry{
		Type:    t,
		Content: content,
	}
}

func (p *Parser) nextChar() {
	if p.position >= len(p.input) {
		p.ch = 0
	} else {
		p.ch = p.input[p.readPosition]
	}

	p.position = p.readPosition
	p.readPosition++
}
func (p *Parser) nextSegment() string {
	pos := p.position
	for p.ch != '\t' {
		p.nextChar()
	}

	return p.input[pos:p.position]

}
