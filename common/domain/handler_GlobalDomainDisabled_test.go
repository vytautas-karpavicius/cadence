// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package domain

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/archiver"
	"github.com/uber/cadence/common/archiver/provider"
	"github.com/uber/cadence/common/clock"
	"github.com/uber/cadence/common/cluster"
	"github.com/uber/cadence/common/log/loggerimpl"
	"github.com/uber/cadence/common/mocks"
	"github.com/uber/cadence/common/persistence"
	persistencetests "github.com/uber/cadence/common/persistence/persistence-tests"
	"github.com/uber/cadence/common/service/config"
	dc "github.com/uber/cadence/common/service/dynamicconfig"
	"github.com/uber/cadence/common/types"
)

type (
	domainHandlerGlobalDomainDisabledSuite struct {
		suite.Suite
		persistencetests.TestBase

		minRetentionDays     int
		maxBadBinaryCount    int
		metadataMgr          persistence.MetadataManager
		mockProducer         *mocks.KafkaProducer
		mockDomainReplicator Replicator
		archivalMetadata     archiver.ArchivalMetadata
		mockArchiverProvider *provider.MockArchiverProvider

		handler *handlerImpl
	}
)

func TestDomainHandlerGlobalDomainDisabledSuite(t *testing.T) {
	s := new(domainHandlerGlobalDomainDisabledSuite)
	suite.Run(t, s)
}

func (s *domainHandlerGlobalDomainDisabledSuite) SetupSuite() {
	if testing.Verbose() {
		log.SetOutput(os.Stdout)
	}

	s.TestBase = persistencetests.NewTestBaseWithCassandra(&persistencetests.TestBaseOptions{
		ClusterMetadata: cluster.GetTestClusterMetadata(false, false),
	})
	s.TestBase.Setup()
}

func (s *domainHandlerGlobalDomainDisabledSuite) TearDownSuite() {
	s.TestBase.TearDownWorkflowStore()
}

func (s *domainHandlerGlobalDomainDisabledSuite) SetupTest() {
	logger := loggerimpl.NewNopLogger()
	dcCollection := dc.NewCollection(dc.NewNopClient(), logger)
	s.minRetentionDays = 1
	s.maxBadBinaryCount = 10
	s.metadataMgr = s.TestBase.MetadataManager
	s.mockProducer = &mocks.KafkaProducer{}
	s.mockDomainReplicator = NewDomainReplicator(s.mockProducer, logger)
	s.archivalMetadata = archiver.NewArchivalMetadata(
		dcCollection,
		"",
		false,
		"",
		false,
		&config.ArchivalDomainDefaults{},
	)
	domainConfig := Config{
		MinRetentionDays:  dc.GetIntPropertyFn(s.minRetentionDays),
		MaxBadBinaryCount: dc.GetIntPropertyFilteredByDomain(s.maxBadBinaryCount),
		FailoverCoolDown:  dc.GetDurationPropertyFnFilteredByDomain(0 * time.Second),
	}
	s.mockArchiverProvider = &provider.MockArchiverProvider{}
	s.handler = NewHandler(
		domainConfig,
		logger,
		s.metadataMgr,
		s.ClusterMetadata,
		s.mockDomainReplicator,
		s.archivalMetadata,
		s.mockArchiverProvider,
		clock.NewRealTimeSource(),
	).(*handlerImpl)
}

func (s *domainHandlerGlobalDomainDisabledSuite) TearDownTest() {
	s.mockProducer.AssertExpectations(s.T())
	s.mockArchiverProvider.AssertExpectations(s.T())
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestRegisterGetDomain_InvalidGlobalDomain() {
	domainName := s.getRandomDomainName()
	description := "some random description"
	email := "some random email"
	retention := int32(7)
	emitMetric := true
	activeClusterName := cluster.TestCurrentClusterName
	clusters := []*types.ClusterReplicationConfiguration{
		&types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(activeClusterName),
		},
	}
	data := map[string]string{"some random key": "some random value"}
	isGlobalDomain := true

	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		Description:                            common.StringPtr(description),
		OwnerEmail:                             common.StringPtr(email),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(emitMetric),
		Clusters:                               clusters,
		ActiveClusterName:                      common.StringPtr(activeClusterName),
		Data:                                   data,
		IsGlobalDomain:                         common.BoolPtr(isGlobalDomain),
	})
	s.NotNil(err)
	s.IsType(&types.BadRequestError{}, err)
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestRegisterGetDomain_InvalidCluster() {
	domainName := s.getRandomDomainName()
	description := "some random description"
	email := "some random email"
	retention := int32(7)
	emitMetric := true
	activeClusterName := cluster.TestAlternativeClusterName
	clusters := []*types.ClusterReplicationConfiguration{
		&types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(activeClusterName),
		},
	}
	data := map[string]string{"some random key": "some random value"}
	isGlobalDomain := false

	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		Description:                            common.StringPtr(description),
		OwnerEmail:                             common.StringPtr(email),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(emitMetric),
		Clusters:                               clusters,
		ActiveClusterName:                      common.StringPtr(activeClusterName),
		Data:                                   data,
		IsGlobalDomain:                         common.BoolPtr(isGlobalDomain),
	})
	s.IsType(&types.BadRequestError{}, err)
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestRegisterGetDomain_AllDefault() {
	domainName := s.getRandomDomainName()
	var clusters []*types.ClusterReplicationConfiguration
	for _, replicationConfig := range persistence.GetOrUseDefaultClusters(s.ClusterMetadata.GetCurrentClusterName(), nil) {
		clusters = append(clusters, &types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(replicationConfig.ClusterName),
		})
	}

	retention := int32(1)
	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
	})
	s.Nil(err)

	resp, err := s.handler.DescribeDomain(context.Background(), &types.DescribeDomainRequest{
		Name: common.StringPtr(domainName),
	})
	s.Nil(err)

	s.NotEmpty(resp.DomainInfo.GetUUID())
	resp.DomainInfo.UUID = common.StringPtr("")
	s.Equal(&types.DomainInfo{
		Name:        common.StringPtr(domainName),
		Status:      types.DomainStatusRegistered.Ptr(),
		Description: common.StringPtr(""),
		OwnerEmail:  common.StringPtr(""),
		Data:        map[string]string{},
		UUID:        common.StringPtr(""),
	}, resp.DomainInfo)
	s.Equal(&types.DomainConfiguration{
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(true),
		HistoryArchivalStatus:                  types.ArchivalStatusDisabled.Ptr(),
		HistoryArchivalURI:                     common.StringPtr(""),
		VisibilityArchivalStatus:               types.ArchivalStatusDisabled.Ptr(),
		VisibilityArchivalURI:                  common.StringPtr(""),
		BadBinaries:                            &types.BadBinaries{Binaries: map[string]*types.BadBinaryInfo{}},
	}, resp.Configuration)
	s.Equal(&types.DomainReplicationConfiguration{
		ActiveClusterName: common.StringPtr(s.ClusterMetadata.GetCurrentClusterName()),
		Clusters:          clusters,
	}, resp.ReplicationConfiguration)
	s.Equal(common.EmptyVersion, resp.GetFailoverVersion())
	s.False(resp.GetIsGlobalDomain())
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestRegisterGetDomain_NoDefault() {
	domainName := s.getRandomDomainName()
	description := "some random description"
	email := "some random email"
	retention := int32(7)
	emitMetric := true
	activeClusterName := cluster.TestCurrentClusterName
	clusters := []*types.ClusterReplicationConfiguration{
		&types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(activeClusterName),
		},
	}
	data := map[string]string{"some random key": "some random value"}
	isGlobalDomain := false

	var expectedClusters []*types.ClusterReplicationConfiguration
	for _, replicationConfig := range persistence.GetOrUseDefaultClusters(s.ClusterMetadata.GetCurrentClusterName(), nil) {
		expectedClusters = append(expectedClusters, &types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(replicationConfig.ClusterName),
		})
	}

	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		Description:                            common.StringPtr(description),
		OwnerEmail:                             common.StringPtr(email),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(emitMetric),
		Clusters:                               clusters,
		ActiveClusterName:                      common.StringPtr(activeClusterName),
		Data:                                   data,
		IsGlobalDomain:                         common.BoolPtr(isGlobalDomain),
	})
	s.Nil(err)

	resp, err := s.handler.DescribeDomain(context.Background(), &types.DescribeDomainRequest{
		Name: common.StringPtr(domainName),
	})
	s.Nil(err)

	s.NotEmpty(resp.DomainInfo.GetUUID())
	resp.DomainInfo.UUID = common.StringPtr("")
	s.Equal(&types.DomainInfo{
		Name:        common.StringPtr(domainName),
		Status:      types.DomainStatusRegistered.Ptr(),
		Description: common.StringPtr(description),
		OwnerEmail:  common.StringPtr(email),
		Data:        data,
		UUID:        common.StringPtr(""),
	}, resp.DomainInfo)
	s.Equal(&types.DomainConfiguration{
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(emitMetric),
		HistoryArchivalStatus:                  types.ArchivalStatusDisabled.Ptr(),
		HistoryArchivalURI:                     common.StringPtr(""),
		VisibilityArchivalStatus:               types.ArchivalStatusDisabled.Ptr(),
		VisibilityArchivalURI:                  common.StringPtr(""),
		BadBinaries:                            &types.BadBinaries{Binaries: map[string]*types.BadBinaryInfo{}},
	}, resp.Configuration)
	s.Equal(&types.DomainReplicationConfiguration{
		ActiveClusterName: common.StringPtr(s.ClusterMetadata.GetCurrentClusterName()),
		Clusters:          expectedClusters,
	}, resp.ReplicationConfiguration)
	s.Equal(common.EmptyVersion, resp.GetFailoverVersion())
	s.Equal(isGlobalDomain, resp.GetIsGlobalDomain())
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestUpdateGetDomain_NoAttrSet() {
	domainName := s.getRandomDomainName()
	description := "some random description"
	email := "some random email"
	retention := int32(7)
	emitMetric := true
	data := map[string]string{"some random key": "some random value"}
	var clusters []*types.ClusterReplicationConfiguration
	for _, replicationConfig := range persistence.GetOrUseDefaultClusters(s.ClusterMetadata.GetCurrentClusterName(), nil) {
		clusters = append(clusters, &types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(replicationConfig.ClusterName),
		})
	}

	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		Description:                            common.StringPtr(description),
		OwnerEmail:                             common.StringPtr(email),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
		EmitMetric:                             common.BoolPtr(emitMetric),
		Clusters:                               clusters,
		ActiveClusterName:                      common.StringPtr(s.ClusterMetadata.GetCurrentClusterName()),
		Data:                                   data,
	})
	s.Nil(err)

	fnTest := func(info *types.DomainInfo, config *types.DomainConfiguration,
		replicationConfig *types.DomainReplicationConfiguration, isGlobalDomain bool, failoverVersion int64) {
		s.NotEmpty(info.GetUUID())
		info.UUID = common.StringPtr("")
		s.Equal(&types.DomainInfo{
			Name:        common.StringPtr(domainName),
			Status:      types.DomainStatusRegistered.Ptr(),
			Description: common.StringPtr(description),
			OwnerEmail:  common.StringPtr(email),
			Data:        data,
			UUID:        common.StringPtr(""),
		}, info)
		s.Equal(&types.DomainConfiguration{
			WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
			EmitMetric:                             common.BoolPtr(emitMetric),
			HistoryArchivalStatus:                  types.ArchivalStatusDisabled.Ptr(),
			HistoryArchivalURI:                     common.StringPtr(""),
			VisibilityArchivalStatus:               types.ArchivalStatusDisabled.Ptr(),
			VisibilityArchivalURI:                  common.StringPtr(""),
			BadBinaries:                            &types.BadBinaries{Binaries: map[string]*types.BadBinaryInfo{}},
		}, config)
		s.Equal(&types.DomainReplicationConfiguration{
			ActiveClusterName: common.StringPtr(s.ClusterMetadata.GetCurrentClusterName()),
			Clusters:          clusters,
		}, replicationConfig)
		s.Equal(common.EmptyVersion, failoverVersion)
		s.False(isGlobalDomain)
	}

	updateResp, err := s.handler.UpdateDomain(context.Background(), &types.UpdateDomainRequest{
		Name: common.StringPtr(domainName),
	})
	s.Nil(err)
	fnTest(
		updateResp.DomainInfo,
		updateResp.Configuration,
		updateResp.ReplicationConfiguration,
		updateResp.GetIsGlobalDomain(),
		updateResp.GetFailoverVersion(),
	)

	getResp, err := s.handler.DescribeDomain(context.Background(), &types.DescribeDomainRequest{
		Name: common.StringPtr(domainName),
	})
	s.Nil(err)
	fnTest(
		getResp.DomainInfo,
		getResp.Configuration,
		getResp.ReplicationConfiguration,
		getResp.GetIsGlobalDomain(),
		getResp.GetFailoverVersion(),
	)
}

func (s *domainHandlerGlobalDomainDisabledSuite) TestUpdateGetDomain_AllAttrSet() {
	domainName := s.getRandomDomainName()
	err := s.handler.RegisterDomain(context.Background(), &types.RegisterDomainRequest{
		Name:                                   common.StringPtr(domainName),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(1),
	})
	s.Nil(err)

	description := "some random description"
	email := "some random email"
	retention := int32(7)
	emitMetric := true
	activeClusterName := cluster.TestCurrentClusterName
	clusters := []*types.ClusterReplicationConfiguration{
		&types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(activeClusterName),
		},
	}
	data := map[string]string{"some random key": "some random value"}

	var expectedClusters []*types.ClusterReplicationConfiguration
	for _, replicationConfig := range persistence.GetOrUseDefaultClusters(s.ClusterMetadata.GetCurrentClusterName(), nil) {
		expectedClusters = append(expectedClusters, &types.ClusterReplicationConfiguration{
			ClusterName: common.StringPtr(replicationConfig.ClusterName),
		})
	}

	fnTest := func(info *types.DomainInfo, config *types.DomainConfiguration,
		replicationConfig *types.DomainReplicationConfiguration, isGlobalDomain bool, failoverVersion int64) {
		s.NotEmpty(info.GetUUID())
		info.UUID = common.StringPtr("")
		s.Equal(&types.DomainInfo{
			Name:        common.StringPtr(domainName),
			Status:      types.DomainStatusRegistered.Ptr(),
			Description: common.StringPtr(description),
			OwnerEmail:  common.StringPtr(email),
			Data:        data,
			UUID:        common.StringPtr(""),
		}, info)
		s.Equal(&types.DomainConfiguration{
			WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
			EmitMetric:                             common.BoolPtr(emitMetric),
			HistoryArchivalStatus:                  types.ArchivalStatusDisabled.Ptr(),
			HistoryArchivalURI:                     common.StringPtr(""),
			VisibilityArchivalStatus:               types.ArchivalStatusDisabled.Ptr(),
			VisibilityArchivalURI:                  common.StringPtr(""),
			BadBinaries:                            &types.BadBinaries{Binaries: map[string]*types.BadBinaryInfo{}},
		}, config)
		s.Equal(&types.DomainReplicationConfiguration{
			ActiveClusterName: common.StringPtr(s.ClusterMetadata.GetCurrentClusterName()),
			Clusters:          expectedClusters,
		}, replicationConfig)
		s.Equal(common.EmptyVersion, failoverVersion)
		s.False(isGlobalDomain)
	}

	updateResp, err := s.handler.UpdateDomain(context.Background(), &types.UpdateDomainRequest{
		Name: common.StringPtr(domainName),
		UpdatedInfo: &types.UpdateDomainInfo{
			Description: common.StringPtr(description),
			OwnerEmail:  common.StringPtr(email),
			Data:        data,
		},
		Configuration: &types.DomainConfiguration{
			WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(retention),
			EmitMetric:                             common.BoolPtr(emitMetric),
			HistoryArchivalStatus:                  types.ArchivalStatusDisabled.Ptr(),
			HistoryArchivalURI:                     common.StringPtr(""),
			VisibilityArchivalStatus:               types.ArchivalStatusDisabled.Ptr(),
			VisibilityArchivalURI:                  common.StringPtr(""),
			BadBinaries:                            &types.BadBinaries{Binaries: map[string]*types.BadBinaryInfo{}},
		},
		ReplicationConfiguration: &types.DomainReplicationConfiguration{
			ActiveClusterName: common.StringPtr(activeClusterName),
			Clusters:          clusters,
		},
	})
	s.Nil(err)
	fnTest(updateResp.DomainInfo, updateResp.Configuration, updateResp.ReplicationConfiguration, updateResp.GetIsGlobalDomain(), updateResp.GetFailoverVersion())

	getResp, err := s.handler.DescribeDomain(context.Background(), &types.DescribeDomainRequest{
		Name: common.StringPtr(domainName),
	})
	s.Nil(err)
	fnTest(
		getResp.DomainInfo,
		getResp.Configuration,
		getResp.ReplicationConfiguration,
		getResp.GetIsGlobalDomain(),
		getResp.GetFailoverVersion(),
	)
}

func (s *domainHandlerGlobalDomainDisabledSuite) getRandomDomainName() string {
	return "domain" + uuid.New()
}
