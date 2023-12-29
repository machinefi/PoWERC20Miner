########################################################################################################################
# Copyright (c) 2019 IoTeX Foundation
# This source code is provided 'as is' and no warranties are given as to title or non-infringement, merchantability
# or fitness for purpose and, to the extent permitted by law, all liability for your use of the code is disclaimed.
# This source code is governed by Apache License 2.0 that can be found in the LICENSE file.
########################################################################################################################

PROJECT_NAME := "depinrc-20"
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: all build clean

all: build

build:
	@echo "Building binary..."
	@cd $(GOBASE)/cmd
	@GOBIN=$(GOBIN) go build -o $(GOBIN) $(GOBASE)/...

