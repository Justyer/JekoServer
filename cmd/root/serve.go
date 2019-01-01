package root

import (
	"log"

	"github.com/Justyer/JekoServer/cmd/load_model"
	"github.com/Justyer/JekoServer/cmd/serve"
	"github.com/Justyer/JekoServer/model"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start serve",
	Long:  `start serve`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("serve called:", model.TaskType, model.Flag_T_Type)
		load_model.InitGlobalConfig()
		serve.Run()
	},
}

func init() {
	model.TaskType = "serve"
	serveCmd.PersistentFlags().StringVarP(&model.Flag_T_Type, "type", "t", "tcp", "serve type")
	RootCmd.AddCommand(serveCmd)
}
