package parser

// EntryType encodes a valid gopher menu type.
type EntryType string

// Entry defines a single line in a gopher menu.
// Server and Port may be empty values, e.g. for INFO lines.
type Entry struct {
	Type     EntryType
	UserName string
	Selector string
	Server   string
	Port     string
}

const (
	// canonical types as defined in https://tools.ietf.org/html/rfc1436

	// TextFile is a file with text content
	TextFile = "TEXT_FILE"
	// Directory is a gopher submenu
	Directory = "DIRECTORY"
	// PhoneBook is a CSO phone-book server
	PhoneBook = "PHONE_BOOK"
	// IndexSearch is an Index-Search server
	IndexSearch = "INDEX_SEARCH"

	// non-canonical, but broadly supported types

	// Info is an info-line
	Info = "INFO"
	// Html is an HTML-file
	Html = "HTML"

	// custom types used in this client implementation

	// Invalid denotes an invalid menu entry
	Invalid = "INVALID"
)

var types = map[byte]EntryType{
	'0': TextFile,
	'1': Directory,
	'2': PhoneBook,
	'7': IndexSearch,
	'h': Html,
	'i': Info,
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
	p := &Parser{
		input:    input,
		position: 0,
	}
	p.nextChar()
	return p
}

// NextLine parses a line from a gopher menu and returns the type according to RFC1436.
func (p *Parser) NextLine() *Entry {
	t, ok := types[p.ch]
	if !ok {
		return &Entry{
			Type:     Invalid,
			UserName: p.input,
		}
	}
	p.nextChar()

	userName := p.nextSegment()
	selector := p.nextSegment()
	server := p.nextSegment()
	port := p.nextSegment()

	return &Entry{
		Type:     t,
		UserName: userName,
		Selector: selector,
		Server:   server,
		Port:     port,
	}
}

func (p *Parser) nextChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0
	} else {
		p.ch = p.input[p.readPosition]
	}

	p.position = p.readPosition
	p.readPosition++
}

func (p *Parser) nextSegment() string {
	pos := p.position

	for p.ch != '\t' && p.ch != 0 {
		p.nextChar()
	}
	buf := p.input[pos:p.position]
	p.nextChar()

	return buf
}
