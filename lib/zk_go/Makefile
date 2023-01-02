ifeq ($(OS),Windows_NT)
	OS_TYPE := Windows
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME),Linux)
		OS_TYPE := Linux
	else ifeq ($(UNAME),Darwin)
		OS_TYPE := Darwin
	else
		OS_TYPE := Unknown
	endif
endif

GO   ?= $(shell which go)
MAKE ?= $(shell which make)

ifndef V
  V := @
else
  V :=
  GO_TEST_FLAGS := -v
endif

zk_core_DIR := $(PWD)/../zk_core
zk_setup_DIR := $(PWD)/../zk_setup
zk_publish_DIR := $(PWD)/../zk_publish

zk_setup := $(zk_setup_DIR)/zk_setup
zk_publish := $(zk_publish_DIR)/zk_publish
ifeq ($(OS_TYPE),Windows)
  zk_setup := $(addsuffix .exe,$(zk_setup))
  zk_publish := $(addsuffix .exe,$(zk_publish))
endif

# -a option of zk_setup
zk_setup_OPTION_A := 1024
# -c option of zk_publish in plain mode
zk_publish_OPTION_C := 1024

CGO_LDFLAGS := -L$(zk_core_DIR)

TEST_DATA_DIR := $(PWD)/test_data
TEST_ECC_PUB := $(TEST_DATA_DIR)/ecc_pub.bin
TEST_ZKSNARK_PUB := $(TEST_DATA_DIR)/zksnark_key
TEST_PLAIN_FILE := $(TEST_DATA_DIR)/plain.txt
TEST_TABLE_FILE := $(TEST_DATA_DIR)/table.csv

TEST_OUTPUT_DIR := $(TEST_DATA_DIR)/output

TEST_PUBLISH_DIR := $(TEST_DATA_DIR)/publish
TEST_PUBLISH_PLAIN_DIR := $(TEST_PUBLISH_DIR)/plain
TEST_PUBLISH_TABLE_DIR := $(TEST_PUBLISH_DIR)/table

PKG_ROOT_NAME := github.com/xuxinlai2002/zkFilzkFil-lib/zk_go

ifeq ($(OS_TYPE), Darwin)
	LIBRARY_PATH := DYLD_LIBRARY_PATH=$(zk_core_DIR)
else
	LIBRARY_PATH := LD_LIBRARY_PATH=$(zk_core_DIR)
endif

GO_TEST := $(LIBRARY_PATH) \
		CGO_LDFLAGS=$(CGO_LDFLAGS) \
		$(GO) test $(GO_TEST_FLAGS)

all:

test: clean test_publish zk_core
	@echo "+ Testing ..."
	$(V) rm -rf $(TEST_OUTPUT_DIR)
	$(V) mkdir -p $(TEST_OUTPUT_DIR)
	$(V) $(GO) clean -testcache ./...
	mv ecc_pub.bin $(TEST_ECC_PUB)
	$(V) $(GO_TEST) ./...
	make clean

test_publish: test_publish_plain test_publish_table

test_publish_plain: $(zk_publish) $(TEST_ZKSNARK_PUB) $(TEST_PLAIN_FILE)
	@echo "+ Publishing " $(TEST_PLAIN_FILE) "..."
	$(V) rm -rf $(TEST_PUBLISH_PLAIN_DIR)
	$(V) mkdir -p $(TEST_PUBLISH_PLAIN_DIR)
	$(V) $(zk_publish) \
		-m plain \
		-f $(TEST_PLAIN_FILE) \
		-o $(TEST_PUBLISH_PLAIN_DIR) \
		-c $(zk_publish_OPTION_C)

test_publish_table: $(zk_publish) $(TEST_ZKSNARK_PUB) $(TEST_TABLE_FILE)
	@echo "+ Publishing " $(TEST_TABLE_FILE) "..."
	$(V) rm -rf $(TEST_PUBLISH_TABLE_DIR)
	$(V) mkdir -p $(TEST_PUBLISH_TABLE_DIR)
	$(V) $(zk_publish) \
		-m table \
		-f $(TEST_TABLE_FILE) \
		-o $(TEST_PUBLISH_TABLE_DIR) \
		-t csv \
		-k 0 1 2 3 4 5

zk_core:
	@echo "+ Building zk_core ..."
	$(V) cd $(zk_core_DIR); $(MAKE) lib

$(zk_setup):
	@echo "+ Building zk_setup ..."
	$(V) cd $(zk_setup_DIR); $(MAKE)

$(zk_publish):
	@echo "+ Building zk_publish ..."
	$(V) cd $(zk_publish_DIR); $(MAKE)

$(TEST_ZKSNARK_PUB): $(zk_setup)
	@echo "+ Generating zksnark_pub"
	@echo "+ "$(zk_setup)" -o "$(TEST_ZKSNARK_PUB)
	$(V) $(zk_setup) -o $(TEST_ZKSNARK_PUB)

clean:
	@echo "+ Cleaning ..."
	$(V) rm -f $(TEST_ECC_PUB)
	$(V) rm -rf $(TEST_ZKSNARK_PUB)
	$(V) rm -rf $(TEST_PUBLISH_DIR)
	$(V) rm -rf $(TEST_OUTPUT_DIR)
