package model

// ResAccount struct 전체 계좌 조회
type ResAccount struct {
	Currency            string `json:"currency"`               // 화폐를 의미하는 영문 대문자 코드
	Balance             string `json:"balance"`                // 주문가능 금액/수량
	Locked              string `json:"locked"`                 // 주문 중 묶여있는 금액/수량
	AvgBuyPrice         string `json:"avg_buy_price"`          // 매수평균가
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"` // 매수평균가 수정 여부
	UsnitCurrency       string `json:"unit_currency"`          // 평단가 기준 화폐
}

// ResChance struct 주문 가능 정보
type ResChance struct {
	BidFee     string  `json:"bid_fee"`     // 매수 수수료 비율
	AskFee     string  `json:"ask_fee"`     // 매도 수수료 비율
	Market     market  `json:"market"`      // 마켓에 대한 정보
	BidAccount account `json:"bid_account"` // 매수 시 사용하는 화폐의 계좌 상태
	AskAccount account `json:"ask_account"` // 매도 시 사용하는 화폐의 계좌 상태
}

type market struct {
	ID         string    `json:"id"`          // 마켓의 유일 키
	Name       string    `json:"name"`        // 마켓 이름
	OrderTypes []string  `json:"order_types"` // 지원 주문 방식
	OrderSides []string  `json:"order_sides"` // 지원 주문 종류
	Bid        condition `json:"bid"`         // 매수 시 제약사항
	Ask        condition `json:"ask"`         // 매도 시 제약사항
	MaxTotal   string    `json:"max_total"`   // 최대 매도/매수 금액
	State      string    `json:"state"`       // 마켓 운영 상태

}

type condition struct {
	Currency  string `json:"currency"`   // 화폐를 의미하는 영문 대문자 코드
	PriceUnit string `json:"price_unit"` // 주문금액 단위
	MinTotal  string `json:"min_total"`  // 최소 매도/매수 금액
}

type account struct {
	Currency            string `json:"currency"`               // 화폐를 의미하는 영문 대문자 코드
	Balance             string `json:"balance"`                // 주문가능 금액/수량
	Locked              string `json:"locked"`                 // 주문 중 묶여있는 금액/수량
	AvgBuyPrice         string `json:"avg_buy_price"`          // 매수평균가
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"` // 매수평균가 수정 여부
	UnitCurrency        string `json:"unit_currency"`          // 평단가 기준 화폐
}

// ResOrderSearch struct 개별 주문 조회
type ResOrderSearch struct {
	UUID            string   `json:"uuid"`             // 주문의 고유 아이디
	Side            string   `json:"side"`             // 주문 종류
	OrdType         string   `json:"ord_type"`         // 주문 방식
	Price           string   `json:"price"`            // 주문 당시 화폐 가격
	State           string   `json:"state"`            // 주문 상태
	Market          string   `json:"market"`           // 마켓의 유일키
	CreatedAt       string   `json:"created_at"`       // 주문 생성 시간
	Volume          string   `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume string   `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     string   `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    string   `json:"remaining_fee"`    // 남은 수수료
	PaidFee         string   `json:"paid_fee"`         // 사용된 수수료
	Locked          string   `json:"locked"`           // 주문의 고유 아이디
	ExecutedVolume  string   `json:"executed_volume"`  // 해당 주문에 걸린 체결 수
	TradeCount      int      `json:"trade_count"`      // 주문의 고유 아이디
	Trades          []trades `json:"trades"`           // 체결
}

type trades struct {
	Market    string `json:"market"`     // 마켓의 유일 키
	UUID      string `json:"uuid"`       // 체결의 고유 아이디
	Price     string `json:"price"`      // 체결 가격
	Volume    string `json:"volume"`     // 체결 양
	Funds     string `json:"funds"`      // 체결된 총 가격
	Side      string `json:"side"`       // 체결 종류
	CreatedAt string `json:"created_at"` // 체결 시각
}

// ResOrdersSearch struct 주문 리스트 조회
type ResOrdersSearch struct {
	UUID            string `json:"uuid"`             // 주문의 고유 아이디
	Side            string `json:"side"`             // 주문 종류
	OrdType         string `json:"ord_type"`         // 주문 방식
	Price           string `json:"price"`            // 주문 당시 화폐 가격
	State           string `json:"state"`            // 주문 상태
	Market          string `json:"market"`           // 마켓의 유일키
	CreatedAt       string `json:"created_at"`       // 주문 생성 시간
	Volume          string `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume string `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     string `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    string `json:"remaining_fee"`    // 남은 수수료
	PaidFee         string `json:"paid_fee"`         // 사용된 수수료
	Locked          string `json:"locked"`           // 주문의 고유 아이디
	ExecutedVolume  string `json:"executed_volume"`  // 해당 주문에 걸린 체결 수
	TradeCount      int    `json:"trade_count"`      // 주문의 고유 아이디
}

// ResDeleteOrder struct 주문 취소 접수
type ResDeleteOrder struct {
	UUID            string `json:"uuid"`             // 주문의 고유 아이디
	Side            string `json:"side"`             // 주문 종류
	OrdType         string `json:"ord_type"`         // 주문 방식
	Price           string `json:"price"`            // 주문 당시 화폐 가격
	State           string `json:"state"`            // 주문 상태
	Market          string `json:"market"`           // 마켓의 유일키
	CreatedAt       string `json:"created_at"`       // 주문 생성 시간
	Volume          string `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume string `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     string `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    string `json:"remaining_fee"`    // 남은 수수료
	PaidFee         string `json:"paid_fee"`         // 사용된 수수료
	Locked          string `json:"locked"`           // 주문의 고유 아이디
	ExecutedVolume  string `json:"executed_volume"`  // 해당 주문에 걸린 체결 수
	TradeCount      int    `json:"trade_count"`      // 주문의 고유 아이디
}

// ResOrders struct 주문하기
type ResOrders struct {
	UUID            string `json:"uuid"`             // 주문의 고유 아이디
	Side            string `json:"side"`             // 주문 종류
	OrdType         string `json:"ord_type"`         // 주문 방식
	Price           string `json:"price"`            // 주문 당시 화폐 가격
	AvgPrice        string `json:"AvgPrice"`         // 주문 당시 화폐 가격
	State           string `json:"state"`            // 주문 상태
	Market          string `json:"market"`           // 마켓의 유일키
	CreatedAt       string `json:"created_at"`       // 주문 생성 시간
	Volume          string `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume string `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     string `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    string `json:"remaining_fee"`    // 남은 수수료
	PaidFee         string `json:"paid_fee"`         // 사용된 수수료
	Locked          string `json:"locked"`           // 주문의 고유 아이디
	ExecutedVolume  string `json:"executed_volume"`  // 해당 주문에 걸린 체결 수
	TradeCount      int    `json:"trade_count"`      // 주문의 고유 아이디
}

// ResMarketCode struct 마켓 코드 조회
type ResMarketCode struct {
	Market        string `json:"market"`         // 업비트에서 제공중인 시장 정보
	KoreanName    string `json:"korean_name"`    // 거래 대상 암호화폐 한글명
	EnglishName   string `json:"english_name"`   // 거래 대상 암호화폐 영문명
	MarketWarning string `json:"market_warning"` // 유의 종목 여부 NONE (해당 사항 없음), CAUTION(투자유의)
}

// ResMinuteCandles struct 시세 캔들 조회(분)
type ResMinuteCandles struct {
	Market               string  `json:"market"`                  // 마켓명
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`    // 캔들 기준 시각(UTC 기준)
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`    // 캔들 기준 시각(KST 기준)
	OpeningPrice         float64 `json:"opening_price"`           // 시가
	HighPrice            float64 `json:"high_price"`              // 고가
	LowPrice             float64 `json:"low_price"`               // 저가
	TradePrice           float64 `json:"trade_price"`             // 종가
	Timestamp            int64   `json:"timestamp"`               // 해당 캔들에서 마지막 틱이 저장된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량
	Unit                 int     `json:"unit"`                    // 분 단위(유닛)
}

// ResDayCandles struct 시세 캔들 조회(일)
type ResDayCandles struct {
	Market               string  `json:"market"`                          // 마켓명
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`            // 캔들 기준 시각(UTC 기준)
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`            // 캔들 기준 시각(KST 기준)
	OpeningPrice         float64 `json:"opening_price"`                   // 시가
	HighPrice            float64 `json:"high_price"`                      // 고가
	LowPrice             float64 `json:"low_price"`                       // 저가
	TradePrice           float64 `json:"trade_price"`                     // 종가
	Timestamp            int64   `json:"timestamp"`                       // 마지막 틱이 저장된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`          // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`         // 누적 거래량
	PrevClosingPrice     float64 `json:"prev_closing_price"`              // 전일 종가(UTC 0시 기준)
	ChangePrice          float64 `json:"change_price"`                    // 전일 종가 대비 변화 금액
	ChangeRate           float64 `json:"change_rate"`                     // 전일 종가 대비 변화량
	ConvertedTradePrice  float64 `json:"converted_trade_price,omitempty"` // 종가 환산 화폐 단위로 환산된 가격(요청에 convertingPriceUnit 파라미터 없을 시 해당 필드 포함되지 않음.)
}

// ResWeekCandles struct 시세 캔들 조회(주)
type ResWeekCandles struct {
	Market               string  `json:"market"`                  // 마켓명
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`    // 캔들 기준 시각(UTC 기준)
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`    // 캔들 기준 시각(KST 기준)
	OpeningPrice         float64 `json:"opening_price"`           // 시가
	HighPrice            float64 `json:"high_price"`              // 고가
	LowPrice             float64 `json:"low_price"`               // 저가
	TradePrice           float64 `json:"trade_price"`             // 종가
	Timestamp            int64   `json:"timestamp"`               // 마지막 틱이 저장된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량
	FirstDayOfPeriod     string  `json:"first_day_of_period"`     // 캔들 기간의 가장 첫 날
}

// ResMonthCandles struct 시세 캔들 조회(월)
type ResMonthCandles struct {
	Market               string  `json:"market"`                  // 마켓명
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`    // 캔들 기준 시각(UTC 기준)
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`    // 캔들 기준 시각(KST 기준)
	OpeningPrice         float64 `json:"opening_price"`           // 시가
	HighPrice            float64 `json:"high_price"`              // 고가
	LowPrice             float64 `json:"low_price"`               // 저가
	TradePrice           float64 `json:"trade_price"`             // 종가
	Timestamp            int64   `json:"timestamp"`               // 마지막 틱이 저장된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량
	FirstDayOfPeriod     string  `json:"first_day_of_period"`     // 캔들 기간의 가장 첫 날
}
