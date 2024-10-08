//go:build master || leaf || game || gate || web || center || match || chat

package main

import (
	"os"
	"superman/start"

	"github.com/urfave/cli/v2"
)

func main() {
	args := os.Args

	// 分开编译，好处是单独代码，即便被破解也只能看到其一代码
	app := &cli.App{
		Name:        "game cluster node",
		Description: "game cluster node examples",
		Commands: []*cli.Command{
			start.VersionCommand(),
			start.Command(),
		},
	}

	// 共用main模式 好处是一份可执行文件 起多个不同进程
	//app := &cli.App{
	//	Name:        "game cluster node",
	//	Description: "game cluster node examples",
	//	Commands: []*cli.Command{
	//		versionCommand(),
	//		masterCommand(),
	//		centerCommand(),
	//		webCommand(),
	//		gateCommand(),
	//		gameCommand(),
	//	},
	//}
	_ = app.Run(args)
}

//
//func versionCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "version",
//		Aliases:   []string{"ver", "v"},
//		Usage:     "view version",
//		UsageText: "game cluster node version",
//		Action: func(c *cli.Context) error {
//			fmt.Println(cherryConst.Version())
//			return nil
//		},
//	}
//}
//
//func masterCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "master",
//		Usage:     "run master node",
//		UsageText: "node master --path=./examples/config/profile-gc.json --node=gc-master",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			master.Run(path, node)
//			return nil
//		},
//	}
//}
//
//func centerCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "center",
//		Usage:     "run center node",
//		UsageText: "node center --path=./examples/config/profile-gc.json --node=gc-center",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			center.Run(path, node)
//			return nil
//		},
//	}
//}
//
//func webCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "web",
//		Usage:     "run web node",
//		UsageText: "node web --path=./examples/config/profile-gc.json --node=gc-web-1",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			web.Run(path, node)
//			return nil
//		},
//	}
//}
//
//func gateCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "gate",
//		Usage:     "run gate node",
//		UsageText: "node gate --path=./examples/config/profile-gc.json --node=gc-gate-1",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			gate.Run(path, node)
//			return nil
//		},
//	}
//}
//
//func gameCommand() *cli.Command {
//	return &cli.Command{
//		Name:      "game",
//		Usage:     "run game node",
//		UsageText: "node game --path=./examples/config/profile-gc.json --node=10001",
//		Flags:     getFlag(),
//		Action: func(c *cli.Context) error {
//			path, node := getParameters(c)
//			game.Run(path, node)
//			return nil
//		},
//	}
//}
//
//func getParameters(c *cli.Context) (path, node string) {
//	path = c.String("path")
//	node = c.String("node")
//	return path, node
//}
//
//func getFlag() []cli.Flag {
//	return []cli.Flag{
//		&cli.StringFlag{
//			Name:     "path",
//			Usage:    "profile config path",
//			Required: false,
//			Value:    "./examples/config/profile-gc.json",
//		},
//		&cli.StringFlag{
//			Name:     "node",
//			Usage:    "node id name",
//			Required: true,
//			Value:    "",
//		},
//	}
//}
