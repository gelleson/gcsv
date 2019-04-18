package parser





type Config struct {
	Documents []Documents `yaml:"documents"`

}
type Documents struct {
	Name    string   `yaml:"name"`
	Columns []Column `yaml:"columns"`
	Count   int      `yaml:"rows"`
}

type Column struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Option []string `yaml:"options,flow"`
}