package utils

var (
	GoodsCategoryVerify           = Rules{"CategoryRank": {NotEmpty()}, "CategoryName": {NotEmpty()}}
	AdminUserRegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	MallUserRegisterVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AdminUserChangePasswordVerify = Rules{"Password": {NotEmpty()}}
	GoodsAddParamVerify           = Rules{"GoodsName": {Le("128")}, "GoodsIntro": {Le("200")}, "GoodsCategoryId": {Ge("1")}, "GoodsCoverImg": {NotEmpty()}, "OriginalPrice": {Ge("1"), Le("1000000")},
		"sellingPrice": {Ge("1"), Le("1000000")}, "stockNum": {Ge("1"), Le("100000")}, "Tag": {Le("16")}, "goodsDetailContent": {NotEmpty()}}
	CarouselAddParamVerify       = Rules{"CarouselUrl": {NotEmpty()}, "RedirectUrl": {NotEmpty()}, "CarouselRank": {NotEmpty(), Ge("0"), Le("200")}}
	IndexConfigAddParamVerify    = Rules{"ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "GoodsId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("200")}}
	IndexConfigUpdateParamVerify = Rules{"ConfigId": {NotEmpty()}, "ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "GoodsId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("200")}}
	SaveOrderParamVerify         = Rules{"CartItemIds": {NotEmpty()}, "AddressId": {NotEmpty()}}
)
