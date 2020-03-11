package page

import (
	"math"
)

type Data struct {
	TotalRows  uint        `json:"total_rows"`
	TotalPages uint        `json:"total_pages"`
	NowPage    uint        `json:"now_page"`
	List       interface{} `json:"list"`
}

func Page(count uint, viewRows uint, p uint) (data Data, limit uint, offset uint) {
	var totalRows uint = 0
	var totalPages uint = 0
	var nowPage uint = 0

	//总记录数量
	totalRows = count
	//总分页数量
	if viewRows == 0 {
		totalPages = 1
	} else {
		totalPages = uint(math.Ceil(float64(count) / float64(viewRows)))
	}
	//当前分页数
	if p == 0 {
		nowPage = 1
	} else if p > totalPages {
		nowPage = totalPages
	} else {
		nowPage = p
	}

	data = Data{
		TotalRows:  totalRows,
		TotalPages: totalPages,
		NowPage:    nowPage,
		List:       nil,
	}
	limit = viewRows
	offset = (nowPage - 1) * limit
	return
}
