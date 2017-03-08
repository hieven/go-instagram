package instagram

import (
	"errors"

	"github.com/hieven/go-instagram/config"
	"github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/session"
	"github.com/hieven/go-instagram/utils"
)

func New(cnf *config.Config) (ig *models.Instagram, err error) {
	if cnf.Username == "" {
		return nil, errors.New("username is required")
	}

	if cnf.Password == "" {
		return nil, errors.New("password is required")
	}

	if cnf.Capacity == 0 {
		cnf.Capacity = 1
	}

	ig = &models.Instagram{
		Config: cnf,
	}

	ig.AgentPool, err = utils.NewSuperAgentPool(cnf.Capacity)

	ig.Session, err = session.NewSession(cnf)

	ig.Inbox = &models.Inbox{
		Instagram: ig,
	}

	ig.TimelineFeed = &models.TimelineFeed{
		Instagram:          ig,
		RankTokenGenerator: utils.RankTokenGenerator{},
	}

	err = ig.Login()

	if err != nil {
		return nil, err
	}

	return ig, nil
}
