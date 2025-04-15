package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var rootCmd = &cobra.Command{

	Use : "app",
	Short : "application detail",
	Long : "what can i say",
}

var subCmd = &cobra.Command {
	Use : "server",
	Short : "do do do",
	Run: func(cmd * cobra.Command,args []string){
		fmt.Println("congratulation to u")
	},
}

func Init(){
	rootCmd.AddCommand(subCmd)
}


func Execute(){
	rootCmd.Execute()
}
