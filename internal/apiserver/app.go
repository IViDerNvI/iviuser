package apiserver

import (
	"fmt"
	"time"

	"github.com/ividernvi/iviuser/internal/apiserver/cache"
	"github.com/ividernvi/iviuser/internal/apiserver/cache/ristretto"
	"github.com/ividernvi/iviuser/internal/apiserver/config"
	"github.com/ividernvi/iviuser/internal/apiserver/objstore"
	"github.com/ividernvi/iviuser/internal/apiserver/objstore/minio"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
	"github.com/ividernvi/iviuser/internal/apiserver/store/mysql"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "iviuser",
	Short: "iviuser is a user management system",
}

var apiserverCmd = &cobra.Command{
	Use:   "apiserver",
	Short: "Run the apiserver",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var cfg = config.DefaultConfig()

		MySQLIns, err := mysql.GetMySQLInstanceOr(cfg.Options.MySQLOpts)
		if err != nil {
			panic(err)
		}
		store.SetFactory(MySQLIns)

		CacheIns := ristretto.GetCacheInstance()
		cache.SetCacheFactory(CacheIns)

		MinioIns, err := minio.GetMinioInstance(cfg.Options.MinioOpts)
		if err != nil {
			panic(err)
		}
		objstore.SetObjStore(minio.NewObjStore(MinioIns))

		go RunGRPCServer(cfg)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return NewServer(config.DefaultConfig()).Setup().Run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of iviuser",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("version: %s\n", version)
		return nil
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run the test",
	RunE: func(cmd *cobra.Command, args []string) error {
		for {
			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiserverCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(testCmd)
}
