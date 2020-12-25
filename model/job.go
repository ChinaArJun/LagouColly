package model

import (
	"log"
	"sync"
)

var (
	mutex sync.Mutex

	jobs []LgJob
)

type LgJob struct {
	// 工作名称
	//JobSkill string `json:"job_skill" gorm:""`
	// 城市名称
	City string
	// 地区
	District string

	// 公司简称
	CompanyShortName string
	// 公司全称
	CompanyFullName string
	// 公司标签
	CompanyLabelList string
	// 公司规模
	CompanySize string
	// 融资阶段
	FinanceStage string

	// 行业领域
	IndustryField string
	// 行业标签
	IndustryLables string

	// 职位名称
	PositionName string
	// 职位标签
	PositionLables string
	// 职位诱惑
	PositionAdvantage string
	// 工作年限
	WorkYear string
	// 学历要求
	Education string
	// 薪资范畴
	Salary string

	// 经度
	Longitude float64
	// 纬度
	Latitude float64
	// 附近的地铁
	Linestaion string

	// 发布时间
	CreateTime int64
	// 新增时间
	AddTime int64

	PositionID            string     `json:"positionId"`
	CompanyID             string     `json:"companyId"`
	CompanyLogo           string  `json:"companyLogo"`
	FirstType             string  `json:"firstType"`
	SecondType            string  `json:"secondType"`
	ThirdType             string  `json:"thirdType"`
	SkillLables           string  `json:"skillLables"`
	FormatCreateTime      string  `json:"formatCreateTime"`
	BusinessZones         string  `json:"businessZones"`
	JobNature             string  `json:"jobNature"`
	ImState               string  `json:"imState"`
	LastLogin             string  `json:"lastLogin"`
	PublisherID           string     `json:"publisherId"`
	Approve               string     `json:"approve"`
	Subwayline            string  `json:"subwayline"`
	Stationname           string  `json:"stationname"`
	Distance              string  `json:"distance"`
	Hitags                string  `json:"hitags"`
	ResumeProcessRate     string     `json:"resumeProcessRate"`
	ResumeProcessDay      string     `json:"resumeProcessDay"`
	Score                 string     `json:"score"`
	NewScore              string     `json:"newScore"`
	MatchScore            string `json:"matchScore"`
	MatchScoreExplain     string  `json:"matchScoreExplain"`
	Query                 string  `json:"query"`
	Explain               string  `json:"explain"`
	IsSchoolJob           string     `json:"isSchoolJob"`
	AdWord                string     `json:"adWord"`
	Plus                  string  `json:"plus"`
	PcShow                string     `json:"pcShow"`
	AppShow               string     `json:"appShow"`
	Deliver               string     `json:"deliver"`
	GradeDescription      string  `json:"gradeDescription"`
	PromotionScoreExplain string  `json:"promotionScoreExplain"`
	IsHotHire             string     `json:"isHotHire"`
	Count                 string     `json:"count"`
	AggregatePositionIds  string  `json:"aggregatePositionIds"`
	PromotionType         string  `json:"promotionType"`
	Is51Job               bool    `json:"is51Job"`
	HunterJob             bool    `json:"hunterJob"`
	DetailRecall          bool    `json:"detailRecall"`
	FamousCompany         bool    `json:"famousCompany"`
}

func NewJobPipeline() *LgJob {
	return &LgJob{}
}

func (obj LgJob) TableName() string {
	//return fmt.Sprintf("job_java_25_01" + obj.City)
	return "sp_jobs"
}

func (j *LgJob) Append(js []LgJob) {
	mutex.Lock()
	jobs = append(jobs, js...)
	mutex.Unlock()

	for _, v := range js {
		if err := DB.Create(v).Error; err != nil {
			log.Println("数据新增报错:", err)
		}
	}
}

func (j *LgJob) Get() []LgJob {
	return jobs
}

func (j *LgJob) Push() error {
	for _, v := range j.Get() {
		if err := DB.Create(v).Error; err != nil {
			return err
		}
	}

	return nil
}
