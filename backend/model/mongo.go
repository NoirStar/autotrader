package model

// User model
type User struct {
	ID       string  `bson:"id"`
	PW       string  `bson:"pw"`
	Email    string  `bson:"email"`
	Nickname string  `bson:"nickname"`
	Birth    string  `bson:"birth"`
	Level    int     `bson:"level"`
	Money    float64 `bson:"money"`
}

// Coin data model
type Coin struct {
	Type             string  `bson:"type"`               // 타입 : trade
	Code             string  `bson:"code"`               // 마켓코드
	TradePrice       float64 `bson:"trade_price"`        // 체결 가격
	TradeVolume      float64 `bson:"trade_volume"`       // 체결량
	AskBid           string  `bson:"ask_bid"`            // 매수/매도 구분 ASK : 매도 BID : 매수
	PrevClosingPrice float64 `bson:"prev_closing_price"` // 전일 종가
	Change           string  `bson:"change"`             // 전일 대비 RISE :  상승 EVEN : 보합 FALL : 하락
	ChangePrice      float64 `bson:"change_price"`       // 부호 없는 전일 대비 값
	TradeDate        string  `bson:"trade_date"`         // 체결 일자(UTC 기준)
	TradeTime        string  `bson:"trade_time"`         // 체결 시각(UTC 기준)
	TradeTimestamp   int     `bson:"trade_timestamp"`    // 체결 타임스탬프 (millisecond)
	Timestamp        int     `bson:"timestamp"`          // 타임스탬프 (millisecond)
	SequentialID     int     `bson:"sequential_id"`      // 체결 번호 (Unique)  - 체결의 유일성 판단을 위한 근거로 쓰일 수 있습니다. 하지만 체결의 순서를 보장하지는 못합니다.
	StreamType       string  `bson:"stream_type"`        // 스트림 타입  SNAPSHOT : 스냅샷 REALTIME : 실시간
}
