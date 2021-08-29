package builder

type Config interface {
	Validate() error
}

type Builder interface {
	Initiate(Config) error
	Build(...string) string
	Validate() error
}
