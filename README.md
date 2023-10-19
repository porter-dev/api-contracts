# API Contracts

API contracts for use across multiple services

- Protobuf contracts in Go generated using `task all`
- gRPC Server using [Connect RPC](https://connectrpc.com/docs/introduction/)
- gRPC Server generated using `task all`

## Contributing

To start working on this, run the following commands:

```shell
# install task
brew install go-task/tap/go-task

# install dependencies
task install-tools
```

## Installing

NPM:

```bash
npm i --legacy-peer-deps @porter-dev/api-contracts
```

Go

```
go get -u github.com/porter-dev/api-contracts
```

## Usage

- Make changes to the `generated/porter/v1` protos as needed
- Run `task all` before committing your changes to the PR. This will run the PR prechecks, as well as generate any new code required.
- Create a PR
- Upon subsequent changes to your PR, ensure you run `task all` or your changes will not be included in the merged release.

Note: if you have created any new files in your release, be sure to export them in the `index.ts` file found in `generated/js` (this file is not in fact generated).

## Creating a release

- Visit [Github Releases](https://github.com/porter-dev/api-contracts/releases)
- Click `Create a new release`
- Click `Choose a tag`, and choose the semantic version that suits your release type. Make sure this begins with `v` and is of the format `vX.Y.Z`
- Click `Create new tag`
- Set the form title to with the tag name that you just created
- Click "Generate release notes", then fill out an additionally required information

## Updating message formats

Please consult this guide for best practices on updating message types for backward- and forward-compatability: https://protobuf.dev/programming-guides/proto3/#updating

## Local development

Assuming that you use `workstation`:

### Porter dashboard

In the folder containing your `package.json`, run the following:
`npm i --save --legacy-peer-deps ../../api-contracts/generated/js`

## Porter server

In the folder containing your `go.mod`, run the following:

```bash
go mod edit -replace github.com/porter-dev/api-contracts=../api-contracts
go mod vendor
go mod tidy
```

### Cluster Control Plane

In the folder containing your `go.mod`, run the following:

```bash
go mod edit -replace github.com/porter-dev/api-contracts=../api-contracts
go mod vendor
go mod tidy
```
