package instagram

import (
	"github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/utils"
)

func Create(username string, password string) (*models.Instagram, error) {
	return BulkCreate(username, password, 1, 0)
}

func BulkCreate(username string, password string, capacity int, loginInterval int) (*models.Instagram, error) {
	pool, err := utils.NewSuperAgentPool(capacity)

	if err != nil {
		return nil, err
	}

	ig := &models.Instagram{
		Username:      username,
		Password:      password,
		LoginInterval: loginInterval,
		AgentPool:     pool,
	}

	ig.Inbox = &models.Inbox{
		Instagram: ig,
	}

	ig.TimelineFeed = &models.TimelineFeed{
		Instagram:          ig,
		RankTokenGenerator: utils.RankTokenGenerator{},
	}

	return ig, nil
}
