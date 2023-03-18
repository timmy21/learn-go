package sql

import (
	"errors"
	"strings"
	"text/scanner"
)

//go:generate goyacc -o parser.go parser.y

func Parse(text string) (StmtNode, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	l := &lex{
		s: s,
	}
	_ = yyParse(l)
	return l.result, l.err
}

type lex struct {
	result StmtNode
	err    error
	s      scanner.Scanner
}

func (l *lex) Lex(lval *yySymType) int {
	token := l.s.Scan()
	switch token {
	case scanner.EOF:
		return 0
	case scanner.Ident:
		text := l.s.TokenText()
		v, ok := isKeyword(text)
		if ok {
			return v
		}
		lval.ident = text
		return IDENT
	default:
		return int(l.s.TokenText()[0])
	}
}

func (l *lex) Error(e string) {
	l.err = errors.New(e)
}

func isKeyword(text string) (int, bool) {
	v, ok := keywordMap[strings.ToUpper(text)]
	return v, ok
}

var keywordMap = map[string]int{
	"SELECT": SELECT,
	"FROM":   FROM,
}
