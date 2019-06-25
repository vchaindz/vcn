# CHANGELOG
All notable changes to this project will be documented in this file. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
<a name="unreleased"></a>
## [Unreleased]

### Bug Fixes
- ca-certificates for Dockerfile
- clean other context when user logs in
- **cmd:** show config file path if not default one
- **docker:** accept images but not other objects
- **extractor:** correct empty files handling

### Changes
- reduce public api surface
- code cleanup and fix comments
- bump version
- **api:** move type Error to its own source file
- **api:** deprecate public publisher APIs
- **api:** deprecate publisher fields
- **api:** allow empty size
- **cmd:** internal types
- **env:** remove VERIFY_ prefix
- **internal:** move cli package
- **output:** silence printing when formatted output
- **verify:** improve error labels

### Features
- vcn user agent
- promote --output to global flag and return formatted errors accordingly
- json and yaml output for sign
- **api:** add blockchain organisation support
- **list:** support for --output json/yaml
- **sign:** signing by --hash
- **verify:** add --org for verify


<a name="0.4.3"></a>
## [0.4.3] - 2019-05-21
### Bug Fixes
- **api:** nil pointer dereference
- **extractor:** nil pointer dereference
- **sign:** flag accessed but not defined: hash
- **sign:** cross platform loading spinner

### Changes
- minor refactoring of login/logout
- generate markdown docs for commands
- **api:** load user's artifact by hash only
- **api:** size must be valid, name can be empty
- **sign:** prompt login if needed when started by explorer

### Features
- profile dir per stage
- add test env
- **sign:** untrust/unsupport by --hash flag


<a name="0.4.2"></a>
## [0.4.2] - 2019-05-14
### Bug Fixes
- print correct config file name
- do not read in env var for config
- **sign:** remove ownership disclaimer
- **verify:** cross platform coloured printing
- **verify:** only print size when available

### Changes
- disable config file message
- reorg publisher fields
- minor printing and marshaling improvements
- **api:** fine-tune blockchain verify funcs
- **build:** make makefile more resilient
- **log:** improve metahash tracing
- **output:** correct WriteTo interface

### Code Refactoring
- **sign:** printing

### Features
- improve printing and add yaml output format
- static build and docker cli support in dockerfile
- **verify:** with multiple key by VCN_VERIFY_KEYS env var
- **verify:** allow multiple --key flags
- **verify:** show asset visibility


<a name="0.4.1"></a>
## [0.4.1] - 2019-05-08
### Bug Fixes
- correct err msg when token has expired
- do not print verification error message if --output=json
- verify fallback
- clear context when logging out
- if logged in always show user own signature
- remove test-resources from dockerignore
- return proper error by checking quota before

### Changes
- executable file sniffer
- go mod tidy
- change metadata naming
- add Platform and Architecture fields
- **build:** minor improvements
- **ci:** initial circleci setup
- **cmd:** use extractors
- **meta:** increse signing timeout
- **meta:** increase tx verification rounds

### Code Refactoring
- asset and metadata fields naming

### Features
- user defined attributes for assets
- json output format (--output=json)
- infer asset version
- get metadata from docker image
- add --hash for verify
- improve mime and metadata extractor for executables
- new kind and mimeType attributes for assets
- modularize metadata extractors
- URI parser for assets
- subscription limit enforcement
- **build:** improved build system


<a name="0.4.0"></a>
## [0.4.0] - 2019-04-30
### Bug Fixes
- dashboard cmd has no args
- prefix for docker assets
- find home dir in the right way
- do not exit early when login is required
- do not quit when executed by Win context menu
- ensure default keystore before looking for keys
- remove log.Fatal in favour of returing errors
- gitignore
- ask for confirmation before quitting
- token deletion when already logged out
- improve err msgs when auth is required
- add missing header comments
- temporary dirs creation at startup
- **login:** silence usage message when errored
- **migrate:** do not create a new key if it already exists
- **store:** correct key funcs and tests
- **tracking:** send events at right place in time

### Changes
- use two-stage build
- make createArtifact private
- printing stuff
- no race test for make install
- add dev stage
- move cli package to internal
- improve naming
- code cleanup and minor improvements
- add store context and key utils
- introduce store package
- move logs to internal
- improve wallet not synced message
- correct printing func
- update gitignore
- switch main to cobra
- partial api logging refactor
- code clean up
- introducing Cobra
- move file hashing func to internal
- use make install for Dockerfile
- **build:** improve makefile
- **cli:** spinner changed and code cleaned up
- **dist:** omit symbol table and debug info
- **internal:** correct errors funcs and tests

### Code Refactoring
- login cmd
- list cmd
- untrust and unsupport cmds
- sign, dashboard cmds and APIs
- docker compose
- package oriented design

### Features
- key flag for sign and verify
- profile migration from v0.3.x
- config file and multi-keys support
- logout command
- upgrade to latest assetsrelay sc
- provide Dockerfile
- **build:** build system with Makefile
- **cli:** improved column printing

### BREAKING CHANGE

this commit changes the usage of verify and sign methods

this commit introduce the config file with multi-keys support, and a huge refactoring of vcn code.


<a name="0.3.6"></a>
## [0.3.6] - 2019-04-08

<a name="0.3.5"></a>
## [0.3.5] - 2019-03-28

<a name="0.3.4"></a>
## [0.3.4] - 2019-03-20

<a name="0.3.3"></a>
## [0.3.3] - 2019-03-13

<a name="0.3.2"></a>
## [0.3.2] - 2019-03-11

<a name="0.3.1"></a>
## [0.3.1] - 2019-03-11

<a name="0.3.0"></a>
## [0.3.0] - 2019-03-08

<a name="0.2.2"></a>
## [0.2.2] - 2019-03-07

<a name="v.0.2-beta.0"></a>
## [v.0.2-beta.0] - 2019-02-25

<a name="v.0.1-beta.2"></a>
## v.0.1-beta.2 - 2019-02-19

[Unreleased]: https://github.com/vchain-us/vcn/compare/0.4.3...HEAD
[0.4.3]: https://github.com/vchain-us/vcn/compare/0.4.2...0.4.3
[0.4.2]: https://github.com/vchain-us/vcn/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/vchain-us/vcn/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/vchain-us/vcn/compare/0.3.6...0.4.0
[0.3.6]: https://github.com/vchain-us/vcn/compare/0.3.5...0.3.6
[0.3.5]: https://github.com/vchain-us/vcn/compare/0.3.4...0.3.5
[0.3.4]: https://github.com/vchain-us/vcn/compare/0.3.3...0.3.4
[0.3.3]: https://github.com/vchain-us/vcn/compare/0.3.2...0.3.3
[0.3.2]: https://github.com/vchain-us/vcn/compare/0.3.1...0.3.2
[0.3.1]: https://github.com/vchain-us/vcn/compare/0.3.0...0.3.1
[0.3.0]: https://github.com/vchain-us/vcn/compare/0.2.2...0.3.0
[0.2.2]: https://github.com/vchain-us/vcn/compare/v.0.2-beta.0...0.2.2
[v.0.2-beta.0]: https://github.com/vchain-us/vcn/compare/v.0.1-beta.2...v.0.2-beta.0