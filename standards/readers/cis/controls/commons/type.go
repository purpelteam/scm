package commons

var (
	// CISControlVersion of CIS Controls
	CISControlVersion string
	// CISControlCfgDir is CIS Controls Config Directory
	CISControlCfgDir string
)

// CISControl contain information about CIS Controls
type CISControl struct {
	ID          string                 `yaml:"id" json:"id"`
	Description string                 `yaml:"description" json:"description"`
	Since       string                 `yaml:"since" json:"since"`
	Author      string                 `yaml:"author" json:"author"`
	Categories  []*Category            `json:"categories"`
	IGs         []*ImplementationGroup `json:"implementation_groups"`
	Controls    []*Control             `json:"controls"`
}

// Category contain information about CIS Controls Category Information
type Category struct {
	ID          string   `yaml:"id" json:"id"`
	Description string   `yaml:"description" json:"description"`
	Controls    []string `yaml:"controls" json:"controls"`
}

// Categories contain list of CIS Controls Category information
type Categories struct {
	Values []*Category `yaml:"categories" json:"categories"`
}

// ImplementationGroup contain information about CIS Controls Implementation Group
type ImplementationGroup struct {
	ID                 string `yaml:"id" json:"id"`
	Title              string `yaml:"title" json:"title"`
	Definitions        string `yaml:"definitions" json:"definitions"`
	QuickStartGuidance string `yaml:"quick_start_guidance" json:"quick_start_guidance"`
}

// ImplementationGroups contain list of CIS Implementation Group information
type ImplementationGroups struct {
	Values []*ImplementationGroup `yaml:"implementation_groups" json:"implementation_groups"`
}

// Control contain CIS Control item information
type Control struct {
	ID                 string        `yaml:"id" json:"id"`
	Title              string        `yaml:"title" json:"title"`
	Description        string        `yaml:"description" json:"description"`
	Explanation        string        `yaml:"explanation" json:"explanation"`
	Category           string        `yaml:"category" json:"category"`
	IGNotes            []*IGNote     `yaml:"igs" json:"igs"`
	ProceduresAndTools string        `yaml:"procedures_and_tools" json:"procedures_and_tools"`
	SystemERD          string        `yaml:"system_erd" json:"system_erd"`
	SubControls        []*SubControl `yaml:"sub_controls" json:"sub_controls"`
}

// Controls contain list of CIS Controls information
type Controls struct {
	Values []*Control `yaml:"controls" json:"controls"`
}

// IGNote contain Implementation Group Note
type IGNote struct {
	ID          string `yaml:"id" json:"id"`
	Description string `yaml:"description" json:"description"`
}

// SubControl contain SubControl information
type SubControl struct {
	ID                        string `yaml:"id" json:"id"`
	Title                     string `yaml:"title" json:"title"`
	Description               string `yaml:"description" json:"description"`
	AssetType                 string `yaml:"asset_type" json:"asset_type"`
	SecurityFunction          string `yaml:"security_function" json:"security_function"`
	ImplementationGroupStart  string `yaml:"ig_start" json:"ig_start"`
	ImplementationGroupStatus struct {
		IG1 bool `yaml:"ig1" json:"ig1"`
		IG2 bool `yaml:"ig2" json:"ig2"`
		IG3 bool `yaml:"ig3" json:"ig3"`
	} `yaml:"ig_status" json:"ig_status"`
}

// ItemRoot containt ItemRoot Information
type ItemRoot struct {
	Parent string        `yaml:"parent" json:"parent"`
	Items  []*SubControl `yaml:"items" json:"items"`
}
