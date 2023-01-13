package servicemarketing

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"errors"
)

type PositionInfo struct {
	CreativeSizeId             string           `json:"creative_size_id"`
	CreativeSizeNameDsp        string           `json:"creative_size_name_dsp"`
	CreativeSizeDescription    string           `json:"creative_size_description"`
	SupportProductType         string           `json:"support_product_type"`
	IsSupportTimePeriod        string           `json:"is_support_time_period"`
	IsSupportMultipleCreatives string           `json:"is_support_multiple_creatives"`
	SupportPriceType           string           `json:"support_price_type"`
	Samples                    []*SampleInfo    `json:"samples"`
	Placements                 []*PlacementInfo `json:"placements"`
}

type SampleInfo struct {
	CreativeSizeSample string `json:"creative_size_ample"`
	PreviewTitle       string `json:"preview_title"`
}

type PlacementInfo struct {
	PlacementSizeId     string `json:"placement_size_id"`
	CreativeSize        string `json:"creative_size"`
	CreativeSizeSubType string `json:"creative_size_sub_type"`
}

func GetPositionDetail(category, productType string, accountId int64) (creativeSizeList []*PositionInfo, err error) {
	positions, err := model.NewPosition(vars.DBMysql).PositionList(category, productType, accountId)
	if err != nil {
		return nil, errors.New("版位获取失败：" + err.Error())
	}
	var (
		creativeSizeIds []string
	)
	for _, position := range positions {
		creativeSizeIds = append(creativeSizeIds, position.CreativeSizeId)
		creativeSizeList = append(creativeSizeList, &PositionInfo{
			CreativeSizeId:             position.CreativeSizeId,
			CreativeSizeNameDsp:        position.CreativeSizeNameDsp,
			CreativeSizeDescription:    position.CreativeSizeDescription,
			SupportProductType:         position.SupportProductType,
			IsSupportTimePeriod:        position.IsSupportTimePeriod,
			IsSupportMultipleCreatives: position.IsSupportMultipleCreatives,
			SupportPriceType:           position.SupportPriceType,
		})
	}
	samples, err := model.NewPositionSample(vars.DBMysql).SampleList(creativeSizeIds)
	if err != nil {
		return nil, errors.New("版位图信息获取失败：" + err.Error())
	}
	placements, err := model.NewPositionPlacement(vars.DBMysql).PlacementList(creativeSizeIds)
	if err != nil {
		return nil, errors.New("版位形式获取失败：" + err.Error())
	}
	sampleTmp := make(map[string][]*SampleInfo)
	for _, sample := range samples {
		sampleTmp[sample.CreativeSizeId] = append(sampleTmp[sample.CreativeSizeId], &SampleInfo{
			CreativeSizeSample: sample.CreativeSizeSample,
			PreviewTitle:       sample.PreviewTitle,
		})
	}
	placementTmp := make(map[string][]*PlacementInfo)
	for _, placement := range placements {
		placementTmp[placement.CreativeSizeId] = append(placementTmp[placement.CreativeSizeId], &PlacementInfo{
			PlacementSizeId:     placement.PlacementSizeId,
			CreativeSize:        placement.CreativeSize,
			CreativeSizeSubType: placement.CreativeSizeSubType,
		})
	}
	for i, creativeSize := range creativeSizeList {
		if v, ok := sampleTmp[creativeSize.CreativeSizeId]; ok {
			creativeSizeList[i].Samples = v
		}
		if v, ok := placementTmp[creativeSize.CreativeSizeId]; ok {
			creativeSizeList[i].Placements = v
		}
	}
	return
}
