package messages

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Template(locale string, code string) (string, error) {
	f, _ := filepath.Abs("pkg/phoval/messages/" + locale + ".txt")
	contents, err := ioutil.ReadFile(f)
	if err != nil {

		return "", err
	}

	resp := strings.Replace(string(contents), "{{code}}", code, 1)

	return resp, nil
}
