package statistic_test

import (
	"fmt"
	"testing"
	"x-men/app/modules/statistic"
	"x-men/test/mocks"
	"x-men/test/testdata"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type statisticTestSuite struct {
	suite.Suite
	repository *mocks.MockRepository
	adapter    *mocks.MockHttpAdapter
	service    *mocks.MockStatisticService
}

func TestStatistic(t *testing.T) {
	suite.Run(t, new(statisticTestSuite))
}

func (suite *statisticTestSuite) SetupTest() {
	repositoryCtrl := gomock.NewController(suite.T())
	adapterCtrl := gomock.NewController(suite.T())
	serviceCtrl := gomock.NewController(suite.T())
	defer repositoryCtrl.Finish()
	defer serviceCtrl.Finish()
	defer adapterCtrl.Finish()

	suite.repository = mocks.NewMockRepository(repositoryCtrl)
	suite.adapter = mocks.NewMockHttpAdapter(adapterCtrl)
	suite.service = mocks.NewMockStatisticService(serviceCtrl)
}

func (suite *statisticTestSuite) TestControllerGetStatistic() {
	controller := statistic.NewController(suite.service)

	for i, testD := range testdata.ControllerGetStatistic() {
		fmt.Printf("controller: index=%v description: %v\n", i, testD.Description)

		suite.service.EXPECT().GetStatistic().Return(testD.Response, testD.Error)
		if testD.Error != nil {
			suite.adapter.EXPECT().SendResponse(testD.Status, testD.Error.Error())
		} else {
			suite.adapter.EXPECT().SendResponse(testD.Status, testD.Response)
		}

		controller.GetStatistic(suite.adapter)
	}
}

func (suite *statisticTestSuite) TestServiceGetStatistic() {
	service := statistic.NewService(suite.repository)

	for _, testD := range testdata.ServiceGetStatistic() {
		fmt.Println(testD.Description)
		suite.repository.EXPECT().GetStatistic().Return(testD.Response.CountMutantDNA, testD.Response.CountHumanDNA, testD.Error)

		response, err := service.GetStatistic()

		suite.Equal(testD.Error, err, "should throw the same error")
		suite.Equal(testD.Response, response, "should throw the same response")
	}
}
