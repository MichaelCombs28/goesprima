package goesprima

type Program struct {
	Name string
	Body []StatementListItem
	*Range
}

type Node struct {
	Location *SourceLocation
	*Range
}

type SourceLocation struct {
	Start  Position
	End    Position
	Source string
}

type Range struct {
	Start int
	End   int
}

type Position struct {
	Line   int
	Column int
}
