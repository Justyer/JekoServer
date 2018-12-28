package root

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	cfgFile = "conf/dev.toml"

	RootCmd = &cobra.Command{
		Use:   "jeko",
		Short: "game server",
		Long:  "server for game with Kaller",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config path")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	// 配置自动更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config is change:", e.String())
	})
	viper.WatchConfig()
}
