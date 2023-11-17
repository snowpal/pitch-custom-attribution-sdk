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

const (
	BlockIDValue         = "A01"
	BlockNameValue       = "Trial Block"
	BlockDescValue       = "This is trial block and this is its description"
	BlockSimpleDescValue = "This is block's simple description"
	BlockTagsValue       = "Tag1,Tag2,Tag3"
	BlockFlagValue       = "white"
	BlockTypeID          = 2 // 2 is for "Advanced Placement" block type
	BlockScaleID         = 1 // 1 is for for "Grades" scale
	BlockScaleValueID    = 2 // 2 is for "B" scale value
	BlockDueDateValue    = "2021-10-08T18:47:39.021103599Z"
	BlockStartDateValue  = "2023-11-18T18:47:39.021103599Z"
	BlockEndDateValue    = "2023-11-28T18:47:39.021103599Z"
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
