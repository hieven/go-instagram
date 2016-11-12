package models_test

import (
	"testing"

	"github.com/hieven/go-instagram/constants"
	. "github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/testUtils"
	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"

	"github.com/stretchr/testify/assert"
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
	assert := assert.New(suite.T())
	ig := suite.ig

	responder := testUtils.NewMockResponder(200, "loginSuccess")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err := ig.Login()

	assert.Nil(err)
}

func (suite *InstagramTestSuite) TestInstagramLoginFailed() {
	assert := assert.New(suite.T())
	ig := suite.ig

	responder := testUtils.NewMockResponder(400, "loginIncorrectPassword")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err := ig.Login()

	assert.EqualError(err, "The password you entered is incorrect. Please try again.")

	httpmock.Reset()

	responder = testUtils.NewMockResponder(400, "loginIncorrectUsername")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err = ig.Login()

	assert.EqualError(err, "The username you entered doesn't appear to belong to an account. Please check your username and try again.")

	httpmock.Reset()

	responder = testUtils.NewMockResponder(400, "loginMissingPassword")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err = ig.Login()

	assert.EqualError(err, "Invalid Parameters")

	httpmock.Reset()

	responder = testUtils.NewMockResponder(400, "loginInvalidParameters")
	httpmock.RegisterResponder("POST", constants.ROUTES.Login, responder)

	err = ig.Login()

	assert.EqualError(err, "Invalid Parameters")

	httpmock.Reset()
}

func (suite *InstagramTestSuite) TestInstagramCreateSignature() {
}

func (suite *InstagramTestSuite) TestInstagramSendRequest() {
}
