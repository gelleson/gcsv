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

type Config struct {
	Documents []Documents `yaml:"documents"`
}

// Documents struct to keep config to generate document
type Documents struct {
	// Name of the output file
	Name string `yaml:"name"`
	// Column configs should be generated
	Columns []Column `yaml:"columns"`
	// WithHeader variable to generate headers
	WithHeader bool `yaml:"with_header"`
	// Count is total row
	Count int `yaml:"rows"`
}

// Column struct to keep config to generate column
type Column struct {
	// Name of the column
	Name string `yaml:"name"`
	// Type of the column will be generated
	Type string `yaml:"type"`
	// Option
	Option []string `yaml:"options,flow"`
	// Kwargs extra configs
	Kwargs map[string]interface{} `yaml:"kwargs"`
}
