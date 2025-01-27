package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type BaseModel struct {
	ID        string     `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// ConvertType is a utility function to convert one type to another.
func convertType[T any, U any](from *T, to *U) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}

// Response is a generic function to convert a struct to a response type.
func Response[T any, U any](input *T) *U {
	var output U
	err := convertType(input, &output)
	if err != nil {
		fmt.Printf("Error converting type: %v\n", err)
		return nil
	}
	return &output
}

func (req *ClientListRequest) Prepare() {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 20
	}
	if req.SortColumn == "" {
		req.SortColumn = "id"
	}
	if req.SortDirection == "" {
		req.SortDirection = "desc"
	}
	endDate := time.Now()
	startDate := endDate.AddDate(0, -1, 0)
	if req.StartDate != "" {
		startDate, _ = time.Parse("2006-01-02", req.StartDate)
	}
	if req.EndDate != "" {
		endDate, _ = time.Parse("2006-01-02", req.EndDate)
	}
	req.StartDate = startDate.Format("2006-01-02")
	req.EndDate = endDate.Format("2006-01-02") + " 23:59:59.999"
}
