//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/google/wire"
	"github.com/minghsu0107/go-random-chat/pkg/chat"
	"github.com/minghsu0107/go-random-chat/pkg/common"
	"github.com/minghsu0107/go-random-chat/pkg/config"
	"github.com/minghsu0107/go-random-chat/pkg/forwarder"
	"github.com/minghsu0107/go-random-chat/pkg/infra"
	"github.com/minghsu0107/go-random-chat/pkg/match"
	"github.com/minghsu0107/go-random-chat/pkg/uploader"
	"github.com/minghsu0107/go-random-chat/pkg/user"
	"github.com/minghsu0107/go-random-chat/pkg/web"
)

func InitializeWebServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewHttpLogrus,
		web.NewGinServer,
		web.NewHttpServer,
		web.NewRouter,
		web.NewInfraCloser,
		common.NewServer,
	)
	return &common.Server{}, nil
}

func InitializeChatServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewHttpLogrus,
		common.NewGrpcLogrus,

		infra.NewRedisClient,
		infra.NewRedisCache,

		infra.NewKafkaPublisher,
		infra.NewKafkaSubscriber,
		infra.NewBrokerRouter,

		infra.NewCassandraSession,

		chat.NewUserClientConn,
		chat.NewForwarderClientConn,

		chat.NewUserRepo,
		chat.NewMessageRepo,
		chat.NewChannelRepo,
		chat.NewForwardRepo,

		chat.NewUserRepoCache,
		chat.NewMessageRepoCache,
		chat.NewChannelRepoCache,

		chat.NewMessageSubscriber,

		common.NewSonyFlake,

		chat.NewUserService,
		chat.NewMessageService,
		chat.NewChannelService,
		chat.NewForwardService,

		chat.NewMelodyChatConn,

		chat.NewGinServer,
		chat.NewHttpServer,
		chat.NewGrpcServer,
		chat.NewRouter,
		chat.NewInfraCloser,
		common.NewServer,
	)
	return &common.Server{}, nil
}

func InitializeForwarderServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewGrpcLogrus,

		infra.NewRedisClient,
		infra.NewRedisCache,

		infra.NewKafkaPublisher,
		infra.NewKafkaSubscriber,
		infra.NewBrokerRouter,

		forwarder.NewForwardRepo,

		forwarder.NewForwardService,

		forwarder.NewMessageSubscriber,

		forwarder.NewGrpcServer,
		forwarder.NewRouter,
		forwarder.NewInfraCloser,
		common.NewServer,
	)
	return &common.Server{}, nil
}

func InitializeMatchServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewHttpLogrus,

		infra.NewRedisClient,
		infra.NewRedisCache,

		infra.NewKafkaPublisher,
		infra.NewKafkaSubscriber,
		infra.NewBrokerRouter,

		match.NewChatClientConn,
		match.NewUserClientConn,

		match.NewChannelRepo,
		match.NewUserRepo,
		match.NewMatchingRepo,

		match.NewMatchSubscriber,

		match.NewUserService,
		match.NewMatchingService,

		match.NewMelodyMatchConn,

		match.NewGinServer,
		match.NewHttpServer,
		match.NewRouter,
		match.NewInfraCloser,
		common.NewServer,
	)
	return &common.Server{}, nil
}

func InitializeUploaderServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewHttpLogrus,
		uploader.NewGinServer,
		uploader.NewHttpServer,
		uploader.NewChannelUploadRateLimiter,
		uploader.NewRouter,
		uploader.NewInfraCloser,
		common.NewServer,
		infra.NewRedisClient,
	)
	return &common.Server{}, nil
}

func InitializeUserServer(name string) (*common.Server, error) {
	wire.Build(
		config.NewConfig,
		common.NewObservabilityInjector,
		common.NewHttpLogrus,
		common.NewGrpcLogrus,

		infra.NewRedisClient,
		infra.NewRedisCache,

		user.NewUserRepo,

		common.NewSonyFlake,

		user.NewUserService,

		user.NewGinServer,
		user.NewHttpServer,
		user.NewGrpcServer,
		user.NewRouter,
		user.NewInfraCloser,
		common.NewServer,
	)
	return &common.Server{}, nil
}
