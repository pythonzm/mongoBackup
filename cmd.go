package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func Cmd() {
	app := &cli.App{
		Name:  "backup",
		Usage: "下载阿里云中mongo的备份文件",
		Action: func(c *cli.Context) error {
			downloadDbs(c)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mongo",
				Aliases: []string{"m"},
				Usage:   "mongo的实例ID，多个实例以英文逗号隔开",
			},
			&cli.StringFlag{
				Name:        "date",
				Aliases:     []string{"d"},
				Value:       time.Now().Format("2006-01-02"),
				Usage:       "获取哪天的备份，默认获取当天凌晨的备份文件，格式：yyyy-MM-dd",
				DefaultText: time.Now().Format("2006-01-02"),
			},
			&cli.StringFlag{
				Name:        "hour",
				Aliases:     []string{"H"},
				Value:       "03:00-04:00",
				Usage:       "在控制台设置的备份时间段，默认是：03:00-04:00",
				DefaultText: "03:00-04:00",
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       "/data/mongo-backup",
				Usage:       "指定备份文件存储路径",
				DefaultText: "/data/mongo-backup",
			},
			&cli.BoolFlag{
				Name:    "internal",
				Aliases: []string{"i"},
				Value:   false,
				Usage:   "是否使用内网下载，默认使用公网地址",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

