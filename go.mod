module github.com/assetsadapterstore/ilcoin-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.11.1
	github.com/blocktree/go-owcdrivers v1.1.16
	github.com/blocktree/go-owcrypt v1.0.3
	github.com/blocktree/openwallet v1.5.5
	github.com/bndr/gotabulate v1.1.2
	github.com/btcsuite/btcd v0.0.0-20190315201642-aa6e0f35703c
	github.com/btcsuite/btcutil v0.0.0-20190316010144-3ac1210f4b38
	github.com/codeskyblue/go-sh v0.0.0-20190328095946-f4ce45e7999e
	github.com/ethereum/go-ethereum v1.8.25
	github.com/graarh/golang-socketio v0.0.0-20170510162725-2c44953b9b5f
	github.com/imroc/req v0.2.3
	github.com/pborman/uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/tidwall/gjson v1.2.1
)

//replace github.com/blocktree/openwallet => ../../openwallet
//replace github.com/blocktree/go-owcdrivers => ../../go-owcdrivers
