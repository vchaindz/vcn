# CodeNotary Timestamp Service https://timestamp.codenotary.com jumpstart

## Table of contents

- [CodeNotary Timestamp](#codenotary-timestamp)
- [Quick start](#quick-start)
- [Usage](#usage)

## CodeNotary Timestamp

Every build has a story to tell - Tamperproof provenance for your code and CI/CD pipeline

vcn is the command interface for [CodeNotary Timestamp](https://timestamp.codenotary.com/).
Notarized assets informations are stored in a tamperproof ledger with cryptographic verification backed by [immudb](https://codenotary.com/technologies/immudb/), the immutable database.
Thanks to this `vcn` is extremely fast and provides full immutability for all data ever stored including its history and is cryptographically verifiable.

There are plenty of different use cases:

* CI/CD integration - notarize all outgoing assets, authenticate all incoming
* Store provenance for you own sources
* Trust or remove trust for digital assets and act on the status (within your script or pipeline)
* Simply timestamp files and sources to doublecheck at any time in the future if these are still the same (backup/restore)

**Simply request your API key and get started within a minute! This service is free forever and was built using Open Source!**

### Obtain an API Key
To provide access to our timestamping service a valid API Key is required. If you don't have one yet, simply signup here [CodeNotary Timestamp](https://timestamp.codenotary.com/)
This API Key is bound to your email address and it's required during vcn login.


## Quick start

1. **Installer** In case you use Linux or macOS, the quickest start is our install script:
```bash
bash <(curl https://getvcn.codenotary.com -L)
```

You can also [download the latest release](https://github.com/vchain-us/vcn/releases/latest)

2. **Login** to timestamp.codenotary.com

```bash
vcn login --lc-host timestamp.codenotary.com # type in your API key if requested

# or setting the API key

VCN_LC_API_KEY=<Your-API-Key vcn login --lc-host timestamp.codenotary.com
```


3. **Notarize existing digital objects** Once you have an account you can start notarizing digital assets to give them an identity.

   ```bash
   vcn n <file|dir://directory|docker://dockerimage|git://gitdirectory>

4. **Authenticate digital objects** You can use the command as a starting point.

   ```bash
   vcn a <file|dir://directory|docker://dockerimage|git://gitdirectory>
   ```

For detailed **command line usage** just run `vcn help`.

## Usage

### Wildcard support and recursive notarization

It's also possible to notarize assets using a wildcard pattern.

With `--recursive` flag the utility can recursively notarize inner directories.
```shell script
./vcn n "*.md" --recursive
```

### Notarization

Start with the `login` command. `vcn` will walk you through login and importing up your secret upon initial use.

```
vcn login --lc-host timestamp.codenotary.com
```

Once your secret is set you can notarize assets like in the following examples:

```
vcn notarize <file>
vcn notarize dir://<directory>
vcn notarize docker://<imageId>
vcn notarize podman://<imageId>
vcn notarize git://<path_to_git_repo>
vcn notarize --hash <hash>
```

Change the asset's status:

```
vcn unsupport <asset>
vcn untrust <asset>
```

### Authentication

```
vcn authenticate <file>
vcn authenticate dir://<directory>
vcn authenticate docker://<imageId>
vcn authenticate podman://<imageId>
vcn authenticate git://<path_to_git_repo>
vcn authenticate --hash <hash>
```

To output results in `json` or `yaml` formats:
```
vcn authenticate --output=json <asset>
vcn authenticate --output=yaml <asset>
```

## Examples

#### Authenticate a Docker image automatically prior to running it

First, you’ll need to pull the image by using:

```
docker pull hello-world
```

Then use the below command to put in place an automatic safety check. It allows only verified images to run.

```
vcn authenticate docker://hello-world && docker run hello-world
```
If an image was not verified, it will not run and nothing will execute.


#### Authenticate multiple assets
You can authenticate multiple assets by piping other command outputs into `vcn`:
```
ls | xargs vcn authenticate
```
> The exit code will be `0` only if all the assets in you other command outputs are verified.

#### Authenticate by a specific signer
By adding `--signerID`, you can authenticate that your asset has been signed by a specific SignerID.
> A SignerID is the signer public address (represented as a 40 hex characters long string prefixed with `0x`).

```
vcn authenticate --signerID 0x8... docker://hello-world
```

#### Authenticate using the asset's hash

If you want to authenticate an asset using only its hash, you can do so by using the command as shown below:

```
vcn authenticate --hash fce...
```

#### Unsupport/untrust an asset you do not have anymore

In case you want to unsupport/untrust an asset of yours that you no longer have, you can do so using the asset hash(es) with the following steps below.

First, you’ll need to get the hash of the asset from your CodeNotary Ledger Compliance dashboard or alternatively you can use the `vcn list` command. Then, in the CLI, use:

```
vcn untrust --hash <asset's hash>
# or
vcn unsupport --hash <asset's hash>
```

#### Notarization within automated environments

First, you’ll need to make `vcn` have access to the `${HOME}/.vcn` folder that holds your secret (the private key).
Then, set up your environment accordingly using the following commands:

```bash
export VCN_LC_API_KEY=Your-API-Key
```

Once done, you can use `vcn` in your non-interactive environment using:

```
vcn login --lc-host timestamp.codenotary.com
vcn notarize <asset>
```

> Other commands like `untrust` and `unsupport` will also work.


#### Add custom metadata when signing assets
The user can upload custom metadata when doing an asset notarization using the `--attr` option, e.g.:

```shell script
vcn n README.md --attr Testme=yes --attr project=5 --attr pipeline=test
```

This command would add the custom asset metadata Testme: yes, project: 5, pipeline: test.

The user can read the metadata back on asset authentication, i.e. using the `jq` utility:

```shell script
vcn a README.md -o json | jq .metadata
```

#### Inspect
Inspect has been extended with the addition of new filter: `--last`, `--first`, `--start` and `--end`.
With `--last` and `--first` are returned the N first or last respectively.

```shell script
vcn inspect document.pdf --last 10
```

With `--start` and `--end` it's possible to use a time range filter:

```shell script
vcn inspect document.pdf --start 2020/10/28-08:00:00 --end 2020/10/28-17:00:00
```

If no filters are provided only maximum 100 items are returned.

#### Signer Identifier
It's possible to filter results by signer identifier:

```shell script
vcn inspect document.pdf --signerID Cyg...
```

