# API Contracts

API contracts for use across multiple services

- Task can be installed with `brew install go-task/tap/go-task`
- Protobuf contracts in Go generated using `task all`
- gRPC Server using [Buf Connect](https://connect.build/docs/introduction)
- gRPC Server generated using `task all`

## Installing

NPM:

```bash
npm i --legacy-peer-deps @porter-dev/api-contracts
```

Go

```
go get -u github.com/porter-dev/api-contracts
```

## Creating a release

Run the following commands, to create a release, filling out the tag field. Make sure that the tag is the next supported tag, and doesnt overwrite a previous one, or else this will cause an issue in Go and Javascript.
`vX.X.X` should be an accepted semantic version of your package

Note: if you have created any new files in your release, be sure to export them in the `index.ts` file found in `generated/js` (this file is not in fact generated).

```bash
task release -- vX.X.X
```

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
