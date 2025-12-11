package approvalProcesses

import (
	"encoding/xml"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
)

const NAME = "ApprovalProcess"

func init() {
	internal.TypeRegistry.Register(NAME, func(path string) (metadata.RegisterableMetadata, error) { return Open(path) })
}

// ApprovalProcess was generated 2025-12-10 by zek
type ApprovalProcess struct {
	metadata.MetadataInfo
	XMLName xml.Name `xml:"ApprovalProcess"`
	Xmlns   string   `xml:"xmlns,attr"`
	Active  struct {
		Text string `xml:",chardata"`
	} `xml:"active"`
	AllowRecall struct {
		Text string `xml:",chardata"`
	} `xml:"allowRecall"`
	AllowedSubmitters []struct {
		Type struct {
			Text string `xml:",chardata"`
		} `xml:"type"`
		Submitter *struct {
			Text string `xml:",chardata"`
		} `xml:"submitter"`
	} `xml:"allowedSubmitters"`
	ApprovalPageFields struct {
		Field []struct {
			Text string `xml:",chardata"`
		} `xml:"field"`
	} `xml:"approvalPageFields"`
	ApprovalStep []struct {
		AllowDelegate struct {
			Text string `xml:",chardata"`
		} `xml:"allowDelegate"`
		ApprovalActions *struct {
			Action struct {
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
			} `xml:"action"`
		} `xml:"approvalActions"`
		AssignedApprover *struct {
			Approver []struct {
				Name *struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
			} `xml:"approver"`
			WhenMultipleApprovers *struct {
				Text string `xml:",chardata"`
			} `xml:"whenMultipleApprovers"`
		} `xml:"assignedApprover"`
		EntryCriteria *struct {
			CriteriaItems []struct {
				Field struct {
					Text string `xml:",chardata"`
				} `xml:"field"`
				Operation struct {
					Text string `xml:",chardata"`
				} `xml:"operation"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"value"`
			} `xml:"criteriaItems"`
			Formula *struct {
				Text string `xml:",chardata"`
			} `xml:"formula"`
			BooleanFilter struct {
				Text string `xml:",chardata"`
			} `xml:"booleanFilter"`
		} `xml:"entryCriteria"`
		IfCriteriaNotMet *struct {
			Text string `xml:",chardata"`
		} `xml:"ifCriteriaNotMet"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"label"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
		RejectionActions *struct {
			Action struct {
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
			} `xml:"action"`
		} `xml:"rejectionActions"`
		RejectBehavior *struct {
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"rejectBehavior"`
	} `xml:"approvalStep"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	EnableMobileDeviceAccess struct {
		Text string `xml:",chardata"`
	} `xml:"enableMobileDeviceAccess"`
	EntryCriteria struct {
		CriteriaItems []struct {
			Field struct {
				Text string `xml:",chardata"`
			} `xml:"field"`
			Operation struct {
				Text string `xml:",chardata"`
			} `xml:"operation"`
			Value struct {
				Text string `xml:",chardata"`
			} `xml:"value"`
		} `xml:"criteriaItems"`
		Formula *struct {
			Text string `xml:",chardata"`
		} `xml:"formula"`
	} `xml:"entryCriteria"`
	FinalApprovalRecordLock *struct {
		Text string `xml:",chardata"`
	} `xml:"finalApprovalRecordLock"`
	FinalRejectionRecordLock *struct {
		Text string `xml:",chardata"`
	} `xml:"finalRejectionRecordLock"`
	Label struct {
		Text string `xml:",chardata"`
	} `xml:"label"`
	RecordEditability struct {
		Text string `xml:",chardata"`
	} `xml:"recordEditability"`
	ShowApprovalHistory struct {
		Text string `xml:",chardata"`
	} `xml:"showApprovalHistory"`
	FinalApprovalActions *struct {
		Action []struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"action"`
	} `xml:"finalApprovalActions"`
	FinalRejectionActions *struct {
		Action []struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"action"`
	} `xml:"finalRejectionActions"`
	InitialSubmissionActions *struct {
		Action struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"action"`
	} `xml:"initialSubmissionActions"`
	EmailTemplate *struct {
		Text string `xml:",chardata"`
	} `xml:"emailTemplate"`
	NextAutomatedApprover *struct {
		UseApproverFieldOfRecordOwner struct {
			Text string `xml:",chardata"`
		} `xml:"useApproverFieldOfRecordOwner"`
		UserHierarchyField struct {
			Text string `xml:",chardata"`
		} `xml:"userHierarchyField"`
	} `xml:"nextAutomatedApprover"`
	RecallActions *struct {
		Action struct {
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			Type struct {
				Text string `xml:",chardata"`
			} `xml:"type"`
		} `xml:"action"`
	} `xml:"recallActions"`
}

func (c *ApprovalProcess) SetMetadata(m metadata.MetadataInfo) {
	c.MetadataInfo = m
}

func (c *ApprovalProcess) Type() metadata.MetadataType {
	return NAME
}

func Open(path string) (*ApprovalProcess, error) {
	p := &ApprovalProcess{}
	return p, metadata.ParseMetadataXml(p, path)
}
