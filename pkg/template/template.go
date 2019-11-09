package template

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"
)

type Template struct {
	opts Options
}

func NewTemplate(options ...Option) *Template {
	opts := newOptions(options...)

	return &Template{opts: opts}
}

func (tpl *Template) Render(templateContext interface{}) ([]byte, error) {
	if !tpl.SourceFileExists() {
		return nil, fmt.Errorf("template file missing: %s", tpl.opts.Path)
	}

	buf, err := ioutil.ReadFile(tpl.opts.Path)
	if err != nil {
	    return nil, errors.Wrap(err, "unable to read template")
	}

	tmpl := template.New(tpl.opts.Name).Funcs(sprig.TxtFuncMap())
	tmpl, err = tmpl.Parse(string(buf))
	if err != nil {
	   	return nil, errors.Wrap(err, "unable to parse template")
	}

	out := bytes.Buffer{}
	if err := tmpl.Execute(&out, templateContext); err != nil {
		return nil, errors.Wrap(err, "template execution failed")
	}

	// format go source files
	if tpl.opts.GoSource {
		formatted, err := format.Source(out.Bytes())
		if err != nil {
		    return nil, errors.Wrap(err, "unable to format template output (maybe not a .go file)")
		}
		out = *bytes.NewBuffer(formatted)
	}

	return out.Bytes(), nil
}

// SourceFileExists returns 'true' if the source template file exists
func (tpl *Template) SourceFileExists() bool {
	if tpl.opts.Path == "" {
		return false
	}

	if _, err := os.Stat(tpl.opts.Path); os.IsNotExist(err) {
		return false
	}
	return true
}

type Option func(options *Options)
