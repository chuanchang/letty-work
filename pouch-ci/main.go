package main

import (
	"context"
	"os/exec"
	"time"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

const (
	//DefaultRepo="https://github.com/alibaba/pouch"
	DefaultRepo  = "pouch"
	DefaultOwner = "Letty5411"
)

type Client struct {
	client *github.Client
	cfg    Config
}

// Config refers
type Config struct {
	Owner       string
	Repo        string
	AccessToken string
}

func main() {
	var cfg Config
	var cmdServe = &cobra.Command{
		Use:  "",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			Run(cfg)
		},
	}

	flagSet := cmdServe.Flags()
	flagSet.StringVarP(&cfg.Owner, "owner", "o", DefaultOwner, "github ID to which connect in GitHub")
	flagSet.StringVarP(&cfg.Repo, "repo", "r", DefaultRepo, "github repo to which connect in GitHub")
	flagSet.StringVarP(&cfg.AccessToken, "token", "t", "", "access token to have some control on resources")

	cmdServe.Execute()
}

func Run(cfg Config) error {

	var c Client
	// Create an authenticated Client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	c.client = github.NewClient(tc)

	c.cfg = cfg

	t := time.Date(2018, time.January, 15, 23, 0, 0, 0, time.UTC)
	//start := time.Now()

	for {
		commit := make([]*github.RepositoryCommit, 100)
		commit, _ = c.GetFilterCommit(t)

		logrus.Println(t)
		logrus.Println(len(commit))

		if len(commit) != 0 {
			RunCI(commit)
		}

		// Get the current time and check if there is any update

		time.Sleep(600*time.Second)
		t = time.Now()
	}

	//baseUrl := *pr[0].GetBase().GetRepo().URL

	return nil
}

func RunCI(commit []*github.RepositoryCommit) {
	logrus.Println("In ci")
	for _, v := range commit {
		logrus.Printf("%s", v.GetSHA())

		cmd := exec.Command("java", "-jar", "jenkins-cli.jar", "-s",
			"http://tester:tester@11.160.112.29:8080/", "build", "-f", "-s", "-v", "-p", "commit="+v.GetSHA())
		logrus.Println(cmd)
		cmd.Start()
		err := cmd.Start()
		if err != nil {
			logrus.Errorf("%s", err)
		}
		logrus.Printf("Waiting for command to finish...")
		err = cmd.Wait()
		logrus.Printf("Command finished with error: %v", err)

	}
}
