package goesprima

import (
	"math/big"
	"strings"
)

func NewGenerator() *Generator {
	return new(Generator)
}

type Generator struct {
	ModuleName string
	Statements []StatementListItem
}

func (g *Generator) AddStatements(ss ...StatementListItem) *Generator {
	g.Statements = append(g.Statements, ss...)
	return g
}

func (g *Generator) AddStatement(s StatementListItem) *Generator {
	g.Statements = append(g.Statements, s)
	return g
}

func (g *Generator) String() string {
	s := make([]string, len(g.Statements))
	for i, st := range g.Statements {
		s[i] = st.String()
	}
	return strings.Join(s, "\n")
}

// Helper Functions

func StringLiteral(s string) *LiteralValueString {
	l := LiteralValueString(s)
	return &l
}

func BoolLiteral(b bool) *LiteralValueBool {
	l := LiteralValueBool(b)
	return &l
}

func NumberLiteral(n interface{}) Literal {
	switch t := n.(type) {
	case int:
		l := LiteralValueNumber(float64(t))
		return &l
	case *int:
		l := LiteralValueNumber(float64(*t))
		return &l
	case *float64:
		l := LiteralValueNumber(*t)
		return &l
	case float64:
		l := LiteralValueNumber(t)
		return &l
	case big.Float:
		l := LiteralValueBigFloat(t)
		return &l
	case *big.Float:
		l := LiteralValueBigFloat(*t)
		return &l
	default:
		panic("Invalid type passed to NumberLiteral")
	}
}

func jsElementsToString[T JSElement](values []T) []string {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = v.String()
	}
	return out
}

// Indentation

type Indentor interface {
	Indent(string) string
	IndentArray([]string) []string
}

type Spaces struct {
	Spaces int
}

func (sp *Spaces) Indent(s string) string {
	strs := sp.IndentArray(strings.Split(s, "\n"))
	return strings.Join(strs, "\n")
}

func (sp *Spaces) IndentArray(strs []string) []string {
	r := strings.Repeat(" ", sp.Spaces)
	for i, str := range strs {
		strs[i] = r + str
	}
	return strs
}

type Tabs struct {
	Tabs int
}

func (t *Tabs) Indent(s string) string {
	strs := t.IndentArray(strings.Split(s, "\n"))
	return strings.Join(strs, "\n")
}

func (t *Tabs) IndentArray(strs []string) []string {
	r := strings.Repeat("\t", t.Tabs)
	for i, str := range strs {
		strs[i] = r + str
	}
	return strs
}
