/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/openw"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate, memo string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, memo, contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract) ([]*openwallet.RawTransaction, error) {

	rawTxArray, err := tm.CreateSummaryTransaction(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMSVS51x4XgWNgd354yxd4yC7CrESsA8dN"
	accountID := "4bpD665trhfF69MeZQUj5hpCqogfetEwWoGTrRMSuxvV"
	//accountID := "4bpD665trhfF69MeZQUj5hpCqogfetEwWoGTrRMSuxvV"
	//to := "APWPMowFEb6KdVn9FRvdgdWZMVo1wmnwvp"
	//accountID := "F7aeTnSdjEA16x4H3n1vPtDEo9Xp5Vus11pwY5QF6K3y"

	address := []string{
		//"AHuTQ8J9cEBKkMdvpMPGjMt5aXNi7kt5gy",
		//"AQNKcZnjeB8Lgc9CkLC3tZyuvgBvRUEiec",
		//"ARppp4adUpoRW9dk7Vn2gGUasNSesb4dhr",
		//"ATAineEreLjeFVfnAzxL9XnKAsHY5PziMJ",
		//"AZ79g3MbL1BR95KYrzGT3WZriPRHzpsikb",
		//"Aa8NVJUW6tnbdoYYRmwYgV5TdFXhDvAJXA",
		"AbVCQPf4PYfkk3yWbAb5V1MXM8uYQ5Ndu6",
		"AKMarHenA3iAmWR5sjA3kVGPrt5rtNgcki",
		"AMP7RdWDQG7Za9xW49crvVwBcEFmcELzwg",
		"AP3yrFWvsor2zAjtjZRG1pkUCBi8oGho3e",
		"Ab8ajWDzw36x6ZTd1AqR39hKhAvijogaok",
		"AKvYaeap1JRPbwqc9u9GTpRPFt8yXU8rYW",
		"AaesSb9pgurkvNkNq4mckRsovTvyUGRZis",
		"AMD68p77gafMPa4q5wSgJwErezCKuviStj",
		"AXt7pouNtvdG4L5p2k3n7KGWNsbZCGNZVE",
		"AScugdDownPbJsfya4wnMUzxDZSDHzGr7H",
		"ASYEVyGNHPN1BHYrx3qA27kd4SGJR9iq6g",

	}

    for i:=0 ; i< 1; i++{

	to := address[i %len(address)]
	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.2", "", "", nil)
	if err != nil {
		return
	}

	log.Std.Info("rawTx: %+v", rawTx)

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}
	}
}

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMSVS51x4XgWNgd354yxd4yC7CrESsA8dN"
	//accountID := "4bpD665trhfF69MeZQUj5hpCqogfetEwWoGTrRMSuxvV"
	accountID := "GF3DqiUyNQi2RBqDC7KPJZ5SpARdy51Qn5tH4E5JBrgu"
	summaryAddress := "AQTdNMzKsvG6Qyw2KAtc7Wx5ZekVfgCM5G"

	//contract := openwallet.SmartContract{
	//	Address:  "eosio.token",
	//	Symbol:   "PIA",
	//	Name:     "PIA",
	//	Token:    "PIA",
	//	Decimals: 4,
	//}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTx := range rawTxArray {
		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}

}

