package v_data

type VPositionQuery struct {
	Category    string `form:"category" binding:"required"`
	AccountId   int64  `form:"account_id" binding:"required"`
	ProductType string `form:"product_type" binding:"required"`
}

type VPositionPlacement struct {
	CreativeSizeId string `form:"creative_size_id" binding:"required"`
}

type VPositionPrice struct {
	CreativeSizeId string `form:"creative_size_id" binding:"required"`
	PriceType      string `form:"price_type" binding:"required"`
}
