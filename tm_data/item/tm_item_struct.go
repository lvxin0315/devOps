package item

type TmItem struct {
	APIStack   interface{} `json:"apiStack"`
	DetailDesc struct {
		NewWapDescDynURL string `json:"newWapDescDynUrl"`
		NewWapDescJSON   []struct {
			Component string `json:"component"`
			Data      []struct {
				Height  interface{} `json:"height"`
				Img     string      `json:"img"`
				Simples []struct {
					ImgURL  string `json:"imgUrl"`
					JumpURL string `json:"jumpUrl"`
					Title   string `json:"title"`
				} `json:"simples"`
				Width interface{} `json:"width"`
			} `json:"data"`
			Enable     bool   `json:"enable"`
			ModuleKey  string `json:"moduleKey"`
			ModuleName string `json:"moduleName"`
			ModuleType int64  `json:"moduleType"`
		} `json:"newWapDescJson"`
	} `json:"detailDesc"`
	Item struct {
		BrandValueID string      `json:"brandValueId"`
		CarShipTime  interface{} `json:"carShipTime"`
		CategoryID   string      `json:"categoryId"`
		CommentCount interface{} `json:"commentCount"`
		Enable       bool        `json:"enable"`
		ExtData      struct {
			AddressLevel int64 `json:"addressLevel"`
		} `json:"extData"`
		Favcount         int64       `json:"favcount"`
		H5moduleDescURL  interface{} `json:"h5moduleDescUrl"`
		Images           []string    `json:"images"`
		ItemID           int64       `json:"itemId"`
		Logo             interface{} `json:"logo"`
		ModuleDescParams interface{} `json:"moduleDescParams"`
		ModuleDescURL    interface{} `json:"moduleDescUrl"`
		RootCategoryID   string      `json:"rootCategoryId"`
		ServItemExtend   interface{} `json:"servItemExtend"`
		SkuText          string      `json:"skuText"`
		Slogon           interface{} `json:"slogon"`
		StandingDate     string      `json:"standingDate"`
		StorePacket      interface{} `json:"storePacket"`
		Subtitle         string      `json:"subtitle"`
		TaobaoPcDescURL  interface{} `json:"taobaoPcDescUrl"`
		ThemeType        interface{} `json:"themeType"`
		Title            string      `json:"title"`
		TitleIcon        interface{} `json:"titleIcon"`
		TmallDescURL     string      `json:"tmallDescUrl"`
		Videos           interface{} `json:"videos"`
	} `json:"item"`
	JumpURL struct {
		Apis struct {
			AddFavURL        string `json:"addFavUrl"`
			BarterURL        string `json:"barterUrl"`
			CartURL          string `json:"cartUrl"`
			DescURL          string `json:"descUrl"`
			HTTPSDescURL     string `json:"httpsDescUrl"`
			LogOutURL        string `json:"logOutUrl"`
			LoginURL         string `json:"loginUrl"`
			LongTailURL      string `json:"longTailUrl"`
			MemberShopURL    string `json:"memberShopUrl"`
			MyURL            string `json:"myUrl"`
			NewSelectCityAPI string `json:"newSelectCityApi"`
			PropertyURL      string `json:"propertyUrl"`
			RateURL          string `json:"rateUrl"`
			ReURL            string `json:"reUrl"`
			RecommendURL     string `json:"recommendUrl"`
			ReqMemberURL     string `json:"reqMemberUrl"`
			SkipEsi          string `json:"skipEsi"`
			StatsURL         string `json:"statsUrl"`
		} `json:"apis"`
		Enable bool `json:"enable"`
	} `json:"jumpUrl"`
	Mock struct {
		Delivery struct {
			AreaID  interface{} `json:"areaId"`
			From    interface{} `json:"from"`
			Postage interface{} `json:"postage"`
			To      interface{} `json:"to"`
		} `json:"delivery"`
		Enable  bool `json:"enable"`
		Feature struct {
			HasSku  bool `json:"hasSku"`
			ShowSku bool `json:"showSku"`
		} `json:"feature"`
		Price struct {
			ExtraPrices interface{} `json:"extraPrices"`
			Price       struct {
				LineThrough interface{} `json:"lineThrough"`
				PriceMoney  interface{} `json:"priceMoney"`
				PriceText   string      `json:"priceText"`
				PriceTitle  interface{} `json:"priceTitle"`
			} `json:"price"`
		} `json:"price"`
		SkuBase interface{} `json:"skuBase"`
		SkuCore struct {
			Sku2info struct {
				Zero struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"0"`
				Four400155540176 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540176"`
				Four400155540177 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540177"`
				Four400155540178 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540178"`
				Four400155540179 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540179"`
				Four400155540180 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540180"`
				Four400155540181 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540181"`
				Four400155540182 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540182"`
				Four400155540183 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540183"`
				Four400155540184 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540184"`
				Four400155540185 struct {
					Price struct {
						LineThrough interface{} `json:"lineThrough"`
						PriceMoney  int64       `json:"priceMoney"`
						PriceText   string      `json:"priceText"`
						PriceTitle  string      `json:"priceTitle"`
					} `json:"price"`
					Quantity int64 `json:"quantity"`
				} `json:"4400155540185"`
			} `json:"sku2info"`
			SkuItem struct {
				HideQuantity bool        `json:"hideQuantity"`
				SkuTitle     interface{} `json:"skuTitle"`
			} `json:"skuItem"`
		} `json:"skuCore"`
		Trade struct {
			BuyEnable      bool        `json:"buyEnable"`
			BuyParam       interface{} `json:"buyParam"`
			BuyText        interface{} `json:"buyText"`
			CartEnable     bool        `json:"cartEnable"`
			CartText       interface{} `json:"cartText"`
			Enable         bool        `json:"enable"`
			HintBanner     interface{} `json:"hintBanner"`
			RedirectURL    interface{} `json:"redirectUrl"`
			StartTime      interface{} `json:"startTime"`
			TradeConfig    interface{} `json:"tradeConfig"`
			TradeParams    interface{} `json:"tradeParams"`
			UseNativeTrade interface{} `json:"useNativeTrade"`
			UseTrade30     interface{} `json:"useTrade30"`
			WaitForStart   interface{} `json:"waitForStart"`
		} `json:"trade"`
	} `json:"mock"`
	Modules struct {
		Enable  bool `json:"enable"`
		Modules []struct {
			Data    struct{}    `json:"data"`
			GroupID string      `json:"groupId"`
			Key     string      `json:"key"`
			Name    string      `json:"name"`
			Path    interface{} `json:"path"`
			Version string      `json:"version"`
		} `json:"modules"`
	} `json:"modules"`
	Props struct {
		Enable     bool                     `json:"enable"`
		GroupProps []map[string]interface{} `json:"groupProps"`
		PropsList  interface{}              `json:"propsList"`
	} `json:"props"`
	Rate struct {
		Enable   bool `json:"enable"`
		Keywords []struct {
			Attribute string `json:"attribute"`
			Count     int64  `json:"count"`
			Type      int64  `json:"type"`
			Word      string `json:"word"`
		} `json:"keywords"`
		RateList []struct {
			Content          string      `json:"content"`
			DateTime         string      `json:"dateTime"`
			HeadExtraPic     string      `json:"headExtraPic"`
			HeadPic          string      `json:"headPic"`
			Images           []string    `json:"images"`
			IsVip            string      `json:"isVip"`
			MemberIcon       string      `json:"memberIcon"`
			MemberLevel      string      `json:"memberLevel"`
			SkuInfo          string      `json:"skuInfo"`
			TmallMemberLevel int64       `json:"tmallMemberLevel"`
			UserIcon         interface{} `json:"userIcon"`
			UserName         string      `json:"userName"`
		} `json:"rateList"`
		TotalCount int64 `json:"totalCount"`
	} `json:"rate"`
	Seller struct {
		AllItemCount    int64       `json:"allItemCount"`
		AllProduct      string      `json:"allProduct"`
		CertIcon        interface{} `json:"certIcon"`
		CertShopList    interface{} `json:"certShopList"`
		CreditLevel     string      `json:"creditLevel"`
		CreditLevelIcon string      `json:"creditLevelIcon"`
		Enable          bool        `json:"enable"`
		EncryptSellerID string      `json:"encryptSellerId"`
		Evaluates       []struct {
			Level string `json:"level"`
			Score string `json:"score"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"evaluates"`
		Fans               interface{} `json:"fans"`
		GoodRatePercentage interface{} `json:"goodRatePercentage"`
		NewItemCount       interface{} `json:"newItemCount"`
		RateSum            int64       `json:"rateSum"`
		SellerNick         string      `json:"sellerNick"`
		SellerType         string      `json:"sellerType"`
		ShopIcon           string      `json:"shopIcon"`
		ShopID             int64       `json:"shopId"`
		ShopName           string      `json:"shopName"`
		ShopTitleIcon      interface{} `json:"shopTitleIcon"`
		ShopType           string      `json:"shopType"`
		ShopURL            string      `json:"shopUrl"`
		Starts             string      `json:"starts"`
		TagIcon            interface{} `json:"tagIcon"`
		UserID             int64       `json:"userId"`
		WeitaoID           interface{} `json:"weitaoId"`
		WwURL              string      `json:"wwUrl"`
	} `json:"seller"`
	SkuBase struct {
		Components interface{} `json:"components"`
		Enable     bool        `json:"enable"`
		Props      []struct {
			Name   string `json:"name"`
			Pid    int64  `json:"pid"`
			Values []struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Vid   int64  `json:"vid"`
			} `json:"values"`
		} `json:"props"`
		SkuExtra struct {
			EnableDuration interface{} `json:"enableDuration"`
			PickupPoints   interface{} `json:"pickupPoints"`
			Tickets        interface{} `json:"tickets"`
		} `json:"skuExtra"`
		Skus []struct {
			Images   interface{} `json:"images"`
			PropPath string      `json:"propPath"`
			SkuID    int64       `json:"skuId"`
		} `json:"skus"`
	} `json:"skuBase"`
	TabBar struct {
		Enable            bool `json:"enable"`
		PropertyTabEnable bool `json:"propertyTabEnable"`
		RateListTabEnable bool `json:"rateListTabEnable"`
	} `json:"tabBar"`
	Tags struct {
		Enable bool `json:"enable"`
		Tags   struct {
			AllowRate                 bool `json:"allowRate"`
			AutoccUser                bool `json:"autoccUser"`
			CanEditInItemDet          bool `json:"canEditInItemDet"`
			Cdn75                     bool `json:"cdn75"`
			EnableAliMedicalComponent bool `json:"enableAliMedicalComponent"`
			GlobalSellItem            bool `json:"globalSellItem"`
			GoNewAuctionFlow          bool `json:"goNewAuctionFlow"`
			Is0YuanBuy                bool `json:"is0YuanBuy"`
			IsAliTelecomNew           bool `json:"isAliTelecomNew"`
			IsAlicomMasterCard        bool `json:"isAlicomMasterCard"`
			IsAllowReport             bool `json:"isAllowReport"`
			IsAreaSell                bool `json:"isAreaSell"`
			IsAutoFinancing           bool `json:"isAutoFinancing"`
			IsAutoYushou              bool `json:"isAutoYushou"`
			IsB2Byao                  bool `json:"isB2Byao"`
			IsBundleItem              bool `json:"isBundleItem"`
			IsCarCascade              bool `json:"isCarCascade"`
			IsCarService              bool `json:"isCarService"`
			IsCarYuEBao               bool `json:"isCarYuEBao"`
			IsChaoshi                 bool `json:"isChaoshi"`
			IsCloudShopItem           bool `json:"isCloudShopItem"`
			IsContractPhoneItem       bool `json:"isContractPhoneItem"`
			IsCustomizedItem          bool `json:"isCustomizedItem"`
			IsCyclePurchase           bool `json:"isCyclePurchase"`
			IsDianZiMendian           bool `json:"isDianZiMendian"`
			IsDownShelf               bool `json:"isDownShelf"`
			IsEnableAppleSku          bool `json:"isEnableAppleSku"`
			IsFullCarSell             bool `json:"isFullCarSell"`
			IsH5NewLogin              bool `json:"isH5NewLogin"`
			IsHasPos                  bool `json:"isHasPos"`
			IsHasQualification        bool `json:"isHasQualification"`
			IsHiddenShopAction        bool `json:"isHiddenShopAction"`
			IsHideAttentionBtn        bool `json:"isHideAttentionBtn"`
			IsHidePoi                 bool `json:"isHidePoi"`
			IsHkDirectSale            bool `json:"isHkDirectSale"`
			IsHkItem                  bool `json:"isHkItem"`
			IsHkO2OItem               bool `json:"isHkO2OItem"`
			IsHouseholdService        bool `json:"isHouseholdService"`
			IsIFCShop                 bool `json:"isIFCShop"`
			IsItemAllowSellerView     bool `json:"isItemAllowSellerView"`
			IsLadderGroupon           bool `json:"isLadderGroupon"`
			IsMeilihui                bool `json:"isMeilihui"`
			IsMeiz                    bool `json:"isMeiz"`
			IsMemberShopItem          bool `json:"isMemberShopItem"`
			IsMenDianInventroy        bool `json:"isMenDianInventroy"`
			IsNABundleItem            bool `json:"isNABundleItem"`
			IsNewAttraction           bool `json:"isNewAttraction"`
			IsNewMedical              bool `json:"isNewMedical"`
			IsNextDayService          bool `json:"isNextDayService"`
			IsO2OStaging              bool `json:"isO2OStaging"`
			IsOnePriceCar             bool `json:"isOnePriceCar"`
			IsOtcDrug                 bool `json:"isOtcDrug"`
			IsPreSellBrand            bool `json:"isPreSellBrand"`
			IsPrescriptionDrug        bool `json:"isPrescriptionDrug"`
			IsPurchaseMallVipBuyer    bool `json:"isPurchaseMallVipBuyer"`
			IsRegionLevel             bool `json:"isRegionLevel"`
			IsRx2                     bool `json:"isRx2"`
			IsRx2Count                bool `json:"isRx2Count"`
			IsSavingEnergy            bool `json:"isSavingEnergy"`
			IsService                 bool `json:"isService"`
			IsSevenDaysRefundment     bool `json:"isSevenDaysRefundment"`
			IsShowContentModuleTitle  bool `json:"isShowContentModuleTitle"`
			IsShowEcityBasicSign      bool `json:"isShowEcityBasicSign"`
			IsShowEcityDesc           bool `json:"isShowEcityDesc"`
			IsShowEcityVerticalSign   bool `json:"isShowEcityVerticalSign"`
			IsShowPreClosed           bool `json:"isShowPreClosed"`
			IsSkuColorShow            bool `json:"isSkuColorShow"`
			IsSkuMemorized            bool `json:"isSkuMemorized"`
			IsTeMai                   bool `json:"isTeMai"`
			IsTmallComboSupport       bool `json:"isTmallComboSupport"`
			IsVaccine                 bool `json:"isVaccine"`
			IsVitual3C                bool `json:"isVitual3C"`
			IsWTContract              bool `json:"isWTContract"`
			IsYY                      bool `json:"isYY"`
			IsYYZY                    bool `json:"isYYZY"`
			IsZhengChe                bool `json:"isZhengChe"`
			LoginBeforeCart           bool `json:"loginBeforeCart"`
			MlhNewDesc                bool `json:"mlhNewDesc"`
			Show9sVideo               bool `json:"show9sVideo"`
			ShowDiscountRecommend     bool `json:"showDiscountRecommend"`
			ShowFushiPoiInfo          bool `json:"showFushiPoiInfo"`
			ShowSuperMarketBuy        bool `json:"showSuperMarketBuy"`
			SupermarketAndQianggou    bool `json:"supermarketAndQianggou"`
			TimeKillAuction           bool `json:"timeKillAuction"`
			TryReportDisable          bool `json:"tryReportDisable"`
			UseTrade30                bool `json:"useTrade30"`
		} `json:"tags"`
	} `json:"tags"`
	TraceDatas struct {
		Mods_module_coupon_index struct {
			Module string `json:"module"`
		} `json:"mods/module-coupon/index"`
		Mods_module_hko2o_index struct {
			Module string `json:"module"`
		} `json:"mods/module-hko2o/index"`
		Mods_module_pintuan_index struct {
			Module string `json:"module"`
		} `json:"mods/module-pintuan/index"`
		Mods_module_prom_index struct {
			Module string `json:"module"`
		} `json:"mods/module-prom/index"`
		Mods_module_title_index struct {
			Module string `json:"module"`
		} `json:"mods/module-title/index"`
	} `json:"traceDatas"`
	Trade struct {
		BuyEnable   bool        `json:"buyEnable"`
		BuyParam    interface{} `json:"buyParam"`
		BuyText     interface{} `json:"buyText"`
		CartEnable  bool        `json:"cartEnable"`
		CartText    interface{} `json:"cartText"`
		Enable      bool        `json:"enable"`
		HintBanner  interface{} `json:"hintBanner"`
		RedirectURL interface{} `json:"redirectUrl"`
		StartTime   interface{} `json:"startTime"`
		TradeConfig struct {
			One string `json:"1"`
			Two string `json:"2"`
		} `json:"tradeConfig"`
		TradeParams struct {
			InputCharset string `json:"_input_charset"`
			BuyNow       string `json:"buyNow"`
		} `json:"tradeParams"`
		UseNativeTrade bool        `json:"useNativeTrade"`
		UseTrade30     bool        `json:"useTrade30"`
		WaitForStart   interface{} `json:"waitForStart"`
	} `json:"trade"`
}
