# vcn - vChain CodeNotary
> Global, de-centralized signing of code and other digital assets

## License

This software is released under [GPL3](https://www.gnu.org/licenses/gpl-3.0.en.html).

## Installation

It's easiest to download the lates release from GitHub:

https://github.com/vchain-us/vcn/releases

## Installation from Source

After having installed [golang](https://golang.org/doc/install) 1.12 or newer clone this 
repository into your working directory.

### Build locally

You can build `vcn` in the working directory using the provided `Makefile`.

```
$> make vcn
```

Then run
```
$> ./vcn
```

### System-wide installation

This will put the `vcn` executable into `GOBIN` which is
accessible throughout the system.

```
$> make install
```

## Usage

Register an account with [codernotary.io](https://codenotary.io) first.

Then start with the `login` verb; the cli will walk you through login
and setting up your local keystore upon initial use.
```
$> vcn --help
$> vcn login
```

You're good to use `verify` without the above registration.

```
$> vcn verify <asset>
$> vcn verify docker:<imageId>
```

Once your public key is known on the blockchain you can sign assets:

```
$> vcn sign <asset>
$> vcn sign docker:<image>
```

By default all assets are signed private, so not much information is disclosed about the signer. If you want to make it public and therefore, more trusted, please use the --public switch.

```
$> vcn sign --public <asset>
$> vcn sign --public docker:<image>
```

Change the asset's status

```
$> vcn unsupport <asset>
$> vcn untrust <asset>
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
# verify multiple assets by piping other commands' outputs into vcn
$> ls | xargs vcn verify
```

```
# work with environment
# get logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC)
$> LOG_LEVEL=TRACE vcn login

# or with a proxy
$> HTTP_PROXY=http://localhost:3128 vcn verify <asset>
```

## Development

### Test Automation
Simply run

```
$> make test
```

## Distribution

### Cross-compiling for various platforms

The C libraries of [go-ethereum](https://github.com/ethereum/go-ethereum) make a more sophisticated cross-compilation
necessary. 
The `make dist` target takes care of all steps by using [xgo](https://github.com/techknowlogick/xgo) and `docker`. 
