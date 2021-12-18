package scrapy

import (
	"encoding/json"
	"fmt"
	"net/http"
	gourl "net/url"
	"os"
)

const url = "https://www.nowcoder.com/school/schedule/data?token=&query=%s&typeId=0&propertyId=0&onlyFollow=false&page=%d"

func New(options ...Option) *NewCoder {
	n := &NewCoder{
		query: "",
	}
	for _, option := range options {
		option.apply(n)
	}
	return n
}

type NewCoder struct {
	query string
}

func (n *NewCoder) SaveJson(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := n.getData()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(f)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(" ", "  ")
	return encoder.Encode(data)
}

func (n *NewCoder) GetPageData(page int) ([]CompanyList, error) {
	data, err := n.getDataPerPage(page)
	if err != nil {
		return nil, err
	}
	return data.Data.CompanyList, nil
}

func (n *NewCoder) getDataPerPage(page int) (*ResponseData, error) {

	url := fmt.Sprintf(url, gourl.QueryEscape(n.query), page)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.57")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r ResponseData
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *NewCoder) RawData() ([]CompanyList, error) {
	return n.getData()
}

func (n *NewCoder) getData() ([]CompanyList, error) {
	var res = make([]CompanyList, 0, 1024)
	r, err := n.getDataPerPage(1)
	if err != nil {
		return nil, err
	}
	totalPage := r.Data.PageInfo.PageCount
	page := 2
	for page <= int(totalPage) {
		r, err := n.getDataPerPage(page)
		if err != nil {
			return nil, err
		}
		res = append(res, r.Data.CompanyList...)
		page++
	}
	return res, nil
}
