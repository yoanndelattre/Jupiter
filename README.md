# Jupiter Load Balancer tester

The **Jupiter Load Balancer tester** is a simple Go application that sends HTTP requests to a specified target URL via a SOCKS5 proxy, if configured. This documentation provides instructions on how to use the pre-built Docker image available at `ghcr.io/yoanndelattre/jupiter:main`.

## Prerequisites

Before using Jupiter Load Balancer tester, ensure that you have the following:

- [Docker](https://www.docker.com/) installed on your system.

## Getting Started

1. Pull the pre-built Docker container image from the GitHub Container Registry (available for amd64, arm64 and armv7):

   ```bash
   docker pull ghcr.io/yoanndelattre/jupiter:main
   ```

2. Run the container with the necessary environment variables:

   ```bash
   docker run -e TARGET_URL="https://your-target-url.com" -e SOCKS5_PROXY="127.0.0.1:1080" ghcr.io/yoanndelattre/jupiter:main
   ```

Make sure to replace `"https://your-target-url.com"` with your desired target URL and `"127.0.0.1:1080"` with the SOCKS5 proxy address and port.

## Environment Variables

You need to set the following environment variables to configure the Jupiter Load Balancer tester when running the pre-built Docker image:

- **TARGET_URL**: The target URL to which the load balancer will send HTTP requests.
- **SOCKS5_PROXY**: The address and port of the SOCKS5 proxy (e.g., "127.0.0.1:1080"). If not set, the load balancer will send requests directly without a proxy.

## Features

- The Jupiter Load Balancer tester is designed to continuously send HTTP requests to the specified target URL.
- If a SOCKS5 proxy is configured, it sends requests through the proxy.
- The load balancer includes a timeout of 5 seconds for each request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.