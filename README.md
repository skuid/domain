# Skuid Shared Domain Library

## Purpose

The purpose of this library is to share business logic (authentication), repeatable logic (command line arguments, logging), cross-platform logic (models, constants, protobuf contract definitions) and utilities across projects (e.g. `marina`, `tides`).


# Skuid NLX

At the top level we have tons of NLX-specific logic. This could be moved down into its own package (like `nlx`).

This includes models, authentication, some http stuff, deployment and retrieval logic as well as some `zip` utilities.

# Authorization

The `auth` package provides Skuid-specific logic to do the two-step handshake to get tokens for web use.

# Constants

This is where you'll find variables that are used over and over again in our code and that represent something immutable.

# Errors

Uniform collection for errors.

# Flags

This is a bit complex. `flags` should be a subpackage of `cmd`, since it is also related to `cli` work. This is where we're going to have generic logic that can attach flags to a command and handle different types and environment variable output as well.

# Generated

See "Protobuf"

# Legacy

This is a handful of holdover code from `skuid-cli` before the refactor.

# Proto

## Protocol Buffers

Protocol Buffers are a contract definition language that allow for the creation of multi-language support controls. You can compile proto files to a multitude of languages and have gRPC across those that use it.

Prerequisites:

```
brew install protobuf
go install github.com/favadi/protoc-go-inject-tag
```
### `.proto` directory:

In this directory you'll find the `*.proto` definitions of the contract between `marina` and `tides`.

All `.proto` files can be compiled into target languages. In this case, we're going to target `go` and target the directory / package `generated`

### Step 1: compile the `.go` code:

```bash
	protoc \
		--go_out=generated \
		--go-grpc_out=generated \
		--go-grpc_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		--proto_path=proto \
		.proto/*.proto
```

Then, since we're going to be injecting annotations / manipulating the generated files, run this step. Without it, we may not see the correct `json` mapping as outlined in the comments of the `.proto` definition.

### Step 2: Inject Tags
```bash
	protoc-go-inject-tag \
		-input="generated/*.pb.go"
```

Now you have the compiled code. 

**Tip:** whenever you change `.proto` files, make sure to commit the `generated` files so that they're consumable downstream.
