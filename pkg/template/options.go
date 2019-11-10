package template

import "text/template"

type Options struct {
	Name     string
	Path     string
	GoSource bool
	FuncMap template.FuncMap
}

func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}

func Path(templatePath string) Option {
	return func(opts *Options) {
		opts.Path = templatePath
	}
}

func GoSource(isGoSource bool) Option {
	return func(opts *Options) {
		opts.GoSource = isGoSource
	}
}

func FuncMap(funcMap template.FuncMap) Option {
    return func(opts *Options) {
        opts.FuncMap = funcMap
    }
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Name:     "default_template",
		FuncMap: template.FuncMap{},
		GoSource: false,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
