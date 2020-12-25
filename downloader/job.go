package downloader

import (
	"encoding/json"
	"fmt"
	"go_jobs/fake"
	"go_jobs/pkg/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	jobsApiUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%s&needAddtionalResult=false"
	//jobsApiUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%E4%B8%8A%E6%B5%B7&needAddtionalResult=false"
)

type ListResult struct {
	Code    int
	Success bool
	Msg     string
	Content Content
}

type Content struct {
	PositionResult PositionResult
	PageNo         int
	PageSize       int
}

type PositionResult struct {
	Result     []Result
	TotalCount int
}

type Result struct {
	City              string
	BusinessZones     []string
	CompanyFullName   string
	CompanyLabelList  []string
	CompanyShortName  string
	CompanySize       string
	CreateTime        string
	District          string
	Education         string
	FinanceStage      string
	FirstType         string
	IndustryField     string
	IndustryLables    []string
	JobNature         string
	Latitude          string
	Longitude         string
	PositionAdvantage string
	PositionId        int32
	PositionLables    []string
	PositionName      string
	Salary            string
	SecondType        string
	Stationname       string
	Subwayline        string
	Linestaion        string
	WorkYear          string

	PositionID            int64    `json:"positionId"`
	CompanyID             int64    `json:"companyId"`
	CompanyLogo           string   `json:"companyLogo"`
	ThirdType             string   `json:"thirdType"`
	SkillLables           []string   `json:"skillLables"`
	FormatCreateTime      string   `json:"formatCreateTime"`
	ImState               string   `json:"imState"`
	LastLogin             string   `json:"lastLogin"`
	PublisherID           int64    `json:"publisherId"`
	Approve               int64    `json:"approve"`
	Distance              string   `json:"distance"`
	Hitags                []string   `json:"hitags"`
	ResumeProcessRate     int64    `json:"resumeProcessRate"`
	ResumeProcessDay      int64    `json:"resumeProcessDay"`
	Score                 int64    `json:"score"`
	NewScore              float64    `json:"newScore"`
	MatchScore            float64  `json:"matchScore"`
	MatchScoreExplain     string   `json:"matchScoreExplain"`
	Query                 string   `json:"query"`
	Explain               string   `json:"explain"`
	IsSchoolJob           int64    `json:"isSchoolJob"`
	AdWord                int64    `json:"adWord"`
	Plus                  string   `json:"plus"`
	PcShow                int64    `json:"pcShow"`
	AppShow               int64    `json:"appShow"`
	Deliver               int64    `json:"deliver"`
	GradeDescription      string   `json:"gradeDescription"`
	PromotionScoreExplain string   `json:"promotionScoreExplain"`
	IsHotHire             int64    `json:"isHotHire"`
	Count                 int64    `json:"count"`
	AggregatePositionIds  []string `json:"aggregatePositionIds"`
	PromotionType         string   `json:"promotionType"`
	Is51Job               bool     `json:"is51Job"`
	HunterJob             bool     `json:"hunterJob"`
	DetailRecall          bool     `json:"detailRecall"`
	FamousCompany         bool     `json:"famousCompany"`
}

type jobService struct {
	City string
}

func NewJobService(city string) *jobService {
	return &jobService{City: city}
}

func (l *jobService) GetUrl() string {
	req := fmt.Sprintf(jobsApiUrl, l.City)
	url, _ := url.Parse(req)
	query := url.Query()
	url.RawQuery = query.Encode()

	return url.String()
}
func getDate() string {
	return time.Now().Format("20060102150405")
}
func (l *jobService) GetJobs(pn int, kd string) (*ListResult, error) {
	//client := fake.ProxyAuth{License: "", SecretKey: ""}.GetProxyClient()
	client := http.Client{}
	postReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d&kd=%s", pn, kd))
	req, err := http.NewRequest("POST", l.GetUrl(), postReader)
	if err != nil {
		log.Printf("http.NewRequest err: %v", err)
	}

	//req.Header.Set("Proxy-Switch-Ip", "yes")

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Languag", "zh,en;q=0.9,zh-TW;q=0.8,zh-CN;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "25")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", "user_trace_token=20201224100334-40bb12e9-8494-40e7-abaf-47f930a0f852; LGUID=20201224100334-b7d5a719-3544-4aac-9848-cf11e2c388a6; _ga=GA1.2.1567927933.1608775415; index_location_city=%E4%B8%8A%E6%B5%B7; _gid=GA1.2.2121908850.1608777451; SEARCH_ID=bc8fd0b88e5d407fb38b7e27560cc129; JSESSIONID=ABAAAECABIEACCAC2BD11DE10E758D7B076947617FBCAAE; WEBTJ-ID=20201225095246-17697992d55440-0c3d9368077f84-c791039-1226944-17697992d56d4a; sensorsdata2015session=%7B%7D; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1608775417,1608861167; TG-TRACK-CODE=index_search; PRE_UTM=; PRE_HOST=; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2F; LGSID=20201225154326-9d5fbb1c-c562-4878-9ab1-9a5ba9a2bb26; PRE_SITE=https%3A%2F%2Fwww.lagou.com; _gat=1; X_HTTP_TOKEN=46fce29de2577a4461728880618603e2dbc4fef2ca; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22176927cb31f38d-00163cff7ec69c-c791039-1235456-176927cb320bef%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24os%22%3A%22Windows%22%2C%22%24browser%22%3A%22Chrome%22%2C%22%24browser_version%22%3A%2287.0.4280.88%22%7D%2C%22%24device_id%22%3A%22176927cb31f38d-00163cff7ec69c-c791039-1235456-176927cb320bef%22%7D; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1608882715; LGRID=20201225155156-2aeddaf1-e4b1-453a-b4d8-fbc8690c980c")
	//req.Header.Add("Cookie", "JSESSIONID=ABAAABAABEIABCI76211F4613E8F3F9A0CDBC51203F3A7F; " +
	//	"WEBTJ-ID="+getDate()+"-1769267f49fb2b-091b185a012074-6c112c7c-2073600-1769267f4a0ab9; " +
	//	"RECOMMEND_TIP=true; sajssdk_2015_cross_new_user=1; sensorsdata2015session=%7B%7D; " +
	//	"_ga=GA1.2.885228844.1608774055; _gid=GA1.2.608180794.1608774055; _gat=1; " +
	//	"PRE_UTM=; PRE_HOST=; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2F; " +
	//	"user_trace_token="+getDate()+"-acc83966-b0e2-46fc-8cca-2ff7f766a74f; " +
	//	"LGSID="+getDate()+"-ef171881-56a2-429e-9742-499f32724956; " +
	//	"PRE_SITE=https%3A%2F%2Fwww.lagou.com; LGUID="+getDate()+"-4f6b3c1e-5f78-41b2-9487-7a5c3d07bf3d; " +
	//	"Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1608774058; index_location_city=%E5%8C%97%E4%BA%AC; " +
	//	"TG-TRACK-CODE=index_search; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1608774064; " +
	//	"SEARCH_ID=76ff2a51e0424c869a6b3af2a641072d; X_HTTP_TOKEN=1ba6de335b5ff0259604778061d5e968478a6dcbb4; " +
	//	"sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221769267f60155-0b0b299e842fc3-6c112c7c-" +
	//	"2073600-1769267f6027fd%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_" +
	//	"source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_" +
	//	"search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24" +
	//	"latest_referrer%22%3A%22%22%2C%22%24os%22%3A%22MacOS%22%2C%22%24browser%22%3A%22Chrome%22%2C%22%24" +
	//	"browser_version%22%3A%2287.0.4280.88%22%7D%2C%22%24device_id%22%3A%221769267f60155-0b0b299e842fc3-6c112c7c-2073600-" +
	//	"1769267f6027fd%22%7D; LGRID=20201224094110-49010a46-cd95-4b79-979c-9c0a34b8301a")
	req.Header.Add("Cookie", "_ga=GA1.2.161331334.1522592243; "+
		"user_trace_token=20201224221723-"+uuid.GetUUID()+"; "+
		"LGUID=20180401221723-"+uuid.GetUUID()+"; "+
		"index_location_city=%E6%B7%B1%E5%9C%B3; "+
		"JSESSIONID="+uuid.GetUUID()+"; "+
		"_gid=GA1.2.1140631185.1523090450; "+
		"Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1522592243,1523090450; "+
		"TG-TRACK-CODE=index_search; _gat=1; "+
		"LGSID=20180407221340-"+uuid.GetUUID()+"; "+
		"PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2F; "+
		"PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fjobs%2Flist_golang%3FlabelWords%3D%26fromSearch%3Dtrue%26suginput%3D; "+
		"Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1523110425; "+
		"LGRID=20180407221344-"+uuid.GetUUID()+"; "+
		"SEARCH_ID="+uuid.GetUUID()+"")
	req.Header.Add("Host", "www.lagou.com")
	req.Header.Add("Origin", "https://www.lagou.com")
	req.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")
	req.Header.Add("User-Agent", fake.GetUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results ListResult
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
