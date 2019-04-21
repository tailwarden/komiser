<p align="center">
    <img src="komiser.png" width="15%"/>
</p>

[![Docker Stars](https://img.shields.io/docker/pulls/mlabouardy/komiser.svg)](https://hub.docker.com/r/mlabouardy/komiser/) 
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![CircleCI](https://circleci.com/gh/mlabouardy/komiser/tree/master.svg?style=svg&circle-token=d35b1c7447995e60909b24fd316fef0988e76bc8)](https://circleci.com/gh/mlabouardy/komiser/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/mlabouardy/komiser)](https://goreportcard.com/report/github.com/mlabouardy/komiser) [![Docker Stars](https://img.shields.io/github/issues/mlabouardy/komiser.svg)](https://github.com/mlabouardy/komiser/issues) [<img src="https://img.shields.io/badge/slack-@komiser-yellow.svg?logo=slack">](https://oresoftware.slack.com/messages/CCAD1H94G)



## Download

Below are the available downloads for the latest version of Komiser (1.0.0). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://s3.us-east-1.amazonaws.com/komiser/1.0.0/linux/komiser
```

### Windows:

```
wget https://s3.us-east-1.amazonaws.com/komiser/1.0.0/windows/komiser
```

### Mac OS X:

```
wget https://s3.us-east-1.amazonaws.com/komiser/1.0.0/osx/komiser
```

_Note_: make sure to add the execution permission to Komiser `chmod +x komiser`

## How to use

* Create an IAM user with the following IAM [policy](https://raw.githubusercontent.com/mlabouardy/komiser/master/policy.json):

```
wget https://s3.amazonaws.com/komiser/policy.json
```

* Add your **Access Key ID** and **Secret Access Key** to *~/.aws/credentials* using this format

``` 
[default]
aws_access_key_id = <access key id>
aws_secret_access_key = <secret access key>
region = us-east-1
```

* That should be it. Try out the following from your command prompt to start the server:

```
komiser start --port 3000 --duration 30
```

* Point your browser to http://localhost:3000

<p align="center">
    <img src="screenshot.png"/>
</p>

## Docker

Use the official Komiser image:

```
docker run -d -p 3000:3000 -e AWS_ACCESS_KEY_ID="" -e AWS_SECRET_ACCESS_KEY="" -e AWS_DEFAULT_REGION="" --name komiser mlabouardy/komiser
```

### Environment Variables

| Environment Variable | Description |
| -------------------- | ----------- |
| PORT | Server Port |
| DURATION | Server Cache Expiration (minutes) |
| AWS_ACCESS_KEY_ID | AWS Access Key |
| AWS_SECRET_ACCESS_KEY | AWS Secret Key |
| AWS_DEFAULT_REGION | AWS Region |
| AWS_PROFILE | AWS Profile ([Docs](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)) |

## Configuring Credentials

When using the CLI you'll generally need your AWS credentials to authenticate with AWS services. Komiser supports multiple methods of supporting these credentials. By default the CLI will source credentials automatically from its default credential chain.

* Environment Credentials - Set of environment variables that are useful when sub processes are created for specific roles.

* Shared Credentials file (~/.aws/credentials) - This file stores your credentials based on a profile name and is useful for local development.

* EC2 Instance Role Credentials - Use EC2 Instance Role to assign credentials to application running on an EC2 instance. This removes the need to manage credential files in production.

## Supported AWS Services

### Compute

* EC2
* Lambda
* Elasticbeanstalk
* ECS
* EKS

### Storage

* S3
* EBS
* RDS
* DocumentDB
* DynamoDB
* Elasticache
* Glacier
* EFS
* Snapshots

### Security

* IAM
* Security Groups
* ACL
* VPC
* KeyPairs
* KMS
* Route Tables

### Network

* Internet Gateway
* Route53
* ELB
* CloudFront
* Nat gateway
* Elastic IP

### Data & AI

* Kinesis

### Monitoring

* CloudWatch
* CloudTrail

### Application Integrations

* SES
* SQS

## Maintainers

* Mohamed Labouardy

## Discussion

Join us on Slack: https://komiser.slack.com

## Tutorial

* [Komiser: AWS Environment Inspector](http://www.blog.labouardy.com/komiser-aws-environment-inspector/)

## License

MIT
