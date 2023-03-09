# API Contracts

API contracts for use across multiple services

- Protobuf contracts in Go generated using `task all`
- gRPC Server using [Buf Connect](https://connect.build/docs/introduction)
- gRPC Server generated using `task all`

## Creating a release

Run the following commands, to create a release, filling out the tag field. Make sure that the tag is the next supported tag, and doesnt overwrite a previous one, or else this will cause an issue in Go and Javascript.

```bash
tag="v0.0.0"
```

```bash
task all
git add .
git commit -m "Releasing tag"
```
