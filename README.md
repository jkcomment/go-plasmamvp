# go-plasmamvp

## Install

```
go get -u github.com/yuzushioh/go-plasmamvp/cmd/plasma
```

## Usage

```
export PRIVATE_KEY=[Private Key]
export CONTRACT_ADDRESS=$(plasma deploy)

# plasma deposit, submitblock and withdraw use CONTRACT_ADDRESS environment
plasma deposit [AMOUNT]
plasma submitblock ?
plasma withdraw ?
```
