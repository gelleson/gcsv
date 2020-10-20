package generator

type TYPE int

const (
	INT TYPE = iota
	FLOAT
	STRING
	DATE
)

type OPTION int

const (
	SEQ OPTION = iota
	UNIQUE
	NAME
	LAST_NAME
	ADDRESS
	SENTENCE
	COMPANY
	NIL
)
