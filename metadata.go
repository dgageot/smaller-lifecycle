package lifecycle

import "path"

type BuildMetadata struct {
	Processes  []Process        `toml:"processes" json:"processes"`
	Buildpacks []Buildpack      `toml:"buildpacks" json:"buildpacks"`
	BOM        []BOMEntry       `toml:"bom" json:"bom"`
	Launcher   LauncherMetadata `toml:"-" json:"launcher"`
	Slices     []Slice          `toml:"slices" json:"-"`
}

type LauncherMetadata struct {
	Version string         `json:"version"`
	Source  SourceMetadata `json:"source"`
}

type SourceMetadata struct {
	Git GitMetadata `json:"git"`
}

type GitMetadata struct {
	Repository string `json:"repository"`
	Commit     string `json:"commit"`
}

func MetadataFilePath(layersDir string) string {
	return path.Join(layersDir, "config", "metadata.toml")
}
