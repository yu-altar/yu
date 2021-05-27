package node

import (
	. "github.com/Lawliet-Chan/yu/blockchain"
	"github.com/Lawliet-Chan/yu/chain_env"
	. "github.com/Lawliet-Chan/yu/common"
	"github.com/Lawliet-Chan/yu/context"
	. "github.com/Lawliet-Chan/yu/tripod"
)

func ExecuteTxns(block IBlock, env *chain_env.ChainEnv, land *Land) error {
	chain := env.Chain
	base := env.Base
	sub := env.Sub

	blockHash := block.GetHeader().GetHash()
	stxns, err := base.GetTxns(blockHash)
	if err != nil {
		return err
	}
	for _, stxn := range stxns {
		ecall := stxn.GetRaw().GetEcall()
		ctx, err := context.NewContext(stxn.GetPubkey().Address(), ecall.Params)
		if err != nil {
			return err
		}
		err = land.Execute(ecall, ctx, env)
		if err != nil {
			return err
		}

		err = chain.UpdateBlock(block)
		if err != nil {
			return err
		}

		for _, event := range ctx.Events {
			event.Height = block.GetHeader().GetHeight()
			event.BlockHash = blockHash
			event.ExecName = ecall.ExecName
			event.TripodName = ecall.TripodName
			event.BlockStage = ExecuteTxnsStage
			event.Caller = stxn.GetRaw().GetCaller()

			if sub != nil {
				sub.Push(event)
			}
		}

		for _, e := range ctx.Errors {
			e.Caller = stxn.GetRaw().GetCaller()
			e.BlockStage = ExecuteTxnsStage
			e.TripodName = ecall.TripodName
			e.ExecName = ecall.ExecName
			e.BlockHash = blockHash
			e.Height = block.GetHeader().GetHeight()

			if sub != nil {
				sub.Push(e)
			}
		}

		err = base.SetEvents(ctx.Events)
		if err != nil {
			return err
		}
		err = base.SetErrors(ctx.Errors)
		if err != nil {
			return err
		}
	}
	return nil
}
