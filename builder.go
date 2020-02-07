package lifecycle

type BuildEnv interface {
	AddRootDir(baseDir string) error
	AddEnvDir(envDir string) error
	List() []string
}

type Process struct {
	Type    string   `toml:"type" json:"type"`
	Command string   `toml:"command" json:"command"`
	Args    []string `toml:"args" json:"args"`
	Direct  bool     `toml:"direct" json:"direct"`
}

type Slice struct {
	Paths []string `tom:"paths"`
}

type BOMEntry struct {
	Require
	Buildpack Buildpack `toml:"buildpack" json:"buildpack"`
}
