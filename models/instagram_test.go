package models_test

import (
	"testing"

	"github.com/hieven/go-instagram/constants"
	. "github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/testUtils"
	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"

	"github.com/stretchr/testify/suite"
)

type InstagramTestSuite struct {
	suite.Suite

	ig    *Instagram
	inbox Inbox
}

func TestInstagramSuite(t *testing.T) {
	suite.Run(t, new(InstagramTestSuite))
}

func (suite *InstagramTestSuite) SetupSuite() {
	gorequest.DisableTransportSwap = true

	suite.ig = &Instagram{
		Username: "even",
		Password: "qweasd",
	}
}

func (suite *InstagramTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (suite *InstagramTestSuite) SetupTest() {
	pool, _ := testUtils.MockAgentPool(1)

	suite.ig.AgentPool = pool
}

func (suite *InstagramTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (suite *InstagramTestSuite) TestInstagramLoginSuccess() {
	ig := suite.ig

	responder := testUtils.NewMockResponder(200, "loginSuccess")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err := ig.Login()

	suite.Nil(err)
}

func (suite *InstagramTestSuite) TestInstagramLoginFailed() {
	ig := suite.ig

	tests := []struct {
		ResponseTemplate string
		ExpecredError    string
	}{
		{
			ResponseTemplate: "loginIncorrectPassword",
			ExpecredError:    "The password you entered is incorrect. Please try again.",
		},
		{
			ResponseTemplate: "loginIncorrectUsername",
			ExpecredError:    "The username you entered doesn't appear to belong to an account. Please check your username and try again.",
		},
		{
			ResponseTemplate: "loginMissingPassword",
			ExpecredError:    "Invalid Parameters",
		},
		{
			ResponseTemplate: "loginInvalidParameters",
			ExpecredError:    "Invalid Parameters",
		},
	}

	for _, test := range tests {
		responder := testUtils.NewMockResponder(400, test.ResponseTemplate)
		httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

		err := ig.Login()

		suite.EqualError(err, test.ExpecredError)

		httpmock.Reset()
	}
}

func (suite *InstagramTestSuite) TestInstagramLikeSuccess() {
	ig := suite.ig
	mediaID := "foo"
	url := constants.GetURL("Like", struct{ ID string }{ID: mediaID})

	responder := testUtils.NewMockResponder(200, "likeSuccess")
	httpmock.RegisterResponder("POST", url, responder)

	err := ig.Like(mediaID)

	suite.Nil(err)
}

func (suite *InstagramTestSuite) TestInstagramLikeLoginRequired() {
	ig := suite.ig
	mediaID := "foo"
	url := constants.GetURL("Like", struct{ ID string }{ID: mediaID})

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("POST", url, responder)

	err := ig.Like(mediaID)

	suite.EqualError(err, "login_required")
}

func (suite *InstagramTestSuite) TestInstagramUnlikeSuccess() {
	ig := suite.ig
	mediaID := "foo"
	url := constants.GetURL("Unlike", struct{ ID string }{ID: mediaID})

	responder := testUtils.NewMockResponder(200, "unlikeSuccess")
	httpmock.RegisterResponder("POST", url, responder)

	err := ig.Unlike(mediaID)

	suite.Nil(err)
}

func (suite *InstagramTestSuite) TestInstagramUnlikeLoginRequired() {
	ig := suite.ig
	mediaID := "foo"
	url := constants.GetURL("Unlike", struct{ ID string }{ID: mediaID})

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("POST", url, responder)

	err := ig.Unlike(mediaID)

	suite.EqualError(err, "login_required")
}

func (suite *InstagramTestSuite) TestInstagramCreateSignature() {
	// TODO: TestInstagramCreateSignature
}

func (suite *InstagramTestSuite) TestInstagramSendRequest() {
	// TODO: TestInstagramSendRequest
}
