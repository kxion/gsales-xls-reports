package model

import (
	"errors"
)

// ReportType int
type ReportType int

// Constants
const (
	MonthlySalesReport ReportType = iota + 1
	PayPeriodReport
)

// ReportStringToType function
func ReportStringToType(rType string) (ReportType, error) {
	var rt ReportType

	switch rType {
	case "monthlysales":
		rt = MonthlySalesReport
	case "payperiod":
		rt = PayPeriodReport
	default:
		rt = 0
	}
	if rt == 0 {
		return rt, errors.New("Invalid report type request")
	}
	return rt, nil
}
