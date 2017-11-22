package ilcd

// LangString is an ILCD multi-language string
type LangString []LangStringItem

// LangStringItem represents an item in an ILCD multi-language string
type LangStringItem struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"lang,attr"`
}

// Get returns the value for the given language code
func (ls LangString) Get(lang string) string {
	if ls == nil {
		return ""
	}
	for _, item := range ls {
		if item.Lang == lang {
			return item.Value
		}
	}
	return ""
}

// Ref is a data set reference to an ILCD data set.
type Ref struct {
	UUID    string     `xml:"refObjectId,attr"`
	Type    string     `xml:"type,attr"`
	URI     string     `xml:"uri,attr"`
	Version string     `xml:"version,attr"`
	Name    LangString `xml:"shortDescription"`
}

// Classification describes an ILCD classification entry in a data set
type Classification struct {
	Name    string  `xml:"name,attr"`
	Classes []Class `xml:"class"`
}

// GetClass returns the class with the given level from the classification.
func (c *Classification) GetClass(level int) *Class {
	if c == nil || c.Classes == nil {
		return nil
	}
	for i, class := range c.Classes {
		if class.Level == level {
			return &c.Classes[i]
		}
	}
	return nil
}

// Class is a category in an ILCD data set classification.
type Class struct {
	Level int    `xml:"level,attr"`
	ID    string `xml:"classId1,attr"`
	Name  string `xml:",chardata"`
}

// CommonDataEntry <dataEntryBy>
type CommonDataEntry struct {
	TimeStamp   string `xml:"timeStamp"`
	DataFormats []Ref  `xml:"referenceToDataSetFormat"`
}

// CommonPublication <publicationAndOwnership>
type CommonPublication struct {
	Version string `xml:"dataSetVersion"`
	URI     string `xml:"permanentDataSetURI"`
}
