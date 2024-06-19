// U2SMTP - smtp server
//
// Licensed under the MIT License for individual use or a commercial
// license for enterprise use. See LICENSE.txt and COMMERCIAL_LICENSE.txt.

package server

import (
	"fmt"
	"strings"
)

type CmdError int

const (
	CmdParseOk CmdError = iota
	CmdParseInvalid
)

// parse command
func parseCommand(line string) (cmd string, arg string, err CmdError) {
	line = strings.TrimRight(line, "\r\n")

	l := len(line)
	switch {
	case strings.HasPrefix(strings.ToUpper(line), "STARTTLS"):
		return "STARTTLS", "", CmdParseOk
	case l == 0:
		return "", "", CmdParseOk
	case l < 4:
		return "", "", CmdParseInvalid
	case l == 4:
		return strings.ToUpper(line), "", CmdParseOk
	case l == 5:
		return "", "", CmdParseInvalid
	}

	if line[4] != ' ' {
		return "", "", CmdParseInvalid
	}

	return strings.ToUpper(line[0:4]), strings.TrimSpace(line[5:]), CmdParseOk
}

type HelloArgError int

const (
	HelloArgOk HelloArgError = iota
	HelloArgDomainInvalid
)

// parse HELO comman arguments
func parseHelloArguments(arg string) (string, HelloArgError) {
	domain := arg
	if idx := strings.IndexRune(arg, ' '); idx >= 0 {
		domain = arg[:idx]
	}
	if domain == "" {
		return "", HelloArgDomainInvalid
	}

	// if domain end with . than remove it
	lastChar := domain[len(domain)-1:]
	if lastChar == "." {
		domain = domain[:len(domain)-1]
	}
	// fmt.Println(lastChar)

	return domain, HelloArgOk
}

// parser parses command arguments defined in RFC 5321 section 4.1.2.
type Parser struct {
	s string
}

func newParser(str string) Parser {
	return Parser{
		s: str,
	}
}

// trim white space
func (p *Parser) trim() {
	p.s = strings.TrimSpace(p.s)
}

func (p *Parser) peekByte() (byte, bool) {
	if len(p.s) == 0 {
		return 0, false
	}
	return p.s[0], true
}

func (p *Parser) readByte() (byte, bool) {
	ch, ok := p.peekByte()
	if ok {
		p.s = p.s[1:]
	}
	return ch, ok
}

func (p *Parser) acceptByte(ch byte) bool {
	got, ok := p.peekByte()
	if !ok || got != ch {
		return false
	}
	p.readByte()
	return true
}

func (p *Parser) expectByte(ch byte) error {
	if !p.acceptByte(ch) {
		if len(p.s) == 0 {
			return fmt.Errorf("expected '%v', got EOF", string(ch))
		} else {
			return fmt.Errorf("expected '%v', got '%v'", string(ch), string(p.s[0]))
		}
	}
	return nil
}

func (p *Parser) parseReversePath() (string, error) {
	if strings.HasPrefix(p.s, "<>") {
		p.s = strings.TrimPrefix(p.s, "<>")
		return "", nil
	}
	return p.parsePath()
}

func (p *Parser) parsePath() (string, error) {
	hasBracket := p.acceptByte('<')
	if p.acceptByte('@') {
		i := strings.IndexByte(p.s, ':')
		if i < 0 {
			return "", fmt.Errorf("malformed a-d-l")
		}
		p.s = p.s[i+1:]
	}
	mbox, err := p.parseMailbox()
	if err != nil {
		return "", fmt.Errorf("in mailbox: %v", err)
	}
	if hasBracket {
		if err := p.expectByte('>'); err != nil {
			return "", err
		}
	}
	return mbox, nil
}

func (p *Parser) parseMailbox() (string, error) {
	localPart, err := p.parseLocalPart()
	if err != nil {
		return "", fmt.Errorf("in local-part: %v", err)
	} else if localPart == "" {
		return "", fmt.Errorf("local-part is empty")
	}

	if err := p.expectByte('@'); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(localPart)
	sb.WriteByte('@')

	for {
		ch, ok := p.peekByte()
		if !ok {
			break
		}
		if ch == ' ' || ch == '\t' || ch == '>' {
			break
		}
		p.readByte()
		sb.WriteByte(ch)
	}

	if strings.HasSuffix(sb.String(), "@") {
		return "", fmt.Errorf("domain is empty")
	}

	return sb.String(), nil
}

func (p *Parser) parseLocalPart() (string, error) {
	var sb strings.Builder

	if p.acceptByte('"') { // quoted-string
		for {
			ch, ok := p.readByte()
			switch ch {
			case '\\':
				ch, ok = p.readByte()
			case '"':
				return sb.String(), nil
			}
			if !ok {
				return "", fmt.Errorf("malformed quoted-string")
			}
			sb.WriteByte(ch)
		}
	} else { // dot-string
		for {
			ch, ok := p.peekByte()
			if !ok {
				return sb.String(), nil
			}
			switch ch {
			case '@':
				return sb.String(), nil
			case '(', ')', '<', '>', '[', ']', ':', ';', '\\', ',', '"', ' ', '\t':
				return "", fmt.Errorf("malformed dot-string")
			}
			p.readByte()
			sb.WriteByte(ch)
		}
	}
}

// cutPrefix is a version of strings.CutPrefix which is case-insensitive.
func (p *Parser) cutPrefix(prefix string) bool {
	if len(p.s) < len(prefix) || !strings.EqualFold(p.s[:len(prefix)], prefix) {
		return false
	}
	p.s = p.s[len(prefix):]
	return true
}
