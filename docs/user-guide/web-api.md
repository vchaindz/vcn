# Web API

All endpoints accepts Basic Auth with user credentials (mantatory for notarization).
If a custom notarization password is needed, add the `x-notarization-password` header, otherwise the login password will be used instead.

## Notarization

**Endpoints**
- POST `/notarize`
- POST `/untrust`
- POST `/unsupport`

**Query params**
- `public` (if present and not empy will set the visibility to public, otherwise private)


**Body request**
```json
{
  "kind": "file", // string, optional
  "name": "filename.pdf", // string
  "hash": "......", // string
  "size": 4096, // int, optional, cannot be < 0
  "contentType": "application/pdf", // string, optional
  "metadata": { // object, optional
    // ...
  }
}
```

**Body response**
Same as authentication, see below.

## Authentication

**Endpoint**
- GET `/authentication/<hash>`

**Query params**
- `signers` comma-separated list of SignerID(s)
- `org` organization ID
> `org`, if present, takes precedence over `signers`

**Body response**
> Results are indentical to `vcn authenticate ... --output=json` ones.
- example of unverified asset:
```json
{
  "kind": "",
  "name": "",
  "hash": "non-existing-hash",
  "verification": {
    "level": 0,
    "owner": "",
    "status": 2,
    "timestamp": ""
  }
}
```

- example of trusted asset with some fields omitted (because empty)
```json
{
  "kind": "",
  "name": "Test vector",
  "hash": "test",
  "metadata": {
    "note": "This hash was signed for testing purpose"
  },
  "visibility": "PUBLIC",
  "createdAt": "2019-06-28T22:46:46.317819",
  "verificationCount": 28,
  "signerCount": 1,
  "signer": "leonardo@vchain.us",
  "company": "vChain",
  "verification": {
    "level": 3,
    "owner": "0x7f66cb537c27251d007bd3c8ec731690c744f5e4",
    "status": 0,
    "timestamp": "2019-06-28T22:46:45Z"
  }
}
```

- example with all field populated
```json
{
  "kind": "file",
  "name": "vcn-v0.6.3-linux-amd64",
  "hash": "cabea5ccdf9380775f1d40fd2a1baec8ee697ecf107f13283bcfc08bd0c9df65",
  "size": 16433816,
  "contentType": "application/x-executable",
  "url": "...",
  "metadata": {
    "architecture": "x86_64",
    "file": {
      "arch": "X86_64",
      "format": "ELF",
      "platform": "GNU/Linux",
      "type": "EXEC",
      "x64": true
    },
    "platform": "GNU/Linux",
    "version": "0.6.3"
  },
  "visibility": "PUBLIC",
  "createdAt": "2019-09-25T14:01:23.159792",
  "verificationCount": 3,
  "signerCount": 1,
  "signer": "leonardo@vchain.us",
  "company": "vChain",
  "website": "https://codenotary.io",
  "verification": {
    "level": 3,
    "owner": "0x068e10d036175b874017320db5a9b852620679c4",
    "status": 0,
    "timestamp": "2019-09-25T14:01:20Z"
  }
}
```

