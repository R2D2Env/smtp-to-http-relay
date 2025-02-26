
# bd := $(shell date +%F_%Z%H.%M.%S)
# cid = $(shell echo xyz-$(bd))
# cid = $(shell docker create shmailr:$(bd))

# .PHONY: help
# help: 
# 	echo "Run `make build` to build a new version of shamilr"
# 
compile: 
	@echo Compile shamilr in docker container
	@docker build -t shmailr_builder:latest --build-arg HTTP_PROXY="$(http_proxy)" --build-arg HTTPS_PROXY="$(https_proxy)" --build-arg NO_PROXY="$(no_proxy)" . 

fetch:
	@echo Fetch compiled shmair from compile container 
	$(eval cid=$(shell docker create shmailr_builder:latest))
	docker cp $(cid):/usr/local/bin/shmailr ./bin/shmailr
	docker rm -v $(cid)
	docker rmi shmailr_builder:latest
	docker image prune -af

build: compile fetch
#	@echo "Building container on date $(bd)"
#	@docker build -t shmailr:$(bd) .
