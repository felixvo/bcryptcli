## Description
Bcrypt CLI encode/check match bcrypt
## Installation
`go get -u github.com/felixvo/bcryptcli`
### Usage
#### Encode
```bash
bcryptcli your-password
$2a$10$Q5psIIboePHnIKI4vBKKNOWdOX3o9tnwVtaSEgQW4rRKNMDftQFVm
```
#### Check match
```bash
bcryptcli -p your-password '$2a$10$hWePochhBRiLLGMdA9c04.0LcUGoobW7bjnbxpFT.ib6jlnCIlJ/6'
MATCH
```