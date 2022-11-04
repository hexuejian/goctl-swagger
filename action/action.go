package action

import (
	"github.com/hexuejian/goctl-swagger/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Generator(ctx *cli.Context) error {
	fileName := ctx.String("filename")

	if len(fileName) == 0 {
		fileName = "rest.swagger.json"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	basepath := ctx.String("basepath")
	host := ctx.String("host")
	return generate.Do(fileName, host, basepath, p)
}

func GeneratorTest(fileName string) error {
	p := &plugin.Plugin{
		ApiFilePath: "user.api",
		Style:       "",
		Dir:         ".",
	}
	api, err := parser.Parse(p.ApiFilePath)
	if err != nil {
		return err
	}
	p.Api = api

	basepath := "/api"
	host := "127.0.0.1"
	return generate.Do(fileName, host, basepath, p)
}
