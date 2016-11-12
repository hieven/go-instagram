package models_test

import (
	"strconv"
	"testing"

	"github.com/hieven/go-instagram/constants"
	. "github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/testUtils"
	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LocationTestSuite struct {
	suite.Suite

	ig       *Instagram
	location Location
	url      string
}

func TestLocationSuite(t *testing.T) {
	suite.Run(t, new(LocationTestSuite))
}

func (suite *LocationTestSuite) SetupSuite() {
	gorequest.DisableTransportSwap = true

	suite.ig = &Instagram{
		Username: "even",
		Password: "qweasd",
	}

	suite.location.Pk = 10

	suite.url = constants.ROUTES.LocationFeed + strconv.FormatInt(suite.location.Pk, 10) + "/"
}

func (suite *LocationTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (suite *LocationTestSuite) SetupTest() {
	pool, _ := testUtils.MockAgentPool(1)

	suite.ig.AgentPool = pool
	suite.location.Instagram = suite.ig
}

func (suite *LocationTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (suite *LocationTestSuite) TestGetRankedMediasSuccess() {
	assert := assert.New(suite.T())
	location := suite.location
	url := suite.url

	responder := testUtils.NewMockResponder(200, "locationGetMedias")
	httpmock.RegisterResponder("GET", url, responder)

	medias, err := location.GetRankedMedias()

	assert.NotEmpty(medias)
	assert.Nil(err)
}

func (suite *LocationTestSuite) TestGetRankedMediasLoginRequired() {
	assert := assert.New(suite.T())
	location := suite.location
	url := suite.url

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("GET", url, responder)

	medias, err := location.GetRankedMedias()

	assert.Len(medias, 0)
	assert.EqualError(err, "login_required")
}

func (suite *LocationTestSuite) TestGetRecentMediasSuccess() {
	assert := assert.New(suite.T())
	location := suite.location
	url := suite.url

	responder := testUtils.NewMockResponder(200, "locationGetMedias")
	httpmock.RegisterResponder("GET", url, responder)

	medias, err := location.GetRecentMedias()

	assert.NotEmpty(medias)
	assert.Nil(err)
}

func (suite *LocationTestSuite) TestGetRecentMediasLoginRequired() {
	assert := assert.New(suite.T())
	location := suite.location
	url := suite.url

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("GET", url, responder)

	medias, err := location.GetRecentMedias()

	assert.Len(medias, 0)
	assert.EqualError(err, "login_required")
}
