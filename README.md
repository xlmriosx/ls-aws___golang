# ☁️ AWS Resource Lister --- Golang

A command-line tool written in Go that validates AWS credentials and lists AWS resources like **S3 buckets**, **EC2 instances**, and **Lambda functions**, using the AWS SDK for Go v2.

## 📌 Project Overview

This CLI tool is useful for quickly verifying your AWS credentials and exploring cloud resources in a given AWS region. It uses AWS STS to validate your session and provides readable output for common AWS services.

## 🚀 Features

- ✅ Validates AWS credentials with STS.
- 📂 Lists all S3 buckets in your account.
- 🖥️ Lists EC2 instances and their states.
- 🧠 Lists Lambda functions with runtime info.
- 🌎 Allows region selection via flag.
- 🧩 Uses official AWS SDK for Go v2.

## 🛠️ Prerequisites

- Go 1.18 or later.
- AWS credentials properly configured (`~/.aws/credentials` or environment variables).

## 🛠️ Run as a programmer
- Ensure you have Go installed (version 1.18+ recommended).
- Clone the repository or copy the code into a .go file (e.g., main.go).
- Enter to the folder.
- Run the program:

```bash
go run main.go -region <region> -operation <operation>
```

## 🍭 Run as a user
- Download artefact from [here]()
- Run executable:
```bash
ping-monitor -config <path>
```

## ⚙️ Flags
| Flag         | Description                                           | Default      |
| ------------ | ----------------------------------------------------- | ------------ |
| `-region`    | AWS region to use                                     | `us-east-1`  |
| `-operation` | Cloud operation: `list-s3`, `list-ec2`, `list-lambda` | *(required)* |

## 🔳 Output
You will see some as:

```bash
AWS credentials validated successfully
S3 Buckets:
- my-logs-bucket (Created: 2023-01-15 12:30:45 +0000 UTC)
- my-app-assets (Created: 2022-11-03 08:21:00 +0000 UTC)
```

## ❗ Errors
If the credentials are invalid or expired:
```bash
failed to validate AWS credentials: UnrecognizedClientException: The security token included in the request is invalid.
```

If an invalid operation is passed:
```bash
Error: Invalid operation. Use list-s3, list-ec2, or list-lambda
```
