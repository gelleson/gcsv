package parser

import (
	"github.com/gelleson/gcsv/pkg/generator"
)

type Parser struct {
	config Config
}

func NewParser(config Config) *Parser {
	return &Parser{config: config}
}
func (p Parser) prepareDocument() []generator.Document {
	var documents []generator.Document
	for _, doc := range p.config.Documents {
		document := *generator.NewDocument(doc.Name)
		document.Rows = doc.Count
		for key, value := range doc.Columns {
			cType, cOptions := p.getColumn(value.Type, value.Option)
			document.Columns = append(document.Columns, generator.Column{
				Name:     value.Name,
				Position: key,
				Field: generator.Field{
					Type:   cType,
					Option: cOptions,
				},
			})
		}
		documents = append(documents, document)
	}
	return documents
}

func (p Parser) getColumn(t string, options []string) (generator.TYPE, []generator.OPTION) {
	columnType := p.columnType(t)
	opts := p.columnOptions(options)
	return columnType, opts
}

func (p Parser) columnType(c string) generator.TYPE {
	switch c {
	case "int":
		{
			return generator.INT
		}
	case "float":
		{
			return generator.FLOAT
		}
	case "date":
		{
			return generator.DATE
		}
	case "string":
		{
			return generator.STRING
		}
	default:
		return generator.STRING
	}
}

func (p Parser) columnOptions(options []string) []generator.OPTION {
	opts := make([]generator.OPTION, len(options))
	for i := 0; i < len(options); i++ {
		switch options[i] {
		case "uniq":
			opts[i] = generator.UNIQUE
		case "address":
		case "name":
			opts[i] = generator.NAME
		case "last_name":
			opts[i] = generator.LAST_NAME
		case "seq":
			opts[i] = generator.SEQ
		case "company":
			opts[i] = generator.COMPANY
		default:
			opts[i] = generator.NIL
		}
	}
	return opts
}

func (p Parser) PrepareDocument() []generator.Document {
	return p.prepareDocument()
}
