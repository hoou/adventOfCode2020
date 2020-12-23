package main

type Stack struct {
	symbols []string
}

func (s *Stack) top() string {
	if len(s.symbols) < 1 {
		return ""
	}

	return s.symbols[len(s.symbols)-1]
}

func (s *Stack) len() int {
	return len(s.symbols)
}

func (s *Stack) pop() string {
	token := s.top()
	s.symbols = s.symbols[:len(s.symbols)-1]
	return token
}

func (s *Stack) push(symbol string) {
	s.symbols = append(s.symbols, symbol)
}
