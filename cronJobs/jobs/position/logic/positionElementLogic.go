package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"bs.mobgi.cc/library/curl"
	"fmt"
	"strconv"
	"strings"
)

type PositionElementLogic struct {
	requestChan chan *requestElementChan
	workThread  int
	maxThread   int
	doneChan    chan struct{}
	elements    []*model.PositionElement
	elementChan chan *model.PositionElement
}

func NewPositionElementLogic() *PositionElementLogic {
	return &PositionElementLogic{
		workThread:  0,
		doneChan:    make(chan struct{}),
		maxThread:   10,
		elements:    make([]*model.PositionElement, 0),
		requestChan: make(chan *requestElementChan),
		elementChan: make(chan *model.PositionElement),
	}
}

func (l *PositionElementLogic) ElementQuery() (err error) {
	tokenMap, accountIds, err := getTokenMap()
	if err != nil {
		return err
	}
	positions, err := model.NewPosition(vars.DBMysql).GetPositionByAccountIds(accountIds)
	if err != nil {
		return err
	}
	params := make([]*requestElementChan, 0)
	for _, position := range positions {
		if token, ok := tokenMap[position.AccountId]; ok { // token
			params = append(params, &requestElementChan{
				Params: statements.PositionElementRequest{
					AdvertiserId: position.AdvertiserId,
					Filtering:    statements.PositionElementFilter{CreativeSizeId: position.CreativeSizeId},
				},
				Token: token,
			})
		}
	}
	if len(params) == 0 {
		return
	}
	go func(params []*requestElementChan) {
		i := 0
		for {
			if l.workThread < l.maxThread {
				l.requestChan <- params[i]
				i++
			}
			if i >= len(params) {
				break
			}
		}
	}(params)

	finishWorker := 0
	for {
		select {
		case <-l.doneChan:
			l.workThread--
			finishWorker++
		case elements, ok := <-l.elementChan:
			if ok {
				l.elements = append(l.elements, elements)
			}
		case param, ok := <-l.requestChan:
			if ok {
				l.workThread++
				go l.getElement(param)
			}
		}

		if l.workThread == 0 && finishWorker == len(params) {
			break
		}
	}

	if len(l.elements) > 0 {
		err = model.NewPositionElement(vars.DBMysql).BatchInsert(l.elements)
	}
	return
}

func (l *PositionElementLogic) getElement(param *requestElementChan) {
	defer func() {
		l.doneChan <- struct{}{}
	}()

	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Tools.PositionElement")).Get().JsonData(param.Params)
	if err != nil {
		fmt.Println("参数生成失败：", err)
		return
	}
	var response statements.PositionElementResponse
	if err = c.Request(&response, curl.Authorization(param.Token)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败：", response.Code, response.Message)
		return
	}
	//_element := make(map[string]struct{})
	for n, list := range response.Data.ElementInfoList {
		//key := fmt.Sprintf("%d-%s", response.Data.CreativeSizeId, list.Subtype)
		//if _, ok := _element[key]; !ok {
		for _, info := range list.ElementInfoList {
			durations := make([]string, len(info.Duration))
			for i, elementDuration := range info.Duration {
				durations[i] = fmt.Sprintf("%d,%d", elementDuration.Min, elementDuration.Max)
			}
			l.elementChan <- &model.PositionElement{
				GroupNumber:     int64(n + 1),
				CreativeSizeId:  strconv.Itoa(response.Data.CreativeSizeId),
				SubType:         list.Subtype,
				ElementId:       strconv.Itoa(info.ElementId),
				ElementName:     info.ElementName,
				ElementTitle:    info.ElementTitle,
				ElementCaption:  info.ElementCaption,
				Width:           info.Width,
				Height:          info.Height,
				MinWidth:        info.MinWidth,
				MinHeight:       info.MinHeight,
				MinLength:       info.MinLength,
				MaxLength:       info.MaxLength,
				FileSizeKbLimit: info.FileSizeKbLimit,
				GifSizeKbLimit:  info.GifSizeKbLimit,
				FileFormat:      info.FileFormat,
				Pattern:         info.Pattern,
				Duration:        strings.Join(durations, "|"),
				MinOccurs:       info.MinOccurs,
				MaxOccurs:       info.MaxOccurs,
			}
		}
		//}
		//_element[key] = struct{}{}
	}

	return
}
