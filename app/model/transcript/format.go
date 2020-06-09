package transcript

import (
	"encoding/json"

	"github.com/kyleu/rituals.dev/app/model/auth"
)

type Format struct {
	Key   string
	Title string
	Icon  string
}

var FormatPrint = Format{
	Key:   "print",
	Title: "Print",
	Icon:  "print",
}

var FormatPDF = Format{
	Key:   "pdf",
	Title: "Download PDF",
	Icon:  "file-pdf",
}

var FormatJSON = Format{
	Key:   "json",
	Title: "Download JSON",
	Icon:  "code",
}

var FormatExcel = Format{
	Key:   "excel",
	Title: "Download Excel",
	Icon:  "file-text",
}

var FormatAsana = Format{
	Key:   "asana",
	Title: "Export to Asana",
	Icon:  "more",
}

var FormatConfluence = Format{
	Key:   "confluence",
	Title: "Export to Confluence",
	Icon:  "file-edit",
}

var FormatGitHub = Format{
	Key:   "github",
	Title: "Export to GitHub",
	Icon:  auth.ProviderGitHub.Icon,
}

var FormatGoogleDocs = Format{
	Key:   "googledocs",
	Title: "Export to Google Docs",
	Icon:  auth.ProviderGoogle.Icon,
}

var FormatSlack = Format{
	Key:   "slack",
	Title: "Export to Slack",
	Icon:  auth.ProviderSlack.Icon,
}

var FormatTrello = Format{
	Key:   "trello",
	Title: "Export to Trello",
	Icon:  "album",
}

var PrintFormats = []Format{FormatPrint, FormatPDF, FormatJSON, FormatExcel}

var AllFormats = []Format{
	FormatPrint, FormatPDF, FormatJSON, FormatExcel,
	FormatAsana, FormatConfluence, FormatGitHub, FormatGoogleDocs, FormatSlack, FormatTrello,
}

func FormatFromString(s string) Format {
	for _, t := range AllFormats {
		if t.Key == s {
			return t
		}
	}
	return FormatPrint
}

func (t *Format) String() string {
	return t.Key
}

func (t Format) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Format) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = FormatFromString(s)
	return nil
}
