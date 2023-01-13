package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"bs.mobgi.cc/cronJobs/jobs"
	"bs.mobgi.cc/library/curl"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type PositionLogic struct {
	logx.Logger
	tokenChan chan *jobs.QueryParam
}

func NewPositionLogic() *PositionLogic {
	return &PositionLogic{
		tokenChan: make(chan *jobs.QueryParam),
	}
}

func (l *PositionLogic) PositionQuery() (err error) {
	go jobs.GetTokens(l.tokenChan)

	for token := range l.tokenChan {
		account, err := model.NewAct(vars.DBMysql).FindAccountById(token.AccountId) // 只有几条数据
		if err != nil {
			fmt.Println("定向包调度失败：", err)
		}

		for category := range vars.CreativeCategory {
			l.query(token, account.AdvertiserId, category)
		}
	}

	return
}

func (l *PositionLogic) query(param *jobs.QueryParam, advertiserId, category string) {
	data := statements.PositionRequest{
		AdvertiserId: advertiserId,
		Filtering:    statements.PositionFiltering{Category: category},
	}
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Tools.Position")).Get().JsonData(data)
	if err != nil {
		fmt.Println("请求生成失败：", err)
		return
	}
	var response statements.PositionResponse
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口返回错误：", response.Message)
		return
	}

	positions, samples, placements := formatPosition(response.Data.CreativeSizeInfoList, param.AccountId, advertiserId, category)
	if err = model.NewPosition(vars.DBMysql).BatchInsert(positions, samples, placements); err != nil {
		fmt.Println("写入数据错误：", err)
	}

	return
}

func formatPosition(src []*statements.CreativeSizeInfo, accountId int64, advertiserId, category string) ([]*model.Position, []*model.PositionSample, []*model.PositionPlacement) {
	var (
		positions          = make([]*model.Position, len(src))
		positionSamples    = make([]*model.PositionSample, 0)
		positionPlacements = make([]*model.PositionPlacement, 0)
	)

	for i, info := range src {
		positions[i] = &model.Position{
			AccountId:                  accountId,
			AdvertiserId:               advertiserId,
			Category:                   category,
			CreativeSizeId:             strconv.Itoa(info.CreativeSizeId),
			CreativeSizeNameDsp:        info.CreativeSizeBaseInfo.CreativeSizeNameDsp,
			CreativeSizeDescription:    info.CreativeSizeBaseInfo.CreativeSizeDescription,
			SupportProductType:         info.CreativeSizeOperationInfo.SupportProductType,
			SupportObjectiveType:       info.CreativeSizeOperationInfo.SupportObjectiveType,
			IsSupportTimePeriod:        info.CreativeSizeOperationInfo.IsSupportTimePeriod,
			IsSupportMultipleCreatives: info.CreativeSizeOperationInfo.IsSupportMultipleCreatives,
			SupportPriceType:           info.CreativeSizePriceInfo.SupportPriceType,
			LastPullTime:               time.Now(),
		}
		for _, sample := range info.CreativeSizeBaseInfo.CreativeSizeSampleList {
			positionSamples = append(positionSamples, &model.PositionSample{
				CreativeSizeId:     strconv.Itoa(info.CreativeSizeId),
				CreativeSizeSample: sample.CreativeSizeSample,
				PreviewTitle:       sample.PreviewTitle,
			})
		}
		for _, placement := range info.CreativeSizeBaseInfo.CreativeSizePlacementList {
			positionPlacements = append(positionPlacements, &model.PositionPlacement{
				CreativeSizeId:             strconv.Itoa(info.CreativeSizeId),
				PlacementSizeId:            placement.PlacementSizeId,
				CreativeSize:               placement.CreativeSize,
				CreativeSizeSubType:        placement.CreativeSizeSubType,
				IsSupportMultipleCreatives: placement.IsSupportMultipleCreatives,
			})
		}
	}
	return positions, positionSamples, positionPlacements
}
