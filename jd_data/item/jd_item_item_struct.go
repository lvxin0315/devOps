package item

type JdItemOnly struct {
	Item struct {
		BrandID   string   `json:"brandId"`
		BrandName string   `json:"brandName"`
		Category  []string `json:"category"`
		//DataFrom     int64    `json:"dataFrom"`
		//Description  string   `json:"description"`
		Height int64 `json:"height"`
		//Image        []string `json:"image"`
		//IsPop        bool     `json:"isPop"`
		Length int64 `json:"length"`
		//Model        string   `json:"model"`
		NewColorSize []struct {
			One         string `json:"1"`
			Two         string `json:"2"`
			Three       string `json:"3"`
			SpecName    string `json:"SpecName"`
			Dim         int64  `json:"dim"`
			SequenceNo1 string `json:"sequenceNo1"`
			SequenceNo2 string `json:"sequenceNo2"`
			SequenceNo3 string `json:"sequenceNo3"`
			SkuID       string `json:"skuId"`
		} `json:"newColorSize"`
		PName string `json:"pName"`
		//PTag        int64  `json:"pTag"`
		ProductArea string `json:"productArea"`
		SaleProp    struct {
			One   string `json:"1"`
			Two   string `json:"2"`
			Three string `json:"3"`
		} `json:"saleProp"`
		SalePropSeq struct {
			One   []string `json:"1"`
			Two   []string `json:"2"`
			Three []string `json:"3"`
		} `json:"salePropSeq"`
		SaleUnit string `json:"saleUnit"`
		SkuID    string `json:"skuId"`
		//SkuMark  string `json:"skuMark"`
		SkuName string `json:"skuName"`
		SkuType string `json:"skuType"`
		//Skumark  int64  `json:"skumark"`
		//SpAttr   struct {
		//	YuShou            string `json:"YuShou"`
		//	Fare              string `json:"fare"`
		//	Is7ToReturn       string `json:"is7ToReturn"`
		//	IsFlashPurchase   string `json:"isFlashPurchase"`
		//	Jgys              string `json:"jgys"`
		//	NationallySetWare string `json:"nationallySetWare"`
		//	Sfkc              string `json:"sfkc"`
		//	Thwa              string `json:"thwa"`
		//	Tsfw              string `json:"tsfw"`
		//	Yhzy              string `json:"yhzy"`
		//} `json:"spAttr"`
		//TimelinessID  interface{}  `json:"timelinessId"`
		//UnLimitCid    int64  `json:"unLimit_cid"`
		//Upc           string `json:"upc"`
		//ValuePayFirst int64  `json:"valuePayFirst"`
		//VenderID      string `json:"venderID"`
		//Warestatus    string `json:"warestatus"`
		Weight string `json:"weight"`
		Width  int64  `json:"width"`
		//Yn            int64  `json:"yn"`
	} `json:"item"`
}

type JdItemInfo struct {
	//AdvertCount struct {
	//	Ad string `json:"ad"`
	//	ID string `json:"id"`
	//} `json:"AdvertCount"`
	//AllOverImg string `json:"allOverImg"`
	//AreaID     string `json:"areaId"`
	//Bankpromo  struct {
	//	ActURL     string `json:"actUrl"`
	//	ActivityID string `json:"activityId"`
	//	Expire     string `json:"expire"`
	//	Title      string `json:"title"`
	//} `json:"bankpromo"`
	//Bigouinfo     string        `json:"bigouinfo"`
	//BizMsg        string        `json:"bizMsg"`
	//BizRetCode    string        `json:"bizRetCode"`
	//ChnImg        []interface{} `json:"chnImg"`
	//Daojia        string        `json:"daojia"`
	//ErrCode       string        `json:"errCode"`
	//Flag          struct{}      `json:"flag"`
	//HasSubscribe  string        `json:"hasSubscribe"`
	//HuanURL       string        `json:"huanUrl"`
	//InfoVideoID   string        `json:"infoVideoId"`
	//IsFestival    string        `json:"isFestival"`
	//IsMaskSku     string        `json:"isMaskSku"`
	//IsNeedEncrypt string        `json:"isNeedEncrypt"`
	//Kanjia        string        `json:"kanjia"`
	//MagicLevel    string        `json:"magicLevel"`
	//MainVideoID   string        `json:"mainVideoId"`
	//Model3DId     string        `json:"model3DId"`
	//Msg           string        `json:"msg"`
	//Pingou        string        `json:"pingou"`
	//PlusFlag      string        `json:"plusFlag"`
	//PlusLimitBuy  struct {
	//	LimitNum      int64    `json:"limitNum"`
	//	LimitText     string   `json:"limitText"`
	//	NoSaleFlag    int64    `json:"noSaleFlag"`
	//	NoSaleText    string   `json:"noSaleText"`
	//	PromotionText string   `json:"promotionText"`
	//	ResultErrMsg  string   `json:"resultErrMsg"`
	//	ResultExt     struct{} `json:"resultExt"`
	//	ResultFlag    bool     `json:"resultFlag"`
	//} `json:"plusLimitBuy"`
	//PlusMemberType string `json:"plusMemberType"`
	Price struct {
		Bp  string `json:"bp"`
		Ext string `json:"ext"`
		Fmp string `json:"fmp"`
		ID  string `json:"id"`
		L   string `json:"l"`
		M   string `json:"m"`
		Nup string `json:"nup"`
		Op  string `json:"op"`
		P   string `json:"p"`
		Pcp string `json:"pcp"`
		Sfp string `json:"sfp"`
		Sp  string `json:"sp"`
		Stp string `json:"stp"`
		Tkp string `json:"tkp"`
		Tpp string `json:"tpp"`
		Up  string `json:"up"`
		Vdp string `json:"vdp"`
	} `json:"price"`
	//Promov2 []struct {
	//	Hit string `json:"hit"`
	//	ID  string `json:"id"`
	//	Pis []struct {
	//		Adurl      string `json:"adurl"`
	//		Customtag  string `json:"customtag"`
	//		D          string `json:"d"`
	//		Ori        string `json:"ori"`
	//		Pid        string `json:"pid"`
	//		St         string `json:"st"`
	//		Subextinfo string `json:"subextinfo"`
	//	} `json:"pis"`
	//} `json:"promov2"`
	//Ptqq    string `json:"ptqq"`
	//RetCode string `json:"retCode"`
	//RuID    string `json:"ruId"`
	//Sence   string `json:"sence"`
	Stock struct {
		ArrivalDate    string        `json:"ArrivalDate"`
		Dc             []interface{} `json:"Dc"`
		Dti            interface{}   `json:"Dti"`
		Ext            string        `json:"Ext"`
		IsPurchase     bool          `json:"IsPurchase"`
		PlusFlagInfo   string        `json:"PlusFlagInfo"`
		PopType        int64         `json:"PopType"`
		StockState     int64         `json:"StockState"`
		StockStateName string        `json:"StockStateName"`
		Ab             string        `json:"ab"`
		Ac             string        `json:"ac"`
		Ad             string        `json:"ad"`
		Ae             string        `json:"ae"`
		Af             string        `json:"af"`
		AfsCode        int64         `json:"afsCode"`
		Ag             string        `json:"ag"`
		Area           struct {
			CityName     string `json:"cityName"`
			CountyName   string `json:"countyName"`
			ProvinceName string `json:"provinceName"`
			Success      bool   `json:"success"`
			TownName     string `json:"townName"`
		} `json:"area"`
		AreaLevel         int64         `json:"areaLevel"`
		Channel           int64         `json:"channel"`
		Cla               []interface{} `json:"cla"`
		Code              int64         `json:"code"`
		DcID              string        `json:"dcId"`
		DcashDesc         string        `json:"dcashDesc"`
		Eb                string        `json:"eb"`
		Ec                string        `json:"ec"`
		FreshEdi          interface{}   `json:"freshEdi"`
		Ir                []interface{} `json:"ir"`
		IsJDexpress       string        `json:"isJDexpress"`
		IsPlus            bool          `json:"isPlus"`
		IsSam             bool          `json:"isSam"`
		IsSopUseSelfStock string        `json:"isSopUseSelfStock"`
		IsWalMar          bool          `json:"isWalMar"`
		JdPrice           struct {
			ID string `json:"id"`
			M  string `json:"m"`
			Op string `json:"op"`
			P  string `json:"p"`
		} `json:"jdPrice"`
		M                 string `json:"m"`
		NationallySetWare string `json:"nationallySetWare"`
		Pr                struct {
			PromiseResult string `json:"promiseResult"`
			ResultCode    int64  `json:"resultCode"`
		} `json:"pr"`
		PromiseMark   string `json:"promiseMark"`
		PromiseResult string `json:"promiseResult"`
		PromiseYX     struct {
			HelpLink   string `json:"helpLink"`
			IconCode   string `json:"iconCode"`
			IconSrc    string `json:"iconSrc"`
			IconTip    string `json:"iconTip"`
			IconType   int64  `json:"iconType"`
			PicURL     string `json:"picUrl"`
			ResultCode int64  `json:"resultCode"`
			ShowName   string `json:"showName"`
		} `json:"promiseYX"`
		RealSkuID int64  `json:"realSkuId"`
		Rfg       int64  `json:"rfg"`
		Rid       string `json:"rid"`
		Rn        int64  `json:"rn"`
		SelfD     struct {
			Cg        string      `json:"cg"`
			ColType   int64       `json:"colType"`
			Deliver   string      `json:"deliver"`
			Df        interface{} `json:"df"`
			ID        int64       `json:"id"`
			Linkphone string      `json:"linkphone"`
			Po        string      `json:"po"`
			Type      int64       `json:"type"`
			URL       string      `json:"url"`
			Vender    string      `json:"vender"`
			Vid       int64       `json:"vid"`
		} `json:"self_D"`
		ServiceInfo string        `json:"serviceInfo"`
		Sid         string        `json:"sid"`
		SidDely     string        `json:"sidDely"`
		SkuID       int64         `json:"skuId"`
		SkuState    int64         `json:"skuState"`
		Sr          interface{}   `json:"sr"`
		StockDesc   string        `json:"stockDesc"`
		Support     []interface{} `json:"support"`
		V           string        `json:"v"`
		Vd          interface{}   `json:"vd"`
		VenderType  interface{}   `json:"venderType"`
		WeightValue string        `json:"weightValue"`
	} `json:"stock"`
	Upc    string `json:"upc"`
	WqAddr string `json:"wq_addr"`
	//Yuyue  struct {
	//	Address          string `json:"address"`
	//	Category         string `json:"category"`
	//	D                int64  `json:"d"`
	//	Etime            string `json:"etime"`
	//	Flag             bool   `json:"flag"`
	//	HasAddress       bool   `json:"hasAddress"`
	//	HidePrice        int64  `json:"hidePrice"`
	//	Info             string `json:"info"`
	//	InsertTime       int64  `json:"insertTime"`
	//	IsBefore         int64  `json:"isBefore"`
	//	IsJ              int64  `json:"isJ"`
	//	Num              int64  `json:"num"`
	//	PlusD            int64  `json:"plusD"`
	//	PlusEtime        string `json:"plusEtime"`
	//	PlusStime        string `json:"plusStime"`
	//	PlusType         int64  `json:"plusType"`
	//	QiangEtime       string `json:"qiangEtime"`
	//	QiangStime       string `json:"qiangStime"`
	//	RiskCheck        string `json:"riskCheck"`
	//	SellWhilePresell string `json:"sellWhilePresell"`
	//	ShowPromoPrice   string `json:"showPromoPrice"`
	//	Sku              int64  `json:"sku"`
	//	State            int64  `json:"state"`
	//	Stime            string `json:"stime"`
	//	Type             string `json:"type"`
	//	URL              string `json:"url"`
	//	YueEtime         string `json:"yueEtime"`
	//	YueStime         string `json:"yueStime"`
	//} `json:"yuyue"`
	//YuyueDraw string `json:"yuyueDraw"`
}
