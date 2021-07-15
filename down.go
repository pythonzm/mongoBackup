package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dds"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func downloadDbs(c *cli.Context) {
	mongoIds := c.String("mongo")
	wg := new(sync.WaitGroup)
	for _, mongoId := range strings.Split(mongoIds, ",") {
		wg.Add(1)
		go func(mongoId string) {
			defer wg.Done()
			url := getUrl(c, strings.TrimSpace(mongoId))
			dbPath := filepath.Join(c.String("path"), c.String("date"))
			e := os.MkdirAll(dbPath, os.ModeDir)
			if e != nil {
				panic(e)
			}
			fmt.Printf("开始下载实例：%s 在 %s 的备份文件\n", mongoId, c.String("date"))
			downloadFile(url, filepath.Join(dbPath, fmt.Sprintf("%s.xb", mongoId)))
			fmt.Printf("实例：%s 在 %s 的备份文件下载完毕\n", mongoId, c.String("date"))
		}(mongoId)
	}
	wg.Wait()
}

func getUrl(c *cli.Context, mongoId string) (url string) {
	client, err := dds.NewClientWithAccessKey(os.Getenv("regionId"), os.Getenv("accessKeyId"), os.Getenv("accessKeySecret"))
	if err != nil {
		panic(err)
	}

	request := dds.CreateDescribeBackupsRequest()
	request.Scheme = "https"

	startTime, endTime := getStartAndEndTime(c.String("date"), c.String("hour"))
	request.StartTime = startTime
	request.EndTime = endTime
	request.DBInstanceId = mongoId

	response, err := client.DescribeBackups(request)
	if err != nil {
		panic(err)
	}

	if len(response.Backups.Backup) == 0 {
		panic(fmt.Errorf("实例：%s 在 %s %s 时段没有备份内容", mongoId, c.String("date"), c.String("hour")))
	}
	if c.Bool("internal") {
		url = response.Backups.Backup[0].BackupIntranetDownloadURL
	} else {
		url = response.Backups.Backup[0].BackupDownloadURL
	}

	return
}

func downloadFile(url, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}

func getStartAndEndTime(date string, hour string) (startTime, endTime string) {
	d := strings.Split(hour, "-")
	startHour := d[0]
	endHour := d[1]
	start := fmt.Sprintf("%s %s", date, startHour)
	end := fmt.Sprintf("%s %s", date, endHour)
	localStart, _ := time.ParseInLocation("2006-01-02 15:04", start, time.Local)
	localEnd, _ := time.ParseInLocation("2006-01-02 15:04", end, time.Local)
	startTime = localStart.UTC().Format("2006-01-02T15:04Z")
	endTime = localEnd.Add(time.Hour * 1).UTC().Format("2006-01-02T15:04Z")
	return
}

