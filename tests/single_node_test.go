package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yu-org/yu/apps/asset"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/keypair"
	"github.com/yu-org/yu/core/startup"
	"github.com/yu-org/yu/core/types"
	cliAsset "github.com/yu-org/yu/example/client/asset"
	"github.com/yu-org/yu/example/client/callchain"
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

func TestSingleNode(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go runChain(t, &wg)
	time.Sleep(2 * time.Second)
	transferAsset(t)
	wg.Wait()
}

func runChain(t *testing.T, wg *sync.WaitGroup) {

	poaCfg := poa.DefaultCfg(0)
	yuCfg := startup.InitDefaultKernelConfig()
	// yuCfg.MaxBlockNum = 10
	yuCfg.IsAdmin = true

	// reset the history data
	os.RemoveAll(yuCfg.DataDir)

	assetTri := asset.NewAsset("yu-coin")
	poaTri := poa.NewPoa(poaCfg)

	chain := startup.InitDefaultKernel(yuCfg, poaTri, assetTri)
	chain.Startup()

	wg.Done()
}

func transferAsset(t *testing.T) {
	pubkey, privkey, err := keypair.GenKeyPair(keypair.Sr25519)
	if err != nil {
		panic("generate key error: " + err.Error())
	}

	toPubkey, _, err := keypair.GenKeyPair(keypair.Sr25519)
	if err != nil {
		panic("generate To Address key error: " + err.Error())
	}

	sub, err := callchain.NewSubscriber()
	if err != nil {
		panic("new subscriber failed: " + err.Error())
	}

	resultCh := make(chan *types.Receipt)
	go sub.SubEvent(resultCh)

	var (
		createAmount uint64 = 500
		transfer1    uint64 = 50
		transfer2    uint64 = 100
	)

	t.Log("-------- send Creating Account --------")
	cliAsset.CreateAccount(privkey, pubkey, createAmount)
	time.Sleep(10 * time.Second)
	balance := cliAsset.QueryAccount(pubkey)
	assert.Equal(t, createAmount, balance)

	t.Log("-------- send Transferring 1 --------")
	cliAsset.TransferBalance(privkey, pubkey, toPubkey.Address(), transfer1, 0)
	time.Sleep(8 * time.Second)

	balance1 := cliAsset.QueryAccount(pubkey)
	toBalance1 := cliAsset.QueryAccount(toPubkey)
	assert.Equal(t, createAmount-transfer1, balance1)
	assert.Equal(t, transfer1, toBalance1)

	t.Log("-------- send Transferring 2 --------")
	cliAsset.TransferBalance(privkey, pubkey, toPubkey.Address(), transfer2, 0)
	time.Sleep(6 * time.Second)

	balance2 := cliAsset.QueryAccount(pubkey)
	toBalance2 := cliAsset.QueryAccount(toPubkey)
	assert.Equal(t, createAmount-transfer1-transfer2, balance2)
	assert.Equal(t, transfer1+transfer2, toBalance2)

	resp, err := http.Get("http://localhost:7999/api/receipts_count?block_number=3")
	assert.NoError(t, err)

	byt, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	apiResp := new(kernel.APIResponse)
	err = json.Unmarshal(byt, apiResp)
	assert.NoError(t, err)
	t.Log("receipt count = ", apiResp.Data)

	_, err = http.Get("http://localhost:7999/api/admin/stop")
	assert.NoError(t, err)
}
