package commons

var (
	// Version of PCI DSS Standard
	Version string
	// CfgDir is Config Directory
	CfgDir string
)

// Standard contain information about PCI DSS Standards
type Standard struct {
	ID           string         `yaml:"id" json:"id"`
	Description  string         `yaml:"description" json:"description"`
	Since        string         `yaml:"since" json:"since"`
	Author       string         `yaml:"author" json:"author"`
	Goals        []*Goal        `json:"goals"`
	Requirements []*Requirement `json:"requirements"`
}

// Goal contain information about PCI DSS Standard Goals Information
type Goal struct {
	ID           string   `yaml:"id" json:"id"`
	Description  string   `yaml:"description" json:"description"`
	Requirements []string `yaml:"requirements" json:"requirements"`
}

// Goals contain list of PCI DSS goal information
type Goals struct {
	Values []*Goal `yaml:"goals" json:"goals"`
}

// Requirement contain information about PCI DSS Standard Requirement Information
type Requirement struct {
	ID              string  `yaml:"id" json:"id"`
	Description     string  `yaml:"description" json:"description"`
	Explanation     string  `yaml:"explanation" json:"explanation"`
	ExplanationNote string  `yaml:"explanation_note" json:"explanation_note"`
	Goal            string  `yaml:"goal" json:"goal"`
	Items           []*Item `yaml:"items" json:"items"`
}

// Requirements contain list of PCI DSS Requirement information
type Requirements struct {
	Values []*Requirement `yaml:"requirements" json:"requirements"`
}

// Item containt Item Information
type Item struct {
	ID               string              `yaml:"id" json:"id"`
	Description      string              `yaml:"description" json:"description"`
	DescriptionNote  string              `yaml:"description_note" json:"description_note"`
	TestingProcedure []*TestingProcedure `yaml:"testing_procedures" json:"testing_procedures"`
	Guidance         string              `yaml:"guidance" json:"guidance"`
	GuidanceNote     string              `yaml:"guidance_note" json:"guidance_note"`
	Items            []*Item             `yaml:"items" json:"items"`
}

// TestingProcedure contain Testing Procedure information
type TestingProcedure struct {
	ID              string `yaml:"id" json:"id"`
	Description     string `yaml:"description" json:"description"`
	DescriptionNote string `yaml:"description_note" json:"description_note"`
}

// ItemRoot containt ItemRoot Information
type ItemRoot struct {
	Parent string  `yaml:"parent" json:"parent"`
	Items  []*Item `yaml:"items" json:"items"`
}
