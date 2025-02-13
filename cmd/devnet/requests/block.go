package requests

import (
	"bytes"
	"fmt"

	"github.com/ledgerwatch/erigon/cmd/devnet/models"
	"github.com/ledgerwatch/erigon/cmd/rpctest/rpctest"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/types"
)

func GetBlockByNumber(reqId int, blockNum uint64, withTxs bool) (rpctest.EthBlockByNumber, error) {
	reqGen := initialiseRequestGenerator(reqId)
	var b rpctest.EthBlockByNumber

	req := reqGen.GetBlockByNumber(blockNum, withTxs)

	res := reqGen.Erigon(models.ETHGetBlockByNumber, req, &b)
	if res.Err != nil {
		return b, fmt.Errorf("error getting block by number: %v", res.Err)
	}

	if b.Error != nil {
		return b, fmt.Errorf("error populating response object: %v", b.Error)
	}

	return b, nil
}

func GetTransactionCount(reqId int, address common.Address, blockNum models.BlockNumber) (rpctest.EthGetTransactionCount, error) {
	reqGen := initialiseRequestGenerator(reqId)
	var b rpctest.EthGetTransactionCount

	if res := reqGen.Erigon(models.ETHGetTransactionCount, reqGen.GetTransactionCount(address, blockNum), &b); res.Err != nil {
		return b, fmt.Errorf("error getting transaction count: %v", res.Err)
	}

	if b.Error != nil {
		return b, fmt.Errorf("error populating response object: %v", b.Error)
	}

	return b, nil
}

func SendTransaction(reqId int, signedTx *types.Transaction) (*common.Hash, error) {
	reqGen := initialiseRequestGenerator(reqId)
	var b rpctest.EthSendRawTransaction

	var buf bytes.Buffer
	if err := (*signedTx).MarshalBinary(&buf); err != nil {
		return nil, fmt.Errorf("failed to marshal binary: %v", err)
	}

	if res := reqGen.Erigon(models.ETHSendRawTransaction, reqGen.SendRawTransaction(buf.Bytes()), &b); res.Err != nil {
		return nil, fmt.Errorf("could not make to request to eth_sendRawTransaction: %v", res.Err)
	}

	return &b.TxnHash, nil
}
