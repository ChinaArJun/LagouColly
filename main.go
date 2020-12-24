package main

import (
	"go_jobs/model"
	"go_jobs/spider"
	"log"
	"sync"
)

var (

	kds = []string{
		"java",
	}
	citys = []string{
		"北京",
		"上海",
		"广州",
		"深圳",
		"杭州",
		"成都",
		"武汉",
		"杭州",
		"长沙",
		"西安",
	}

	initResults = []spider.InitResult{}
	loopResults = []spider.LoopResult{}
	jobPipeline = model.NewJobPipeline()

	wg sync.WaitGroup
)

func main() {
	for _, kd := range kds {
		for _, city := range citys {
			wg.Add(1)
			go func(city string, kd string) {
				defer wg.Done()
				initResult, err := spider.InitJobs(city, 1, kd)
				if err != nil {
					log.Fatalln(err)
				}

				initResults = append(initResults, initResult...)
				loopResults = append(loopResults, spider.LoopJobs())
			}(city, kd)
		}
	}

	wg.Wait()
	log.Println("爬取完成")
	//pJobs := []downloader.Result{
	//	downloader.Result{
	//	City:              "1",
	//	BusinessZones:     "",
	//	CompanyFullName:   "2",
	//	CompanyLabelList:  nil,
	//	CompanyShortName:  "",
	//	CompanySize:       "",
	//	CreateTime:        "",
	//	District:          "",
	//	Education:         "",
	//	FinanceStage:      "",
	//	FirstType:         "",
	//	IndustryField:     "",
	//	IndustryLables:    nil,
	//	JobNature:         "",
	//	Latitude:          "",
	//	Longitude:         "",
	//	PositionAdvantage: "",
	//	PositionId:        0,
	//	PositionLables:    nil,
	//	PositionName:      "",
	//	Salary:            "",
	//	SecondType:        "",
	//	Stationname:       "",
	//	Subwayline:        "",
	//	Linestaion:        "",
	//	WorkYear:          "",
	//}}
	//
	//jobPipeline.Append(convert.ToPipelineJobs(pJobs))
	//jobPipeline.Push()

	log.Printf("Init Results: %v", initResults)
	log.Printf("Loop Results: %v", loopResults)
}
