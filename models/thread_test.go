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

type ThreadTestSuite struct {
	suite.Suite

	ig     *Instagram
	thread *Thread
}

func TestThreadSuite(t *testing.T) {
	suite.Run(t, new(ThreadTestSuite))
}

func (suite *ThreadTestSuite) SetupSuite() {
	gorequest.DisableTransportSwap = true

	suite.ig = &Instagram{
		Username: "even",
		Password: "qweasd",
	}

	suite.thread = &Thread{
		ID: "ID",
	}
}

func (suite *ThreadTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (suite *ThreadTestSuite) SetupTest() {
	pool, _ := testUtils.MockAgentPool(1)

	suite.ig.AgentPool = pool
	suite.thread.Instagram = suite.ig
}

func (suite *ThreadTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (suite *ThreadTestSuite) TestThreadBroadcastTextSuccess() {
	assert := assert.New(suite.T())
	thread := suite.thread

	responder := testUtils.NewMockResponder(200, "locationGetMedias")
	httpmock.RegisterResponder("POST", constants.ROUTES.ThreadsBroadcastText, responder)

	err := thread.BroadcastText("something")

	assert.Nil(err)
}

func (suite *ThreadTestSuite) TestThreadBroadcastTextMissingText() {
	assert := assert.New(suite.T())
	thread := suite.thread

	responder := testUtils.NewMockResponder(200, "threadBroadcastTextMissingText")
	httpmock.RegisterResponder("POST", constants.ROUTES.ThreadsBroadcastText, responder)

	err := thread.BroadcastText("")

	assert.EqualError(err, "Text is missing")
}

func (suite *ThreadTestSuite) TestThreadBroadcastTextLoginRequired() {
	assert := assert.New(suite.T())
	thread := suite.thread

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("POST", constants.ROUTES.ThreadsBroadcastText, responder)

	err := thread.BroadcastText("something")

	assert.EqualError(err, "login_required")
}

func (suite *ThreadTestSuite) TestThreadShowSuccess() {
	assert := assert.New(suite.T())
	thread := suite.thread

	url := constants.ROUTES.ThreadsShow + thread.ID + "/"

	responder := testUtils.NewMockResponder(200, "threadShowSuccess")
	httpmock.RegisterResponder("GET", url, responder)

	thread, err := thread.Show()

	assert.NotEmpty(thread.Items)
	assert.Nil(err)
}

func (suite *ThreadTestSuite) TestThreadShowLoginRequired() {
	assert := assert.New(suite.T())
	thread := suite.thread

	url := constants.ROUTES.ThreadsShow + thread.ID + "/"

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("GET", url, responder)

	thread, err := thread.Show()

	assert.Nil(thread)
	assert.EqualError(err, "login_required")
}
