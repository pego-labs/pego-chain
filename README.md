## PEGO Chain

PEGO Chain is a Web3 infrastructure that provides high autonomy, scalability, and
sustainability for decentralized applications. PEGO has designed a fully
community-driven autonomous Web3 infrastructure, where all on-chain parameterscan be flexibly adjusted through proposals initiated by the community. Accordingtothe PPOS consensus, the established composite economic incentive mechanismcanpromote the sustainable development of the PEGO ecosystem. Additionally, thePEGO blockchain is fully compatible with the EVM system, making it an ideal choicefor developers aiming to build scalable decentralized applications.

More details in [White Paper](https://pego.network/pego/whitepaper_en.pdf).

PEGO Chain starts its development based on go-ethereum and BSC fork. So you may see many toolings, binaries and also docs are based on Ethereum ones, such as the name “geth”.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)

## Key features

### Proof of Staked Authority

Although Proof-of-Work (PoW) has been approved as a practical mechanism to implement a decentralized network, it is not friendly to the environment and also requires a large size of participants to maintain the security.

Proof-of-Authority(PoA) provides some defense to 51% attack, with improved efficiency and tolerance to certain levels of Byzantine players (malicious or hacked).
Meanwhile, the PoA protocol is most criticized for being not as decentralized as PoW, as the validators, i.e. the nodes that take turns to produce blocks, have all the authorities and are prone to corruption and security attacks.

Other blockchains, such as EOS and Cosmos both, introduce different types of Deputy Proof of Stake (DPoS) to allow the token holders to vote and elect the validator set. It increases the decentralization and favors community governance.

To combine DPoS and PoA for consensus, Pego Chain implement a novel consensus engine called Parlia that:

1. Blocks are produced by a limited set of validators.
2. Validators take turns to produce blocks in a PoA manner, similar to Ethereum's Clique consensus engine.
3. Validator set are elected in and out based on a staking based governance on PEGO Chain.
4. The validator set change is relayed via a cross-chain communication mechanism.
5. Parlia consensus engine will interact with a set of system contracts to achieve liveness slash, revenue distributing and validator set renewing func.

## Native Token

**PG** will run on PEGO Chain in the same way as ETH runs on Ethereum so that it remains as `native token` for PEGO. This means,
PG will be used to:

1. pay `gas` to deploy or invoke Smart Contract on PEGO Chain.
2. perform cross-chain operations, such as transfer token assets across PEGO Chain.

## Building the source

Many of the below are the same as or similar to go-ethereum.

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth).

Building `geth` requires both a Go (version 1.19 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The pego project comes with several wrappers/executables found in the `cmd`
directory.

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| :--------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`geth`** | Main PEGO Chain client binary. It is the entry point into the PEGO network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It has the same and more RPC and other interface as go-ethereum and can be used by other processes as a gateway into the PEGO network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI page](https://geth.ethereum.org/docs/interface/command-line-options) for command line options. |
|   `clef`   | Stand-alone signing tool, which can be used as a backend signer for `geth`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
|  `devp2p`  | Utilities to interact with nodes on the networking layer, without running a full blockchain.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|  `abigen`  | Source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://geth.ethereum.org/docs/dapp/native-bindings) page for details.                                                                                               |
| `bootnode` | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                                                                                                                                                                            |
|   `evm`    | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                                                                                            |
| `rlpdump`  | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://eth.wiki/en/fundamentals/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                                                                                                                 |

## Running `geth`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://geth.ethereum.org/docs/fundamentals/command-line-options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance.

### Hardware Requirements

The hardware must meet certain requirements to run a full node on mainnet:
- VPS running recent versions of Mac OS X, Linux, or Windows.
- IMPORTANT! 300 GB of free disk space, solid-state drive(SSD), gp3, 8k IOPS, 250 MB/S throughput, read latency <1ms. (if node is started with snap sync, it will need NVMe SSD)
- 8 cores of CPU and 32 GB of memory (RAM)
- Suggest m6i.2xlarge instance type on AWS, c2-standard-8 on Google cloud.
- A broadband Internet connection with upload/download speeds of 5 MB/S


### Steps to Run a Full Node / Validator Node

#### 1. Build geth Client
```shell
git clone https://github.com/pego-labs/pego-chain.git
cd pego-chain
make geth  ## build geth client
cp ./build/bin/geth /usr/local/bin/   ## copy geth to your local bin path 
```

#### 2. Setup Node Directory 
```shell
mkdir -p ~/pego_node/data  ## create node directory
cd ~/pego_node
cp ~/pego-chain/config/* .  ## copy config file to node dir
```

#### 3. Initialize Node Directory
```shell
geth --datadir ./data init genesis.json  ## init node with genesis block
```

#### 4. Start the node

**For full node**
```shell
geth --config ./config.toml --datadir ./data  
```

**Or For validator node**

Prepare your wallet address and keystore file to be used as validator. Put your keystore file to *pego_node/data/keystore/*

```shell
## Start validator node
geth  --config ./config.toml --datadir ./data  --syncmode full -unlock <your_wallet_address> --password <password_file_for_keystore> --mine --allow-insecure-unlock  --port 30311 --verbosity 5 --cache 18000
```

#### 5. Monitor node status

Monitor the log from **./data/pego.log** by default. When the node has started syncing, should be able to see the following output:
```shell
t=2023-05-30T17:23:22+0000 lvl=info msg="Imported new chain segment"             blocks=1    txs=0 mgas=0.000 elapsed="497.765µs" mgasps=0.000  number=37376 hash=0x4ec4457886dfec05c212da9bd78705de1c8fe36315278c135e5547917632c377 dirty="56.94 KiB"
t=2023-05-30T17:23:27+0000 lvl=info msg="Imported new chain segment"             blocks=1    txs=0 mgas=0.000 elapsed="557.456µs" mgasps=0.000  number=37377 hash=0xb90903cd6e46e293c5936937bde947eb6ff0353ef56f879d1878c0a317d3476b dirty="56.94 KiB"
```

#### 6. Interact with fullnode
Start up `geth`'s built-in interactive [JavaScript console](https://geth.ethereum.org/docs/interface/javascript-console),
(via the trailing `console` subcommand) through which you can interact using [`web3` methods](https://web3js.readthedocs.io/en/)
(note: the `web3` version bundled within `geth` is very old, and not up to date with official docs),
as well as `geth`'s own [management APIs](https://geth.ethereum.org/docs/rpc/server).
This tool is optional and if you leave it out you can always attach to an already running
`geth` instance with `geth attach`.

#### 7. More

More details about **running a node**(coming soon) and **becoming a validator**(coming soon)

*Note: Although there are some internal protective measures to prevent transactions from
crossing over between the main network and test network, you should make sure to always
use separate accounts for play-money and real-money. Unless you manually move
accounts, `geth` will by default correctly separate the two networks and will not make any
accounts available between them.*

### Configuration

As an alternative to passing the numerous flags to the `geth` binary, you can also pass a
configuration file via:

```shell
$ geth --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ geth --your-favourite-flags dumpconfig
```

### Programmatically interfacing `geth` nodes

As a developer, sooner rather than later you'll want to start interacting with `geth` and the
PEGO network via your own programs and not manually through the console. To aid
this, `geth` has built-in support for a JSON-RPC based APIs ([standard APIs](https://ethereum.github.io/execution-apis/api-documentation/)
and [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `geth`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

* `--http` Enable the HTTP-RPC server
* `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
* `--http.port` HTTP-RPC server listening port (default: `8545`)
* `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
* `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
* `--ws` Enable the WS-RPC server
* `--ws.addr` WS-RPC server listening interface (default: `localhost`)
* `--ws.port` WS-RPC server listening port (default: `8546`)
* `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
* `--ws.origins` Origins from which to accept WebSocket requests
* `--ipcdisable` Disable the IPC-RPC server
* `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,txpool,web3`)
* `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `geth` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
PEGO nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to pego, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. If you wish to submit
more complex changes though, please check up with the core devs first on [our telegram channel](https://t.me/pegoofficial)
to ensure those changes are in line with the general philosophy of the project and/or get
some early feedback which can make both your efforts much lighter as well as our review
and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

* Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
  guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
* Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
  guidelines.
* Pull requests need to be based on and opened against the `master` branch.
* Commit messages should be prefixed with the package(s) they modify.
    * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://geth.ethereum.org/docs/developers)
for more details on configuring your environment, managing project dependencies, and
testing procedures.

## License

The PEGO library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The PEGO binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
