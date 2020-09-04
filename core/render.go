package core

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func envVars() (map[string]interface{}, error) {
	resultMapping := make(map[string]interface{})
	envMapping := make(map[string]interface{})
	resultMapping["env"] = envMapping
	for _, value := range os.Environ() {
		vs := strings.SplitN(value, "=", 2)
		if len(vs) == 2 {
			envMapping[vs[0]] = vs[1]
		} else {
			envMapping[vs[0]] = ""
		}
	}
	return resultMapping, nil
}

func process(tmpl *template.Template, vars map[string]interface{}, buf io.Writer) (bool, error) {
	err := tmpl.Execute(buf, vars)
	if err != nil {
		logrus.Errorf("template.Execute(): %v", err)
		return false, err
	}
	return true, nil
}

func Render(srcFile, destFile string) (error, int) {
	vars, err := envVars()
	if err != nil {
		return errors.Annotate(err, "envVars"), 10
	}

	f, err := os.Stat(srcFile)
	if err != nil {
		return errors.Annotate(err, "statSrcFile"), 11
	}
	tmpl := template.New(f.Name()).Option("missingkey=error").Funcs(sprig.TxtFuncMap())
	tmpl, _ = tmpl.Parse(
		`{{- define "ext.scsv-quote" }}
		{{- . | replace "\"" "" | replace "'" "" | nospace | splitList "," | toJson | replace "[" "" | replace "]" "" }}
		{{- end }}`)
	tmpl, err = tmpl.ParseFiles(srcFile)
	if err != nil {
		return errors.Annotate(err, "parseSrcFile"), 11
	}

	buf := bytes.Buffer{}
	_, err = process(tmpl, vars, &buf)
	if err != nil {
		return errors.Annotate(err, "processTemplate"), 12
	}

	err = ioutil.WriteFile(destFile, buf.Bytes(), 0644)
	if err != nil {
		return errors.Annotate(err, "writeFile"), 13
	}
	return nil, 0
}
