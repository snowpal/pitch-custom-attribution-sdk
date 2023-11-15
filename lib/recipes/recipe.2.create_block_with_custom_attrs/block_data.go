package blockattrs

const (
	BlockID = iota
	Name
	Description
	SimpleDescription
	Tags
	Flag
	BlockType
	ScaleAndScaleValue
	DueDate
	StartDate
	EndDate
)

var BlockAttrs = map[int]struct {
	Name      string
	Required  bool
	List      bool
	Metadata  bool
	ChildName string
}{
	BlockID: {
		Name:     "Block ID",
		Required: false,
		List:     false,
		Metadata: false,
	},
	Name: {
		Name:     "Name",
		Required: true,
		List:     false,
		Metadata: false,
	},
	Description: {
		Name:     "Description",
		Required: false,
		List:     false,
		Metadata: false,
	},
	SimpleDescription: {
		Name:     "Simple Description",
		Required: false,
		List:     false,
		Metadata: false,
	},
	Tags: {
		Name:     "Tags",
		Required: false,
		List:     false,
		Metadata: false,
	},
	Flag: {
		Name:     "Flag",
		Required: false,
		List:     false,
		Metadata: false,
	},
	BlockType: {
		Name:     "Block Type",
		Required: false,
		List:     false,
		Metadata: false,
	},
	ScaleAndScaleValue: {
		Name:      "Scale",
		Required:  false,
		List:      false,
		Metadata:  false,
		ChildName: "Scale Value",
	},
	DueDate: {
		Name:     "Due Date",
		Required: false,
		List:     false,
		Metadata: false,
	},
	StartDate: {
		Name:     "Start Date",
		Required: false,
		List:     false,
		Metadata: false,
	},
	EndDate: {
		Name:     "End Date",
		Required: false,
		List:     false,
		Metadata: false,
	},
}

const (
	BlockAttrsBagName = "Block Attributes"
)
