# grpc-cue-envoy-rbac

## Skaffold
```bash
skaffold dev
```

| service | port |
| --- | --- |
| [example](deployments/example.yaml) | 9000 |
| [envoy](deployments/envoy.yaml) | 8080 |

```bash
grpcurl -plaintext localhost:9000 list
grpcurl -plaintext localhost:8080 list
```
