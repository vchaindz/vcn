# vcn - vChain CodeNotary [![CircleCI](https://circleci.com/gh/vchain-us/vcn.svg?style=svg)](https://circleci.com/gh/vchain-us/vcn)
> code signing in 1 simple step

## How it works
![vcn How it works](https://raw.githubusercontent.com/vchain-us/vcn/master/docs/vcn_hiwb.png "How it works")

## Installation

It's easiest to download the latest version from the [relase page](
https://github.com/vchain-us/vcn/releases).

### Installation from Source

After having installed [golang](https://golang.org/doc/install) 1.12 or newer clone this 
repository into your working directory.

### Build locally

You can build `vcn` in the working directory using the provided `Makefile`.

```
make vcn
```

Then run
```
./vcn
```

### System-wide installation

This will put the `vcn` executable into `GOBIN` which is
accessible throughout the system.

```
make install
```

## Usage

Detailed **commands usage** can be found [here](docs/cmd/vcn.md).

Furthermore, check out our list of **integrations**:

* [Docker](docs/DOCKERINTEGRATION.md)

### Basic usage

Register an account with [codernotary.io](https://codenotary.io) first.

Then start with the `login` command. `vcn` will walk you through login and setting up your local keystore upon initial use.
```
vcn --help
vcn login
```

You're good to use `verify` without the above registration.

```
vcn verify <asset>
vcn verify docker://<imageId>
```

Output results in `json` or `yaml` format:
```
vcn verify --output=json <asset>
vcn verify --output=yaml <asset>
```

Once your public key is known on the blockchain you can sign assets:

```
vcn sign <asset>
vcn sign docker://<image>
```
> By default all assets are signed private, so not much information is disclosed about the signer. If you want to make it public and therefore, more trusted, please use the `--public` flag.

```
vcn sign --public <asset>
vcn sign --public docker://<image>
```

Change the asset's status

```
vcn unsupport <asset>
vcn untrust <asset>
```

Fetch all assets you've signed:

```
vcn list
```

Have a look at analytics and extended functionality on the dashboard (browser needed):

```
vcn dashboard
```

### Advanced usage 

You're good to start doing really cool things, e.g.

```
# run a Docker image only when it can be successfully verified
vcn verify docker://hello-world && docker run hello-world
```

```
# by adding `--key` you can verify that your asset has been signed by a specific public key 
vcn verify --key 0x8f2d1422aed72df1dba90cf9a924f2f3eb3ccd87 docker://hello-world
```
`vcn verify` also:
- accept multiple keys by multiple by flag, usage `--key 0x0...1 --key 0x0...2` or comma separated `--key 0x0...1,0x0...2`
- accept multiple keys by env var `VCN_VERIFY_KEYS`, usage: space separated `VCN_VERIFY_KEYS=0x0...1 0x0...2`
- `--key` takes precedence over  `VCN_VERIFY_KEYS`

```
# verify multiple assets by piping other commands' outputs into vcn
ls | xargs vcn verify
```

```
# work with environment
# get logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC)
LOG_LEVEL=TRACE vcn login

# or with a proxy
HTTP_PROXY=http://localhost:3128 vcn verify <asset>
```

## Testing
```
make test
```

## Cross-compiling for various platforms

The C libraries of [go-ethereum](https://github.com/ethereum/go-ethereum) make a more sophisticated cross-compilation
necessary. 
The `make dist` target takes care of all steps by using [xgo](https://github.com/techknowlogick/xgo) and [docker](https://github.com/docker). 

## License

This software is released under [GPL3](https://www.gnu.org/licenses/gpl-3.0.en.html).
