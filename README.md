**Komiser is back 🎉 and we're looking for maintainers to work on the new [roadmap](https://oraculi.canny.io/), if you're interested, join us on our <a href="https://discord.oraculi.io">Discord</a> community**

---

<h1 align="center"><img src="https://cdn.komiser.io/images/komiser-readme.banner.png" alt="Amp Logo"></h1>

<h4 align="center">
    <a href="https://discord.oraculi.io">Discord</a> |
    <a href="https://github.com/mlabouardy/komiser/discussions">Discussions</a> |
    <a href="https://komiser.io/">Site</a><br/><br/>
    <a href="https://docs.komiser.io/">Guide</a> |
    <a href="https://docs.komiser.io/docs/docs/how-to-komiser/alerts">How to Komiser</a> |
    <a href="https://docs.komiser.io/docs/intro">Docs</a><br/><br/>
    <a href="https://docs.komiser.io/docs/introduction/contribute">Contribute</a> | 
    <a href="https://docs.komiser.io/docs/FAQ/">FAQs</a><br/><br/>
</h4>

[![Price](https://img.shields.io/badge/price-FREE-0098f7.svg)](https://github.com/mlabouardy/komiser/blob/master/LICENSE) [![Docker Stars](https://img.shields.io/docker/pulls/mlabouardy/komiser.svg)](https://hub.docker.com/r/mlabouardy/komiser) 
[![ELv2 License](https://img.shields.io/badge/license-ELv2-green)](LICENSE) [![CircleCI](https://circleci.com/gh/mlabouardy/komiser/tree/master.svg?style=svg&circle-token=d35b1c7447995e60909b24fd316fef0988e76bc8)](https://circleci.com/gh/mlabouardy/komiser/tree/master) [![Docker Stars](https://img.shields.io/github/issues/mlabouardy/komiser.svg)](https://github.com/mlabouardy/komiser/issues) [![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.oraculi.io/)

Stay under budget by uncovering hidden costs, monitoring increases in spend, and making impactful changes based on custom recommendations.

**Discuss it on [Product Hunt](https://www.producthunt.com/posts/komiser) 🦄**

*Cloud version is available in private beta test stage, sign in for free at [https://cloud.oraculi.io](https://cloud.oraculi.io)*

[![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/fold_left.svg?style=social&label=Follow%20%40Komiser)](https://twitter.com/komiseree) [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Optimize%20Cost%20and%20Security%20on%20AWS&url=https://github.com/mlabouardy/komiser&via=mlabouardy&hashtags=komiser,aws,gcp,cloud,serverless,devops) 

## Komiser CLI, try it out!
---
The easiest way to get started with Komiser is to install the [Komiser CLI](https://docs.komiser.io/docs/overview/introduction/getting-started) by running the `Homebrew` commands below. Don't have Homebrew? Install it [here](https://docs.brew.sh/Installation).

```
brew tap HelloOraculi/komiser
brew install komiser
```

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Komiser?](#what-is-komiser)
  - [Who is using it?](#who-is-using-it)
- [Getting Started](#getting-started)
  - [Download](#download)
  - [Installation on AWS](#installation-on-aws)
  - [Installation on GCP](#installtion-on-gcp)
  - [Installation on Azure](#instalation-on-azure)
  - [Installation on Digital Ocean](#installation-on-digital-ocean)
  - [Installation on OVH](#installation-on-ovh)
- [Documentation](#documentation)
  - [Jump right in](#jump-right-in)
- [Bugs and feature requests](#bugs-and-feature-requests)
- [Users](#users)
- [Versioning](#versioning)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# What is Komiser?
Komiser is an open source project created in 2017. Created to **analyse** and **manage cloud cost**, **usage**, **security** and **governance** all in one place. With komiser you can also: 
* Control your **resource usage** and gain visibility across all used services to achieve maximum cost-effectiveness.
* Detect **potential vulnerabilities** that could put your cloud environment at risk.
* Get a deep understanding of **how you spend** on the AWS, GCP, OVH, DigitalOcean and Azure.

## Who is using it?
Komiser was built with every Cloud Engineer, Developer, DevOps engineer and SRE in mind. We understand that tackling cost savings, security improvements and resource usage analyse efforts can be hard, sometimes just knowing where to start, can be the most challenging part at tomes. Komiser is here to help those cloud practitioners see their cloud resources and accounts much more clearly. Only with clear insight can timely and efficient actions take place.

# Getting started

## Download

You can run Komiser locally, as a Docker container or running it inside a Kubernetes cluster.
Below are the available downloads for the latest version of Komiser (2.9.0). Please [download](https://docs.komiser.io/docs/overview/introduction/getting-started) the appropriate package for your operating system and architecture. 

## Installation on AWS

### Connect Komiser CLI to your AWS account. 
* Connect a [local deployment](https://docs.komiser.io/docs/Cloud%20Providers/aws#local-komiser-cli-single-account) of Komiser CLI to you AWS account

### Deploy Komiser to single account access EKS cluster (Helm chart)
* If you want to connect a single AWS account follow the documentation [here](https://docs.komiser.io/docs/Cloud%20Providers/aws#eks-installation-single-account).

### Deploy Komiser to a multi account access EKS cluster (Helm chart)
* If you are would like to connect various AWS accounts to a Komiser deployment in a Management EKS cluster, follow the steps [here](https://docs.komiser.io/docs/Cloud%20Providers/aws#multiple-account-eks-helm-chart-installation).

## Installtion on GCP 

Connect a local deployment of Komiser CLI to your [**GCP**](https://docs.komiser.io/docs/Cloud%20Providers/gcp) account.

## Instalation on Azure

Connect a local deployment of Komiser CLI to you [**Azure**](https://docs.komiser.io/docs/Cloud%20Providers/azure) account.

## Installation on Digital Ocean

Connect a local deployment of Komiser CLI to you [**Digital Ocean**](https://docs.komiser.io/docs/Cloud%20Providers/digital-ocean) account.

## Installation on OVH

Connect a local deployment of Komiser CLI to you [**OVH**](https://docs.komiser.io/docs/Cloud%20Providers/ovh) account

# Documentation

Head over to the official `Komiser` documentation at [docs.komiser.io](https://docs.komiser.io). The source repository for the documentation website is [helloOraculi/docs](https://github.com/helloOraculi/docs). 

We know that writing docs isn't usually at the top of too many peoples "What I like to do for fun" list, but if you feel so inclined, by all means, consider [contributing](https://docs.komiser.io/docs/introduction/contribute) to our documentation repository, we will be very grateful. It's built using [Docusaurus](https://docusaurus.io/). 

## Jump right in:
* [Documentation overview](https://docs.komiser.io/docs/intro)
* [Installation](https://docs.komiser.io/docs/overview/introduction/getting-started)
* [FAQs](https://docs.komiser.io/docs/FAQ/)
* Video series:
    * [How to: Komiser](https://docs.komiser.io/docs/docs/how-to-komiser/alerts)
    * [Cloud cost savings tips](https://docs.komiser.io/docs/Quickstarts/overview)

# Bugs and feature requests

Have a bug or a feature request? Please first read the issue guidelines and search for existing and closed issues. If your problem or idea is not addressed yet, [please open a new issue](https://github.com/mlabouardy/komiser/issues/new).

# Roadmap and Contributing

We are very excited about what is in store in the coming weeks and months, take a look at the [public roadmap](https://oraculi.canny.io/) to stay on top of what's coming down the pipeline. 

Komiser is written in `Golang` and is `Elv2 licensed` - contributions are always welcome whether that means providing feedback, be it through GitHub, through the `#feedback` channel on our [Discord server](https://discord.oraculi.io) or testing existing and new features. Feel free to check out our [contributor guidelines](./CONTRIBUTING.md) and consider becoming a **contributor** today. 

# Users

If you'd like to have your company represented and are using `Komiser` please give formal written permission below via a comment on this [thread](https://github.com/mlabouardy/komiser/issues/76) or via email to contact@oraculi.io.

We will need a URL to a svg or png logo, a text title and a company URL.

# Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/mlabouardy/komiser/tags). 
