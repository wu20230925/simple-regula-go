# simple-regula-go

A simple go SDK for regula.

# 1. How to use
See the example.

## 2. Update model

Clone repo [DocumentReader-web-openapi](https://github.com/regulaforensics/DocumentReader-web-openapi)
to local.

### 2.1 Bundle scheme to single.yml file

For generate code in Go by [oapi-codegen](https://github.com/deepmap/oapi-codegen/tree/master), here are some changes:

- Edit **rt.yml** - remove **components.schemas.ResultItem.discriminator** node
- Edit **rt-authenticity.yml** - remove **components.schemas.AuthenticityCheckResultItem.discriminator**
- Edit index.yml:

```
  /api/v2/transaction/{transaction_id}/process  ->  /api/v2/transaction/{transactionId}/process
  /api/v2/transaction/{transaction_id} -> /api/v2/transaction/{transactionId}
  /api/v2/transaction/{transaction_id}/results -> /api/v2/transaction/{transactionId}/results
  /api/v2/transaction/{transaction_id}/file -> /api/v2/transaction/{transactionId}/file
  /api/v2/tag/{tag_id} -> /api/v2/tag/{tagId}
```

```
Then run next command:
```bash
npx @redocly/openapi-cli@latest bundle index.yml > single.yml
```



it will generate single.yml.

### 2.3 generate Go code.

```bash
oapi-codegen -package model -generate "types,skip-prune" [path of single.yml]> model/model.gen.go
```

**Notion**: You need to install [oapi-codegen](https://github.com/deepmap/oapi-codegen/tree/master) firstly.
