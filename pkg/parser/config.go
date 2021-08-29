/*
 * Copyright (c) 2020. gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package parser

import (
	"github.com/gelleson/gcsv/pkg/builder/types"
	"github.com/gelleson/gcsv/pkg/generator"
)

// Parser uses to parse yaml document
type Parser struct {
	config Config
}

// NewParser construct instance of Parser
func NewParser(config Config) *Parser {
	return &Parser{config: config}
}

func (p Parser) prepareDocument() []generator.Document {
	var documents []generator.Document
	for _, doc := range p.config.Documents {
		document := generator.NewDocument(doc.Name)
		document.WithHeader = doc.WithHeader
		document.Rows = doc.Count
		for columnIndex, column := range doc.Columns {
			cType := p.columnType(column.Type)
			document.Columns = append(document.Columns, generator.Column{
				Name:     column.Name,
				Position: columnIndex,
				Field: generator.Field{
					Type: cType,
				},
				Kwargs: column.Kwargs,
			})
		}
		documents = append(documents, document)
	}
	return documents
}

func (p Parser) columnType(c string) types.TYPE {
	switch c {
	case "int":
		return types.INT
	case "float":
		return types.FLOAT
	case "date":
		return types.DATE
	case "seq":
		return types.SEQ
	case "string":
		return types.PERSONAL
	case "personal":
		return types.PERSONAL
	default:
		return types.STRING
	}
}

// PreparedDocument are all documents will generate
func (p Parser) PreparedDocument() []generator.Document {
	return p.prepareDocument()
}
