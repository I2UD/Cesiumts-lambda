# Cesium Terrain Server on AWS Lambda

## Table of Contents

- [Overview](#overview)
- [Installation & Deployment](#installation--deployment)
  - [Prerequisites](#prerequisites)
  - [Deployment Steps](#deployment-steps)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Development & Testing](#development--testing)
  - [Running Locally](#running-locally)
  - [Modifying the Code](#modifying-the-code)
- [Contribution Guidelines](#contribution-guidelines)
- [License](#license)
- [For Non-Technical Users](#for-non-technical-users)

---

## Overview

The **Cesium Terrain Server (CTS) on AWS Lambda** allows you to deploy a terrain server using AWS Lambda. This setup leverages the **AWS Lambda Go API Proxy** for efficient request handling and is designed to work seamlessly with **CesiumJS** applications.

This deployment is particularly useful for rendering terrain data dynamically and reducing infrastructure costs by utilizing AWS Lambda’s serverless architecture.

---

## Installation & Deployment

### Prerequisites

Ensure you have the following installed:
- [AWS CLI](https://aws.amazon.com/cli/)
- [AWS SAM CLI](https://aws.amazon.com/serverless/sam/)
- [Docker](https://docs.docker.com/get-docker/) *(for containerized execution)*
- Go 1.x *(if modifying the Go code)*

### Deployment Steps

1. **Build the project:**
   ```sh
   sam build
   ```

2. **Validate the AWS CloudFormation template:**
   ```sh
   sam validate
   ```

3. **Deploy using guided setup:**
   ```sh
   sam deploy --guided
   ```
   If you use AWS profiles, specify one using `--profile <profile-name>`.

4. **Retrieve the API Gateway URL:**
   After deployment, AWS SAM will output the API Gateway URL that can be used to access the Cesium Terrain Server.

---

## Configuration

This project uses **AWS CloudFormation** templates (`template.yaml`) to define resources.
The following **parameters** can be modified:
- `BaseTerrainUrl`: Base URL prefix under which tilesets are served (default: `/tilesets`).
- `TILESET_ROOT`: Mounted path for terrain tilesets (default: `/mnt/lambda`).

You can modify `template.yaml` to adjust networking settings, IAM roles, or function timeout values.

---

## API Endpoints

The deployed AWS Lambda function exposes the following endpoints:
- **Ping health check:**
  ```sh
  GET /ping
  ```
- **Fetch terrain metadata:**
  ```sh
  GET /tilesets/{tileset}/layer.json
  ```
- **Fetch terrain tiles:**
  ```sh
  GET /tilesets/{tileset}/{z}/{x}/{y}.terrain
  ```

These endpoints serve terrain data in a format compatible with **CesiumJS**.

---

## Development & Testing

### Running Locally

You can test the application locally using **AWS SAM CLI**:
```sh
sam local start-api
```
This will start a local API Gateway for testing the endpoints before deploying.

### Modifying the Code

The main server logic is implemented in `main.go`. If you need to modify functionality:
1. Update the Go code in `main.go`.
2. Rebuild the image using:
   ```sh
   sam build
   ```
3. Redeploy using:
   ```sh
   sam deploy
   ```

---

## Contribution Guidelines

Bug reports and feature requests are welcome on the [issues page](https://github.com/dymaxionlabs/cesiumts-lambda). Contributions should follow these steps:

1. **Fork the repository**.
2. **Create a feature branch** (`git checkout -b feature-branch`).
3. **Commit your changes** (`git commit -m "Description of changes"`).
4. **Push the branch** (`git push origin feature-branch`).
5. **Submit a pull request**.

---

## License

This project is licensed under **Apache 2.0**. See the [LICENSE](LICENSE) file for details.

---

## For Non-Technical Users

If you're not a developer but need to deploy this project, this section will guide you through the basic steps to get the **Cesium Terrain Server** running without deep technical knowledge.

### **What is this project for?**
The Cesium Terrain Server allows applications like **CesiumJS** to visualize real-world terrain data dynamically. If you need to serve **high-quality elevation data** over the web without managing dedicated servers, this solution is for you. The server is designed to scale automatically, reducing operational overhead while ensuring seamless terrain rendering.

### **How to deploy this project?**
You only need to:
1. **Install AWS CLI and AWS SAM** (Amazon's Serverless Application Model).
2. **Run the deployment commands** as outlined in the [Deployment Steps](#deployment-steps) section.
3. **Copy the API Gateway URL** provided at the end of the deployment process.
4. **Use this URL** in your CesiumJS-based application to fetch terrain tiles.

This solution is perfect for municipalities, researchers, and GIS professionals who need access to **scalable terrain data** without deep cloud computing expertise.
