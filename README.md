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

### /nokamoto.github.ExampleService/CreateExample

```bash
grpcurl -plaintext -H x-username:nokamoto@example.com localhost:8080 nokamoto.github.ExampleService/CreateExample
```

example.proto
```protobuf
service ExampleService {
  rpc CreateExample(CreateExampleRequest) returns (Example) {
    option (nokamoto.github.authz) = {
      permission: "example.example.create"
    };
  }
}
```

rbac.json
```json
{
    "rules":[
        {
            "path":"/nokamoto.github.ExampleService/CreateExample",
            "authorization":{
                "permission":"example.example.create"
            }
        }
    ]
}
```

```mermaid
flowchart LR
    proto([example.proto])
    json([rbac.json])
    grpcurl["grpcurl\nlocalhost:8080\n-H x-username:nokamoto@example.com\nnokamoto.github.ExampleService/CreateExample"]
    skaffold[[skaffold portForward\nlocalhost:8080]]
    envoy(envoy:8080)
    authorization(authorization:9001)
    example(example:9000)
    rbac(rbac:9002)

    subgraph configuration
        proto -- protoc-gen-rbac --> json
    end

    json -- volumeMounts\n/etc/authorization/rbac.json --- authorization

    grpcurl -- "username = nokamoto@example.com" --> skaffold
    skaffold --> envoy

    subgraph minikube
        envoy -- "/envoy.service.auth.v3.Authorization/Check\nusername = nokamoto@example.com" --> authorization
        authorization -- "/nokamoto.github.RBACService/AuthorizeUser\nusername = nokamoto@example.com\npermission = example.example.create" --> rbac
        envoy -- /nokamoto.github.ExampleService/CreateExample --> example
    end
```

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
