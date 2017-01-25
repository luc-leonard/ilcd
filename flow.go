package ilcd

import (
	"encoding/xml"
)

// Flow represents an ILCD flow data set
type Flow struct {
	XMLName                 xml.Name          `xml:"flowDataSet"`
	Info                    *FlowInfo         `xml:"flowInformation>dataSetInformation"`
	ReferenceFlowPropertyID int               `xml:"flowInformation>quantitativeReference>referenceToReferenceFlowProperty"`
	Type                    string            `xml:"modellingAndValidation>LCIMethod>typeOfDataSet"`
	Publication             *FlowPublication  `xml:"administrativeInformation>publicationAndOwnership"`
	FlowProperties          []FlowPropertyRef `xml:"flowProperties>flowProperty"`
}

// ReferenceFlowProperty returns the reference to the reference flow property of
// the flow.
func (f *Flow) ReferenceFlowProperty() *FlowPropertyRef {
	if f == nil || f.FlowProperties == nil {
		return nil
	}
	for i, ref := range f.FlowProperties {
		if ref.ID == f.ReferenceFlowPropertyID {
			return &f.FlowProperties[i]
		}
	}
	return nil
}

// FlowInfo contains the general flow information
type FlowInfo struct {
	UUID            string           `xml:"UUID"`
	Name            *FlowName        `xml:"name"`
	Synonyms        LangString       `xml:"synonyms"`
	Classifications []Classification `xml:"classificationInformation>classification"`
	CAS             string           `xml:"CASNumber"`
	Comment         LangString       `xml:"generalComment"`
}

// FlowName contains the name fields of a flow.
type FlowName struct {
	BaseName       LangString `xml:"baseName"`
	Treatment      LangString `xml:"treatmentStandardsRoutes"`
	MixAndLocation LangString `xml:"mixAndLocationTypes"`
	Properties     LangString `xml:"flowProperties"`
}

// FlowPublication contains the information about publication and ownership of a
// flow.
type FlowPublication struct {
	Version string `xml:"dataSetVersion"`
}

// FlowPropertyRef describes a flow property of a flow.
type FlowPropertyRef struct {
	ID           int        `xml:"dataSetInternalID,attr"`
	FlowProperty *Ref       `xml:"referenceToFlowPropertyDataSet"`
	Mean         float64    `xml:"meanValue"`
	Comment      LangString `xml:"generalComment"`
}
