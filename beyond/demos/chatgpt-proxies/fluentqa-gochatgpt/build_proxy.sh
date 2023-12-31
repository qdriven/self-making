# env GOOS=linux GOARCH=amd64
go build -o build/chatproxy cmd/proxy/main.go



# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

# GOOS - Target Operating System	GOARCH - Target Platform
# android	arm
# darwin	386
# darwin	amd64
# darwin	arm
# darwin	arm64
# dragonfly	amd64
# freebsd	386
# freebsd	amd64
# freebsd	arm
# linux	386
# linux	amd64
# linux	arm
# linux	arm64
# linux	ppc64
# linux	ppc64le
# linux	mips
# linux	mipsle
# linux	mips64
# linux	mips64le
# netbsd	386
# netbsd	amd64
# netbsd	arm
# openbsd	386
# openbsd	amd64
# openbsd	arm
# plan9	386
# plan9	amd64
# solaris	amd64
# windows	386
# windows	amd64