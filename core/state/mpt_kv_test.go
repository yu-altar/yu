package state

import (
	. "github.com/yu-org/yu/common"
	"github.com/yu-org/yu/config"
	"github.com/yu-org/yu/infra/storage/kv"
	"os"
	"testing"
)

var TestStateKvCfg = &config.MptKvConf{
	IndexDB: config.KVconf{KvType: "bolt", Path: "./state_index.db", Hosts: nil},
	NodeBase: config.KVconf{
		KvType: "bolt",
		Path:   "./state_base.db",
		Hosts:  nil,
	},
}

type TestTripod struct{}

func (tt *TestTripod) Name() string {
	return "test-tripod"
}

func TestKvCommit(t *testing.T) {
	kvdb, err := kv.NewKvdb(&config.KVconf{
		KvType: "bolt",
		Path:   "./test-mpt-kv.db",
		Hosts:  nil,
	})
	statekv := NewMptKV(kvdb)

	tri := &TestTripod{}

	statekv.Set(tri, []byte("dayu-key"), []byte("dayu-value"))

	statekv.NextTxn()

	stateRoot, err := statekv.Commit()
	if err != nil {
		t.Fatalf("commit state-kv error: %s", err.Error())
	}

	statekv.FinalizeBlock(NullHash)

	t.Logf("state-root is %s", stateRoot.String())

	value, err := statekv.Get(tri, []byte("dayu-key"))
	if err != nil {
		t.Fatalf("get state-kv error: %s", err.Error())
	}
	t.Logf("get value is %s", string(value))

	value, err = statekv.GetByBlockHash(tri, []byte("dayu-key"), NullHash)
	if err != nil {
		t.Fatalf("get state-kv by blockHash error: %s", err.Error())
	}
	t.Logf("get value by blockHash is %s", string(value))

	removeTestDB()
}

func removeTestDB() {
	os.RemoveAll(TestStateKvCfg.NodeBase.Path)
	os.RemoveAll(TestStateKvCfg.IndexDB.Path)
}
