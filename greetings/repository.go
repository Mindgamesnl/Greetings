package greetings

import "strings"

type LanguageInfo struct {
	Name string
	Meta Language
}

type Repository struct {
	Name        string
	Description string
	Language    LanguageInfo
	URL         string
	Owner       string
}

func (repository Repository) ToHtml() string {
	html := Read("resources/project.html")

	html = strings.ReplaceAll(html, "{{account_handle}}", You.Handle)
	html = strings.ReplaceAll(html, "{{repo_name}}", repository.Name)
	html = strings.ReplaceAll(html, "{{repo_description}}", repository.Description)
	html = strings.ReplaceAll(html, "{{language_color}}", repository.Language.Meta.Color)
	html = strings.ReplaceAll(html, "{{language_name}}", repository.Language.Name)

	return html
}
