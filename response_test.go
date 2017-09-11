package oaicrawl

import (
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"strings"
	"testing"
)

func TestIdentify(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/Identify-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp IdentifyResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T10:12:18Z" {
		t.Error("wrong response date")
	}
	if resp.Identify.AdminEmail[0] != "tosho-joho@tufs.ac.jp" {
		t.Error("could not unmarshal admin email")
	}
	if len(resp.Identify.Compression) != 2 {
		t.Error("wrong compression count")
	}
	if resp.Identify.RepositoryName != "Prometheus-Academic Collections" {
		t.Error("wrong repository name")
	}
}

func TestListIdentifiers(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ListIdentifiers-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp ListIdentifiersResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T12:09:54Z" {
		t.Error("wrong response date")
	}
	if len(resp.ListIdentifiers.Headers) != 5327 {
		t.Errorf("wrong number of headers: want %v", len(resp.ListIdentifiers.Headers))
	}
	if resp.ListIdentifiers.Headers[5149].Status != "deleted" {
		t.Errorf("expected status: deleted")
	}
}

func TestListIdentifiersWithToken(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ListIdentifiers-01.xml")
	if err != nil {
		t.Error(err)
	}
	var resp ListIdentifiersResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T07:23:49Z" {
		t.Error("wrong response date")
	}
	if len(resp.ListIdentifiers.Headers) != 20 {
		t.Errorf("wrong number of headers: want %v", len(resp.ListIdentifiers.Headers))
	}
	if resp.ListIdentifiers.ResumptionToken != "-_--_-oai_dc-_--_-20" {
		t.Errorf("wrong token: %s", resp.ListIdentifiers.ResumptionToken)
	}
}
func TestListSets(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ListSets-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp ListSetsResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T07:39:19Z" {
		t.Error("wrong response date")
	}
	if len(resp.ListSets.Sets) != 28 {
		t.Errorf("wrong number of sets: want %v", len(resp.ListSets.Sets))
	}
	if resp.ListSets.Sets[3].SetName != "Language: Czech" {
		t.Errorf("wrong set name")
	}
	if resp.ListSets.Sets[3].SetSpec != "Language:Czech" {
		t.Errorf("wrong set spec")
	}
}

func TestListMetadataFormats(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ListMetadataFormats-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp ListMetadataFormatsResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T07:42:56Z" {
		t.Error("wrong response date")
	}
	if len(resp.ListMetadataFormats.MetadataFormats) != 4 {
		t.Errorf("wrong number of formats, want %v", len(resp.ListMetadataFormats.MetadataFormats))
	}
	if resp.ListMetadataFormats.MetadataFormats[3].MetadataPrefix != "lar" {
		t.Errorf("wrong prefix, want %v", resp.ListMetadataFormats.MetadataFormats[3].MetadataPrefix)
	}
}

func TestListRecords(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/ListRecords-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp ListRecordsResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T07:58:25Z" {
		t.Error("wrong response date")
	}
	if len(resp.ListRecords.Records) != 20 {
		t.Errorf("expected 20 records, got %d", len(resp.ListRecords.Records))
	}
	if resp.ListRecords.Records[0].Header.Identifier != "oai:amser.org:AMSER-2" {
		t.Errorf("wrong identifier, got %v, want oai:amser.org:AMSER-2",
			resp.ListRecords.Records[0].Header.Identifier)
	}
	if resp.ListRecords.Records[0].Header.DateStamp != "2004-08-26" {
		t.Errorf("wrong identifier, got %v, want 2004-08-26",
			resp.ListRecords.Records[0].Header.DateStamp)
	}
	data := strings.Replace(`
CiAgICAgICAgPHJlY29yZCB4bWxucz0iaHR0cDovL25zLm5zZGwub3JnL25jcy9sYXIiIHhtbG5zOnhz
aT0iaHR0cDovL3d3dy53My5vcmcvMjAwMS9YTUxTY2hlbWEtaW5zdGFuY2UiIHhzaTpzY2hlbWFMb2Nh
dGlvbj0iaHR0cDovL25zLm5zZGwub3JnL25jcy9sYXIgaHR0cDovL25zLm5zZGwub3JnL25jcy9sYXIv
MS4wMC9zY2hlbWFzL2xhci54c2QiPgogICAgICAgICAgPHJlY29yZElEPm9haTphbXNlci5vcmc6QU1T
RVItMjwvcmVjb3JkSUQ+CiAgICAgICAgICA8cmVjb3JkRGF0ZT4yMDA1LTExLTAyPC9yZWNvcmREYXRl
PgogICAgICAgICAgPGlkZW50aWZpZXI+aHR0cHM6Ly93d3cubmlhaWQubmloLmdvdi88L2lkZW50aWZp
ZXI+CiAgICAgICAgICA8dGl0bGU+TmF0aW9uYWwgSW5zdGl0dXRlIG9mIEFsbGVyZ3kgYW5kIEluZmVj
dGlvdXMgRGlzZWFzZXM8L3RpdGxlPgogICAgICAgICAgPGRlc2NyaXB0aW9uPkNyZWF0ZWQgb3ZlciBm
aWZ0eSB5ZWFycyBhZ28sIHRoZSBOYXRpb25hbCBJbnN0aXR1dGUgb2YgQWxsZXJneSBhbmQgSW5mZWN0
aW91cyBEaXNlYXNlcyAoTkFJRCkgJmFtcDtxdW90O2NvbmR1Y3RzIGFuZCBzdXBwb3J0cyBiYXNpYyBh
bmQgYXBwbGllZCByZXNlYXJjaCB0byBiZXR0ZXIgdW5kZXJzdGFuZCwgdHJlYXQsIGFuZCB1bHRpbWF0
ZWx5IHByZXZlbnQgaW5mZWN0aW91cywgaW1tdW5vbG9naWMsIGFuZCBhbGxlcmdpYyBkaXNlYXNlcy4m
YW1wO3F1b3Q7IEluIHJlY2VudCB5ZWFycywgdGhlIHNjb3BlIG9mIHRoZSBpbnN0aXR1dGUncyByZXNl
YXJjaCBhY3Rpdml0aWVzIGhhcyBleHBhbmRlZCB0byBpbmNsdWRlIGVtZXJnaW5nIGlzc3VlcyBzdWNo
IGFzIHRoZSBwb3NzaWJpbGl0eSBvZiBiaW90ZXJyb3Jpc20gYW5kIFdlc3QgTmlsZSB2aXJ1cy4gVGhl
IHNpdGUgY29udGFpbnMgYSB3ZWFsdGggb2YgaW5mb3JtYXRpb24gb24gdGhlIGFjdGl2aXRpZXMgb2Yg
TkFJRCwgc3VjaCBhcyB0aGUgbW9zdCByZWNlbnQgcHVibGljYXRpb25zLCBvcmdhbml6YXRpb25hbCBo
aWVyYXJjaHksIGFuZCBmdW5kaW5nIG9wcG9ydHVuaXRpZXMgZm9yIHJlc2VhcmNoZXJzIGFuZCBzY2hv
bGFycy4gVGhlIG5ld3Nyb29tIGFyZWEgaXMgcXVpdGUgdGhvcm91Z2gsIGFzIHZpc2l0b3JzIGhhdmUg
YWNjZXNzIHRvIHRoZSBkYXRhYmFzZSBvZiBuZXdzIHJlbGVhc2VzIGRhdGluZyBiYWNrIHRvIDE5OTUg
YW5kIGFjY2VzcyB0byBTY2lCaXRlcywgd2hpY2ggZmVhdHVyZXMgYnJpZWYgc3VtbWFyaWVzIG9mIGFy
dGljbGVzIGFib3V0IE5BSUQtZnVuZGVkIHJlc2VhcmNoLCB1cGRhdGVkIHdlZWtseS4gVGhlIHNpdGUg
aXMgbm90YWJsZSBmb3IgaXRzIGV4dGVuc2l2ZSBzcGVjaWFsIHNlY3Rpb24gb24gdGhlIGdyb3dpbmcg
YmF0dGVyeSBvZiByZXNlYXJjaCBvbiBiaW9kZWZlbnNlIHN0cmF0ZWdpZXMuPC9kZXNjcmlwdGlvbj4K
ICAgICAgICAgIDxzdWJqZWN0IHhtbG5zPSJodHRwOi8vbnMubnNkbC5vcmcvbmNzL2xhciI+SGVhbHRo
L01lZGljaW5lPC9zdWJqZWN0PgogICAgICAgICAgPGxhbmd1YWdlPmVuLVVTPC9sYW5ndWFnZT4KICAg
ICAgICAgIDxmb3JtYXQgeG1sbnM9Imh0dHA6Ly9ucy5uc2RsLm9yZy9uY3MvbGFyIj5hcHBsaWNhdGlv
bi9wZGY8L2Zvcm1hdD4KICAgICAgICAgIDxhdWRpZW5jZVJlZmluZW1lbnQgeG1sbnM9Imh0dHA6Ly9u
cy5uc2RsLm9yZy9uY3MvbGFyIj5MZWFybmVyPC9hdWRpZW5jZVJlZmluZW1lbnQ+CiAgICAgICAgICA8
dHlwZSB4bWxucz0iaHR0cDovL25zLm5zZGwub3JnL25jcy9sYXIiPkluZm9ybWF0aXZlIFRleHQ8L3R5
cGU+CiAgICAgICAgICA8ZGF0ZSB0eXBlPSJQdWJsaXNoZWQiPjIwMDUtMTEtMDI8L2RhdGU+CiAgICAg
ICAgICA8Y29udHJpYnV0b3Igcm9sZT0iUHVibGlzaGVyIj5OYXRpb25hbCBJbnN0aXR1dGUgb2YgQWxs
ZXJneSBhbmQgSW5mZWN0aW91cyBEaXNlYXNlczwvY29udHJpYnV0b3I+CiAgICAgICAgICA8YWNjZXNz
UmVzdHJpY3Rpb25zPkZyZWUgYWNjZXNzPC9hY2Nlc3NSZXN0cmljdGlvbnM+CiAgICAgICAgICA8bGlj
ZW5zZT4KICAgICAgICAgICAgPG5hbWU+VW5rbm93bjwvbmFtZT4KICAgICAgICAgICAgPHByb3BlcnR5
PlRlcm1zIG9mIHVzZSB1bmtub3duPC9wcm9wZXJ0eT4KICAgICAgICAgIDwvbGljZW5zZT4KICAgICAg
ICA8L3JlY29yZD4KICAgICAg`, "\n", "", -1)
	encoded := base64.StdEncoding.EncodeToString(resp.ListRecords.Records[0].Metadata.Body)
	if encoded != data {
		t.Errorf("wrong metadata field value")
	}
}

func TestGetRecord(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/GetRecord-00.xml")
	if err != nil {
		t.Error(err)
	}
	var resp GetRecordResponse
	if err := xml.Unmarshal(b, &resp); err != nil {
		t.Error(err)
	}
	if resp.ResponseDate != "2017-09-11T13:56:47Z" {
		t.Error("wrong response date")
	}
	if resp.GetRecord.Record.Header.SetSpec[0] != "druckschriften.dl.ub.uni.freiburg.de" {
		t.Errorf("wrong set spec")
	}
	if len(resp.GetRecord.Record.Metadata.Body) != 11191 {
		t.Errorf("wrong metadata length, expected: 11191")
	}
}
