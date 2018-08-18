GOPKG=github.com/alyyousuf7/libimobiledevice-go
IMAGE=libimobiledevice

.PHONY: all image shell binaries clean

all: binaries

image: Dockerfile
	@docker build -t $(IMAGE) .

shell: image
	@docker run --rm -it -v $(shell pwd):/go/src/$(GOPKG) $(IMAGE)

binaries: bin/idevice_id bin/ideviceinfo

clean:
	@rm -rf bin/

bin/%: cmd/% **/*.go **/**/*.go
	@go build -o $@ $(GOPKG)/cmd/$(@F)/...
