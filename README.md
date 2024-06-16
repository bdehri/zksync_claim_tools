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

If you set only_approve to true, no claim will be made, only the approval will be done.

```bash
./zksync_claim_tools --config <config_path>
```

If you want to claim tokens, you need to set only_approve to false. If you want to gather all tokens from multi-accounts to target address, you need to set gather_all to true and target address should be set.

If you want to claim tokens and give approval to specific contracts, you need to set approve to true and set the target contract to intended router address.