# API Contracts

API contracts for use across multiple services

- Protobuf contracts in Go generated using `task all`
- gRPC Server using [Buf Connect](https://connect.build/docs/introduction)
- gRPC Server generated using `task all`

## Creating a release

Run the following commands, to create a release, filling out the tag field. Make sure that the tag is the next supported tag, and doesnt overwrite a previous one, or else this will cause an issue in Go and Javascript.
`vX.X.X` should be an accepted semantic version of your package

```bash
task release -- vX.X.X
```
