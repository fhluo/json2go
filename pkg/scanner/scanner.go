package scanner

import (
	"fmt"
	"github.com/fhluo/json2go/pkg/stack"
	"github.com/fhluo/json2go/pkg/token"
	"unicode/utf8"
)

const eof = -1

type Scanner struct {
	buffer []byte

	character     rune
	offset        int
	readingOffset int

	stack *stack.Stack[token.Token]
}

func New(data []byte) *Scanner {
	scanner := new(Scanner)
	scanner.buffer = data
	scanner.stack = stack.New[token.Token]()
	scanner.next()
	return scanner
}

func (s *Scanner) next() {
	if s.readingOffset >= len(s.buffer) {
		s.character = eof
		return
	}

	r, size := rune(s.buffer[s.readingOffset]), 1
	if r >= utf8.RuneSelf {
		r, size = utf8.DecodeRune(s.buffer[s.readingOffset:])
	}

	s.character = r
	s.offset = s.readingOffset
	s.readingOffset += size
}

func (s *Scanner) peek() byte {
	if s.readingOffset >= len(s.buffer) {
		return 0
	}
	return s.buffer[s.readingOffset]
}

func (s *Scanner) skipWhitespace() {
	for s.character == ' ' || s.character == '\n' || s.character == '\r' || s.character == '\t' {
		s.next()
	}
}

func (s *Scanner) scanString() (string, error) {
	s.next()
	start := s.offset

	for s.character != '"' {
		switch s.character {
		case eof:
			return "", fmt.Errorf("")
		case '\\':
			s.next()
			s.next()
		default:
			s.next()
		}
	}

	end := s.offset
	s.next()

	return string(s.buffer[start:end]), nil
}

func (s *Scanner) scanTrue() error {
	if s.character != 't' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'r' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'u' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'e' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	return nil
}

func (s *Scanner) scanFalse() error {
	if s.character != 'f' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'a' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'l' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 's' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'e' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	return nil
}

func (s *Scanner) scanNull() error {
	if s.character != 'n' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'u' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'l' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	if s.character != 'l' {
		return fmt.Errorf("illegal character %c", s.character)
	}
	s.next()

	return nil
}

func (s *Scanner) More() bool {
	s.skipWhitespace()
	if s.character == ',' {
		s.next()
	}
	s.skipWhitespace()

	return s.character != eof && s.character != '}' && s.character != ']'
}

func (s *Scanner) Token() (token.Token, string, error) {
scanAgain:
	s.skipWhitespace()

	switch s.character {
	case '{':
		s.stack.Push(token.LeftBrace)
		s.next()
		return token.LeftBrace, "", nil
	case ':':
		s.next()
		goto scanAgain
	case '}':
		if s.stack.IsEmpty() {
			return token.Null, "", fmt.Errorf("brackets do not match")
		} else {
			if s.stack.Top() == token.LeftBrace {
				s.stack.Pop()
			} else {
				return token.Illegal, "", fmt.Errorf("expecting '%s', got '%s", token.LeftBrace, s.stack.Top())
			}
		}
		s.next()
		return token.RightBrace, "", nil
	case '[':
		s.stack.Push(token.LeftBracket)
		s.next()
		return token.LeftBracket, "", nil
	case ']':
		if s.stack.IsEmpty() {
			return token.Illegal, "", fmt.Errorf("brackets do not match")
		} else {
			if s.stack.Top() == token.LeftBracket {
				s.stack.Pop()
			} else {
				return token.Illegal, "", fmt.Errorf("expecting '%s', got '%s", token.LeftBracket, s.stack.Top())
			}
		}
		s.next()
		return token.RightBracket, "", nil
	case '"':
		literal, err := s.scanString()
		return token.String, literal, err
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		t := token.Int
		s.next()
		for '0' <= s.character && s.character <= '9' {
			s.next()
		}
		if s.character == '.' {
			t = token.Float
			s.next()
			for '0' <= s.character && s.character <= '9' {
				s.next()
			}
		}
		if s.character == 'e' || s.character == 'E' {
			t = token.Float
			s.next()
			if s.character == '+' || s.character == '-' {
				s.next()
			}
			for '0' <= s.character && s.character <= '9' {
				s.next()
			}
		}
		return t, "", nil
	case 't':
		return token.Bool, "", s.scanTrue()
	case 'f':
		return token.Bool, "", s.scanFalse()
	case 'n':
		return token.Null, "", s.scanNull()
	case ',':
		s.next()
		goto scanAgain
	case eof:
		return token.EOF, "", nil
	default:
		return token.Illegal, "", fmt.Errorf("illegal character %c", s.character)
	}
}
