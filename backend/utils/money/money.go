package money

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Money struct. formats money data
type Money struct {
	amount string
}

// NewMoney function. float64 데이터 타입을 업비트 원단위 포멧으로 변경
func NewMoney(amount float64) (*Money, error) {

	m := Money{}
	switch {
	case amount > 0 && amount < 10:
		m.amount = fmt.Sprintf("%.2f", amount)
	case amount >= 10 && amount < 100:
		m.amount = fmt.Sprintf("%.2f", amount)
	case amount >= 100 && amount < 1000:
		m.amount = strconv.Itoa(int(amount))
	case amount >= 1000:
		p := message.NewPrinter(language.English)
		m.amount = p.Sprintf("%d", uint32(amount))
	default:
		return &m, errors.New("Money is lower than 0")
	}
	return &m, nil
}

// Display function. 원단위 포멧 string 반환
func (m *Money) Display() string {
	return m.amount
}
