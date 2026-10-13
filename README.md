# gRPC Truss Example

This project demonstrates a gRPC service using Truss for model deployment.

## Prerequisites

- Docker CLI installed
- Go installed
- Protocol Buffer compiler (`protoc`) and the Go plugins `protoc-gen-go` and `protoc-gen-go-grpc`, if you regenerate the included bindings
- Access to a container registry (Docker Hub, Google Container Registry, etc.)

## Setup

Get the repository with:
```bash
git clone https://github.com/basetenlabs/truss-grpc-example.git
```

Install Truss with:
```bash
pip install --upgrade truss
```

## Deployment

### 0. Generate Protobuf files

Generated Go bindings are included. Run this step only after editing `example.proto`, with the Protocol Buffer compiler and both Go plugins installed.

```bash
protoc --go_out=. --go-grpc_out=. example.proto
```
For more information about Protobuf, refer to the [Protobuf documentation](https://protobuf.dev/reference/go/go-generated/).

### 1. Build and Push Docker Image

First, build the Docker image using the provided Dockerfile:

```bash
docker build -t your-registry/truss-grpc-demo:latest . --platform linux/amd64
docker push your-registry/truss-grpc-demo:latest
```

### 2. Update Configuration

Update the `config.yaml` file to use your newly built image. Keep `server_port: 50051` to match the server. Baseten uses the gRPC health service; the HTTP endpoint fields are ignored.

```yaml
base_image:
  image: your-registry/truss-grpc-demo:latest  # Replace with your image
[...]
```

### 3. Push Model with Truss

Deploy your model using the Truss CLI:

```bash
truss push . --promote
```

For more detailed information about Truss deployment, refer to the [truss push documentation](https://docs.baseten.co/reference/cli/truss/push).

### 4. Call the Model
Set your model ID and [Baseten API key](https://docs.baseten.co/organization/api-keys) in the environment:

```bash
export BASETEN_MODEL_ID="YOUR_MODEL_ID"
export BASETEN_API_KEY="YOUR_BASETEN_API_KEY"
go run ./client
```

The client connects with TLS to `model-MODEL_ID.grpc.api.baseten.co:443`, sends the required routing and authentication metadata, and prints each streamed reply. Normal stream completion exits successfully; server errors fail the command. The request has a 60-second timeout.

See the [gRPC deployment guide](https://docs.baseten.co/development/model/grpc) for the public endpoint and deployment contract.
