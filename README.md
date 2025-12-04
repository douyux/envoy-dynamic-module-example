# Envoy Dynamic Module Example

This project demonstrates the usage of Envoy's dynamic modules feature with Go, including HTTP filtering, external authentication, and response compression.

## Overview

The example sets up an Envoy proxy with dynamic modules that routes traffic to a httpbin backend service. It showcases two different listener configurations to demonstrate the behavior differences when using external authentication filters.

### Starting the Sandbox

1. Navigate to the sandbox directory:

```bash
cd sandbox
```

2. Start all services using Docker Compose:

```bash
docker-compose up --build
```

### Stopping the Sandbox

```bash
docker-compose down
```

## Port Configuration

### Port 1062 - With External Authorization

This listener includes the following HTTP filter chain:

1. **Compressor** (gzip): Compresses responses larger than 1024 bytes
2. **External Auth** (ext_authz): Validates requests via gRPC (timeout: 0.25s)
3. **Dynamic Module** (passthrough): Custom Go-based HTTP filter
4. **Router**: Routes to backend

**Behavior**: When receiving large POST requests, this port will **block and timeout** because the ext_authz filter needs to inspect the request headers before allowing the request to proceed. However, with the compressor filter placed before it, the request processing may be delayed waiting for the full request body, causing the 0.25s timeout to be exceeded.

### Port 1063 - Without External Authorization

This listener includes a simplified HTTP filter chain:

1. **Compressor** (gzip): Compresses responses larger than 1024 bytes
2. **Dynamic Module** (passthrough): Custom Go-based HTTP filter
3. **Router**: Routes to backend

**Behavior**: This port processes requests **normally** without the ext_authz filter, allowing large POST requests to flow through without blocking or timeout issues.

## Testing Different Port Behaviors

### Testing Port 1063 (Normal Behavior)

Send a 1MB file to port 1063:

```bash
curl --http2-prior-knowledge -X POST http://localhost:1063/post --data-binary '@testfile/1m'
```

**Expected Result**: The request completes successfully. The httpbin service responds with the posted data information.

### Testing Port 1062 (Blocking Behavior)

Send a 1MB file to port 1062:

```bash
curl --http2-prior-knowledge -X POST http://localhost:1062/post --data-binary '@testfile/1m'
```
