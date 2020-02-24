# ilcoin-adapter

ilcoin-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 项目依赖库

- [go-owcrypt](https://github.com/blocktree/go-owcrypt.git)
- [go-owcdrivers](https://github.com/blocktree/.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建ILC.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 1
# node api url, if RPC Server Type = 0, use ilcoin core full node
;serverAPI = "http://127.0.0.1:8333/"
# node api url, if RPC Server Type = 1, use bitbay insight-api
serverAPI = "https://ilcoinexplorer.com/api/"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = false
# support omnicore
omniSupport = false
# Omni Core RPC API
omniCoreAPI = "http://ip:port"
# Omni Core RPC Authentication Username
omniRPCUser = "user"
# Omni Core RPC Authentication Password
omniRPCPassword = "password"
# Omni token transfer minimum cost
omniTransferCost = "0.00000546"
# support segWit
supportSegWit = false
# minimum transaction fees
minFees = "0.00001"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```
