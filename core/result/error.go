package result

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	. "github.com/yu-org/yu/common"
)

type Error struct {
	Caller      Address  `json:"caller"`
	BlockStage  string   `json:"block_stage"`
	BlockHash   Hash     `json:"block_hash"`
	Height      BlockNum `json:"height"`
	TripodName  string   `json:"tripod_name"`
	WritingName string   `json:"writing_name"`
	Err         string   `json:"err"`
}

func (e *Error) Type() ResultType {
	return ErrorType
}

func (e *Error) Hash() (Hash, error) {
	byt, err := e.Encode()
	if err != nil {
		return NullHash, err
	}
	hash := sha256.Sum256(byt)
	return hash, nil
}

func (e *Error) Error() (str string) {
	if e.BlockStage == ExecuteTxnsStage {
		str = fmt.Sprintf(
			"[Error] Caller(%s) call Tripod(%s) Writing(%s) in Block(%s) on Height(%d): %s",
			e.Caller.String(),
			e.TripodName,
			e.WritingName,
			e.BlockHash.String(),
			e.Height,
			e.Err,
		)
	} else {
		str = fmt.Sprintf(
			"[Error] %s Block(%s) on Height(%d) in Tripod(%s): %s",
			e.BlockStage,
			e.BlockHash.String(),
			e.Height,
			e.TripodName,
			e.Err,
		)
	}
	return
}

func (e *Error) Encode() ([]byte, error) {
	byt, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	byt = append(ErrorTypeByt, byt...)
	return byt, nil
}

func (e *Error) Decode(data []byte) error {
	return json.Unmarshal(data[ResultTypeBytesLen:], e)
}

//
//type Errors []Error
//
//func ToErrors(errors []Error) Errors {
//	var es Errors
//	es = append(es, errors...)
//	return es
//}
//
//func (es Errors) ToArray() []Error {
//	return es[:]
//}
//
//func (es Errors) Encode() ([]byte, error) {
//	return GobEncode(es)
//}
