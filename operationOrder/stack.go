package main

type Stack struct {
	tokens []*Token
}

func (s *Stack) top() *Token {
	if len(s.tokens) < 1 {
		return nil
	}

	return s.tokens[len(s.tokens)-1]
}

func (s *Stack) len() int {
	return len(s.tokens)
}

func (s *Stack) pop() *Token {
	token := s.top()
	s.tokens = s.tokens[:len(s.tokens)-1]
	return token
}

func (s *Stack) push(token *Token) {
	s.tokens = append(s.tokens, token)
}

func (s *Stack) tokensInBracket() int {
	result := 0
	for i := len(s.tokens) - 1; i >= 0; i-- {
		if s.tokens[i].name == "bracket" {
			break
		}
		result++
	}
	return result
}
