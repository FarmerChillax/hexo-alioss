build-win:
	CGO_ENABLE=0 GOOS=windows go build -o bin/win_x64/hexo_ali_oss.exe

build-darwin:
	CGO_ENABLE=0 GOOS=darwin go build -o bin/darwin/hexo_ali_oss

build-linux:
	CGO_ENABLE=0 GOOS=linux go build -o bin/linux/hexo_ali_oss

# GOARCH=arm

build-all: build-darwin build-linux build-win

release-linux:
	tar -c bin/linux -f release/hexo_ali_oss_0.2.0_linux.tar

release-darwin:
	tar -c bin/linux -f release/hexo_ali_oss_0.2.0_darwin.tar

release-win:
	zip -r release/hexo_ali_oss_0.2.0_win.zip  bin/win 

release-all: release-linux release-darwin release-win

test:
	go test ./...

export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat linux:riscv64

version =  v0.2.1

all: build

build: app

app:
	@$(foreach n, $(os-archs),\
		os=$(shell echo "$(n)" | cut -d : -f 1);\
		arch=$(shell echo "$(n)" | cut -d : -f 2);\
		gomips=$(shell echo "$(n)" | cut -d : -f 3);\
		target_suffix=$${os}_$${arch};\
		echo "Build $${os}-$${arch}...";\
		env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} GOMIPS=$${gomips} go build -trimpath -ldflags "$(LDFLAGS)" -o ./release/"$(version)"/bin/hexo_ali_oss_"$(version)"_$${target_suffix} ;\
		echo "Build $${os}-$${arch} done";\
	)

# # tar -c ./release/"$(version)"/bin/hexo_ali_oss_"$(version)"_$${target_suffix} -f ./release/"$(version)"/hexo_ali_oss_"$(version)"_$${target_suffix}.tar \
