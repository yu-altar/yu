package result

import (
	"fmt"
	. "github.com/Lawliet-Chan/yu/common"
	. "github.com/Lawliet-Chan/yu/utils/codec"
)

type Event struct {
	Caller     Address
	BlockStage string
	BlockHash  Hash
	Height     BlockNum
	TripodName string
	ExecName   string
	Value      string
}

func (e *Event) Encode() ([]byte, error) {
	return GlobalCodec.EncodeToBytes(e)
}

func (e *Event) Decode(data []byte) error {
	return GlobalCodec.DecodeBytes(data, e)
}

func (e *Event) Type() ResultType {
	return EventType
}

func (e *Event) Sprint() (str string) {
	if e.BlockStage == ExecuteTxnsStage {
		str = fmt.Sprintf(
			"[Event] Caller(%s) call Tripod(%s) Execution(%s) in Block(%s) on Height(%d): %v",
			e.Caller.String(),
			e.TripodName,
			e.ExecName,
			e.BlockHash,
			e.Height,
			e.Value,
		)
	} else {
		str = fmt.Sprintf(
			"[Event] %s Block(%s) on Height(%d) in Tripod(%s): %v",
			e.BlockStage,
			e.BlockHash,
			e.Height,
			e.TripodName,
			e.Value,
		)
	}
	return
}

//type Events []Event
//
//func ToEvents(events []Event) Events {
//	var es Events
//	es = append(es, events...)
//	return es
//}
//
//func (es Events) ToArray() []Event {
//	return es[:]
//}
//
//func (es Events) Encode() ([]byte, error) {
//	return GobEncode(es)
//}
