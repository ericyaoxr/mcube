package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/ericyaoxr/mcube/cmd/mcube/enum"
)

// EnumCmd 枚举生成器
var EnumCmd = &cobra.Command{
	Use:   "enum",
	Short: "枚举生成器",
	Long:  `枚举生成器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		code, err := enum.G.Generate()
		if err != nil {
			return err
		}

		if len(code) == 0 {
			return nil
		}

		fname := os.Getenv("GOFILE")
		genFile := fname[0:len(fname)-len(path.Ext(fname))] + "_enum_generate.go"

		return ioutil.WriteFile(genFile, code, 0644)
	},
}

func init() {
	RootCmd.AddCommand(EnumCmd)
}

func init() {
	EnumCmd.PersistentFlags().BoolVarP(&enum.G.Marshal, "marshal", "m", false, "is generate json MarshalJSON and UnmarshalJSON method")
	EnumCmd.PersistentFlags().BoolVarP(&enum.G.ProtobufExt, "protobuf_ext", "p", false, "is generate protobuf extention method")
}
