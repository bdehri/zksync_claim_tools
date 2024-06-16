# ZkSync Claim Tools

You can use this tool to claim your zkSync tokens from the zkSync contract to the L2 network, pre-approve target contracts, and gather all tokens from multi-accounts to target address.

## Installation

```bash
git clone https://github.com/bdehri/zksync_claim_tools.git
cd zksync_claim_tools
```

Then,

```
make build
```

if make is not installed, you can run the following command instead:

```bash
go build -o zksync_claim_tools main.go
```

## Usage

### Claim

This command can be used to claim your zkSync tokens from the zkSync contract to the L2 network.

```bash
./zksync_claim_tools claim --config <config_path> 
```

If you want to set gas price manually in gwei, you can use the following command:

```bash
./zksync_claim_tools claim --config <config_path> --gas-price <gas_price>

```

If you want to gather all tokens to single account, you can use the following command:

```bash
./zksync_claim_tools claim --config <config_path> --gather true --target <target_address>
```

If you want to approve tokens to target contract after claiming, you can use the following command:

```bash
./zksync_claim_tools claim --config <config_path> --approve true --target <target_address>
```

### Approve

This command can be used to pre-approve zksync token to target contracts. 

```bash
./zksync_claim_tools approve --config <config_path> --target <target_address>
```

