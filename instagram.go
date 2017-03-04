package instagram

import (
	"errors"
	"time"

	"github.com/hieven/go-instagram/config"
	"github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/utils"
)

func New(config *config.Config) (*models.Instagram, error) {
	if config.Username == "" {
		return nil, errors.New("username is required")
	}

	if config.Password == "" {
		return nil, errors.New("password is required")
	}

	if config.Capacity == 0 {
		config.Capacity = 1
	}

	if config.LoginInterval == 0 {
		config.LoginInterval = 1 * time.Second
	}

	pool, err := utils.NewSuperAgentPool(config.Capacity)

	if err != nil {
		return nil, err
	}

	ig := &models.Instagram{
		Config:    config,
		AgentPool: pool,
	}

	ig.Inbox = &models.Inbox{
		Instagram: ig,
	}

	ig.TimelineFeed = &models.TimelineFeed{
		Instagram:          ig,
		RankTokenGenerator: utils.RankTokenGenerator{},
	}

	if err := ig.Login(); err != nil {
		return nil, err
	}

	return ig, nil
}
