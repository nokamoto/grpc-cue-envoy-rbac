# grpc-cue-envoy-rbac

## Skaffold
```bash
make # create deployments/rbac.json (protoc-gen-rbac)
skaffold dev
```

| service | port |
| --- | --- |
| [envoy](deployments/envoy.yaml) | 8080 |
| [example](deployments/example.yaml) | 9000 |
| [authorization](deployments/authorization.yaml) | 9001 |
| [rbac](deployments/rbac.yaml) | 9002 |
| [reflection](deployments/reflection.yaml) | 9003 |

### /grpc.reflection.v1alpha.ServerReflection

```bash
grpcurl -plaintext localhost:8080 list
```

```mermaid
flowchart LR
    grpcurl[grpcurl -plaintext\nlocalhost:8080 list]
    skaffold[[skaffold portForward\nlocalhost:8080]]
    envoy(envoy:8080)
    authorization(authorization:9001)
    reflection(reflection:9003)

    grpcurl --> skaffold
    skaffold --> envoy

    subgraph minikube
        envoy -- /envoy.service.auth.v3.Authorization/Check --> authorization
        envoy -- /grpc.reflection.v1alpha.ServerReflection --> reflection
    end
```

### todo
```bash
grpcurl -plaintext localhost:9000 list
grpcurl -plaintext localhost:9001 list
grpcurl -plaintext localhost:9002 list
grpcurl -plaintext localhost:9003 list
```
