# vcn - vChain CodeNotary

Global, de-centralized signing of code and other digital assets

## License

This software is released under [GPL3](https://www.gnu.org/licenses/gpl-3.0.en.html).

## Installation

It's easiest to download the lates release from GitHub:

https://github.com/vchain-us/vcn/releases

## Installation from Source

After having installed [golang](https://golang.org/doc/install) clone this 
repository into your `GOPATH`, usually this is `$HOME/go/src/` on Unix-like
operating systems.

### PATH

Set the `GOPATH`

```
$> export GOPATH=$HOME/go
$> export GOBIN=$GOPATH/bin
$> PATH=$PATH:$GOPATH:$GOBIN
$> export PATH
```

or simply put it to `$HOME/.bash_profile` once.

### System-wide installation

This will put the executable into `GOBIN` which is
accessible throughout the system.

```
$> cd vcn/vcn
$> go install
$> vcn
```

Alternatively you can build `vcn` in the working directory.

```
$> go get ./...
$> go build
$> ./vcn
```

## Usage

Register an account with vChain.us first.

Then start with the `login` verb; the cli will walk you through login
and setting up your local keystore upon initial use.
```
$> vcn --help
$> vcn login
```

You're good to use `verify` without the above registration.

```
$> vcn verify <file>
$> vcn verify docker:<imageId>
```

Once your public key is known on the blockchain you can sign assets:

```
$> vcn sign <file>
$> vcn sign docker:<image>
```

By default all assets are signed private, so not much information is disclosed about the signer. If you want to make it public and therefore, more trusted, please use the -public switch.

```
$> vcn sign -public <file>
$> vcn sign -public docker:<image>
```

Change the asset's status

```
$> vcn unsupport <file>
$> vcn untrust <file>
```

Have a look at analytics and extended functionality on the dashboard (browser needed):

```
$> vcn dashboard
```

Fetch all assets you've signed:

```
$> vcn list
```

### Advanced usage 

You're good to start doing really cool things, e.g.

```
# run a Docker image only when it can be successfully verified
$> vcn verify docker:hello-world && docker run hello-world
```

```
# verify multiple files by piping other commands' outputs into vcn
$> ls | xargs vcn verify
```

```
# work with environment
# get logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC)
$> LOG_LEVEL=TRACE vcn login

# or with a proxy
$> HTTP_PROXY=http://localhost:3128 vcn verify <file>
```

## Development

### Test Automation
Simply run

```
$> go test
```

### Smart contracts

```
$ solc --abi Proof.sol -o build
$ abigen --abi=./Proof.abi --pkg proof --out=Proof.go
```
Attention: depends on where your *.sol and *.abi are located. The resulting
autogenerated go sources must go into `vcn/vcn/`.

## Distribution

### Cross-compiling for various platforms

The C libraries of go-etehreum make a more sophisticated cross-compilation
necessary. Make sure you have `xgo` installed:

```
go get -u github.com/karalabe/xgo
```

`./build-multiplatform.sh` in `dist` takes care of the rest.

