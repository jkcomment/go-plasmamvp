# go-plasmamvp

## Install

```
go get -u github.com/yuzushioh/go-plasmamvp/cmd/plasma
```

## Usage

```
export CONTRACT_ADDRESS=$(plasma deploy [Private Key])

# plasma deposit, submitblock and withdraw use CONTRACT_ADDRESS environment
plasma deposit [ADDRESS] [AMOUNT]
plasma submitblock ?
plasma withdraw ?
```
