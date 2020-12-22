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

func (s *Stack) topNonOperand() (*Token, int) {
	for i := len(s.tokens) - 1; i >= 0; i-- {
		if s.tokens[i].name != "operand" {
			return s.tokens[i], i
		}
	}
	return nil, -1
}

func (s *Stack) insert(index int, value *Token) {
	if len(s.tokens) == index { // nil or empty slice or after last element
		s.tokens = append(s.tokens, value)
	}
	s.tokens = append(s.tokens[:index+1], s.tokens[index:]...) // index < len(a)
	s.tokens[index] = value
}

func (s *Stack) reduce() {
	token := s.top()
	if token.name == "bracket" && token.value == ")" {
		s.pop() // )
		tokenInsideBrackets := s.pop()
		s.pop() // (
		s.push(tokenInsideBrackets)
	} else {
		operand2 := s.pop()
		operator := s.pop()
		operand1 := s.pop()
		operand1.value = doOperation(operator.value.(string), operand1.value.(int), operand2.value.(int))
		s.push(operand1)
	}
}
