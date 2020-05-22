package commons

var (
	// MappingCfgDir is Mapping Config Directory
	MappingCfgDir string
)

// MappingDirectories contain Mapping Directories information
type MappingDirectories struct {
	Indexs   []*Index   `yaml:"indexs" json:"indexs"`
	Mappings []*Mapping `yaml:"mappings" json:"mappings"`
}

// Index contain index information
type Index struct {
	ID          string `yaml:"id" json:"id"`
	Description string `yaml:"description" json:"description"`
	Version     string `yaml:"version" json:"version"`
}

// Mapping contain mapping information
type Mapping struct {
	ID          string `yaml:"id" json:"id"`
	References  string `yaml:"references" json:"references"`
	From        string `yaml:"from" json:"from"`
	To          string `yaml:"to" json:"to"`
	Location    string `yaml:"location" json:"location"`
	Description string `yaml:"description" json:"description"`
}

// MappingDefinitions contain mapping definition
type MappingDefinitions struct {
	References string         `yaml:"references" json:"references"`
	From       string         `yaml:"from" json:"from"`
	To         string         `yaml:"to" json:"to"`
	Items      []*MappingItem `yaml:"items" json:"items"`
}

// MappingItem contain mapping item
type MappingItem struct {
	ID      string `yaml:"id" json:"id"`
	StdFrom string `yaml:"std_from" json:"std_from"`
	StdTo   string `yaml:"std_to" json:"std_to"`
}
