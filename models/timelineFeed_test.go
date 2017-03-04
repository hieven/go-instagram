package models_test

import (
	"strconv"
	"testing"

	"github.com/hieven/go-instagram/config"
	"github.com/hieven/go-instagram/constants"
	. "github.com/hieven/go-instagram/models"
	"github.com/hieven/go-instagram/testUtils"
	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/suite"
)

type TimelineFeedTestSuite struct {
	suite.Suite

	ig   *Instagram
	feed *TimelineFeed
}

type fakeRankTokenGenerator struct{}

func (generator fakeRankTokenGenerator) GenerateRankToken(userID string) string {
	return userID
}

func TestTimelineSuite(t *testing.T) {
	suite.Run(t, new(TimelineFeedTestSuite))
}

func (suite *TimelineFeedTestSuite) SetupSuite() {
	gorequest.DisableTransportSwap = true

	config := &config.Config{
		Username: "even",
		Password: "password",
	}

	suite.ig = &Instagram{
		Config: config,
	}

	suite.ig.Pk = 1

	suite.feed = &TimelineFeed{
		RankTokenGenerator: fakeRankTokenGenerator{},
	}
}

func (suite *TimelineFeedTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (suite *TimelineFeedTestSuite) SetupTest() {
	pool, _ := testUtils.MockAgentPool(1)

	suite.ig.AgentPool = pool
	suite.feed.Instagram = suite.ig
}

func (suite *TimelineFeedTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (suite *TimelineFeedTestSuite) TestTimelineFeedSuccess() {
	feed := suite.feed

	userID := strconv.FormatInt(suite.ig.Pk, 10)
	rankToken := feed.GenerateRankToken(userID)
	url := constants.GetURL("TimelineFeed", struct {
		MaxID     string
		RankToken string
	}{MaxID: feed.Cursor, RankToken: rankToken})

	responder := testUtils.NewMockResponder(200, "timelineFeed")
	httpmock.RegisterResponder("GET", url, responder)

	items, err := feed.Get()

	suite.Len(items, 2)
	suite.NoError(err)
}

func (suite *TimelineFeedTestSuite) TestTimelineFeedLoginRequired() {
	userID := strconv.FormatInt(suite.ig.Pk, 10)
	rankToken := suite.feed.GenerateRankToken(userID)
	url := constants.GetURL("TimelineFeed", struct {
		MaxID     string
		RankToken string
	}{
		MaxID:     suite.feed.Cursor,
		RankToken: rankToken,
	})

	responder := testUtils.NewMockResponder(400, "loginRequired")
	httpmock.RegisterResponder("GET", url, responder)

	items, err := suite.feed.Get()

	suite.Nil(items)
	suite.EqualError(err, "login_required")
}

func (suite *TimelineFeedTestSuite) TestTimelineFeedSetCursor() {
	suite.feed.SetCursor("test")

	suite.EqualValues("test", suite.feed.Cursor)

	suite.feed.Cursor = ""
}

func (suite *TimelineFeedTestSuite) TestTimelineFeedGetCursor() {
	cursor := suite.feed.GetCursor()

	suite.EqualValues(suite.feed.Cursor, cursor)
}

func (suite *TimelineFeedTestSuite) TestTimelineFeedIsMoreAvailable() {
	moreAvailable := suite.feed.IsMoreAvailable()

	suite.EqualValues(suite.feed.MoreAvailable, moreAvailable)
}
