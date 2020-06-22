package greetings

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"os"
	"strconv"
)

var (
	You          User
	Repositories []Repository
)

type User struct {
	Handle      string
	Name        string
	Description string
	Avatar      string
}

func LoadGithubData() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: LoadedInstance.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, "")
	You = User{
		Name:        user.GetName(),
		Description: user.GetBio(),
		Avatar:      user.GetAvatarURL(),
		Handle:      user.GetLogin(),
	}

	if err != nil {
		logrus.Error("Invalid github access token.")
		os.Exit(1)
	}

	logrus.Info("Generating for " + user.GetName())

	//var options &github.RepositoryListOptions
	options := &github.RepositoryListOptions{
		Sort:        "pushed",
		Affiliation: "owner",
		ListOptions: github.ListOptions{
			PerPage: LoadedInstance.RepoCount,
		},
	}

	repos, _, _ := client.Repositories.List(ctx, "", options)

	logrus.Info("Repos: " + strconv.Itoa(len(repos)))

	for i := range repos {
		repo := repos[i]
		if (LoadedInstance.OnlyPublic && repo.GetPrivate()) || (LoadedInstance.HideEmptyDesc && repo.GetDescription() == "") {
			continue
		}

		var languageMeta LanguageInfo
		if repo.GetLanguage() == "" {
			languageMeta = LanguageInfo{
				Name: "Plain text / Unknown",
				Meta: Language{
					Color: "#777",
				},
			}
		} else {
			languageMeta = LanguageInfo{
				Name: repo.GetLanguage(),
				Meta: KnownLanguages.Languages[repo.GetLanguage()],
			}
		}

		Repositories = append(Repositories, Repository{
			Name:        repo.GetName(),
			Description: repo.GetDescription(),
			Language:    languageMeta,
			URL:         repo.GetURL(),
			Owner:       repo.GetOwner().GetName(),
		})
	}

}
