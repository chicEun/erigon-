// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package jsonrpc

import (
	"context"

	"github.com/erigontech/erigon-lib/chain"
	"github.com/erigontech/erigon-lib/kv"
	"github.com/erigontech/erigon-lib/log/v3"
	"github.com/erigontech/erigon-lib/types"
	"github.com/erigontech/erigon/core/tracing"
	"github.com/erigontech/erigon/execution/exec3"
)

type GenericTracer interface {
	TracingHooks() *tracing.Hooks
	SetTransaction(tx types.Transaction)
	Found() bool
}

func (api *OtterscanAPIImpl) genericTracer(tx kv.TemporalTx, ctx context.Context, blockNum, txnID uint64, txIndex int, chainConfig *chain.Config, tracer GenericTracer) error {
	executor := exec3.NewTraceWorker(tx, chainConfig, api.engine(), api._blockReader, tracer)
	defer executor.Close()

	// if block number changed, calculate all related field
	header, err := api._blockReader.HeaderByNumber(ctx, tx, blockNum)
	if err != nil {
		return err
	}
	if header == nil {
		log.Warn("[rpc] header is nil", "blockNum", blockNum)
		return nil
	}
	executor.ChangeBlock(header)

	txn, err := api._txnReader.TxnByIdxInBlock(ctx, tx, blockNum, txIndex)
	if err != nil {
		return err
	}
	if txn == nil {
		log.Warn("[rpc genericTracer] txn is nil", "blockNum", blockNum, "txIndex", txIndex)
		return nil
	}
	err = executor.ExecTxn(txnID, txIndex, txn, false)
	if err != nil {
		return err
	}
	return nil
}
