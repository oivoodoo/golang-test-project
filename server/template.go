package server

import (
	"path/filepath"
	"io"
	"io/ioutil"
	"os"
	"strings"
	html "html/template"
	/*"fmt"*/
)

type TemplateLoader struct {
	paths []string
	templates map[string]Template
}

type Template interface {
	Name() string
	Content() string
	Render(wr io.Writer, args interface{}) error
}

func NewTemplateLoader(paths []string) *TemplateLoader {
	loader := &TemplateLoader{
		paths: paths,
	}
	return loader
}

func (loader *TemplateLoader) Refresh() *Error {
	loader.templates = map[string]Template{}

	for _, basePath := range loader.paths {
		// Walk only returns an error if the template loader is completely unusable
		// (namely, if one of the TemplateFuncs does not have an acceptable signature).
		filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				ERROR.Add(err.Error())
				return nil
			}

			if info.IsDir() {
				return nil
			}

			templateName := path[len(basePath)+1:]
			if os.PathSeparator == '\\' {
				templateName = strings.Replace(templateName, `\`, `/`, -1) // `
			}

			if _, ok := loader.templates[templateName]; ok {
				return nil
			}

			fileBytes, err := ioutil.ReadFile(path)
			if err != nil {
				ERROR.Add(err.Error())
				return nil
			}

			content := string(fileBytes)
			template, err := html.New(templateName).Parse(content)
			if err != nil {
				ERROR.Add(err.Error())
				return nil
			}

			loader.templates[templateName] = NewTemplate(templateName, content, template)

			return nil
		})
	}

	return nil
}

type HtmlTemplate struct {
	name string
	content string
	template *html.Template
}

func (template *HtmlTemplate) Content() string {
	return template.content
}

func (template *HtmlTemplate) Name() string {
	return template.name
}

func (t *HtmlTemplate) Render(wr io.Writer, args interface{}) error {
	t.template.Execute(wr, args)
	return nil
}

func NewTemplate(name string, content string, template *html.Template) Template {
	return Template(&HtmlTemplate {
		name: name,
		content: content,
		template: template,
	})
}
