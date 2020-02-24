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

package ilcoin

import (
	"github.com/blocktree/openwallet/common"
	"github.com/blocktree/openwallet/log"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/shopspring/decimal"
	"net/url"
	"testing"
)

//http://192.168.32.107:20003/insight-api/

func TestGetBlockHeightByExplorer(t *testing.T) {
	height, err := tw.getBlockHeightByExplorer()
	if err != nil {
		t.Errorf("getBlockHeightByExplorer failed unexpected error: %v\n", err)
		return
	}
	t.Logf("getBlockHeightByExplorer height = %d \n", height)
}

func TestGetBlockHashByExplorer(t *testing.T) {
	hash, err := tw.getBlockHashByExplorer(343211)
	if err != nil {
		t.Errorf("getBlockHashByExplorer failed unexpected error: %v\n", err)
		return
	}
	t.Logf("getBlockHashByExplorer hash = %s \n", hash)
}

func TestGetBlockByExplorer(t *testing.T) {
	block, err := tw.getBlockByExplorer("0000000000000539a2055b519687e61f216a993b5115eb08e927bca8419c9e7a")
	if err != nil {
		t.Errorf("GetBlock failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetBlock = %v \n", block)
}

func TestListUnspentByExplorer(t *testing.T) {
	list, err := tw.listUnspentByExplorer(0, "1PeGx8FWC5eXmGEv8mpZ9xxgFtrdetg277")
	if err != nil {
		t.Errorf("listUnspentByExplorer failed unexpected error: %v\n", err)
		return
	}
	for i, unspent := range list {
		t.Logf("listUnspentByExplorer[%d] = %v \n", i, unspent)
	}

}

func TestGetTransactionByExplorer(t *testing.T) {
	raw, err := tw.getTransactionByExplorer("2732d60a1bac9e31c50f65a8a6d7e33529742788fae4043ba1a14512275c86ab")
	if err != nil {
		t.Errorf("getTransactionByExplorer failed unexpected error: %v\n", err)
		return
	}
	t.Logf("getTransactionByExplorer = %v \n", raw)
}

func TestGetBalanceByExplorer(t *testing.T) {
	raw, err := tw.getBalanceByExplorer("ms9NeTGFtaMcjrqRyRogkHqRoR8b1sQwu3")
	if err != nil {
		t.Errorf("getBalanceByExplorer failed unexpected error: %v\n", err)
		return
	}
	t.Logf("getBalanceByExplorer = %v \n", raw)
}

func TestGetMultiAddrTransactionsByExplorer(t *testing.T) {
	list, err := tw.getMultiAddrTransactionsByExplorer(0, 15, "2N7Mh6PLX39japSF76r2MAf7wT7WKU5TdpK")
	if err != nil {
		t.Errorf("getMultiAddrTransactionsByExplorer failed unexpected error: %v\n", err)
		return
	}
	for i, tx := range list {
		t.Logf("getMultiAddrTransactionsByExplorer[%d] = %v \n", i, tx)
	}

}


func TestSocketIO(t *testing.T) {

	var (
		room = "inv"
		endRunning = make(chan bool, 1)
	)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("192.168.32.107", 20003, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		return
	}

	err = c.On("tx", func(h *gosocketio.Channel, args interface{}) {
		log.Info("New transaction received: ", args)
		txMap, ok := args.(map[string]interface{})
		if ok {
			txid := txMap["txid"].(string)
			trx, _ := tw.GetTransaction(txid)
			log.Std.Info("trx: %+v", trx)
		}
	})
	if err != nil {
		return
	}

	err = c.On("block", func(h *gosocketio.Channel, args interface{}) {
		log.Info("New block received: ", args)
	})
	if err != nil {
		return
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Info("Disconnected")
	})
	if err != nil {
		return
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Info("Connected")
		h.Emit("subscribe", room)
	})
	if err != nil {
		return
	}

	<- endRunning
}

func TestEstimateFeeRateByExplorer(t *testing.T) {
	feeRate, _ := tw.estimateFeeRateByExplorer()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.String())
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.String())
}

func TestURLParse(t *testing.T) {
	apiUrl, err := url.Parse("http://192.168.32.107:20003/insight-api/")
	if err != nil {
		t.Errorf("url.Parse failed unexpected error: %v\n", err)
		return
	}
	domain := apiUrl.Hostname()
	port := common.NewString(apiUrl.Port()).Int()
	t.Logf("%s : %d", domain, port)
}

func TestDecimalAdd(t *testing.T) {
	unconfirmBalance, _ := decimal.NewFromString("-5.1")
	confirmBalance, _ := decimal.NewFromString("5.1")
	balance := confirmBalance.Add(unconfirmBalance)
	log.Infof("balance = %s", balance.String())
}