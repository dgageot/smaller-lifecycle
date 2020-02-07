package lifecycle

type Buildpack struct {
	ID       string `toml:"id" json:"id"`
	Version  string `toml:"version" json:"version"`
	Optional bool   `toml:"optional,omitempty" json:"optional,omitempty"`
}

func (bp Buildpack) dir() string {
	return escapeID(bp.ID)
}

type BuildpackGroup struct {
	Group []Buildpack `toml:"group"`
}

type Require struct {
	Name     string                 `toml:"name" json:"name"`
	Version  string                 `toml:"version" json:"version"`
	Metadata map[string]interface{} `toml:"metadata" json:"metadata"`
}

type BuildpackOrder []BuildpackGroup
