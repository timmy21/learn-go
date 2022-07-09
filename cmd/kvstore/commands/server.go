package commands

import (
	"context"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/timmy21/learn-go/pkg/test/kvstore"
	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, _ := zap.NewDevelopment()
		goGroup, rootCtx := errgroup.WithContext(cmd.Context())

		backend := kvstore.NewMemBackend()
		srv := kvstore.NewService(backend, logger)

		{
			viper.SetDefault("http.listen.addr", "127.0.0.1")
			viper.SetDefault("http.listen.port", 8007)

			addr := fmt.Sprintf("%s:%d", viper.GetString("http.listen.addr"), viper.GetInt("http.listen.port"))
			server := http.Server{
				Addr:     addr,
				Handler:  kvstore.NewMux(srv),
				ErrorLog: zap.NewStdLog(logger),
			}

			goGroup.Go(func() error {
				logger.Info("Listening and serving HTTP", zap.String("addr", addr))
				err := server.ListenAndServe()
				if errors.Is(err, http.ErrServerClosed) {
					return nil
				}
				return err
			})

			goGroup.Go(func() (_ error) {
				<-rootCtx.Done()
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				err := server.Shutdown(ctx)
				if err == nil {
					return
				}

				logger.Info("shutdown http server timeout", zap.Error(err))
				err = server.Close()
				if err != nil {
					logger.Info("close http server error", zap.Error(err))
				}
				return
			})
		}

		{
			viper.SetDefault("grpc.listen.addr", "127.0.0.1")
			viper.SetDefault("grpc.listen.port", 8008)

			addr := fmt.Sprintf("%s:%d", viper.GetString("grpc.listen.addr"), viper.GetInt("grpc.listen.port"))
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				return errors.WithStack(err)
			}
			server := grpc.NewServer()
			kvstorepb.RegisterKVStoreServer(server, srv)

			goGroup.Go(func() error {
				logger.Info("Listening and serving GRPC", zap.String("addr", addr))
				err = server.Serve(lis)
				if errors.Is(err, grpc.ErrServerStopped) {
					return nil
				}
				return err
			})

			goGroup.Go(func() (_ error) {
				<-rootCtx.Done()
				ch := make(chan struct{})
				go func() {
					server.GracefulStop()
					close(ch)
				}()

				select {
				case <-ch:
					return
				case <-time.After(time.Second * 5):
					server.Stop()
					logger.Info("force stop grpc server")
					return
				}
			})
		}
		return goGroup.Wait()
	},
}
