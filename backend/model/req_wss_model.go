package model

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

// ReqInfoWSS 데이터를 요청하기 위한 필요 필드들의 배열
type ReqInfoWSS []interface{}

// ReqTicketFieldWSS struct 시세를 수신하는 대상을 식별하며 되도록 유니크한 값을 사용하도록 권장
type ReqTicketFieldWSS struct {
	Ticket uuid.UUID `json:"ticket"`
}

// ReqTypeFieldWSS struct 수신하고 싶은 시세 정보를 나열
type ReqTypeFieldWSS struct {
	Type           string   `json:"type"`  // 수신할 시세 타입 현재가 : ticker , 체결 : trade, 호가 : orderbook
	Codes          []string `json:"codes"` // 수신할 시세 종목 정보.
	IsOnlySnapshot bool     `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime bool     `json:"isOnlyRealtime,omitempty"`
}

// ReqFormatFieldWSS struct 요청 포맷 정보 생략 가능
type ReqFormatFieldWSS struct {
	Format string `json:"format,omitempty"`
}

// NewReqForInfoWSS 정보 요청을 위한 구조체 생성
func NewReqForInfoWSS(reqType string, codes []string, isRealTime bool) *ReqInfoWSS {
	u := uuid.New()
	tkf := ReqTicketFieldWSS{Ticket: u}
	tyf := ReqTypeFieldWSS{}
	// ff := ReqFormatField{}
	if isRealTime {
		tyf = ReqTypeFieldWSS{
			Type:           reqType,
			Codes:          codes,
			IsOnlyRealtime: true,
		}
	} else {
		tyf = ReqTypeFieldWSS{
			Type:           reqType,
			Codes:          codes,
			IsOnlySnapshot: true,
		}
	}
	return &ReqInfoWSS{tkf, tyf}
}

// ReqForInfoJSON object 요청을 json 데이터로 변환
func (ri *ReqInfoWSS) ReqForInfoJSON() string {
	a, err := json.Marshal(ri)
	if err != nil {
		log.Fatalln(err)
	}
	return string(a)
}
