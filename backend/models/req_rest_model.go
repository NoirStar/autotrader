package models

// ReqChance struct 주문 가능 정보
type ReqChance struct {
	Market string `json:"market"` // 마켓 ID
}

// ReqOrderSearch struct 개별 주문 조회
// uuid 혹은 identifier 둘 중 하나의 값이 반드시 포함되어야 합니다.
type ReqOrderSearch struct {
	UUID       string `json:"uuid,omitempty"`       // 주문 UUID
	Identifier string `json:"identifier,omitempty"` // 조회용 사용자 지정 값
}

// ReqOrdersSearch 주문 리스트 조회
type ReqOrdersSearch struct {
	Market      string   `json:"market"`             // 마켓 ID
	UUIDS       []string `json:"uuids"`              // 주문 UUID의 목록
	Identifiers []string `json:"identifiers"`        // 주문 identifier의 목록
	State       string   `json:"state"`              // wait : 체결 대기 (default) watch : 예약주문 대기 done : 전체 체결 완료 cancel : 주문 취소
	States      []string `json:"states"`             // 주문 상태의 목록 * 미체결 주문(wait, watch)과 완료 주문(done, cancel)은 혼합하여 조회하실 수 없습니다.
	Page        int      `json:"page,omitempty"`     // 페이지 수, default: 1
	Limit       int      `json:"limit,omitempty"`    // 요청 개수, default: 100
	OrderBy     int      `json:"order_by,omitempty"` // asc: 오름차순, desc: 내림차순 (default)
}

// ReqDeleteOrder 주문 취소 접수
// uuid 혹은 identifier 둘 중 하나의 값이 반드시 포함되어야 합니다.
type ReqDeleteOrder struct {
	UUID       string `json:"uuid,omitempty"`       // 주문 UUID
	Identifier string `json:"identifier,omitempty"` // 조회용 사용자 지정 값
}

// ReqOrders struct 주문하기
type ReqOrders struct {
	Market     string `json:"market"`               // 마켓 ID (필수)
	Side       string `json:"side"`                 // 주문 종류 (필수)
	Volume     string `json:"volume,omitempty"`     // 주문량 (지정가, 시장가 매도 시 필수)
	Price      string `json:"price,omitempty"`      // 주문 가격. (지정가, 시장가 매수 시 필수)
	OrdType    string `json:"ord_type"`             // 주문 타입 (필수) / limit : 지정가 주문, price : 시장가 주문(매수), market : 시장가 주문(매도)
	Identifier string `json:"identifier,omitempty"` // 조회용 사용자 지정값 (선택)
}

// ReqMinuteCandles struct 시세 캔들 조회(분)
type ReqMinuteCandles struct {
	Unit   string `json:"unit"`   // 분 단위. 가능한 값 : 1, 3, 5, 15, 10, 30, 60, 240
	Market string `json:"market"` // 마켓 코드 (ex. KRW-BTC)
	To     string `json:"to"`     // 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss. 비워서 요청시 가장 최근 캔들
	Count  uint32 `json:"count"`  // 캔들 개수(최대 200개까지 요청 가능)
}

// ReqDayCandles struct 시세 캔들 조회(일)
type ReqDayCandles struct {
	Market              string `json:"market"`                        // 마켓 코드 (ex. KRW-BTC)
	To                  string `json:"to"`                            // 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss. 비워서 요청시 가장 최근 캔들
	Count               uint32 `json:"count"`                         // 캔들 개수
	ConvertingPriceUnit string `json:"convertingPriceUnit,omitempty"` // 종가 환산 화폐 단위 (생략 가능, KRW로 명시할 시 원화 환산 가격을 반환.)
}

// ReqWeekCandles struct 시세 캔들 조회(주)
type ReqWeekCandles struct {
	Market string `json:"market"` // 마켓 코드 (ex. KRW-BTC)
	To     string `json:"to"`     // 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss. 비워서 요청시 가장 최근 캔들
	Count  uint32 `json:"count"`  // 캔들 개수
}

// ReqMonthCandles struct 시세 캔들 조회(월)
type ReqMonthCandles struct {
	Market string `json:"market"` // 마켓 코드 (ex. KRW-BTC)
	To     string `json:"to"`     // 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss. 비워서 요청시 가장 최근 캔들
	Count  uint32 `json:"count"`  // 캔들 개수
}
