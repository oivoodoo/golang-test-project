package music

import (
	"path/filepath"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
			loader.templates[templateName] = NewTemplate(templateName, content)

			return nil
		})
	}

	return nil
}

type HtmlTemplate struct {
	name string
	content string
}

func (template *HtmlTemplate) Content() string {
	return template.content
}

func (template *HtmlTemplate) Name() string {
	return template.name
}

func (template *HtmlTemplate) Render(wr io.Writer, args interface{}) error {
	return nil
}

func NewTemplate(name string, content string) Template {
	template := &HtmlTemplate { name: name, content: content }
	return Template(template)
}

