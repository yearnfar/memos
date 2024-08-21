package parser

import (
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
)

type TableParser struct{}

func NewTableParser() *TableParser {
	return &TableParser{}
}

func (*TableParser) Match(tokens []*tokenizer.Token) (ast.Node, int) {
	rawRows := tokenizer.Split(tokens, tokenizer.NewLine)
	if len(rawRows) < 3 {
		return nil, 0
	}

	headerTokens := rawRows[0]
	if len(headerTokens) < 3 {
		return nil, 0
	}

	delimiterTokens := rawRows[1]
	if len(delimiterTokens) < 3 {
		return nil, 0
	}

	// Check header.
	if len(headerTokens) < 5 {
		return nil, 0
	}
	headerCells, ok := matchTableCellTokens(headerTokens)
	if headerCells == 0 || !ok {
		return nil, 0
	}

	// Check delimiter.
	if len(delimiterTokens) < 5 {
		return nil, 0
	}
	delimiterCells, ok := matchTableCellTokens(delimiterTokens)
	if delimiterCells != headerCells || !ok {
		return nil, 0
	}
	for index, t := range tokenizer.Split(delimiterTokens, tokenizer.Pipe) {
		if index == 0 || index == headerCells {
			if len(t) != 0 {
				return nil, 0
			}
			continue
		}
		// Each delimiter cell should be like ` --- `, ` :-- `, ` --: `, ` :-: `.
		if len(t) < 5 {
			return nil, 0
		}

		delimiterTokens := t[1 : len(t)-1]
		if len(delimiterTokens) < 3 {
			return nil, 0
		}
		if (delimiterTokens[0].Type != tokenizer.Colon &&
			delimiterTokens[0].Type != tokenizer.Hyphen) ||
			(delimiterTokens[len(delimiterTokens)-1].Type != tokenizer.Colon &&
				delimiterTokens[len(delimiterTokens)-1].Type != tokenizer.Hyphen) {
			return nil, 0
		}
		for _, token := range delimiterTokens[1 : len(delimiterTokens)-1] {
			if token.Type != tokenizer.Hyphen {
				return nil, 0
			}
		}
	}

	// Check rows.
	rows := rawRows[2:]
	matchedRows := 0
	for _, rowTokens := range rows {
		cells, ok := matchTableCellTokens(rowTokens)
		if cells != headerCells || !ok {
			break
		}
		matchedRows++
	}
	if matchedRows == 0 {
		return nil, 0
	}
	rows = rows[:matchedRows]

	headerNodes := make([]ast.Node, 0)
	delimiter := make([]string, 0)
	rowsNodes := make([][]ast.Node, 0)

	cols := len(tokenizer.Split(headerTokens, tokenizer.Pipe)) - 2
	for _, t := range tokenizer.Split(headerTokens, tokenizer.Pipe)[1 : cols+1] {
		if len(t) < 3 {
			headerNodes = append(headerNodes, &ast.Text{})
		} else {
			cellTokens := t[1 : len(t)-1]
			nodes, err := ParseBlockWithParsers(cellTokens, []BlockParser{NewHeadingParser(), NewParagraphParser()})
			if err != nil {
				return nil, 0
			}
			if len(nodes) != 1 {
				return nil, 0
			}
			headerNodes = append(headerNodes, nodes[0])
		}
	}
	for _, t := range tokenizer.Split(delimiterTokens, tokenizer.Pipe)[1 : cols+1] {
		if len(t) < 3 {
			delimiter = append(delimiter, "")
		} else {
			delimiter = append(delimiter, tokenizer.Stringify(t[1:len(t)-1]))
		}
	}
	for _, row := range rows {
		rowNodes := make([]ast.Node, 0)
		for _, t := range tokenizer.Split(row, tokenizer.Pipe)[1 : cols+1] {
			if len(t) < 3 {
				rowNodes = append(rowNodes, &ast.Text{})
			} else {
				nodes, err := ParseBlockWithParsers(t[1:len(t)-1], []BlockParser{NewHeadingParser(), NewParagraphParser()})
				if err != nil {
					return nil, 0
				}
				if len(nodes) != 1 {
					return nil, 0
				}
				rowNodes = append(rowNodes, nodes[0])
			}
		}
		rowsNodes = append(rowsNodes, rowNodes)
	}

	size := len(headerTokens) + len(delimiterTokens) + 2
	for _, row := range rows {
		size += len(row)
	}
	size = size + len(rows) - 1

	return &ast.Table{
		Header:    headerNodes,
		Delimiter: delimiter,
		Rows:      rowsNodes,
	}, size
}

func matchTableCellTokens(tokens []*tokenizer.Token) (int, bool) {
	if len(tokens) == 0 {
		return 0, false
	}

	pipes := 0
	for _, token := range tokens {
		if token.Type == tokenizer.Pipe {
			pipes++
		}
	}
	cells := tokenizer.Split(tokens, tokenizer.Pipe)
	if len(cells) != pipes+1 {
		return 0, false
	}
	if len(cells[0]) != 0 || len(cells[len(cells)-1]) != 0 {
		return 0, false
	}
	for _, cellTokens := range cells[1 : len(cells)-1] {
		if len(cellTokens) == 0 {
			return 0, false
		}
		if cellTokens[0].Type != tokenizer.Space {
			return 0, false
		}
		if cellTokens[len(cellTokens)-1].Type != tokenizer.Space {
			return 0, false
		}
	}

	return len(cells) - 1, true
}
