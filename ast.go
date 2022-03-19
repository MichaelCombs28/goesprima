package goesprima

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
)

var (
	indentor Indentor = &Spaces{2}
	mu       sync.Mutex
)

func SetGlobaIndentor(i Indentor) {
	mu.Lock()
	defer mu.Unlock()
	indentor = i
}

type JSElement interface {
	String() string
}

// Compiler checks
var (
	// ArgumentListElement
	_ ArgumentListElement = new(SpreadElement)

	// Binding Patterns
	_ BindingPattern = new(ArrayPattern)
	_ BindingPattern = new(ObjectPattern)

	// LiteralValues
	LiteralValueUndefined Literal = &literalValueUndefined{}
	LiteralValueNull      Literal = &literalValueNull{}
	_                     Literal = new(LiteralValueString)
	_                     Literal = new(LiteralValueBool)
	_                     Literal = new(LiteralValueNumber)
	_                     Literal = new(LiteralValueBigFloat)

	// Expressions
	_ Expression = new(ArrayExpression)
	_ Expression = new(ArrowFunctionExpression)
	_ Expression = new(AssignmentExpression)
	_ Expression = new(AwaitExpression)
	_ Expression = new(BinaryExpression)
	_ Expression = new(LogicalExpression)
	_ Expression = new(CallExpression)
	_ Expression = new(ChainExpression)
	_ Expression = new(ClassExpression)
	_ Expression = new(ComputedMemberExpression)
	_ Expression = new(ConditionalExpression)
	_ Expression = new(Identifier)
	_ Expression = new(FunctionExpression)
	_ Expression = new(NewExpression)
	_ Expression = new(ObjectExpression)
	_ Expression = new(SequenceExpression)
	_ Expression = new(StaticMemberExpression)
	_ Expression = new(TaggedTemplateExpression)
	_ Expression = new(UnaryExpression)
	_ Expression = new(UpdateExpression)
	_ Expression = new(YieldExpression)

	// Declarations
	_ Declaration = new(ClassDeclaration)
	_ Declaration = new(FunctionDeclaration)
	_ Declaration = new(ImportDeclaration)
	_ Declaration = new(VariableDeclaration)

	// Statements
	_ Statement = new(BreakStatement)
	_ Statement = new(ContinueStatement)
	_ Statement = new(DebuggerStatement)
	_ Statement = new(DoWhileStatement)
	_ Statement = new(EmptyStatement)
	_ Statement = new(ExpressionStatement)
	_ Statement = new(Directive)
	_ Statement = new(ForStatement)
	_ Statement = new(ForInStatement)
	_ Statement = new(ForOfStatement)
	_ Statement = new(IfStatement)
	_ Statement = new(ReturnStatement)
	_ Statement = new(SwitchStatement)
	_ Statement = new(ThrowStatement)
	_ Statement = new(TryStatement)
	_ Statement = new(WhileStatement)
	_ Statement = new(WithStatement)
	_ Statement = new(BlockStatement)

	// StatementListItems
	_ StatementListItem = new(ClassDeclaration)
	_ StatementListItem = new(FunctionDeclaration)
	_ StatementListItem = new(ImportDeclaration)
	_ StatementListItem = new(VariableDeclaration)
	_ StatementListItem = new(BreakStatement)
	_ StatementListItem = new(ContinueStatement)
	_ StatementListItem = new(DebuggerStatement)
	_ StatementListItem = new(DoWhileStatement)
	_ StatementListItem = new(EmptyStatement)
	_ StatementListItem = new(ExpressionStatement)
	_ StatementListItem = new(Directive)
	_ StatementListItem = new(ForStatement)
	_ StatementListItem = new(ForInStatement)
	_ StatementListItem = new(ForOfStatement)
	_ StatementListItem = new(IfStatement)
	_ StatementListItem = new(ReturnStatement)
	_ StatementListItem = new(SwitchStatement)
	_ StatementListItem = new(ThrowStatement)
	_ StatementListItem = new(TryStatement)
	_ StatementListItem = new(WhileStatement)
	_ StatementListItem = new(WithStatement)
	_ StatementListItem = new(BlockStatement)

	// ArrayPatternElements
	_ ArrayPatternElement = new(AssignmentPattern)
	_ ArrayPatternElement = new(Identifier)
	_ ArrayPatternElement = new(ArrayPattern)
	_ ArrayPatternElement = new(ObjectPattern)
	_ ArrayPatternElement = new(RestElement)

	// ChainElement
	_ ChainElement = new(CallExpression)
	_ ChainElement = new(ComputedMemberExpression)
	_ ChainElement = new(StaticMemberExpression)

	// ExportableDefaultDeclaration
	_ ExportableDefaultDeclaration = new(Identifier)
	_ ExportableDefaultDeclaration = new(ClassDeclaration)
	_ ExportableDefaultDeclaration = new(FunctionDeclaration)

	// ExportableNamedDeclaration
	_ ExportableNamedDeclaration = new(ClassDeclaration)
	_ ExportableNamedDeclaration = new(FunctionDeclaration)
	_ ExportableNamedDeclaration = new(VariableDeclaration)

	// FunctionParameter
	_ FunctionParameter = new(Identifier)

	// ImportDeclarationSpecifiers
	_ ImportDeclarationSpecifier = new(ImportDefaultSpecifier)
	_ ImportDeclarationSpecifier = new(ImportNamespaceSpecifier)
	_ ImportDeclarationSpecifier = new(ImportSpecifier)

	// ObjectExpressionProperties
	_ ObjectExpressionProperty = new(Property)
	_ ObjectExpressionProperty = new(SpreadElement)

	// ObjectPatternProperties
	_ ObjectPatternProperty = new(PropertyPattern)
	_ ObjectPatternProperty = new(RestElement)

	// PropertyKey
	_ PropertyKey = new(Identifier)

	// PropertyValue
	_ PropertyValue = new(AssignmentPattern)
	_ PropertyValue = new(FunctionExpression)
	_ PropertyValue = new(Identifier)

	// ExportDeclarations
	_ ExportDeclaration = new(ExportAllDeclaration)
	_ ExportDeclaration = new(ExportDefaultDeclaration)
	_ ExportDeclaration = new(ExportNamedDeclaration)

	// ClassProperty
	_ ClassProperty = new(MethodDefinition)
	_ ClassProperty = new(PropertyDefinition)
)

// Interfaces

type ArgumentListElement interface {
	JSElement
	argumentListElement()
}

type ArrayExpressionElement interface {
	JSElement
	arrayExpressionElement()
}

type ExportDeclaration interface {
	exportDeclaration()
	Declaration
}

type Literal interface {
	JSElement
	literal()

	PropertyKey
	ArgumentListElement
	ArrayExpressionElement
	Expression
	ExpressionOrImport
	ExportableDefaultDeclaration
}

// Deconstruction patterns eg. ({one, two}) => ...
type BindingPattern interface {
	bindingPattern()

	JSElement
	ExportableDefaultDeclaration
	BindingIdentifierOrPattern
	FunctionParameter
	PropertyValue
}

type BindingIdentifierOrPattern interface {
	JSElement
	bindingIdentifierOrPattern()
}

type Expression interface {
	expression()

	JSElement
	ExportableDefaultDeclaration
	ArgumentListElement
	ArrayExpressionElement
	ExpressionOrImport
}

type Declaration interface {
	declaration()

	JSElement
	StatementListItem
}

type Statement interface {
	JSElement
	statement()
}

type StatementListItem interface {
	JSElement
	statementListItem()
}

type ArrayPatternElement interface {
	JSElement
	arrayPatternElement()
}

type ObjectPatternProperty interface {
	JSElement
	objectPatternProperty()
}

type PropertyKey interface {
	JSElement
	propertyKey()
}

type PropertyValue interface {
	JSElement
	propertyValue()
}

type FunctionParameter interface {
	JSElement
	functionParameter()
}

type ChainElement interface {
	JSElement
	chainElementToString() string
}

type ExportableDefaultDeclaration interface {
	JSElement
	exportableDefaultDeclaration()
}

type ClassProperty interface {
	JSElement
	classProperty()
}

type ExportableNamedDeclaration interface {
	JSElement
	exportableNamedDeclaration()
}

type ImportDeclarationSpecifier interface {
	JSElement
	importDeclarationSpecifier()
}

type ObjectExpressionProperty interface {
	JSElement
	objectExpressionProperty()
}

type ExpressionOrImport interface {
	JSElement
	expressionOrImport()
}

// Structs

type ExportAllDeclaration struct {
	Source Literal
	*Node
}

func (e *ExportAllDeclaration) String() string {
	panic("ExportAllDeclaration not implemented")
}

type ExportDefaultDeclaration struct {
	Declaration ExportableDefaultDeclaration
	*Node
}

func (e *ExportDefaultDeclaration) String() string {
	return "export default " + e.Declaration.String()
}

type ExportNamedDeclaration struct {
	Declaration ExportableNamedDeclaration
	Specifiers  []ExportSpecifier
	*Node
}

func (e *ExportNamedDeclaration) String() (s string) {
	s = "export "
	if e.Declaration != nil {
		s += e.Declaration.String()
	} else {
		spec := jsElementsToString(e.Specifiers)

		s += "{ " + strings.Join(spec, ", ") + " }"
	}
	s += ";"
	return
}

type ExportSpecifier struct {
	Exported *Identifier
	Local    *Identifier
	*Node
}

func (e ExportSpecifier) String() (s string) {
	s = e.Exported.String()
	if e.Local != nil {
		s += " as " + e.Local.String()
	}
	return
}

type BlockStatement struct {
	Items []Statement
	*Node
}

func (b *BlockStatement) String() string {
	r := jsElementsToString(b.Items)
	return strings.Join(r, "\n")
}

type ArrayPattern struct {
	Elements []ArrayPatternElement
	*Node
}

func (a *ArrayPattern) String() (s string) {
	if len(a.Elements) == 0 {
		return "[]"
	}
	s = "[\n"
	r := jsElementsToString(a.Elements)
	s += indentor.Indent(strings.Join(r, ", ")) + "\n]"
	return
}

type ObjectPattern struct {
	Properties []ObjectPatternProperty
	*Node
}

func (o *ObjectPattern) String() (s string) {
	l := len(o.Properties)
	if l == 0 {
		return "{}"
	}
	s = "{\n"
	props := indentor.IndentArray(jsElementsToString(o.Properties))
	s += strings.Join(props, ",\n") + ",\n}"
	return
}

type Identifier struct {
	Name string
	*Node
}

func (i *Identifier) String() string {
	return i.Name
}

type AssignmentPattern struct {
	Left  BindingIdentifierOrPattern
	Right Expression
}

func (a *AssignmentPattern) String() string {
	return a.Left.String() + "=" + a.Right.String()
}

type literalValueNull struct{}

func (l *literalValueNull) String() string {
	return "null"
}

type literalValueUndefined struct{}

func (l *literalValueUndefined) String() string {
	return "undefined"
}

type LiteralValueString string

func (l *LiteralValueString) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

type LiteralValueBool bool

func (l *LiteralValueBool) String() string {
	return strconv.FormatBool(bool(*l))
}

type LiteralValueNumber float64

func (l *LiteralValueNumber) String() string {
	return strconv.FormatFloat(float64(*l), 'f', 6, 64)
}

type LiteralValueBigFloat big.Float

func (l *LiteralValueBigFloat) String() string {
	n := big.Float(*l)
	return n.String()
}

// Expressions
type ArrayExpression struct {
	Elements []ArrayExpressionElement
	*Node
}

func (a *ArrayExpression) String() (s string) {
	if len(a.Elements) == 0 {
		return "[]"
	}
	s = "[\n"
	r := jsElementsToString(a.Elements)
	s += indentor.Indent(strings.Join(r, ", ")) + ",\n]"
	return
}

type ArrowFunctionExpression struct {
	Params []FunctionParameter
	Body   BlockStatement
	Async  bool
	*Node
}

func (a *ArrowFunctionExpression) String() (s string) {
	if a.Async {
		s = "async "
	}

	params := jsElementsToString(a.Params)
	s += "(" + strings.Join(params, ", ") + ") => {\n" + indentor.Indent((&a.Body).String()) + "\n}"
	return
}

type AwaitExpression struct {
	Arguement Expression
	*Node
}

func (a *AwaitExpression) String() string {
	return "await " + a.Arguement.String()
}

type AssignmentExpression struct {
	Operator assignmentOperator
	Left     Expression
	Right    Expression
	*Node
}

func (a *AssignmentExpression) String() string {
	return a.Left.String() + " " + string(a.Operator) + " " + a.Right.String()
}

type assignmentOperator string

const (
	AssignmentOperatorEq     assignmentOperator = "="
	AssignmentOperatorPlus   assignmentOperator = "+="
	AssignmentOperatorMinus  assignmentOperator = "-="
	AssignmentOperatorTimes  assignmentOperator = "*="
	AssignmentOperatorDivide assignmentOperator = "/="
	AssignmentOperatorMod    assignmentOperator = "%="
)

type BinaryExpression struct {
	Operator binaryOperator
	Left     Expression
	Right    Expression
	*Node
}

func (b *BinaryExpression) String() string {
	return b.Left.String() + " " + string(b.Operator) + " " + b.Right.String()
}

type binaryOperator string

const (
	BinaryOperatorADD                binaryOperator = "+"
	BinaryOperatorMinus              binaryOperator = "-"
	BinaryOperatorMultiply           binaryOperator = "*"
	BinaryOperatorExponent           binaryOperator = "*"
	BinaryOperatorDivide             binaryOperator = "/"
	BinaryOperatorModulus            binaryOperator = "%"
	BinaryOperatorAND                binaryOperator = "&"
	BinaryOperatorOR                 binaryOperator = "|"
	BinaryOperatorXOR                binaryOperator = "^"
	BinaryOperatorNOT                binaryOperator = "~"
	BinaryOperatorSHIFTLEFT          binaryOperator = "<<"
	BinaryOperatorSHIFTRIGHT         binaryOperator = ">>"
	BinaryOperatorZEROFILLSHIFTRIGHT binaryOperator = ">>>"
)

type LogicalExpression struct {
	Operator logicalOperator
	Left     Expression
	Right    Expression
	*Node
}

func (l *LogicalExpression) String() string {
	return l.Left.String() + " " + string(l.Operator) + " " + l.Right.String()
}

type logicalOperator string

const (
	LogicalOperatorOr                logicalOperator = "||"
	LogicalOperatorAnd               logicalOperator = "&&"
	LogicalOperatorNullishCoelescing logicalOperator = "??"
)

type CallExpression struct {
	Callee    Expression
	Arguments []ArgumentListElement
	Optional  bool
	*Node
}

func (c *CallExpression) String() string {
	return c.Callee.String() + c.argsToString()
}

func (c *CallExpression) argsToString() string {
	return "(" + argListToString(c.Arguments) + ")"
}

type CatchClause struct {
	BindingIdentifierOrPattern
	Body BlockStatement
	*Node
}

func (c CatchClause) String() string {
	return "catch (" + c.BindingIdentifierOrPattern.String() + ") {\n" + indentor.Indent(c.Body.String()) + "\n}"
}

type Import struct {
	*Node
}

type ChainExpression struct {
	Expression ChainElement
	*Node
}

func (c *ChainExpression) String() (s string) {
	return c.Expression.chainElementToString()
}

type ClassExpression struct {
	ID         *Identifier
	SuperClass *Identifier
	Body       *ClassBody
	*Node
}

func (c *ClassExpression) String() (s string) {
	s = "class "

	if c.ID != nil {
		s += c.ID.String() + " "
	}

	if c.SuperClass != nil {
		s += "extends " + c.SuperClass.String() + " "
	}
	s += "{\n" + indentor.Indent(c.Body.String()) + "\n}"
	return
}

type ComputedMemberExpression struct {
	Object   Expression
	Property Expression
	Optional bool
	*Node
}

func (c *ComputedMemberExpression) String() string {
	return c.Object.String() + "[" + c.Property.String() + "]"
}

type ConditionalExpression struct {
	Test       Expression
	Consequent Expression
	Alternate  Expression
	*Node
}

func (c *ConditionalExpression) String() string {
	return c.Test.String() + "? " + c.Consequent.String() + ": " + c.Alternate.String()
}

type FunctionExpression struct {
	ID     *Identifier
	Params []FunctionParameter
	Body   BlockStatement
	FunctionType
	*Node
}

func (f *FunctionExpression) String() (s string) {
	switch f.FunctionType {
	case FunctionTypeAsync:
		s = "async function "
	case FunctionTypeGenerator:
		s = "function* "
	default:
		s = "function "
	}
	if f.ID != nil {
		s += f.ID.String()
	}
	s += "(" + functionParametersToString(f.Params) + ") {\n" + f.Body.String() + "\n}"
	return
}

type NewExpression struct {
	Callee    Expression
	Arguments []ArgumentListElement
	*Node
}

func (n *NewExpression) String() string {
	return "new " + n.calleeToString() + "(" + argListToString(n.Arguments) + ")"
}

func (n *NewExpression) calleeToString() string {
	switch v := n.Callee.(type) {
	case *Identifier:
		return v.String()
	case *StaticMemberExpression:
		return v.String()
	default:
		return "(" + n.Callee.String() + ")"
	}
}

type ObjectExpression struct {
	Properties []ObjectExpressionProperty
	*Node
}

func (o *ObjectExpression) String() string {
	args := objectExpressionPropertiesToString(o.Properties)
	return "{\n" + indentor.Indent(args) + "\n}"
}

type SequenceExpression struct {
	Expressions []Expression
	*Node
}

func (s *SequenceExpression) String() string {
	out := jsElementsToString(s.Expressions)
	return strings.Join(out, ", ")
}

type StaticMemberExpression struct {
	Object   Expression
	Property Expression
	*Node
}

func (s *StaticMemberExpression) String() string {
	return s.objectToString() + "." + s.propertyToString()
}

func (s *StaticMemberExpression) objectToString() string {
	switch v := s.Object.(type) {
	case *CallExpression:
		return v.String()
	case *Identifier:
		return v.String()
	case *StaticMemberExpression:
		return v.String()
	case *ComputedMemberExpression:
		return v.String()
	}
	return "(" + s.Object.String() + ")"
}

func (s *StaticMemberExpression) propertyToString() string {
	switch v := s.Property.(type) {
	case *Identifier:
		return v.Name
	case *StaticMemberExpression:
		return v.String()
	case *CallExpression:
		return v.String()
	default:
		return "(" + v.String() + ")"
	}
}

type SwitchCase struct {
	Test       Expression
	Consequent BlockStatement
}

func (s SwitchCase) String() string {
	return "case " + s.Test.String() + ":\n" + indentor.Indent(s.Consequent.String())
}

type TaggedTemplateExpression struct {
	Tag   Expression
	Quasi TemplateLiteral
	*Node
}

func (t *TaggedTemplateExpression) String() string {
	panic("Not Implemented")
}

type TemplateLiteral struct {
	Quasis      []TemplateElement
	Expressions []Expression
	*Node
}

type TemplateElement struct {
}

type UnaryExpression struct {
	Operator UnaryOperatorType
	Argument Expression
	*Node
}

func (u *UnaryExpression) String() string {
	var expr string
	switch v := u.Argument.(type) {
	case *Identifier:
		expr = v.String()
	case *StaticMemberExpression:
		expr = v.String()

	default:
		expr = "(" + v.String() + ")"
	}
	return fmt.Sprintf(string(u.Operator), expr)
}

type UnaryOperatorType string

const (
	UnaryOperatorTypePlus             UnaryOperatorType = "+%s"
	UnaryOperatorTypeMinus            UnaryOperatorType = "-%s"
	UnaryOperatorTypeIncrementPrefix  UnaryOperatorType = "++%s"
	UnaryOperatorTypeIncrementPostfix UnaryOperatorType = "%s++"
	UnaryOperatorTypeDecrementPrefix  UnaryOperatorType = "--%s"
	UnaryOperatorTypeDecrementPostfix UnaryOperatorType = "%s--"
)

type UpdateExpression struct {
	Argument Expression
	*Node
}

func (u *UpdateExpression) String() string {
	return u.Argument.String()
}

type YieldExpression struct {
	Argument Expression
	Delegate bool
	*Node
}

func (y *YieldExpression) String() (s string) {
	s = "yield"
	if y.Delegate {
		s += "*"
	}
	if y.Argument != nil {
		s += " " + y.Argument.String()
	}
	return
}

// Declarations
type ClassDeclaration struct {
	ID         *Identifier
	SuperClass Expression
	Body       *ClassBody
	*Node
}

func (c *ClassDeclaration) String() (s string) {
	s = "class "

	if c.ID != nil {
		s += c.ID.String() + " "
	}

	if c.SuperClass != nil {
		s += superClassToString(c.SuperClass) + " "
	}
	s += "{\n" + indentor.Indent(c.Body.String()) + "\n}"
	return
}

func superClassToString(exp Expression) (s string) {
	s = "extends "
	var computed bool
	switch exp.(type) {
	case *BinaryExpression:
		computed = true
	case *LogicalExpression:
		computed = true
	case Literal:
		computed = true
	}
	if computed {
		s += "("
	}
	s += exp.String()
	if computed {
		s += ")"
	}
	return
}

type ClassBody struct {
	Properties []ClassProperty
	*Node
}

func (c *ClassBody) String() string {
	props := jsElementsToString(c.Properties)
	return strings.Join(props, "\n")
}

type MethodDefinition struct {
	Static bool
	Key    PropertyKey
	Value  FunctionExpression
	*Node
}

func (m *MethodDefinition) String() (s string) {
	if m.Static {
		s = "static "
	}

	switch m.Value.FunctionType {
	case FunctionTypeAsync:
		s = "async "
	case FunctionTypeGenerator:
		s = "*"
	}
	s += m.keyToString() + m.valueToString()
	return
}

func (m *MethodDefinition) valueToString() (s string) {
	s += "(" + functionParametersToString(m.Value.Params) + ") {\n" + m.Value.Body.String() + "\n}"
	return
}

func (m *MethodDefinition) keyToString() string {
	switch v := m.Key.(type) {
	case *Identifier:
		return v.String()
	default:
		return "[" + v.String() + "]"
	}
}

type PropertyDefinition struct {
	Static bool
	Key    PropertyKey
	Value  Expression
	*Node
}

func (p *PropertyDefinition) String() (s string) {
	if p.Static {
		s = "static "
	}
	s += p.keyToString() + " = " + p.Value.String() + ";"
	return
}

func (p *PropertyDefinition) keyToString() string {
	switch v := p.Key.(type) {
	case *Identifier:
		return v.String()
	default:
		return "[" + v.String() + "]"
	}
}

type PropertyPattern struct {
	Key       PropertyKey
	Computed  bool
	Value     PropertyValue
	Kind      string
	Method    bool
	ShortHand bool
	*Node
}

func (p *PropertyPattern) String() (s string) {
	if p.Computed {
		s += "["
	}
	s += p.keyToString()
	if p.Computed {
		s += "]"
	}

	if p.Value != nil {
		s += p.Value.String()
	}
	return
}

func (p *PropertyPattern) keyToString() string {
	switch v := p.Key.(type) {
	case *Identifier:
		return v.String()
	default:
		return "[" + v.String() + "]"
	}
}

type Property struct {
	Key       PropertyKey
	Value     Expression
	Kind      string
	Method    bool
	ShortHand bool
	*Node
}

func (p *Property) String() (s string) {
	s = p.keyToString()

	if p.Value != nil {
		s += ": " + p.Value.String()
	}
	return
}

func (p *Property) keyToString() string {
	switch v := p.Key.(type) {
	case *Identifier:
		return v.String()
	default:
		return "[" + v.String() + "]"
	}
}

type FunctionDeclaration struct {
	ID     *Identifier
	Params []FunctionParameter
	Body   BlockStatement
	FunctionType
	*Node
}

func (f *FunctionDeclaration) String() (s string) {
	switch f.FunctionType {
	case FunctionTypeAsync:
		s = "async function "
	case FunctionTypeGenerator:
		s = "function* "
	default:
		s = "function "
	}
	if f.ID != nil {
		s += f.ID.String()
	}
	s += "(" + functionParametersToString(f.Params) + ") {\n" + f.Body.String() + "\n}"
	return
}

type FunctionType string

const (
	FunctionTypeNormal    FunctionType = "normal"
	FunctionTypeAsync     FunctionType = "async"
	FunctionTypeGenerator FunctionType = "generator"
)

type ImportDeclaration struct {
	Specifiers []ImportDeclarationSpecifier
	Source     string
	*Node
}

func (i *ImportDeclaration) String() (s string) {
	s = "import "
	if l := len(i.Specifiers); l > 0 {
		var defaultImport string
		var namespaceImport string
		var namedImports []string

		for _, spec := range i.Specifiers {
			switch v := spec.(type) {
			case *ImportDefaultSpecifier:
				defaultImport = v.String()
			case *ImportNamespaceSpecifier:
				namespaceImport = v.String()
			case *ImportSpecifier:
				named := jsElementsToString(v.NamedImports)
				namedImports = append(namedImports, named...)
			}
		}

		if defaultImport != "" {
			s += defaultImport
		}

		if namespaceImport != "" {
			if defaultImport != "" {
				s += ", "
			}
			s += namespaceImport
		}

		if len(namedImports) > 0 {
			named := strings.Join(namedImports, ", ")
			if defaultImport != "" || namespaceImport != "" {
				s += ", "
			}
			s += "{ " + named + " }"
		}
		s += " from "
	}
	b, _ := json.Marshal(i.Source)
	s += string(b) + ";"
	return
}

type VariableDeclaration struct {
	Declarations []VariableDeclarator
	Kind         VariableDeclarationType
	*Node
}

func (v *VariableDeclaration) String() (s string) {
	decl := make([]string, len(v.Declarations))
	for n, x := range v.Declarations {
		decl[n] = x.ID.String() + " = " + x.Init.String()
	}
	s = string(v.Kind) + " " + strings.Join(decl, ",")
	return
}

type VariableDeclarationType string

const (
	VariableDeclarationTypeConst VariableDeclarationType = "const"
	VariableDeclarationTypeLet   VariableDeclarationType = "let"
	VariableDeclarationTypeVar   VariableDeclarationType = "var"
)

type VariableDeclarator struct {
	ID   BindingIdentifierOrPattern
	Init Expression
}

// Statements
type BreakStatement struct {
	Label *Identifier
	*Node
}

func (b *BreakStatement) String() (s string) {
	s = "break"
	if b.Label != nil {
		s += " " + b.Label.String()
	}
	s += ";"
	return
}

type ContinueStatement struct {
	Label *Identifier
	*Node
}

func (c *ContinueStatement) String() (s string) {
	s = "continue"
	if c.Label != nil {
		s += " " + c.Label.String()
	}
	s += ";"
	return
}

type DebuggerStatement struct {
	*Node
}

func (d *DebuggerStatement) String() string {
	return "debugger;"
}

type DoWhileStatement struct {
	Body BlockStatement
	Test Expression
	*Node
}

func (d *DoWhileStatement) String() string {
	return "do {\n" + indentor.Indent(d.Body.String()) + "\n} while(" + d.Test.String() + ");"
}

type EmptyStatement struct {
	*Node
}

func (e *EmptyStatement) String() string {
	return ";"
}

type ExpressionStatement struct {
	Expression
	*Node
}

func (e *ExpressionStatement) String() string {
	return e.Expression.String() + ";"
}

type Directive struct {
	Expression
	Directive string
	*Node
}

func (d *Directive) String() string {
	panic("Not Supported")
}

type ForStatement struct {
	Init   Expression
	Test   Expression
	Update Expression
	Body   BlockStatement
	*Node
}

func (f *ForStatement) String() string {
	return "for(" + f.Init.String() + ", " + f.Test.String() + ", " + f.Update.String() + "){\n" + indentor.Indent(f.Body.String()) + "\n}"
}

type ForInStatement struct {
	Left  Expression
	Right Expression
	Body  BlockStatement
	Each  bool
	*Node
}

func (f *ForInStatement) String() string {
	return "for(" + f.Left.String() + " in " + f.Right.String() + "){\n" + indentor.Indent(f.Body.String()) + "\n}"
}

type ForOfStatement struct {
	Await bool
	Left  Expression
	Right Expression
	Body  BlockStatement
	*Node
}

func (f *ForOfStatement) String() (s string) {
	s = "for"
	if f.Await {
		s += " await"
	}
	s += "(" + f.Left.String() + " in " + f.Right.String() + "){\n" + indentor.Indent(f.Body.String()) + "\n}"
	return
}

type IfStatement struct {
	Test       Expression
	Consequent Statement
	Alternate  Statement
	*Node
}

func (f *IfStatement) String() (s string) {
	s = "if (" + f.Test.String() + ") {\n" + indentor.Indent(f.Consequent.String()) + "\n}"
	if f.Alternate != nil {
		s += " else "
		var isIf bool
		if _, isIf := f.Alternate.(*IfStatement); !isIf {
			s += "{\n"
		}
		s += indentor.Indent(f.Alternate.String())
		if !isIf {
			s += "\n}"
		}
	}
	return s
}

type ReturnStatement struct {
	Argument Expression
	*Node
}

func (r *ReturnStatement) String() (s string) {
	s = "return"
	if r.Argument != nil {
		s += " " + r.Argument.String()
	}
	s += ";"
	return
}

type SwitchStatement struct {
	Discriminant Expression
	Cases        []SwitchCase
	*Node
}

func (r *SwitchStatement) String() string {
	cases := jsElementsToString(r.Cases)
	return "switch {\n" + indentor.Indent(strings.Join(cases, "\n")) + "\n}"
}

type ThrowStatement struct {
	Argument Expression
	*Node
}

func (t *ThrowStatement) String() string {
	return "throw " + t.Argument.String() + ";"
}

type TryStatement struct {
	Block     BlockStatement
	Handler   CatchClause
	Finalizer *BlockStatement
	*Node
}

func (t *TryStatement) String() (s string) {
	s = "try {\n" + indentor.Indent(t.Block.String()) + "\n} " + t.Handler.String()
	if t.Finalizer != nil {
		s += " finally {\n" + indentor.Indent(t.Finalizer.String()) + "\n}"
	}
	return
}

type WhileStatement struct {
	Test Expression
	Body Statement
	*Node
}

func (w *WhileStatement) String() string {
	return "while (" + w.Test.String() + ") {\n" + indentor.Indent(w.Body.String()) + "\n}"
}

type WithStatement struct {
	Object Expression
	Body   Statement
	*Node
}

func (w *WithStatement) String() string {
	return "with (" + w.Object.String() + ") {\n" + indentor.Indent(w.Body.String()) + "\n}"
}

// Imports
type ImportDefaultSpecifier struct {
	Local *Identifier
	*Node
}

func (i *ImportDefaultSpecifier) String() string {
	return i.Local.String()
}

type ImportNamespaceSpecifier struct {
	Local *Identifier
	*Node
}

func (i *ImportNamespaceSpecifier) String() string {
	return "* as " + i.Local.String()
}

type ImportSpecifier struct {
	NamedImports []NamedImport
	*Node
}

func (i ImportSpecifier) String() (s string) {
	if len(i.NamedImports) == 0 {
		return "{}"
	}

	return "{ " + strings.Join(i.namedImports(), ", ") + " }"
}

func (i *ImportSpecifier) namedImports() []string {
	named := make([]string, len(i.NamedImports))
	for n, ni := range i.NamedImports {
		out := ni.Imported.Name
		if ni.Local != nil {
			out += " as " + ni.Local.Name
		}
		named[n] = out
	}
	return named
}

type NamedImport struct {
	Local    *Identifier
	Imported *Identifier
	*Node
}

func (n NamedImport) String() (s string) {
	s = n.Imported.Name
	if n.Local != nil {
		s += " as " + n.Local.Name
	}
	return
}

type LabeledStatement struct {
	Label Identifier
	Body  Statement
	*Node
}

type MetaProperty struct {
	Meta     Identifier
	Property Identifier
	*Node
}

// misc
type RestElement struct {
	Argument BindingIdentifierOrPattern
	*Node
}

func (r *RestElement) String() string {
	return "..." + r.Argument.String()
}

type SpreadElement struct {
	Argument Expression
	*Node
}

func (s *SpreadElement) String() string {
	return "..." + s.Argument.String()
}

type Super struct {
	*Node
}

// Type safety

// ArgumentListElements
func (s *SpreadElement) argumentListElement()            {}
func (s *Identifier) argumentListElement()               {}
func (s *ArrayExpression) argumentListElement()          {}
func (s *ArrowFunctionExpression) argumentListElement()  {}
func (s *AwaitExpression) argumentListElement()          {}
func (s *AssignmentExpression) argumentListElement()     {}
func (s *BinaryExpression) argumentListElement()         {}
func (s *LogicalExpression) argumentListElement()        {}
func (s *CallExpression) argumentListElement()           {}
func (s *ChainExpression) argumentListElement()          {}
func (s *ClassExpression) argumentListElement()          {}
func (s *ComputedMemberExpression) argumentListElement() {}
func (s *ConditionalExpression) argumentListElement()    {}
func (s *FunctionExpression) argumentListElement()       {}
func (s *NewExpression) argumentListElement()            {}
func (s *ObjectExpression) argumentListElement()         {}
func (s *SequenceExpression) argumentListElement()       {}
func (s *StaticMemberExpression) argumentListElement()   {}
func (s *TaggedTemplateExpression) argumentListElement() {}
func (s *UnaryExpression) argumentListElement()          {}
func (s *UpdateExpression) argumentListElement()         {}
func (s *YieldExpression) argumentListElement()          {}
func (s *literalValueUndefined) argumentListElement()    {}
func (s *literalValueNull) argumentListElement()         {}
func (s *LiteralValueString) argumentListElement()       {}
func (s *LiteralValueBool) argumentListElement()         {}
func (s *LiteralValueNumber) argumentListElement()       {}
func (s *LiteralValueBigFloat) argumentListElement()     {}

// ArrayExpressionElements
func (s *SpreadElement) arrayExpressionElement()            {}
func (s *Identifier) arrayExpressionElement()               {}
func (s *ArrayExpression) arrayExpressionElement()          {}
func (s *ArrowFunctionExpression) arrayExpressionElement()  {}
func (s *AssignmentExpression) arrayExpressionElement()     {}
func (s *AwaitExpression) arrayExpressionElement()          {}
func (s *BinaryExpression) arrayExpressionElement()         {}
func (s *LogicalExpression) arrayExpressionElement()        {}
func (s *CallExpression) arrayExpressionElement()           {}
func (s *ChainExpression) arrayExpressionElement()          {}
func (s *ClassExpression) arrayExpressionElement()          {}
func (s *ComputedMemberExpression) arrayExpressionElement() {}
func (s *ConditionalExpression) arrayExpressionElement()    {}
func (s *FunctionExpression) arrayExpressionElement()       {}
func (s *NewExpression) arrayExpressionElement()            {}
func (s *ObjectExpression) arrayExpressionElement()         {}
func (s *SequenceExpression) arrayExpressionElement()       {}
func (s *StaticMemberExpression) arrayExpressionElement()   {}
func (s *TaggedTemplateExpression) arrayExpressionElement() {}
func (s *UnaryExpression) arrayExpressionElement()          {}
func (s *UpdateExpression) arrayExpressionElement()         {}
func (s *YieldExpression) arrayExpressionElement()          {}
func (s *literalValueUndefined) arrayExpressionElement()    {}
func (s *literalValueNull) arrayExpressionElement()         {}
func (s *LiteralValueString) arrayExpressionElement()       {}
func (s *LiteralValueBool) arrayExpressionElement()         {}
func (s *LiteralValueNumber) arrayExpressionElement()       {}
func (s *LiteralValueBigFloat) arrayExpressionElement()     {}

// Literal Value
func (l *literalValueNull) literal()      {}
func (l *literalValueUndefined) literal() {}
func (l *LiteralValueBool) literal()      {}
func (l *LiteralValueNumber) literal()    {}
func (l *LiteralValueBigFloat) literal()  {}
func (l LiteralValueString) literal()     {}

// Patterns
func (a *ArrayPattern) bindingPattern()              {}
func (o *ObjectPattern) bindingPattern()             {}
func (i *Identifier) bindingIdentifierOrPattern()    {}
func (a *ArrayPattern) bindingIdentifierOrPattern()  {}
func (o *ObjectPattern) bindingIdentifierOrPattern() {}

// Expressions
func (n *Identifier) expression()               {}
func (n *ArrayExpression) expression()          {}
func (n *ArrowFunctionExpression) expression()  {}
func (n *AssignmentExpression) expression()     {}
func (s *AwaitExpression) expression()          {}
func (n *BinaryExpression) expression()         {}
func (n *LogicalExpression) expression()        {}
func (n *CallExpression) expression()           {}
func (n *ChainExpression) expression()          {}
func (n *ClassExpression) expression()          {}
func (n *ComputedMemberExpression) expression() {}
func (n *ConditionalExpression) expression()    {}
func (n *FunctionExpression) expression()       {}
func (n *NewExpression) expression()            {}
func (n *ObjectExpression) expression()         {}
func (n *SequenceExpression) expression()       {}
func (n *StaticMemberExpression) expression()   {}
func (n *TaggedTemplateExpression) expression() {}
func (n *UnaryExpression) expression()          {}
func (n *UpdateExpression) expression()         {}
func (n *YieldExpression) expression()          {}
func (s *literalValueUndefined) expression()    {}
func (s *literalValueNull) expression()         {}
func (s *LiteralValueString) expression()       {}
func (s *LiteralValueBool) expression()         {}
func (s *LiteralValueNumber) expression()       {}
func (s *LiteralValueBigFloat) expression()     {}

// Declarations
func (s *FunctionDeclaration) declaration()      {}
func (s *VariableDeclaration) declaration()      {}
func (s *ClassDeclaration) declaration()         {}
func (s *ImportDeclaration) declaration()        {}
func (s *ExportAllDeclaration) declaration()     {}
func (s *ExportDefaultDeclaration) declaration() {}
func (s *ExportNamedDeclaration) declaration()   {}

// Statements
func (s *BreakStatement) statement()      {}
func (s *ContinueStatement) statement()   {}
func (s *DebuggerStatement) statement()   {}
func (s *DoWhileStatement) statement()    {}
func (s *EmptyStatement) statement()      {}
func (s *ExpressionStatement) statement() {}
func (s *Directive) statement()           {}
func (s *ForStatement) statement()        {}
func (s *ForInStatement) statement()      {}
func (s *ForOfStatement) statement()      {}
func (s *IfStatement) statement()         {}
func (s *ReturnStatement) statement()     {}
func (s *SwitchStatement) statement()     {}
func (s *ThrowStatement) statement()      {}
func (s *TryStatement) statement()        {}
func (s *WhileStatement) statement()      {}
func (s *WithStatement) statement()       {}
func (s *BlockStatement) statement()      {}

// StatementListItems
func (s *FunctionDeclaration) statementListItem()      {}
func (s *VariableDeclaration) statementListItem()      {}
func (s *ClassDeclaration) statementListItem()         {}
func (s *ImportDeclaration) statementListItem()        {}
func (s *BreakStatement) statementListItem()           {}
func (s *ContinueStatement) statementListItem()        {}
func (s *DebuggerStatement) statementListItem()        {}
func (s *DoWhileStatement) statementListItem()         {}
func (s *EmptyStatement) statementListItem()           {}
func (s *ExpressionStatement) statementListItem()      {}
func (s *Directive) statementListItem()                {}
func (s *ForStatement) statementListItem()             {}
func (s *ForInStatement) statementListItem()           {}
func (s *ForOfStatement) statementListItem()           {}
func (s *IfStatement) statementListItem()              {}
func (s *ReturnStatement) statementListItem()          {}
func (s *SwitchStatement) statementListItem()          {}
func (s *ThrowStatement) statementListItem()           {}
func (s *TryStatement) statementListItem()             {}
func (s *WhileStatement) statementListItem()           {}
func (s *WithStatement) statementListItem()            {}
func (s *BlockStatement) statementListItem()           {}
func (s *ExportAllDeclaration) statementListItem()     {}
func (s *ExportDefaultDeclaration) statementListItem() {}
func (s *ExportNamedDeclaration) statementListItem()   {}

// ArrayPatternElements
func (s *AssignmentPattern) arrayPatternElement() {}
func (s *Identifier) arrayPatternElement()        {}
func (s *ArrayPattern) arrayPatternElement()      {}
func (s *ObjectPattern) arrayPatternElement()     {}
func (s *RestElement) arrayPatternElement()       {}

// ChainElements
func (s *CallExpression) chainElementToString() string {
	return s.Callee.String() + "?." + s.argsToString()
}
func (s *ComputedMemberExpression) chainElementToString() string {
	return s.Object.String() + "?.[" + s.Property.String() + "]"
}
func (s *StaticMemberExpression) chainElementToString() string {
	return s.Object.String() + "?." + s.propertyToString()
}

// ExportableDefaultDeclarations
func (s *Identifier) exportableDefaultDeclaration()               {}
func (s *ClassDeclaration) exportableDefaultDeclaration()         {}
func (s *FunctionDeclaration) exportableDefaultDeclaration()      {}
func (s *ArrayPattern) exportableDefaultDeclaration()             {}
func (s *ObjectPattern) exportableDefaultDeclaration()            {}
func (n *ArrayExpression) exportableDefaultDeclaration()          {}
func (n *ArrowFunctionExpression) exportableDefaultDeclaration()  {}
func (n *AssignmentExpression) exportableDefaultDeclaration()     {}
func (s *AwaitExpression) exportableDefaultDeclaration()          {}
func (n *BinaryExpression) exportableDefaultDeclaration()         {}
func (n *LogicalExpression) exportableDefaultDeclaration()        {}
func (n *CallExpression) exportableDefaultDeclaration()           {}
func (n *ChainExpression) exportableDefaultDeclaration()          {}
func (n *ClassExpression) exportableDefaultDeclaration()          {}
func (n *ComputedMemberExpression) exportableDefaultDeclaration() {}
func (n *ConditionalExpression) exportableDefaultDeclaration()    {}
func (n *FunctionExpression) exportableDefaultDeclaration()       {}
func (n *NewExpression) exportableDefaultDeclaration()            {}
func (n *ObjectExpression) exportableDefaultDeclaration()         {}
func (n *SequenceExpression) exportableDefaultDeclaration()       {}
func (n *StaticMemberExpression) exportableDefaultDeclaration()   {}
func (n *TaggedTemplateExpression) exportableDefaultDeclaration() {}
func (n *UnaryExpression) exportableDefaultDeclaration()          {}
func (n *UpdateExpression) exportableDefaultDeclaration()         {}
func (n *YieldExpression) exportableDefaultDeclaration()          {}
func (s *literalValueUndefined) exportableDefaultDeclaration()    {}
func (s *literalValueNull) exportableDefaultDeclaration()         {}
func (s *LiteralValueString) exportableDefaultDeclaration()       {}
func (s *LiteralValueBool) exportableDefaultDeclaration()         {}
func (s *LiteralValueNumber) exportableDefaultDeclaration()       {}
func (s *LiteralValueBigFloat) exportableDefaultDeclaration()     {}

// ExportableNamedDeclarations
func (n *ClassDeclaration) exportableNamedDeclaration()    {}
func (n *FunctionDeclaration) exportableNamedDeclaration() {}
func (n *VariableDeclaration) exportableNamedDeclaration() {}

// FunctionParameters
func (s *ArrayPattern) functionParameter()  {}
func (s *ObjectPattern) functionParameter() {}
func (s *Identifier) functionParameter()    {}

// ImportDeclarationSpecifiers
func (s *ImportDefaultSpecifier) importDeclarationSpecifier()   {}
func (s *ImportNamespaceSpecifier) importDeclarationSpecifier() {}
func (s *ImportSpecifier) importDeclarationSpecifier()          {}

// ObjectExpressionProperties
func (s *Property) objectExpressionProperty()      {}
func (s *SpreadElement) objectExpressionProperty() {}

// ObjectPatternProperty
func (s *PropertyPattern) objectPatternProperty() {}
func (s *RestElement) objectPatternProperty()     {}

// PropertyKeys
func (s *Identifier) propertyKey()            {}
func (s *literalValueUndefined) propertyKey() {}
func (s *literalValueNull) propertyKey()      {}
func (s *LiteralValueString) propertyKey()    {}
func (s *LiteralValueBool) propertyKey()      {}
func (s *LiteralValueNumber) propertyKey()    {}
func (s *LiteralValueBigFloat) propertyKey()  {}

// PropertyValues
func (s *Identifier) propertyValue()         {}
func (s *FunctionExpression) propertyValue() {}
func (a *AssignmentPattern) propertyValue()  {}
func (a *ArrayPattern) propertyValue()       {}
func (a *ObjectPattern) propertyValue()      {}

// ExpressionOrImport
func (n *Identifier) expressionOrImport()               {}
func (n *ArrayExpression) expressionOrImport()          {}
func (n *ArrowFunctionExpression) expressionOrImport()  {}
func (n *AssignmentExpression) expressionOrImport()     {}
func (s *AwaitExpression) expressionOrImport()          {}
func (n *BinaryExpression) expressionOrImport()         {}
func (n *LogicalExpression) expressionOrImport()        {}
func (n *CallExpression) expressionOrImport()           {}
func (n *ChainExpression) expressionOrImport()          {}
func (n *ClassExpression) expressionOrImport()          {}
func (n *ComputedMemberExpression) expressionOrImport() {}
func (n *FunctionExpression) expressionOrImport()       {}
func (n *ConditionalExpression) expressionOrImport()    {}
func (n *NewExpression) expressionOrImport()            {}
func (n *ObjectExpression) expressionOrImport()         {}
func (n *SequenceExpression) expressionOrImport()       {}
func (n *StaticMemberExpression) expressionOrImport()   {}
func (n *TaggedTemplateExpression) expressionOrImport() {}
func (n *UnaryExpression) expressionOrImport()          {}
func (n *UpdateExpression) expressionOrImport()         {}
func (n *YieldExpression) expressionOrImport()          {}
func (s *literalValueUndefined) expressionOrImport()    {}
func (s *literalValueNull) expressionOrImport()         {}
func (s *LiteralValueString) expressionOrImport()       {}
func (s *LiteralValueBool) expressionOrImport()         {}
func (s *LiteralValueNumber) expressionOrImport()       {}
func (s *LiteralValueBigFloat) expressionOrImport()     {}

// ExportDeclaration
func (s *ExportAllDeclaration) exportDeclaration()     {}
func (s *ExportDefaultDeclaration) exportDeclaration() {}
func (s *ExportNamedDeclaration) exportDeclaration()   {}

// ClassProperty
func (s *MethodDefinition) classProperty()   {}
func (s *PropertyDefinition) classProperty() {}

// Helpers
func argListToString(args []ArgumentListElement) string {
	out := jsElementsToString(args)
	return strings.Join(out, ", ")
}

func objectExpressionPropertiesToString(props []ObjectExpressionProperty) (s string) {
	pl := len(props)
	result := jsElementsToString(props)
	s = strings.Join(result, ",\n")
	if pl > 0 {
		s += ","
	}
	return
}

func functionParametersToString(params []FunctionParameter) string {
	out := jsElementsToString(params)
	return strings.Join(out, ", ")
}
