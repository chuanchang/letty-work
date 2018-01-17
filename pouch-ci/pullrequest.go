package main

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

func (c *Client) GetFullPullRequest(opt *github.PullRequestListOptions) ([]*github.PullRequest, error) {

	pr, _, err := c.client.PullRequests.List(context.Background(), c.cfg.Owner, c.cfg.Repo, opt)

	//pullRequest, _, err := c.client.PullRequests.Get(context.Background(), c.cfg.Owner, c.cfg.Repo, number)
	if err != nil {
		logrus.Errorf("failed to get pull request in repo %s: %v", c.cfg.Repo, err)
		return nil, err
	}

	return pr, nil
}

// Get the latest number created Pull request.
func (c *Client) GetLatestNumCreatedPR(num int) ([]*github.PullRequest, error) {

	opt := &github.PullRequestListOptions{
		Sort:      "created",
		Direction: "asc",
	}

	pr, err := c.GetFullPullRequest(opt)
	return pr[0:num], err
}
