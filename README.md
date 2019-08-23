# vcn - vChain CodeNotary [![CircleCI](https://circleci.com/gh/vchain-us/vcn.svg?style=svg)](https://circleci.com/gh/vchain-us/vcn)


# The Trust and Integrity platform for the Cloud native environment
Give any digital asset a meaningful, globally-unique, immutable identity that is authentic, verifiable, traceable from anywhere. 

When using CodeNotary vcn in source code, release, deployment or at runtime, you allow a continuous trust verification that can be used to detect unusual or unwanted activity in your workload and act on it. 
Powered by CodeNotary's digital identity infrastructure, vcn lets you notarize all of your digital assets that add a trust level of your choice, custom attributes and a meaningful status without touching or appending anything (unlike digital certificates).
That allows change and revocation post-release without breaking any customer environment.

Everything is done in a global, collaborative way to break the common silo solution architecture. Leveraging an immutable, always-on DLT platform allows you to avoid complex setup of Certificate authorities or digital certificates (that are unfit for DevOps anyway). 

# DevSecOps in mind
Codenotary vcn is a solution written by a devops-obsessed engineers for Devops engineers to bring better trust and security to the the CloudNative source to deployment process 

# What kind of behaviors can CodeNotary vcn detect?
vcn (and its extensions for Docker, Kubernetes, documents or CI/CD) can detect, authenticate and alert on any behavior that involves using unauthentic digital assets. vcn verification can be embedded anywhere and can be used to trigger alerts, updates or workflows.

vcn is so versatile, it can help detecting or acting on the following (but not limited to):
* Enable application version checks and actions
* Buggy or rogue libraries can be traced by simple revoke or unsupport
* Revoke or unsupport your build or build version post-deployment (no complex certificate revocation that includes delivery of newly signed builds)
* Stop unwanted containers from being launched
* Make revocation part of the remediation process
* Use revocation without impairing customer environments
* Trace source code to build to deployment by integration into CI/CD or manual workflow
* Tag your applications for specific use cases (alpha, beta - non-commercial aso).

## How it works
![vcn How it works](https://raw.githubusercontent.com/vchain-us/vcn/master/docs/vcn_hiwb.png "How it works")

## Installation

### Binary (Cross-platform)

It's easiest to download the latest version for your platform from the [release page](
https://github.com/vchain-us/vcn/releases).

Once downloaded, you can rename the binary to `vcn`, then run it from anywhere.
> For Linux and macOS you need to mark the file as executable: `chmod +x vcn`

### Homebrew / Linuxbrew

If you are on macOS and using [Homebrew](https://brew.sh/) (or on Linux and using [Linuxbrew](https://linuxbrew.sh/)), you can install `vcn` with the following:

```
brew tap vchain-us/brew
brew install vcn
```

### From Source

After having installed [golang](https://golang.org/doc/install) 1.12 or newer clone this 
repository into your working directory.

#### Build locally

You can build `vcn` in the working directory using the provided `Makefile`.

```
make vcn
```

Then run
```
./vcn
```

#### System-wide

This will put the `vcn` executable into `GOBIN` which is
accessible throughout the system.

```
make install
```

## Usage

For detailed **command line usage** see [docs/cmd/vcn.md](docs/cmd/vcn.md) or just run `vcn help`.


Furthermore, check out our list of **integrations**:

* [docker](docs/DOCKERINTEGRATION.md) - Out of the box support for signing and verify Docker images.
* [hub.docker.com/r/codenotary/vcn](https://hub.docker.com/r/codenotary/vcn) - The `vcn`'s DockerHub repository. 
* [kube-notary](https://github.com/vchain-us/kube-notary) - A Kubernetes watchdog for verifying image trust with CodeNotary.
* [vcn-watchdog](https://github.com/vchain-us/vcn-watchdog) - Continuous verification with CodeNotary for Docker.
* [jsvcn](https://github.com/vchain-us/jsvcn) - CodeNotary JavaScript Client.
* [jvcn](https://github.com/vchain-us/jvcn) - CodeNotary Java Bindings.
* [jvcn-maven-plugin](https://github.com/vchain-us/jvcn-maven-plugin) - Maven dependency verification and enforcement.

### Basic usage

Register an account with [codernotary.io](https://codenotary.io) first.

Then start with the `login` command. `vcn` will walk you through login and setting up your local keystore upon initial use.
```
vcn login
```

You're good to use `verify` without the above registration.

```
vcn verify <file>
vcn verify dir://<directory>
vcn verify docker://<imageId>
vcn verify podman://<imageId>
vcn verify --hash <hash>
```

Output results in `json` or `yaml` formats:
```
vcn verify --output=json <asset>
vcn verify --output=yaml <asset>
```
> Check out the [user guide](docs/user-guide/formatted-output.md) for further details.

Once your public key is known on the blockchain you can sign assets:

```
vcn sign <file>
vcn sign dir://<directory>
vcn sign docker://<imageId>
vcn sign podman://<imageId>
vcn sign --hash <hash>
```

By default all assets are signed private, so not much information is disclosed about the signer. If you want to make it public and therefore, more trusted, please use the `--public` flag.

```
vcn sign --public <asset>
```

Change the asset's status:

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

### Examples

#### Verify a Docker image automatically prior to running it

First, you’ll need to pull the image by using: 

```
docker pull hello-world
```

Then use the below command to put in place an automatic safety check. It allows only verified images to run. 

```
vcn verify docker://hello-world && docker run hello-world
```
If an image was not verified, it will not run and nothing will execute. 


#### Verify multiple assets
You can verify multiple assets by piping other command outputs into `vcn`:
```
ls | xargs vcn verify
```
> The exit code will be `0` only if all the assets in you other command outputs are verified.

#### Verify by a specific signer
By adding `--key`, you can verify that your asset has been signed by a specific signer’s public key.

```
vcn verify --key 0x8f2d1422aed72df1dba90cf9a924f2f3eb3ccd87 docker://hello-world
```

#### Verify by a list of signers

If an asset you or your organization wants to trust needs to be verified against a list of signers as a prerequisite, then use the `vcn verify` command and the following syntax:

- Add a `--key` flag in front of each key you want to add  
(eg. `--key 0x0...1 --key 0x0...2`)
- Or set the env var `VCN_KEY` correctly by using a space to separate each key (eg. `VCN_KEY=0x0...1 0x0...2`)
> Be aware that using the `--key` flag will take precedence over `VCN_KEY`.

The asset verification will succeed only if the asset has been signed by at least one of the signers.

#### Verify using the asset's hash

If you want to verify an asset using only its hash, you can do so by using the command as shown below:

```
vcn verify --hash fce289e99eb9bca977dae136fbe2a82b6b7d4c372474c9235adc1741675f587e
```

#### Unsupport/untrust an asset you do not have anymore

In case you want to unsupport/untrust an asset of yours that you no longer have, you can do so using the asset hash(es) with the following steps below.

First, you’ll need to get the hash of the asset from your CodeNotary [dashboard](https://dashboard.codenotary.io/) or alternatively you can use the `vcn list` command. Then, in the CLI, use:

```
vcn untrust --hash <asset's hash>
# or 
vcn unsupport --hash <asset's hash>
```

#### Signing within automated environments

First, you’ll need to make `vcn` have access to the `${HOME}/.vcn` folder that holds your private keys.
Then, set up your environment accordingly using the following commands:
```
export VCN_USER=<email>
export VCN_PASSWORD=<password>
export KEYSTORE_PASSWORD=<passphrase>
```

Once done, you can use `vcn` in your non-interactive environment using:

```
vcn login
vcn sign --key <your key> <asset>
```
> Other commands like `untrust` and `unsupport` will also work.


#### Working with Docker and Kubernetes

Check out our integrations:

* [Docker](docs/DOCKERINTEGRATION.md)
* [vcn-watchdog](https://github.com/vchain-us/vcn-watchdog)
* [vcn-k8s](https://github.com/vchain-us/vcn-k8s)


## Configuration
See [docs/user-guide/configuration.md](docs/user-guide/configuration.md).

## Environments
See [docs/user-guide/environments.md](docs/user-guide/environments.md).

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
