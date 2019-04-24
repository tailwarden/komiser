<a href="https://amphp.org/">
  <img src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/logo.png" width="200" align="right" alt="Amp Logo">
</a>

<a href="https://amphp.org/"><img alt="Amp" src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/icon.png
" width="120" valign="middle"></a>

Stay under budget by uncovering hidden costs, monitoring increases in spend, and making impactful changes based on custom recommendations.

[![Docker Stars](https://img.shields.io/docker/pulls/mlabouardy/komiser.svg)](https://hub.docker.com/r/mlabouardy/komiser/) 
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![CircleCI](https://circleci.com/gh/mlabouardy/komiser/tree/master.svg?style=svg&circle-token=d35b1c7447995e60909b24fd316fef0988e76bc8)](https://circleci.com/gh/mlabouardy/komiser/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/mlabouardy/komiser)](https://goreportcard.com/report/github.com/mlabouardy/komiser) [![Docker Stars](https://img.shields.io/github/issues/mlabouardy/komiser.svg)](https://github.com/mlabouardy/komiser/issues) [<img src="https://img.shields.io/badge/slack-@komiser-yellow.svg?logo=slack">](https://komiser.slack.com/messages/C9SQPU4Q0/)

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

### AWS

* Create an IAM user with the following IAM [policy](https://raw.githubusercontent.com/mlabouardy/komiser/master/policy.json):

```
wget https://s3.amazonaws.com/komiser/aws/policy.json
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
    <img src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/dashboard.png"/>
</p>

## Documentation

Documentation can be found on [komiser.io](https://docs.komiser.io) as well as in the [`./docs`](./docs) directory.

## Examples

<p align="center">

[![IMAGE ALT TEXT HERE](https://s3.eu-west-3.amazonaws.com/komiser-assets/images/thumbnail.png)](https://www.youtube.com/watch?v=DDWf2KnvgE8)

</p>

## License

The MIT License (MIT). Please see [`LICENSE`](./LICENSE) for more information.