# Binary name
BINARY_NAME=zksync_claim_tools

all: abigen build
build-darwin: 
    GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe
build:
	go build -o $(BINARY_NAME)
clean:
	rm -f $(BINARY_NAME)
run: clean build
	./$(BINARY_NAME)
abigen:
	abigen --abi pkg/abi/distributor.abi --pkg contracts --type ZkSyncMerkleDistributor --out pkg/contracts/zksyncmerkledistributor.go
	abigen --abi pkg/abi/erc20.abi --pkg contracts --type Erc20 --out pkg/contracts/erc20.go