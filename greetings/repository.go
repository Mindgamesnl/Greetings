package greetings

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
	html := `<a class="project-item" href="https://github.com/` + You.Handle + "/" + repository.Name + `">
        <div class="item-inner">
          <div class="content">
            <h3>` + repository.Name + `</h3>
            <p>` + repository.Description + `</p>
          </div>
          <div class="item-footer">
            <b><div class="item-lang" style="color:` + repository.Language.Meta.Color + `"><span>` + repository.Language.Name + `</span></div></b>
          </div>
        </div></a>`

	return html
}
