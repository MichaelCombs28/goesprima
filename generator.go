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

// Indentation

type Indentor interface {
	Indent(string) string
}

type Spaces struct {
	Spaces int
}

func (sp *Spaces) Indent(s string) string {
	r := strings.Repeat(" ", sp.Spaces)
	strs := strings.Split(s, "\n")
	for i, str := range strs {
		strs[i] = r + str
	}
	return strings.Join(strs, "\n")
}

type Tabs struct {
	Tabs int
}

func (t *Tabs) Indent(s string) string {
	r := strings.Repeat("\t", t.Tabs)
	strs := strings.Split(s, "\n")
	for i, str := range strs {
		strs[i] = r + str
	}
	return strings.Join(strs, "\n")
}
