package overwrites

// the direction of the rule defines in which direction the destination is pointing.
type OverwriteRuleDirection int

// the strategy defines on which properties, the overwrites are based.
type OverwriteStrategy int

const (
	UP OverwriteRuleDirection = 1 << iota
	DOWN
	INTERN // this means the overwrite rule will overwrite stuff inside the same level
)

const (
	PATH OverwriteStrategy = 1 << iota // PATH will overwrite files with the exact paths given
	NAME                               // NAME will overwrite files with the same name in the given direction
)

type Overwrite struct {
	OriginFileId      string `json:"originFileId"`
	OriginFileLine    int    `json:"originFileLine"`
	OverwriteFileId   string `json:"overwriteFileId"`
	OverwriteFileLine int    `json:"overwriteFileLine"`
}

type OverwriteConfig struct {
	// list of all rules applying to containing location.
	// order of the elements is important because it also defines
	// the priority of the rule within the rule set. This means
	// that the first rule in the list is the most important and
	// the last rule in the list is the least important.
	Rules []*OverwriteRule `yaml:"rules" json:"rules"`
}

// An overwrite rule is defined relative to the containing artefact.
type OverwriteRule struct {
	// in which way the rule shall be applied / point inside the structure
	Direction OverwriteRuleDirection `yaml:"direction" json:"direction"`
	// the strategy of the rule
	Strategy OverwriteStrategy `yaml:"strategy" json:"strategy"`
	// the path of the file which shall overwrite
	Origin string `yaml:"origin,omitempty" json:"origin,omitempty"`
	// the path of the file which shall be overwritten. The destination path is relative to the origin path
	Destination string `yaml:"destination,omitempty" json:"destination,omitempty"`
}
