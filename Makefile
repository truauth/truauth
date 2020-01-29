build:
	echo "Building binaries"
	go build -o ./bin/authentication github.com/truauth/truauth/cmd/tru-authentication
	go build -o ./bin/authorization github.com/truauth/truauth/cmd/tru-authorization
	go build -o ./bin/identity github.com/truauth/truauth/cmd/tru-identity


run:
	./bin/authentication
	./bin/authorization
	./bin/identity


build-debug:
	echo "Building binaries for debug mode"
	go build -o ./bin/debug/authentication -gcflags "all=-N -l" github.com/truauth/truauth/cmd/tru-authentication
	go build -o ./bin/debug/authorization -gcflags "all=-N -l" github.com/truauth/truauth/cmd/tru-authorization
	go build -o ./bin/debug/identity -gcflags "all=-N -l" github.com/truauth/truauth/cmd/tru-identity

run-debug:
	dlv --listen=localhost:32100 --headless=true --api-version=2 --accept-multiclient exec ./bin/debug/authentication &
	dlv --listen=localhost:32101 --headless=true --api-version=2 --accept-multiclient exec ./bin/debug/authorization &
	dlv --listen=localhost:32102 --headless=true --api-version=2 --accept-multiclient exec ./bin/debug/identity &
	echo "You can now connect for debugging purposes."