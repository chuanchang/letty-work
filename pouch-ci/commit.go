package main

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

func (c *Client) GetFullCommit(opt *github.CommitsListOptions) ([]*github.RepositoryCommit, error) {
	cmit, resp, err := c.client.Repositories.ListCommits(context.Background(), c.cfg.Owner, c.cfg.Repo, opt)
	if err != nil {
		logrus.Errorf("get commit list fromt repo %s failed with\n error:%s\n response :%s\n", c.cfg.Repo, err, resp)
		return nil, err
	}
	return cmit, nil
}

// Get commit by the filter
func (c *Client) GetFilterCommit(time time.Time) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{
		SHA:   "master",
		Since: time,
	}

	cm, err := c.GetFullCommit(opt)
	if err != nil {
		logrus.Errorf("get commit failed with error:%s", err)
		return nil, err
	}

	// reverse the commit so it is sorted by date.
	for i := 0; i < len(cm)/2; i++ {
		j := len(cm) - i - 1
		cm[i], cm[j] = cm[j], cm[i]
	}

	return cm, nil
}
