package source

const (
	GoDaddySource   = "gds"
	NameCheapSource = "ncs"
)

// Source is the interface for domain search sources
type Source interface {
	IsAvailable(string) (bool, error)
}

// Get returns a search source
func Get(config interface{}, sourceType string) Source {
	switch sourceType {
	case GoDaddySource:
		return NewGoDaddy(config.(*GoDaddyConfig)).(Source)
	case NameCheapSource:
		return NewNameCheap(config.(*NameCheapConfig)).(Source)
	default:
		panic("Invalid source: " + sourceType)
	}
}
