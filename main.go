package main

import (
	"github.com/Mindgamesnl/Greetings/greetings"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	greetings.LoadConfiguration()
	greetings.LoadLanguages()
	greetings.LoadGithubData()
	logrus.Info("Initialized data for " + strconv.Itoa(len(greetings.Repositories)))

	logrus.Info("Building HTML")
	projectsHtml := ""
	for i := range greetings.Repositories {
		projectsHtml += greetings.Repositories[i].ToHtml()
	}

	template := greetings.Read("resources/template.html")

	// basic variables
	template = strings.ReplaceAll(template, "{{picture}}", greetings.You.Avatar)
	template = strings.ReplaceAll(template, "{{user_name}}", greetings.You.Name)
	template = strings.ReplaceAll(template, "{{user_bio}}", greetings.You.Description)
	template = strings.ReplaceAll(template, "{{projects}}", projectsHtml)

	socials := ""
	for i := range greetings.LoadedInstance.Socials {
		element := greetings.LoadedInstance.Socials[i]
		socials += `<a class="badge" href="` + element.URL + `">` + element.Name + `</a>`
	}

	template = strings.ReplaceAll(template, "{{socials}}", socials)

	greetings.Write("out/index.html", template)
}
