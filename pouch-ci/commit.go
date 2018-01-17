package main

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"time"
)

func (c *Client)GetFullCommit(opt *github.CommitsListOptions) ([]*github.RepositoryCommit, error) {
	cmit, resp, err := c.client.Repositories.ListCommits(context.Background(), c.cfg.Owner, c.cfg.Repo, opt)
	if err != nil {
		logrus.Errorf("get commit list fromt repo %s failed with\n error:%s\n response :%s\n", c.cfg.Repo, err, resp)
		return nil, err
	}
	return cmit, nil
}

// Get commit by the filter
func (c *Client)GetFilterCommit(time time.Time) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{
		Since: time,
	}

	cm, err := c.GetFullCommit(opt)
	if err != nil {
		logrus.Errorf("get commit failed with error:%s", err)
		return nil, err
	}

	return cm, nil
}
