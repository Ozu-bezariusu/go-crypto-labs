# DL_pw_7

Go wrapper for library that works with algebra on elliptic curves and implements method to work with points on ECC Curve interface from crypto/elliptic.

## Requirements 

```bash
* go ^1.19
* github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0
```

## Installation 

1. Clone repository
```bash
git clone https://github.com/Ozu-bezariusu/DL_pw_7.git
```
```bash
cd DL_pw_7
```
```bash
go get .
```

2. Build and Run project
```bash
go build
```
```bash
./DL_pw_7
```
![Alt text](image.png)

## Usage 

```bash
import (
	ec "DL_pw_7/ec"
    "math/big"
)

var G ec.ECPoint = ec.BasePointGGet()
k := big.NewInt(12345)
H := ec.ScalarMult(*k, G)
```
