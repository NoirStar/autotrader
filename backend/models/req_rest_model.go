package models

// ReqChance struct 주문 가능 정보
type ReqChance struct {
	Market string `json:"market"` // 마켓 ID
}

// ReqOrderSearch struct 개별 주문 조회
// uuid 혹은 identifier 둘 중 하나의 값이 반드시 포함되어야 합니다.
type ReqOrderSearch struct {
	UUID       string `json:"uuid"`       // 주문 UUID
	Identifier string `json:"identifier"` // 조회용 사용자 지정 값
}

// ReqOrdersSearch 주문 리스트 조회
type ReqOrdersSearch struct {
	Market      string   `json:"market"`      // 마켓 ID
	UUIDS       []string `json:"uuids"`       // 주문 UUID의 목록
	Identifiers []string `json:"identifiers"` // 주문 identifier의 목록
	State       string   `json:"state"`       // wait : 체결 대기 (default) watch : 예약주문 대기 done : 전체 체결 완료 cancel : 주문 취소
	States      []string `json:"states"`      // 주문 상태의 목록 * 미체결 주문(wait, watch)과 완료 주문(done, cancel)은 혼합하여 조회하실 수 없습니다.
	Page        int      `json:"page"`        // 페이지 수, default: 1
	Limit       int      `json:"Limit"`       // 요청 개수, default: 100
	OrderBy     int      `json:"order_by"`    // asc: 오름차순, desc: 내림차순 (default)
}
