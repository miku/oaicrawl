package oaicrawl

import (
	"encoding/xml"
	"fmt"
)

// GenericResponse describes the field contained in any response.
type GenericResponse struct {
	XMLName      xml.Name    `xml:"OAI-PMH,omitempty" json:"oai-pmh,omitempty"`
	ResponseDate string      `xml:"responseDate,omitempty" json:"responseDate,omitempty"`
	Request      RequestNode `xml:"request,omitempty" json:"request,omitempty"`
	Error        OAIError    `xml:"error,omitempty" json:"error,omitempty"`
}

// WithResumptionToken can be added to other structs, which can be harvested in batches.
type WithResumptionToken struct {
	ResumptionToken struct {
		XMLName          xml.Name `xml:"resumptionToken,omitempty" json:"resumptionToken,omitempty"`
		Value            string   `xml:",chardata"`
		CompleteListSize string   `xml:"completeListSize,attr"`
		Cursor           string   `xml:"cursor,attr"`
	}
}

// IdentifyResponse reports information about a repository.
type IdentifyResponse struct {
	GenericResponse
	Identify struct {
		XMLName           xml.Name      `xml:"Identify,omitempty" json:"Identify,omitempty"`
		RepositoryName    string        `xml:"repositoryName,omitempty" json:"repositoryName,omitempty"`
		BaseURL           string        `xml:"baseURL,omitempty" json:"baseURL,omitempty"`
		ProtocolVersion   string        `xml:"protocolVersion,omitempty" json:"protocolVersion,omitempty"`
		AdminEmail        []string      `xml:"adminEmail,omitempty" json:"adminEmail,omitempty"`
		EarliestDatestamp string        `xml:"earliestDatestamp,omitempty" json:"earliestDatestamp,omitempty"`
		DeletedRecord     string        `xml:"deletedRecord,omitempty" json:"deletedRecord,omitempty"`
		Granularity       string        `xml:"granularity,omitempty" json:"granularity,omitempty"`
		Description       []Description `xml:"description,omitempty" json:"description,omitempty"`
		Compression       []string      `xml:"compression,omitempty" json:"compression,omitempty"`
	}
}

// ListSetsResponse lists available sets. TODO(miku): resumptiontoken can have expiration date, etc.
type ListSetsResponse struct {
	GenericResponse
	ListSets struct {
		XMLName xml.Name `xml:"ListSets,omitempty" json:"ListSets,omitempty"`
		Sets    []Set    `xml:"set,omitempty"  json:"set,omitempty"`
		WithResumptionToken
	}
}

// A Set has a spec, name and description.
type Set struct {
	SetSpec        string      `xml:"setSpec,omitempty" json:"setSpec,omitempty"`
	SetName        string      `xml:"setName,omitempty" json:"setName,omitempty"`
	SetDescription Description `xml:"setDescription,omitempty" json:"setDescription,omitempty"`
}

// A Header is part of other requests.
type Header struct {
	Status     string   `xml:"status,attr" json:"status,omitempty"`
	Identifier string   `xml:"identifier,omitempty" json:"identifier,omitempty"`
	DateStamp  string   `xml:"datestamp,omitempty" json:"datestamp,omitempty"`
	SetSpec    []string `xml:"setSpec,omitempty" json:"setSpec,omitempty"`
}

// Metadata contains the actual metadata, conforming to various schemas.
type Metadata struct {
	Body []byte `xml:",innerxml"`
}

// GoString is a formatter for Metadata content.
func (md Metadata) GoString() string { return fmt.Sprintf("%s", md.Body) }

// About has addition record information.
type About struct {
	Body []byte `xml:",innerxml" json:"body,omitempty"`
}

// GoString is a formatter for About content.
func (ab About) GoString() string { return fmt.Sprintf("%s", ab.Body) }

// ListIdentifiersResponse lists headers only.
type ListIdentifiersResponse struct {
	GenericResponse
	ListIdentifiers struct {
		XMLName xml.Name `xml:"ListIdentifiers,omitempty" json:"ListIdentifiers,omitempty"`
		Headers []Header `xml:"header,omitempty" json:"header,omitempty"`
		WithResumptionToken
	}
}

// ListRecordsResponse lists records.
type ListRecordsResponse struct {
	GenericResponse
	ListRecords struct {
		XMLName xml.Name `xml:"ListRecords,omitempty" json:"ListRecords,omitempty"`
		Records []Record `xml:"record" json:"record"`
		WithResumptionToken
	}
}

// GetRecordResponse returns a single record.
type GetRecordResponse struct {
	GenericResponse
	GetRecord struct {
		XMLName xml.Name `xml:"GetRecord,omitempty" json:"GetRecord,omitempty"`
		Record  Record   `xml:"record,omitempty" json:"record,omitempty"`
	}
}

// Record represents a single record.
type Record struct {
	Header   Header   `xml:"header,omitempty" json:"header,omitempty"`
	Metadata Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	About    About    `xml:"about,omitempty" json:"about,omitempty"`
}

// RequestNode carries the request information into the response.
type RequestNode struct {
	Verb           string `xml:"verb,attr" json:"verb,omitempty"`
	Set            string `xml:"set,attr" json:"set,omitempty"`
	MetadataPrefix string `xml:"metadataPrefix,attr" json:"metadataPrefix,omitempty"`
}

// OAIError is an OAI protocol error.
type OAIError struct {
	Code    string `xml:"code,attr" json:"code,omitempty"`
	Message string `xml:",chardata" json:"message,omitempty"`
}

// Error formats code and message.
func (e OAIError) Error() string {
	return fmt.Sprintf("oai: %s %s", e.Code, e.Message)
}

// MetadataFormat holds information about a format. Schema and MetadataNamespace
// can contain multiple strings separated by space.
type MetadataFormat struct {
	MetadataPrefix    string `xml:"metadataPrefix,omitempty" json:"metadataPrefix,omitempty"`
	Schema            string `xml:"schema,omitempty" json:"schema,omitempty"`
	MetadataNamespace string `xml:"metadataNamespace,omitempty" json:"metadataNamespace,omitempty"`
}

// ListMetadataFormatsResponse lists supported metadata formats.
type ListMetadataFormatsResponse struct {
	GenericResponse
	ListMetadataFormats struct {
		XMLName         xml.Name         `xml:"ListMetadataFormats,omitempty" json:"ListMetadataFormats,omitempty"`
		MetadataFormats []MetadataFormat `xml:"metadataFormat,omitempty" json:"metadataFormat,omitempty"`
	}
}

// Description holds information about a set. Typically, this includes an
// oai-identifier (OAIIdentifier) node and sometimes more information via eprints
// or oai-toolkit.
type Description struct {
	Body []byte `xml:",innerxml"`
}

// GoString is a formatter for Description content.
func (desc Description) GoString() string { return fmt.Sprintf("%s", desc.Body) }

// OAIIdentifier might occur inside a description, http://www.openarchives.org/OAI/2.0/oai-identifier.xsd.
type OAIIdentifier struct {
	XMLName              xml.Name `xml:"oai-identifier"`
	Scheme               string   `xml:"scheme"`
	RepositoryIdentifier string   `xml:"repositoryIdentifier"`
	Delimiter            string   `xml:"delimiter"`
	SampleIdentifier     string   `xml:"sampleIdentifier"`
}

// Toolkit might occur inside a description, http://oai.dlib.vt.edu/OAI/metadata/toolkit.xsd.
type Toolkit struct {
	XMLName xml.Name `xml:"toolkit"`
	Title   string   `xml:"title"`
	Author  struct {
		Name        string `xml:"name"`
		Email       string `xml:"email"`
		Institution string `xml:"institution"`
	} `xml:"author"`
	ToolkitIcon string `xml:"toolkitIcon"`
	Version     string `xml:"version"`
	URL         string `xml:"URL"`
}

// Eprints might occur inside a description, http://www.openarchives.org/OAI/1.1/eprints.xsd.
type Eprints struct {
	XMLName          xml.Name `xml:"eprints"`
	Content          string   `xml:"content"`
	MetadataPolicy   string   `xml:"metadataPolicy"`
	DataPolicy       string   `xml:"dataPolicy"`
	SubmissionPolicy string   `xml:"submissionPolicy"`
	Comments         []string `xml:"comment"`
}

// Friends might occur inside a descitption, http://www.openarchives.org/OAI/2.0/friends.xsd.
type Friends struct {
	XMLName xml.Name `xml:"friends"`
	BaseURL []string `xml:"baseURL"`
}

// SearchInfo might occur in an about field, https://scout.wisc.edu/XML/searchInfo.xsd.
type SearchInfo struct {
	XMLName               xml.Name `xml:"searchInfo"`
	FullRecordLink        string   `xml:"fullRecordLink"`
	SearchScore           string   `xml:"searchScore"`
	SearchScoreScale      string   `xml:"searchScoreScale"`
	CumulativeRating      string   `xml:"cumulativeRating"`
	CumulativeRatingScale string   `xml:"cumulativeRatingScale"`
}
