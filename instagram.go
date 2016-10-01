package instagram

import (
	"log"

	"github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/utils"
	"github.com/parnurzeal/gorequest"
)

type Instagram struct {
	Username string
	Password string
	Request  *gorequest.SuperAgent
	Inbox    models.Inbox
	Thread   models.Thread
}

func Create(username string, password string) *Instagram {
	ig := Instagram{
		Username: username,
		Password: password,
		Request:  gorequest.New(),
	}

	ig.Login()

	return &ig
}

func (ig Instagram) Login() {
	log.Println("---------------->")
	log.Println("Method: Login")

	uuid := utils.GenerateUUID()

	login := models.Login{
		Csrftoken:         "missing",
		DeviceID:          "android-b256317fd493b848",
		UUID:              uuid,
		UserName:          ig.Username,
		Password:          ig.Password,
		LoginAttemptCount: 0,
		Request:           ig.Request,
	}

	login.Login()
}

func (ig Instagram) GetInboxFeed() []*models.Thread {
	inbox := &models.Inbox{
		Request: ig.Request,
	}

	return inbox.GetFeed()
}
