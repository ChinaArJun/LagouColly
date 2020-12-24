package convert

import (
	"go_jobs/downloader"
	"go_jobs/model"
	"strconv"
	"strings"

	"time"
)

func ToPipelineJobs(dJobs []downloader.Result) []model.LgJob {
	var pJobs []model.LgJob
	for _, v := range dJobs {
		longitude, _ := strconv.ParseFloat(v.Longitude, 64)
		latitude, _ := strconv.ParseFloat(v.Latitude, 64)
		pJobs = append(pJobs, model.LgJob{
			City:     v.City,
			District: v.District,

			CompanyShortName: v.CompanyShortName,
			CompanyFullName:  v.CompanyFullName,
			CompanyLabelList: strings.Join(v.CompanyLabelList, ","),
			CompanySize:      v.CompanySize,
			FinanceStage:     v.FinanceStage,

			PositionName:      v.PositionName,
			PositionLables:    strings.Join(v.PositionLables, ","),
			PositionAdvantage: v.PositionAdvantage,
			WorkYear:          v.WorkYear,
			Education:         v.Education,
			Salary:            v.Salary,

			IndustryField:  v.IndustryField,
			IndustryLables: strings.Join(v.IndustryLables, ","),

			Longitude:  longitude,
			Latitude:   latitude,
			Linestaion: v.Linestaion,

			CreateTime: MustDateToUnix(v.CreateTime),
			AddTime:    time.Now().Unix(),

			PositionID:            strconv.FormatInt(v.PositionID, 10),
			CompanyID:            strconv.FormatInt( v.CompanyID, 10),
			CompanyLogo:           v.CompanyLogo,
			FirstType:             v.FirstType,
			SecondType:            v.SecondType,
			ThirdType:             v.ThirdType,
			SkillLables:           strings.Join(v.SkillLables, ","),
			FormatCreateTime:      v.FormatCreateTime,
			BusinessZones:          strings.Join(v.BusinessZones, ","),
			JobNature:             v.JobNature,
			ImState:               v.ImState,
			LastLogin:             v.LastLogin,
			PublisherID:           strconv.FormatInt(v.PublisherID, 10),
			Approve:              strconv.FormatInt( v.Approve, 10),
			Subwayline:            v.Subwayline,
			Stationname:           v.Stationname,
			Distance:              v.Distance,
			Hitags:                strings.Join(v.Hitags, ","),
			ResumeProcessRate:     strconv.FormatInt( v.ResumeProcessRate, 10),
			ResumeProcessDay:      strconv.FormatInt( v.ResumeProcessDay, 10),
			Score:                 strconv.FormatInt(v.Score, 10),
			NewScore:              strconv.FormatFloat(v.NewScore, 'f', 10, 64),
			MatchScore:            strconv.FormatFloat(v.MatchScore, 'f', 10, 64),
			MatchScoreExplain:     v.MatchScoreExplain,
			Query:                 v.Query,
			Explain:               v.Explain,
			IsSchoolJob:           strconv.FormatInt(v.IsSchoolJob, 10),
			AdWord:                strconv.FormatInt(v.AdWord,10),
			Plus:                  v.Plus,
			PcShow:                strconv.FormatInt(v.PcShow,10 ),
			AppShow:               strconv.FormatInt(v.AppShow, 10),
			Deliver:               strconv.FormatInt(v.Deliver, 10),
			GradeDescription:      v.GradeDescription,
			PromotionScoreExplain: v.PromotionScoreExplain,
			IsHotHire:             strconv.FormatInt(v.IsHotHire, 10),
			Count:                 strconv.FormatInt(v.Count, 10),
			AggregatePositionIds:  strings.Join(v.AggregatePositionIds, ","),
			PromotionType:         v.PromotionType,
			Is51Job:               v.Is51Job,
			HunterJob:             v.HunterJob,
			DetailRecall:          v.DetailRecall,
			FamousCompany:         v.FamousCompany,
		})
	}

	return pJobs
}
