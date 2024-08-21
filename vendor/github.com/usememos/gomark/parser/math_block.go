package parser

import (
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
)

type MathBlockParser struct{}

func NewMathBlockParser() *MathBlockParser {
	return &MathBlockParser{}
}

func (*MathBlockParser) Match(tokens []*tokenizer.Token) (ast.Node, int) {
	rows := tokenizer.Split(tokens, tokenizer.NewLine)
	if len(rows) < 3 {
		return nil, 0
	}
	firstRow := rows[0]
	if len(firstRow) != 2 {
		return nil, 0
	}
	if firstRow[0].Type != tokenizer.DollarSign || firstRow[1].Type != tokenizer.DollarSign {
		return nil, 0
	}

	contentRows := [][]*tokenizer.Token{}
	matched := false
	for _, row := range rows[1:] {
		if len(row) == 2 && row[0].Type == tokenizer.DollarSign && row[1].Type == tokenizer.DollarSign {
			matched = true
			break
		}
		contentRows = append(contentRows, row)
	}
	if !matched {
		return nil, 0
	}

	contentTokens := []*tokenizer.Token{}
	for index, row := range contentRows {
		contentTokens = append(contentTokens, row...)
		if index != len(contentRows)-1 {
			contentTokens = append(contentTokens, &tokenizer.Token{
				Type:  tokenizer.NewLine,
				Value: "\n",
			})
		}
	}
	return &ast.MathBlock{
		Content: tokenizer.Stringify(contentTokens),
	}, 3 + len(contentTokens) + 3
}
