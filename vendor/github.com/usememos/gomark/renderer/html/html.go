package html

import (
	"bytes"
	"fmt"

	"github.com/usememos/gomark/ast"
)

type RendererContext struct {
}

// HTMLRenderer is a simple renderer that converts AST to HTML.
type HTMLRenderer struct {
	output  *bytes.Buffer
	context *RendererContext
}

// NewHTMLRenderer creates a new HTMLRender.
func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{
		output:  new(bytes.Buffer),
		context: &RendererContext{},
	}
}

// RenderNode renders a single AST node to HTML.
func (r *HTMLRenderer) RenderNode(node ast.Node) {
	switch n := node.(type) {
	case *ast.LineBreak:
		r.renderLineBreak(n)
	case *ast.Paragraph:
		r.renderParagraph(n)
	case *ast.CodeBlock:
		r.renderCodeBlock(n)
	case *ast.Heading:
		r.renderHeading(n)
	case *ast.HorizontalRule:
		r.renderHorizontalRule(n)
	case *ast.Blockquote:
		r.renderBlockquote(n)
	case *ast.UnorderedList:
		r.renderUnorderedList(n)
	case *ast.OrderedList:
		r.renderOrderedList(n)
	case *ast.TaskList:
		r.renderTaskList(n)
	case *ast.MathBlock:
		r.renderMathBlock(n)
	case *ast.Table:
		r.renderTable(n)
	case *ast.EmbeddedContent:
		r.renderEmbeddedContent(n)
	case *ast.Text:
		r.renderText(n)
	case *ast.Bold:
		r.renderBold(n)
	case *ast.Italic:
		r.renderItalic(n)
	case *ast.BoldItalic:
		r.renderBoldItalic(n)
	case *ast.Code:
		r.renderCode(n)
	case *ast.Image:
		r.renderImage(n)
	case *ast.Link:
		r.renderLink(n)
	case *ast.AutoLink:
		r.renderAutoLink(n)
	case *ast.Tag:
		r.renderTag(n)
	case *ast.Strikethrough:
		r.renderStrikethrough(n)
	case *ast.EscapingCharacter:
		r.renderEscapingCharacter(n)
	case *ast.Math:
		r.renderMath(n)
	case *ast.Highlight:
		r.renderHighlight(n)
	case *ast.Subscript:
		r.renderSubscript(n)
	case *ast.Superscript:
		r.renderSuperscript(n)
	case *ast.ReferencedContent:
		r.renderReferencedContent(n)
	case *ast.Spoiler:
		r.renderSpoiler(n)
	case *ast.HTMLElement:
		r.renderHTMLElement(n)
	default:
		// Handle other block types if needed.
	}
}

// RenderNodes renders a slice of AST nodes to HTML.
func (r *HTMLRenderer) RenderNodes(nodes []ast.Node) {
	var prevNode ast.Node
	var skipNextLineBreakFlag bool
	for _, node := range nodes {
		if node.Type() == ast.LineBreakNode && skipNextLineBreakFlag {
			if prevNode != nil && ast.IsBlockNode(prevNode) {
				skipNextLineBreakFlag = false
				continue
			}
		}

		r.RenderNode(node)
		prevNode = node
		skipNextLineBreakFlag = true
	}
}

// Render renders the AST to HTML.
func (r *HTMLRenderer) Render(astRoot []ast.Node) string {
	r.RenderNodes(astRoot)
	return r.output.String()
}

func (r *HTMLRenderer) renderLineBreak(*ast.LineBreak) {
	r.output.WriteString("<br>")
}

func (r *HTMLRenderer) renderParagraph(node *ast.Paragraph) {
	r.output.WriteString("<p>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</p>")
}

func (r *HTMLRenderer) renderCodeBlock(node *ast.CodeBlock) {
	r.output.WriteString("<pre><code>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</code></pre>")
}

func (r *HTMLRenderer) renderHeading(node *ast.Heading) {
	element := fmt.Sprintf("h%d", node.Level)
	r.output.WriteString(fmt.Sprintf("<%s>", element))
	r.RenderNodes(node.Children)
	r.output.WriteString(fmt.Sprintf("</%s>", element))
}

func (r *HTMLRenderer) renderHorizontalRule(_ *ast.HorizontalRule) {
	r.output.WriteString("<hr>")
}

func (r *HTMLRenderer) renderBlockquote(node *ast.Blockquote) {
	r.output.WriteString("<blockquote>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</blockquote>")
}

func (r *HTMLRenderer) renderUnorderedList(node *ast.UnorderedList) {
	prevSibling, nextSibling := ast.FindPrevSiblingExceptLineBreak(node), ast.FindNextSiblingExceptLineBreak(node)
	if prevSibling == nil || prevSibling.Type() != ast.UnorderedListNode {
		r.output.WriteString("<ul>")
	}
	r.output.WriteString("<li>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</li>")
	if nextSibling == nil || nextSibling.Type() != ast.UnorderedListNode {
		r.output.WriteString("</ul>")
	}
}

func (r *HTMLRenderer) renderOrderedList(node *ast.OrderedList) {
	prevSibling, nextSibling := ast.FindPrevSiblingExceptLineBreak(node), ast.FindNextSiblingExceptLineBreak(node)
	if prevSibling == nil || prevSibling.Type() != ast.OrderedListNode {
		r.output.WriteString("<ol>")
	}
	r.output.WriteString("<li>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</li>")
	if nextSibling == nil || nextSibling.Type() != ast.OrderedListNode {
		r.output.WriteString("</ol>")
	}
}

func (r *HTMLRenderer) renderTaskList(node *ast.TaskList) {
	prevSibling, nextSibling := ast.FindPrevSiblingExceptLineBreak(node), ast.FindNextSiblingExceptLineBreak(node)
	if prevSibling == nil || prevSibling.Type() != ast.TaskListNode {
		r.output.WriteString("<ul>")
	}
	r.output.WriteString("<li>")
	r.output.WriteString("<input type=\"checkbox\"")
	if node.Complete {
		r.output.WriteString(" checked")
	}
	r.output.WriteString(" disabled>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</li>")
	if nextSibling == nil || nextSibling.Type() != ast.TaskListNode {
		r.output.WriteString("</ul>")
	}
}

func (r *HTMLRenderer) renderMathBlock(node *ast.MathBlock) {
	r.output.WriteString("<pre><code>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</code></pre>")
}

func (r *HTMLRenderer) renderTable(node *ast.Table) {
	r.output.WriteString("<table>")
	r.output.WriteString("<thead>")
	r.output.WriteString("<tr>")
	for _, cell := range node.Header {
		r.output.WriteString("<th>")
		r.RenderNodes([]ast.Node{cell})
		r.output.WriteString("</th>")
	}
	r.output.WriteString("</tr>")
	r.output.WriteString("</thead>")
	r.output.WriteString("<tbody>")
	for _, row := range node.Rows {
		r.output.WriteString("<tr>")
		for _, cell := range row {
			r.output.WriteString("<td>")
			r.RenderNodes([]ast.Node{cell})
			r.output.WriteString("</td>")
		}
		r.output.WriteString("</tr>")
	}
	r.output.WriteString("</tbody>")
	r.output.WriteString("</table>")
}

func (r *HTMLRenderer) renderEmbeddedContent(node *ast.EmbeddedContent) {
	r.output.WriteString("<div>")
	r.output.WriteString(node.ResourceName)
	if node.Params != "" {
		r.output.WriteString("?")
		r.output.WriteString(node.Params)
	}
	r.output.WriteString("</div>")
}

func (r *HTMLRenderer) renderText(node *ast.Text) {
	r.output.WriteString(node.Content)
}

func (r *HTMLRenderer) renderBold(node *ast.Bold) {
	r.output.WriteString("<strong>")
	r.RenderNodes(node.Children)
	r.output.WriteString("</strong>")
}

func (r *HTMLRenderer) renderItalic(node *ast.Italic) {
	r.output.WriteString("<em>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</em>")
}

func (r *HTMLRenderer) renderBoldItalic(node *ast.BoldItalic) {
	r.output.WriteString("<strong><em>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</em></strong>")
}

func (r *HTMLRenderer) renderCode(node *ast.Code) {
	r.output.WriteString("<code>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</code>")
}

func (r *HTMLRenderer) renderImage(node *ast.Image) {
	r.output.WriteString(`<img src="`)
	r.output.WriteString(node.URL)
	r.output.WriteString(`" alt="`)
	r.output.WriteString(node.AltText)
	r.output.WriteString(`" />`)
}

func (r *HTMLRenderer) renderLink(node *ast.Link) {
	r.output.WriteString(`<a href="`)
	r.output.WriteString(node.URL)
	r.output.WriteString(`">`)
	r.output.WriteString(node.Text)
	r.output.WriteString("</a>")
}

func (r *HTMLRenderer) renderAutoLink(node *ast.AutoLink) {
	r.output.WriteString(`<a href="`)
	r.output.WriteString(node.URL)
	r.output.WriteString(`">`)
	r.output.WriteString(node.URL)
	r.output.WriteString("</a>")
}

func (r *HTMLRenderer) renderTag(node *ast.Tag) {
	r.output.WriteString(`<span>`)
	r.output.WriteString(`#`)
	r.output.WriteString(node.Content)
	r.output.WriteString(`</span>`)
}

func (r *HTMLRenderer) renderStrikethrough(node *ast.Strikethrough) {
	r.output.WriteString(`<del>`)
	r.output.WriteString(node.Content)
	r.output.WriteString(`</del>`)
}

func (r *HTMLRenderer) renderEscapingCharacter(node *ast.EscapingCharacter) {
	r.output.WriteString("\\")
	r.output.WriteString(node.Symbol)
}

func (r *HTMLRenderer) renderMath(node *ast.Math) {
	r.output.WriteString("<code>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</code>")
}

func (r *HTMLRenderer) renderHighlight(node *ast.Highlight) {
	r.output.WriteString("<mark>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</mark>")
}

func (r *HTMLRenderer) renderSubscript(node *ast.Subscript) {
	r.output.WriteString("<sub>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</sub>")
}

func (r *HTMLRenderer) renderSuperscript(node *ast.Superscript) {
	r.output.WriteString("<sup>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</sup>")
}

func (r *HTMLRenderer) renderReferencedContent(node *ast.ReferencedContent) {
	r.output.WriteString("<div>")
	r.output.WriteString(node.ResourceName)
	if node.Params != "" {
		r.output.WriteString("?")
		r.output.WriteString(node.Params)
	}
	r.output.WriteString("</div>")
}

func (r *HTMLRenderer) renderSpoiler(node *ast.Spoiler) {
	r.output.WriteString("<details><summary>")
	r.output.WriteString(node.Content)
	r.output.WriteString("</summary></details>")
}

func (r *HTMLRenderer) renderHTMLElement(node *ast.HTMLElement) {
	r.output.WriteString(fmt.Sprintf("<%s >", node.TagName))
}
