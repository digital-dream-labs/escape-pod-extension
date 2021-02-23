build:
	CGO_ENABLED=0 go build \
	-ldflags "$(BUILDINFO_FLAGS) -extldflags "-static"" \
	-trimpath \
	-o escapepodextension cmd/main.go

