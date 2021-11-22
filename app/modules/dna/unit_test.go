package dna_test

import (
	"fmt"
	"testing"
	"x-men/app/modules/dna"
	"x-men/test/mocks"
	"x-men/test/testdata"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type dnaTestSuite struct {
	suite.Suite
	repository *mocks.MockRepository
	adapter    *mocks.MockHttpAdapter
	service    *mocks.MockDnaService
}

func TestDna(t *testing.T) {
	suite.Run(t, new(dnaTestSuite))
}

func (suite *dnaTestSuite) SetupTest() {
	repositoryCtrl := gomock.NewController(suite.T())
	adapterCtrl := gomock.NewController(suite.T())
	serviceCtrl := gomock.NewController(suite.T())
	defer repositoryCtrl.Finish()
	defer serviceCtrl.Finish()
	defer adapterCtrl.Finish()

	suite.repository = mocks.NewMockRepository(repositoryCtrl)
	suite.adapter = mocks.NewMockHttpAdapter(adapterCtrl)
	suite.service = mocks.NewMockDnaService(serviceCtrl)
}

func (suite *dnaTestSuite) TestControllerIsMutant() {
	controller := dna.NewController(suite.service)

	for i, testD := range testdata.ControllerIsMutant() {
		fmt.Printf("controller: index=%v description: %v\n", i, testD.Description)

		if testD.BodyIsOk {
			suite.adapter.EXPECT().GetBody().Return([]byte(testdata.BodyIsMutant))
			suite.service.EXPECT().IsMutant(gomock.Any()).Return(testD.IsMutant, testD.Error)
		} else {
			suite.adapter.EXPECT().GetBody().Return([]byte(""))
		}
		suite.adapter.EXPECT().SendResponse(testD.Status, gomock.Any())

		controller.IsMutant(suite.adapter)
	}
}

func (suite *dnaTestSuite) TestServiceIsMutant() {
	service := dna.NewService(suite.repository)

	for i, testD := range testdata.ServiceIsMutant() {
		fmt.Printf("service: index=%v description: %v\n", i, testD.Description)

		if testD.Error == nil {
			suite.repository.EXPECT().SaveDNA(testD.Dna, testD.IsMutant)
		}

		isMutant, err := service.IsMutant(testD.Dna)

		if (err != nil) || (testD.Error != nil) {
			suite.Equal(testD.Error, err, "should throw the same error")
		}

		suite.Equal(testD.Error, err, "should throw the same error")
		suite.Equal(testD.IsMutant, isMutant, "should throw the same response")
	}
}

func (suite *dnaTestSuite) TestModelReadDna() {
	dna := dna.DNA{}
	dna.LoadDNA(testdata.DnaNNN)
	for i, testD := range testdata.ModelReadDna() {
		fmt.Printf("model: index=%v description: read the %v position of the dna\n", i, i)
		suite.Equal(testD.CrossDiagonalInvert, dna.ReadCrossDiagonalDnaInvert(testD.Position), "CrossDiagonalInvert")
		suite.Equal(testD.CrossDiagonal, dna.ReadCrossDiagonalDna(testD.Position), "CrossDiagonal")
		suite.Equal(testD.DiagonalInvert, dna.ReadDiagonalDnaInvert(testD.Position), "DiagonalInvert")
		suite.Equal(testD.Diagonal, dna.ReadDiagonalDna(testD.Position), "Diagonal")
		suite.Equal(testD.Column, dna.ReadColumnDna(testD.Position), "Column")
		suite.Equal(testD.Row, dna.ReadRowDna(testD.Position), "Row")
	}
}
