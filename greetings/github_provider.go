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

	if err != nil {
		logrus.Error("Could not load github request. Is your token valid?")
		os.Exit(1)
	}

	You = User{
		user.GetLogin(),
		user.GetName(),
		user.GetBio(),
		user.GetAvatarURL(),
	}

	logrus.Info("Generating for " + user.GetName())

	var options = &github.RepositoryListOptions{
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
				"Plain text / Unknown",
				Language{
					Color: "#777",
				},
			}
		} else {
			languageMeta = LanguageInfo{
				repo.GetLanguage(),
				KnownLanguages.Languages[repo.GetLanguage()],
			}
		}

		Repositories = append(Repositories, Repository{
			repo.GetName(),
			repo.GetDescription(),
			languageMeta,
			repo.GetURL(),
			repo.GetOwner().GetName(),
		})
	}

}
