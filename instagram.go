package instagram

import (
	"github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/utils"
)

func Create(username string, password string) (*models.Instagram, error) {
	pool, err := utils.NewSuperAgentPool(1)
	if err != nil {
		return nil, err
	}

	ig := &models.Instagram{
		Username:  username,
		Password:  password,
		AgentPool: pool,
	}

	ig.Inbox = &models.Inbox{Instagram: ig}
	ig.TimelineFeed = &models.TimelineFeed{
		Instagram:          ig,
		RankTokenGenerator: utils.RankTokenGenerator{},
	}

	return ig, nil
}
