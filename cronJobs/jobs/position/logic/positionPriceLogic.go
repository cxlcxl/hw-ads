package logic

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/vars"
	"bs.mobgi.cc/cronJobs/job_data/statements"
	"bs.mobgi.cc/library/curl"
	"fmt"
	"strings"
)

type PositionPriceLogic struct {
	maxThread   int
	workThread  int
	doneChan    chan struct{}
	requestChan chan *requestPriceChan
	priceChan   chan *model.PositionPrice
	basePrices  []*model.PositionPrice
}

func NewPositionPriceLogic() *PositionPriceLogic {
	return &PositionPriceLogic{
		maxThread:   10,
		workThread:  0,
		doneChan:    make(chan struct{}),
		priceChan:   make(chan *model.PositionPrice),
		basePrices:  make([]*model.PositionPrice, 0),
		requestChan: make(chan *requestPriceChan),
	}
}

func (l *PositionPriceLogic) PriceQuery() (err error) {
	tokenMap, accountIds, err := getTokenMap()
	if err != nil {
		return err
	}
	positions, err := model.NewPosition(vars.DBMysql).GetPositionByAccountIds(accountIds)
	if err != nil {
		return err
	}

	tmpPriceTypes := make(map[string]string)
	for s, s2 := range vars.Pricing {
		tmpPriceTypes[s2] = s
	}
	params := make([]*requestPriceChan, 0)
	for _, position := range positions {
		if token, ok := tokenMap[position.AccountId]; ok { // token
			priceTypes := strings.Split(position.SupportPriceType, ",")
			for _, priceType := range priceTypes {
				if t, _ok := tmpPriceTypes[priceType]; _ok {
					params = append(params, &requestPriceChan{
						Params: statements.CreativeSizePriceRequest{
							AdvertiserId: position.AdvertiserId,
							Filtering:    statements.CreativeSizePriceFilter{CreativeSizeId: position.CreativeSizeId, PriceType: t},
						},
						Token: token,
					})
				}
			}
		}
	}

	if len(params) == 0 {
		return
	}

	go func(params []*requestPriceChan) {
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
		case basePrice, ok := <-l.priceChan:
			if ok {
				l.basePrices = append(l.basePrices, basePrice)
			}
		case param, ok := <-l.requestChan:
			if ok {
				l.workThread++
				go l.getBasePrice(param)
			}
		}

		if l.workThread == 0 && finishWorker == len(params) {
			break
		}
	}

	if len(l.basePrices) > 0 {
		err = model.NewPositionPrice(vars.DBMysql).BatchInsert(l.basePrices)
	}
	return
}

func (l *PositionPriceLogic) getBasePrice(param *requestPriceChan) {
	defer func() {
		l.doneChan <- struct{}{}
	}()
	c, err := curl.New(vars.YmlConfig.GetString("MarketingApis.Tools.PositionPrice")).Get().JsonData(param.Params)
	if err != nil {
		fmt.Println("参数生成失败：", err)
		return
	}
	var response statements.CreativeSizePriceResponse
	if err = c.Request(&response, curl.Authorization(param.Token)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败：", response.Code, response.Message)
		return
	}
	l.priceChan <- &model.PositionPrice{
		CreativeSizeId: param.Params.Filtering.CreativeSizeId,
		PriceType:      param.Params.Filtering.PriceType,
		BasePrice:      response.Data.FloorPrice,
	}
	return
}
