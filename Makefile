.PHONY: build

# build
build-app:
	@echo " > Building [app]..."
	@go build
	@echo " > Finished building [app]"

# run
run-app: build-app
	@echo " > Running [app]..."
	@./dummy_app
	@echo " > Finished running [app]"