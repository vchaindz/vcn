## Generating smart contracts on linux

Clone https://github.com/ethereum/go-ethereum and compile `abigen` command in ./cmd/abigen

Download solc at https://github.com/ethereum/solidity/releases?after=v-1.5.0 and chmod +x on it and copy in /usr/bin
solc version must be:
```bash
solc --version
solc, the solidity compiler commandline interface
Version: -1.4.24+commit.e67f0147.Linux.g++
```

Place at the root of the contracts repo and generate go files with:
```
abigen --sol organisations/contracts/OrganisationsRelay.sol --pkg blockchain --out organisationsrelay.go
abigen --sol assets/contracts/AssetsRelay.sol --pkg blockchain --out assetsrelay.go
```

## Cross-compiling for various platforms

The C libraries of [go-ethereum](https://github.com/ethereum/go-ethereum) make a more sophisticated cross-compilation
necessary.
The `make dist` target takes care of all steps by using [xgo](https://github.com/techknowlogick/xgo) and [docker](https://github.com/docker).
