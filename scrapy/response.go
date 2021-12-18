package scrapy

type ResponseData struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	CompanyList []CompanyList `json:"companyList"`
	PageInfo    PageInfo      `json:"pageInfo"`
}

type CompanyList struct {
	OfficalURL       string     `json:"officalUrl"`
	XcxLogo          string     `json:"xcxLogo"`
	InterviewTime    string     `json:"interviewTime"`
	Subject          *string    `json:"subject,omitempty"`
	UnactiveLogo     string     `json:"unactiveLogo"`
	MiniLogo         string     `json:"miniLogo"`
	OfferTime        string     `json:"offerTime"`
	NeituiTime       string     `json:"neituiTime"`
	OfficalEncodeURL string     `json:"officalEncodeUrl"`
	CreateTime       int64      `json:"createTime"`
	Schedules        []Schedule `json:"schedules"`
	Name             string     `json:"name"`
	Logo             string     `json:"logo"`
	Rank             int64      `json:"rank"`
	LogoRadius       string     `json:"logoRadius"`
	TestTime         string     `json:"testTime"`
	End              int64      `json:"end"`
	ID               int64      `json:"id"`
	Guide            []Guide    `json:"guide"`
	WangshenTime     string     `json:"wangshenTime"`
}

type Guide struct {
	BtnName   string `json:"btnName"`
	Name      string `json:"name"`
	EncodeURL string `json:"encodeUrl"`
	URL       string `json:"url"`
}

type Schedule struct {
	Name      string `json:"name"`
	Time      string `json:"time"`
	EncodeURL string `json:"encodeUrl"`
	Content   string `json:"content"`
	URL       string `json:"url"`
}

type PageInfo struct {
	PageCount    int64 `json:"pageCount"`
	PageSize     int64 `json:"pageSize"`
	ElementCount int64 `json:"elementCount"`
	PageCurrent  int64 `json:"pageCurrent"`
}

// type BtnName string

// const (
// 	剑指Offer BtnName = "剑指offer"
// 	查看 BtnName = "查看"
// )

// type GuideName string
// const (
// 	Name刷公司往年真题 GuideName = "刷公司往年真题 "
// 	Name面试常考点必备 GuideName = "面试常考点必备 "
// 	刷公司往年真题 GuideName = "刷公司往年真题"
// 	校招Fqa GuideName = "校招FQA"
// 	看笔经面经 GuideName = "看笔经面经"
// 	面试常考点必备 GuideName = "面试常考点必备"
// )

// type ScheduleName string
// const (
// 	Empty ScheduleName = ""
// 	寻找内部员工内推 ScheduleName = "寻找内部员工内推"
// 	投递方式 ScheduleName = "投递方式"
// 	投递链接 ScheduleName = "投递链接"
// 	招聘详情 ScheduleName = "招聘详情"
// )
