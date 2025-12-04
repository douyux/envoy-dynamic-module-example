FROM --platform=$BUILDPLATFORM golang:1.24.2 AS builder

ARG TARGETPLATFORM
ARG ZIG_VERSION=0.15.1

# Configure Chinese apt mirror and install dependencies
RUN if [ -f /etc/apt/sources.list.d/debian.sources ]; then \
        sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources; \
    fi && \
    if [ -f /etc/apt/sources.list ]; then \
        sed -i 's|http://deb.debian.org|http://mirrors.tuna.tsinghua.edu.cn|g' /etc/apt/sources.list && \
        sed -i 's|http://security.debian.org|http://mirrors.tuna.tsinghua.edu.cn|g' /etc/apt/sources.list; \
    fi && \
    apt update && apt install -y curl xz-utils

# Install zig based on build platform
RUN ARCH=$(uname -m) && \
    curl -L "https://ziglang.org/download/${ZIG_VERSION}/zig-${ARCH}-linux-${ZIG_VERSION}.tar.xz" | tar -J -x -C /usr/local && \
    ln -s "/usr/local/zig-${ARCH}-linux-${ZIG_VERSION}/zig" /usr/local/bin/zig

WORKDIR /build

RUN go env -w GOPROXY=https://athens.qingteng.cn,https://goproxy.cn,direct && \
    go env -w GONOSUMDB=git.qingteng.cn/* && \
    go env -w GOPRIVATE=git.qingteng.cn && \
    go env -w GONOPROXY=abc

COPY go.mod ./
RUN go mod download

COPY ./ ./

RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
        CC="zig cc -target x86_64-linux-gnu" CXX="zig c++ -target x86_64-linux-gnu" CGO_ENABLED=1 GOARCH=amd64 go build -buildmode=c-shared -o /build/libgo_module.so .; \
    elif [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        CC="zig cc -target aarch64-linux-gnu" CXX="zig c++ -target aarch64-linux-gnu" CGO_ENABLED=1 GOARCH=arm64 go build -buildmode=c-shared -o /build/libgo_module.so .; \
    else \
        echo "Unsupported platform: $TARGETPLATFORM" && exit 1; \
    fi

FROM envoyproxy/envoy:v1.36.2

ENV ENVOY_DYNAMIC_MODULES_SEARCH_PATH=/usr/local/lib
COPY --from=builder /build/libgo_module.so /usr/local/lib/libgo_module.so