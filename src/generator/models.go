package generator

type Document struct {
	Name string
	Rows int
	Columns []Column
}

func NewDocument(name string) *Document {
	return &Document{Name: name}
}

type Column struct {
	Name string
	Field Field
	Position int
}
type Field struct {
	Type TYPE
	Option []OPTION
}
