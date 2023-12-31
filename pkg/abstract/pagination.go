package abstract

import (
	"fmt"
	"math"
	"strconv"
)

const defaultSize = 10

type Pagination struct {
	Size    int    `json:"size"`
	Page    int    `json:"page"`
	OrderBy string `json:"order_by"`
}

func NewPagination(size, page int) Pagination {
	return Pagination{
		Size: size,
		Page: page,
	}
}

func (p *Pagination) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		p.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	p.Size = n
	return nil
}

func (p *Pagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		p.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	p.Page = n
	return nil
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetOrderBy() string {
	return p.OrderBy
}
func (p *Pagination) GetOffset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *Pagination) GetSize() int {
	return p.Size
}

func (p *Pagination) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", p.GetPage(), p.GetSize(), p.GetOrderBy())
}

func (p *Pagination) GetTotalPages(totalCount int) int {
	d := float64(totalCount) / float64(p.GetSize())
	return int(math.Ceil(d))
}

func (p *Pagination) GetHasMore(totalCount int) bool {
	return p.GetPage() < totalCount/p.GetSize()
}
