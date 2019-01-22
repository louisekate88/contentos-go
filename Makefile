PACKAGES= github.com/coschain/contentos-go/cmd/wallet-cli/commands \
	github.com/coschain/contentos-go/cmd/wallet-cli/wallet \
	github.com/coschain/contentos-go/common/encoding/kope \
	github.com/coschain/contentos-go/common/encoding/vme \
	github.com/coschain/contentos-go/common/logging \
	github.com/coschain/contentos-go/db/storage \
	github.com/coschain/contentos-go/db/table \
	github.com/coschain/contentos-go/economist \
	github.com/coschain/contentos-go/p2p/message/types \
	github.com/coschain/contentos-go/tests/db \
	github.com/coschain/contentos-go/vm/contract/abi \
	github.com/coschain/contentos-go/vm/contract/table \
	github.com/coschain/contentos-go/vm


COSD = github.com/coschain/contentos-go/cmd/cosd
WALLET = github.com/coschain/contentos-go/cmd/wallet-cli

test:
	@echo "--> Running go test"
	@GO111MODULE=on go test -coverprofile=cc0.txt $(PACKAGES)
	@echo "--> Total code coverage"
	@GO111MODULE=on go run utils/totalcov/main.go . cc0.txt >coverage.txt

test_cover:
	@echo "--> Running go test with coverage"
	@GO111MODULE=on go test -cover $(PACKAGES)

test_detail:
	@echo "--> Running go test"
	@GO111MODULE=on go test -v $(PACKAGES)

build:
	@echo "--> build all"
	@GO111MODULE=on go build $(COSD)
	@GO111MODULE=on go build $(WALLET)

build_cosd:
	@echo "--> build cosd"
	@GO111MODULE=on go build $(COSD)

build_wallet:
	@echo "--> build wallet"
	@GO111MODULE=on go build $(WALLET)

install:
	@echo "--> build all"
	@GO111MODULE=on go install $(COSD)
	@GO111MODULE=on go install $(WALLET)

install_cosd:
	@echo "--> install cosd"
	@GO111MODULE=on go install $(COSD)

install_wallet:
	@echo "--> build wallet"
	@GO111MODULE=on go install $(WALLET)

collect-cover-data:
	@echo "collect cover data"
	rm coverage-all.out
	echo "mode: set" >> coverage-all.out
	$(foreach pkg, $(PACKAGES),\
	go test -coverprofile=coverage.out $(pkg) || exit $$?;\
	if [ -fcoverage.out ]; then \
	tail -n +2 coverage.out >> coverage-all.out;\
	fi\
	;)

test-cover:
	@echo "cover html"
	go tool cover -func=coverage-all.out

test-cover-html:
	@echo "cover html"
	go tool cover -html=coverage-all.out -o coverage.html
