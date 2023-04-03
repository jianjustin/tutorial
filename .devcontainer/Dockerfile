# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.20, 1.19, 1-bullseye, 1.20-bullseye, 1.19-bullseye, 1-buster, 1.20-buster, 1.19-buster
ARG VARIANT=1.20-bullseye
FROM golang:${VARIANT}

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1