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
}
