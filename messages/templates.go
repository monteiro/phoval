package messages

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type TemplateFolderRender struct {
	// folder where the templates are
	TemplateFolder string
}

func (m TemplateFolderRender) Render(locale string, code string) (string, error) {
	f, err := filepath.Abs(m.TemplateFolder + "/" + locale + ".txt")
	if err != nil {
		return "", err
	}
	contents, err := ioutil.ReadFile(f)
	if err != nil {

		return "", err
	}

	resp := strings.Replace(string(contents), "{{code}}", code, 1)

	return resp, nil
}
