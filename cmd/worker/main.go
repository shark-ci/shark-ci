package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/shark-ci/shark-ci/internal/config"
	"github.com/shark-ci/shark-ci/internal/message_queue"
	pb "github.com/shark-ci/shark-ci/internal/proto"
	"github.com/shark-ci/shark-ci/internal/worker"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	err := config.LoadWorkerConfigFromEnv()
	if err != nil {
		slog.Error("Loading config from environment failed.", "err", err)
		os.Exit(1)
	}

	compressedReposPath, err := worker.CreateTmpDir()
	if err != nil {
		slog.Error("Creating TMP dir failed.", "err", err)
		os.Exit(1)
	}

	slog.Info("Connecting to RabbitMQ.")
	rabbitMQ, err := message_queue.NewRabbitMQ(config.WorkerConf.MQ.URI)
	if err != nil {
		slog.Error("Connecting to RabbitMQ failed", "err", err)
		os.Exit(1)
	}
	defer rabbitMQ.Close(context.TODO())
	slog.Info("RabbitMQ connected.")

	slog.Info("Creating gRPC client.")
	conn, err := grpc.Dial(config.WorkerConf.CIServerHost+":"+config.WorkerConf.CIServerGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Connecting to gRPC server failed.", "addr", config.WorkerConf.CIServerHost+":"+config.WorkerConf.CIServerGRPCPort, "err", err)
		os.Exit(1)
	}
	defer conn.Close()
	gRPCClient := pb.NewPipelineReporterClient(conn)
	slog.Info("gRPC client created.")

	err = worker.Run(rabbitMQ, gRPCClient, compressedReposPath)
	if err != nil {
		slog.Error("Running worker failed", "err", err)
		os.Exit(1)
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
}
