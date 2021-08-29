package types

// Config is interface which help to pass value and validate it
type Config interface {
	Validate() error
}
