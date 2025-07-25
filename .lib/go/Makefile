.PHONY: */* */*/*

# Rule to build each exercise
*/*/*:
	@echo "Building $@..."
	-cd $@ &&	go mod init $@
	cd $@ && go test

*/*:
	@echo "Building $@..."
	-cd $@ &&	go mod init
	cd $@ &&	go test
