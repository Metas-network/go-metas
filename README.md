# Metas 

EVM-compatible chain secured by the Lachesis consensus algorithm.

## Building the source

Building `metas` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make metas
```
The build output is ```build/metas``` executable.

## Running `metas`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `metas` instance.

### Launching a network

Launching `metas` for a network:

```shell
$ metas --genesis /path/to/genesis.g
```

### Configuration

As an alternative to passing the numerous flags to the `metas` binary, you can also pass a
configuration file via:

```shell
$ metas --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ metas --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `metas validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ metas --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`metas` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ metas --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ metas --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ metas --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ metas --datadir /path/to/datadir account new # create new account
$ metas --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Lachesis has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/Metas-network/go-metas/app	0.001s
?   	github.com/Metas-network/go-metas/cmd/cmdtest	[no test files]
ok  	github.com/Metas-network/go-metas/cmd/metas	13.890s
?   	github.com/Metas-network/go-metas/cmd/metas/metrics	[no test files]
?   	github.com/Metas-network/go-metas/cmd/metas/tracing	[no test files]
?   	github.com/Metas-network/go-metas/crypto	[no test files]
?   	github.com/Metas-network/go-metas/debug	[no test files]
?   	github.com/Metas-network/go-metas/ethapi	[no test files]
?   	github.com/Metas-network/go-metas/eventcheck	[no test files]
?   	github.com/Metas-network/go-metas/eventcheck/basiccheck	[no test files]
?   	github.com/Metas-network/go-metas/eventcheck/gaspowercheck	[no test files]
?   	github.com/Metas-network/go-metas/eventcheck/heavycheck	[no test files]
?   	github.com/Metas-network/go-metas/eventcheck/parentscheck	[no test files]
ok  	github.com/Metas-network/go-metas/evmcore	2.311s
?   	github.com/Metas-network/go-metas/gossip	[no test files]
?   	github.com/Metas-network/go-metas/gossip/emitter	[no test files]
ok  	github.com/Metas-network/go-metas/gossip/filters	0.250s
?   	github.com/Metas-network/go-metas/gossip/gasprice	[no test files]
?   	github.com/Metas-network/go-metas/gossip/occuredtxs	[no test files]
?   	github.com/Metas-network/go-metas/gossip/piecefunc	[no test files]
ok  	github.com/Metas-network/go-metas/integration	10.640s
```

Also it is tested with [fuzzing](./FUZZING.md).


### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ metas --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ metas --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ metas --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ metas --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the Metas processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
