package txpool

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core/types"
)

type orderedTxns struct {
	txns []*SignedTxn
	idx  map[Hash]*SignedTxn

	// order map[int]Hash
}

func newOrderedTxns() *orderedTxns {
	return &orderedTxns{
		txns: make([]*SignedTxn, 0),
		idx:  make(map[Hash]*SignedTxn),
		// order: make(map[int]Hash),
	}
}

func (ot *orderedTxns) Insert(input *SignedTxn) {
	logrus.WithField("txpool", "ordered-txns").
		Tracef("Insert txn(%s) to Txpool, txn content: %v", input.TxnHash, input.Raw.WrCall)

	// fmt.Println("*** txpool insert: ", input.TxnHash.String())

	ot.idx[input.TxnHash] = input
	ot.txns = append(ot.txns, input)
}

func (ot *orderedTxns) SetOrder(order map[int]Hash) {
	panic("implement me")
	//for i, hash := range order {
	//	ot.setOrder(i, hash)
	//}
}

//func (ot *orderedTxns) setOrder(num int, txHash Hash) {
//	if _, ok := ot.order[num]; ok {
//		ot.setOrder(num+1, txHash)
//	}
//	ot.order[num] = txHash
//}

func (ot *orderedTxns) delete(txnHash Hash) {
	if stxn, ok := ot.idx[txnHash]; ok {
		logrus.WithField("txpool", "ordered-txns").
			Tracef("DELETE txn(%s) from txpool, txn content: %v", stxn.TxnHash.String(), stxn.Raw.WrCall)
		delete(ot.idx, txnHash)
		//delete(ot.order, txnHash)
	}
	for i := 0; i < len(ot.txns); i++ {
		if ot.txns[i].TxnHash == txnHash {
			// fmt.Println("*** txpool delete: ", txnHash.String())
			ot.txns = append(ot.txns[:i], ot.txns[i+1:]...)
			i--
		}
	}

	//for i, hash := range ot.order {
	//	if hash == txnHash {
	//		delete(ot.order, i)
	//
	//		// ot.order = append(ot.order[:i], ot.order[i+1:]...)
	//	}
	//}
}

func (ot *orderedTxns) Deletes(txnHashes []Hash) {
	fmt.Println("txpool count = ", ot.Size())
	for _, txn := range ot.txns {
		params := make(map[string]any)
		err := txn.BindJsonParams(&params)
		if err != nil {
			panic(err)
		}
		fmt.Println("txn is: ", params)
	}
	fmt.Println("delete txs count = ", len(txnHashes))
	for _, hash := range txnHashes {
		ot.delete(hash)
	}
}

func (ot *orderedTxns) Exist(txnHash Hash) bool {
	_, ok := ot.idx[txnHash]
	return ok
}

func (ot *orderedTxns) Get(txnHash Hash) *SignedTxn {
	return ot.idx[txnHash]
}

func (ot *orderedTxns) Gets(numLimit uint64, filter func(txn *SignedTxn) bool) []*SignedTxn {
	if numLimit > uint64(ot.Size()) {
		numLimit = uint64(ot.Size())
	}

	txns := make([]*SignedTxn, 0)

	//ot.excludeEmptyOrder()
	//
	//for num, hash := range ot.order {
	//	// FIXME: if num > numLimit, panic here
	//	if filter(ot.idx[hash]) {
	//		txns[num] = ot.idx[hash]
	//	}
	//}

	fmt.Printf("numLimit = %d, length of txpool is %d\n", numLimit, ot.Size())

	for i := 0; i < int(numLimit); i++ {
		//if txns[i] != nil {
		//	continue
		//}

		// fmt.Println("index of packing tx: ", i)

		txn := ot.txns[i]

		if filter(txn) {
			logrus.WithField("txpool", "ordered-txns").
				Tracef("Pack txn(%s) from Txpool, txn content: %v", txn.TxnHash, txn.Raw.WrCall)
			params := make(map[string]any)
			err := txn.BindJsonParams(&params)
			if err != nil {
				panic(err)
			}
			fmt.Println("pack txn: ", params)
			txns = append(txns, txn)
			// txns[i] = txn
		}
	}
	return txns
}

//func (ot *orderedTxns) excludeEmptyOrder() {
//	for num, hash := range ot.order {
//		if _, ok := ot.idx[hash]; !ok {
//			delete(ot.order, num)
//		}
//	}
//}

func (ot *orderedTxns) SortTxns(fn func(txns []*SignedTxn) []*SignedTxn) {
	orderedTxs := fn(ot.txns)
	ot.txns = orderedTxs
}

func (ot *orderedTxns) GetAll() []*SignedTxn {
	return ot.txns
}

func (ot *orderedTxns) Size() int {
	return len(ot.txns)
}
