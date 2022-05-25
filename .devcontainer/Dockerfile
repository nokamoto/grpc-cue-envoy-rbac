# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.234.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends clang-format

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip \
    && unzip protoc-3.20.1-linux-x86_64.zip -d /usr/local \
    && chmod +x /usr/local/bin/protoc \
    && rm protoc-3.20.1-linux-x86_64.zip

RUN curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 \
    && install skaffold /usr/local/bin/

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
