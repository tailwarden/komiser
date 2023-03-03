**Komiser is back 🎉 and we're looking for maintainers to work on the new [roadmap](https://roadmap.tailwarden.com/), if you're interested, join us on our <a href="https://discord.tailwarden.com">Discord</a> community**

---

<h1 align="center"><img src="https://cdn.komiser.io/images/readme-komiser-header.png?version=latest" alt="Amp Logo"></h1>

<h4 align="center">
    <a href="https://discord.tailwarden.com">Discord</a> |
    <a href="https://github.com/tailwarden/komiser/discussions">Discussions</a> |
    <a href="https://komiser.io?utm_source=github&utm_medium=social/">Site</a><br/><br/>
    <a href="https://docs.komiser.io?utm_source=github&utm_medium=social/">Guide</a> |
    <a href="https://docs.komiser.io/docs/docs/how-to-komiser/alerts?utm_source=github&utm_medium=social">How to Komiser</a> |
    <a href="https://docs.komiser.io/docs/intro">Docs</a><br/><br/>
    <a href="https://docs.komiser.io/docs/introduction/contribute?utm_source=github&utm_medium=social">Contribute</a> | 
    <a href="https://roadmap.tailwarden.com">Roadmap</a><br/><br/>
</h4>

[![Price](https://img.shields.io/badge/price-FREE-0098f7.svg)](https://github.com/tailwarden/komiser/blob/master/LICENSE) [![Docker Stars](https://img.shields.io/docker/pulls/mlabouardy/komiser.svg)](https://hub.docker.com/r/mlabouardy/komiser) 
[![ELv2 License](https://img.shields.io/badge/license-ELv2-green)](LICENSE) [![CircleCI](https://circleci.com/gh/tailwarden/komiser/tree/master.svg?style=svg&circle-token=d35b1c7447995e60909b24fd316fef0988e76bc8)](https://circleci.com/gh/tailwarden/komiser/tree/master) [![Docker Stars](https://img.shields.io/github/issues/tailwarden/komiser.svg)](https://github.com/tailwarden/komiser/issues) [![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.tailwarden.com/)

Komiser is an open-source cloud-agnostic resource manager. It integrates with multiple cloud providers (including AWS, Azure, OCI, DigitalOcean, Kubernetes and CIVO, Linode, Scaleway and Tencent), builds a cloud asset inventory, and helps you break down your cost at the resource level 💰

<h1 align="center"><img src="https://cdn.komiser.io/gifs/readme-komiser-repo.gif?version=latest" alt="Komiser gif"></h1>

*Cloud version is available in private beta, sign in for free at [https://cloud.tailwarden.com](https://cloud.tailwarden.com?utm_source=github&utm_medium=social)*

[![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/fold_left.svg?style=social&label=Follow%20%40Komiser)](https://twitter.com/komiseree) [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Optimize%20Cost%20and%20Security%20on%20AWS&url=https://github.com/tailwarden/komiser&via=mlabouardy&hashtags=komiser,aws,gcp,cloud,serverless,devops) 

## Komiser CLI, try it out! 🚀
---
The easiest way to get started with Komiser is to install the [Komiser CLI](https://docs.komiser.io/docs/overview/introduction/getting-started?utm_source=github&utm_medium=social) by running the `Homebrew` commands below. Don't have Homebrew? Install it [here](https://docs.brew.sh/Installation).

```
brew tap tailwarden/komiser
brew install komiser
```

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Komiser?](#what-is-komiser)
  - [Who is using it?](#who-is-using-it)
- [Getting started](#getting-started)
  - [Download](#download)
  - [Installation on AWS](#installation-on-aws)
    - [Connect Komiser CLI to your AWS account.](#connect-komiser-cli-to-your-aws-account)
    - [Deploy Komiser to single account access EKS cluster (Helm chart)](#deploy-komiser-to-single-account-access-eks-cluster-helm-chart)
    - [Deploy Komiser to a multi account access EKS cluster (Helm chart)](#deploy-komiser-to-a-multi-account-access-eks-cluster-helm-chart)
  - [Installation on Azure](#installation-on-azure)
  - [Installation on Civo](#installation-on-civo)
  - [Installation on OCI](#installation-on-oci)
  - [Installation on Digital Ocean](#installation-on-digital-ocean)
  - [Installation on Linode](#installation-on-linode)
  - [Installation on Tencent](#installation-on-tencent-cloud)
  - [Installation on Scaleway](#installation-on-scaleway)
- [Documentation](#documentation)
  - [Jump right in:](#jump-right-in)
- [Bugs and feature requests](#bugs-and-feature-requests)
- [Roadmap and Contributing](#roadmap-and-contributing)
- [Users](#users)
- [Versioning](#versioning)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# What is Komiser? 🤷
Komiser is an open source project created to **analyse** and **manage cloud cost**, **usage**, **security** and **governance** all in one place. With komiser you can also: 
* Control your **resource usage** and gain visibility across all used services to achieve maximum cost-effectiveness.
* Detect **potential vulnerabilities** that could put your cloud environment at risk.
* Get a deep understanding of **how you spend** on the AWS, Civo, OVH, DigitalOcean and OCI.

<h1 align="center"><img src="https://cdn.komiser.io/images/komiser-readme-wireframe.png?version=latest" alt="Amp Logo"></h1>

## Who is using it?
Komiser was built with every Cloud Engineer, Developer, DevOps engineer and SRE in mind. We understand that tackling cost savings, security improvements and resource usage analyse efforts can be hard, sometimes just knowing where to start, can be the most challenging part at times. Komiser is here to help those cloud practitioners see their cloud resources and accounts much more clearly. Only with clear insight can timely and efficient actions take place.

# Getting started 👇

## Download

You can run Komiser locally, as a Docker container or by running it inside a Kubernetes cluster.
Below are the available downloads for the latest version of Komiser (3.0.0). Please [download](https://docs.komiser.io/docs/overview/introduction/getting-started?utm_source=github&utm_medium=social) the appropriate package for your operating system and architecture. 

## Installation on AWS

### Connect Komiser CLI to your AWS account. 
* Connect a [local deployment](https://docs.komiser.io/docs/cloud-providers/aws#local-komiser-cli-single-account?utm_source=github&utm_medium=social) of Komiser CLI to you AWS account

### Deploy Komiser to single account access EKS cluster (Helm chart)
* If you want to connect a single AWS account follow the documentation [here](https://docs.komiser.io/docs/cloud-providers/aws#eks-installation-single-account?utm_source=github&utm_medium=social). 

Watch the installation [video here](https://www.youtube.com/watch?v=4veDmJpui44&t)

### Deploy Komiser to a multi account access EKS cluster (Helm chart)
* If you would like to connect various AWS accounts to a Komiser deployment in a Management EKS cluster, follow the steps [here](https://docs.komiser.io/docs/cloud-providers/aws#multiple-account-eks-helm-chart-installation?utm_source=github&utm_medium=social).

## Installation on Azure

Connect a local deployment of Komiser CLI to you [**Azure**](https://docs.komiser.io/docs/cloud-providers/azure?utm_source=github&utm_medium=social) account.

## Installation on Civo

Connect a local deployment of Komiser CLI to your [**Civo**](https://docs.komiser.io/docs/cloud-providers/civo?utm_source=github&utm_medium=social) account.

## Installation on OCI

Connect a local deployment of Komiser CLI to your [**OCI**](https://docs.komiser.io/docs/cloud-providers/oci?utm_source=github&utm_medium=social) account.

## Installation on Digital Ocean

Connect a local deployment of Komiser CLI to you [**Digital Ocean**](https://docs.komiser.io/docs/cloud-providers/digital-ocean?utm_source=github&utm_medium=social) account.

## Installation on Linode

Connect a local deployment of Komiser CLI to you [**Linode**](https://docs.komiser.io/docs/cloud-providers/linode?utm_source=github&utm_medium=social) account.

## Installation on Tencent Cloud

Connect a local deployment of Komiser CLI to you [**Tencent**](https://docs.komiser.io/docs/cloud-providers/tencent?utm_source=github&utm_medium=social) account.

## Installation on Scaleway

Connect a local deployment of Komiser CLI to you [**Scaleway**](https://docs.komiser.io/docs/cloud-providers/scaleway?utm_source=github&utm_medium=social) account.

Watch the installation [video here](https://www.youtube.com/watch?v=NBbEpoW-kVs)

# Documentation 📖

Head over to the official `Komiser` documentation at [docs.komiser.io](https://docs.komiser.io?utm_source=github&utm_medium=social). The source repository for the documentation website is [tailwarden/docs](https://github.com/tailwarden/docs). 

We know that writing docs isn't usually at the top of too many peoples "What I like to do for fun" list, but if you feel so inclined, by all means, consider [contributing](https://docs.komiser.io/docs/introduction/contribute?utm_source=github&utm_medium=social) to our documentation repository, we will be very grateful. It's built using [Docusaurus](https://docusaurus.io/). 

## Jump right in:
* [Documentation overview](https://docs.komiser.io/docs/intro)
* [Installation](https://docs.komiser.io/docs/overview/introduction/getting-started)
* [FAQs](https://docs.komiser.io/docs/FAQ/)
* Video series: 📹
    * [How to: Komiser](https://www.youtube.com/watch?v=9pCimmIT-HQ&list=PLFIcIMmOFDZeMzcvOi7bPd4I6xUNq3A5R/alerts)
    * [Cloud cost savings tips](https://docs.komiser.io/docs/Quickstarts/overview)
    * [Installation videos](https://www.youtube.com/watch?v=urxi9z2IUf4&list=PLFIcIMmOFDZfaO_WmUF_qnF8akCII7Uk_) 

# Bugs and feature requests 🐞

Have a bug or a feature request? Please first read the issue guidelines and search for existing and closed issues. If your problem or idea is not addressed yet, [please open a new issue](https://github.com/tailwarden/komiser/issues).

# Roadmap and Contributing 🛣️

We are very excited about what is in store in the coming weeks and months, take a look at the [public roadmap](https://tailwarden.canny.io/) to stay on top of what's coming down the pipeline. 

Komiser is written in `Golang` and is `Elv2 licensed` - contributions are always welcome whether that means providing feedback, be it through GitHub, through the `#feedback` channel on our [Discord server](https://discord.tailwarden.com) or testing existing and new features. Feel free to check out our [contributor guidelines](./CONTRIBUTING.md) and consider becoming a **contributor** today. 

### Watch! :
Learn how to contribute with this walkthrough [video](https://www.youtube.com/watch?v=Vn5uc2elcVg)

# Users 🧑‍🤝‍🧑

If you'd like to have your company represented and are using `Komiser` please give formal written permission below via a comment on this [thread](https://github.com/tailwarden/komiser/issues/76) or via email to contact@tailwarden.com.

We will need a URL to a svg or png logo, a text title and a company URL.

# Versioning 🧮

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/tailwarden/komiser/tags). 

# Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
<tbody>
<tr>
<td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/mlabouardy"><img src="https://avatars.githubusercontent.com/u/10320205?v=4" width="100px;" alt="mlabouardy" /><br /><sub><b>mlabouardy</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=mlabouardy" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/victorgaard"><img src="https://avatars.githubusercontent.com/u/13384559?v=4" width="100px;" alt="victorgaard" /><br /><sub><b>victorgaard</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=victorgaard" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/narasago"><img src="https://avatars.githubusercontent.com/u/36790706?v=4" width="100px;" alt="narasago" /><br /><sub><b>narasago</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=narasago" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/mahendrabagul"><img src="https://avatars.githubusercontent.com/u/6149190?v=4" width="100px;" alt="mahendrabagul" /><br /><sub><b>mahendrabagul</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=mahendrabagul" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/jebinjeb"><img src="https://avatars.githubusercontent.com/u/33154308?v=4" width="100px;" alt="jebinjeb" /><br /><sub><b>jebinjeb</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=jebinjeb" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/nemca"><img src="https://avatars.githubusercontent.com/u/10352850?v=4" width="100px;" alt="nemca" /><br /><sub><b>nemca</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=nemca" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/apps/dependabot"><img src="https://avatars.githubusercontent.com/in/29110?v=4" width="100px;" alt="dependabot[bot]" /><br /><sub><b>dependabot[bot]</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=dependabot[bot]" title="Code">code</a> </td></tr><tr><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/vijeyash1"><img src="https://avatars.githubusercontent.com/u/91282703?v=4" width="100px;" alt="vijeyash1" /><br /><sub><b>vijeyash1</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=vijeyash1" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/ekand"><img src="https://avatars.githubusercontent.com/u/52926724?v=4" width="100px;" alt="ekand" /><br /><sub><b>ekand</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=ekand" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/ShivamTyagi12345"><img src="https://avatars.githubusercontent.com/u/60812924?v=4" width="100px;" alt="ShivamTyagi12345" /><br /><sub><b>ShivamTyagi12345</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=ShivamTyagi12345" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/kranurag7"><img src="https://avatars.githubusercontent.com/u/81210977?v=4" width="100px;" alt="kranurag7" /><br /><sub><b>kranurag7</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=kranurag7" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/snyk-bot"><img src="https://avatars.githubusercontent.com/u/19733683?v=4" width="100px;" alt="snyk-bot" /><br /><sub><b>snyk-bot</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=snyk-bot" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/voyc-geoffrey"><img src="https://avatars.githubusercontent.com/u/108534125?v=4" width="100px;" alt="voyc-geoffrey" /><br /><sub><b>voyc-geoffrey</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=voyc-geoffrey" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/arunnvrs"><img src="https://avatars.githubusercontent.com/u/68091787?v=4" width="100px;" alt="arunnvrs" /><br /><sub><b>arunnvrs</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=arunnvrs" title="Code">code</a> </td></tr><tr><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/ChristianWitts"><img src="https://avatars.githubusercontent.com/u/156934?v=4" width="100px;" alt="ChristianWitts" /><br /><sub><b>ChristianWitts</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=ChristianWitts" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/davidjeddy"><img src="https://avatars.githubusercontent.com/u/6232455?v=4" width="100px;" alt="davidjeddy" /><br /><sub><b>davidjeddy</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=davidjeddy" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/fpapadamis"><img src="https://avatars.githubusercontent.com/u/38500451?v=4" width="100px;" alt="fpapadamis" /><br /><sub><b>fpapadamis</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=fpapadamis" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/jakepage91"><img src="https://avatars.githubusercontent.com/u/38757612?v=4" width="100px;" alt="jakepage91" /><br /><sub><b>jakepage91</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=jakepage91" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/jazimabbas"><img src="https://avatars.githubusercontent.com/u/44744039?v=4" width="100px;" alt="jazimabbas" /><br /><sub><b>jazimabbas</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=jazimabbas" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/jbleduigou"><img src="https://avatars.githubusercontent.com/u/1489214?v=4" width="100px;" alt="jbleduigou" /><br /><sub><b>jbleduigou</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=jbleduigou" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/omerh"><img src="https://avatars.githubusercontent.com/u/1240184?v=4" width="100px;" alt="omerh" /><br /><sub><b>omerh</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=omerh" title="Code">code</a> </td></tr><tr><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/fossabot"><img src="https://avatars.githubusercontent.com/u/29791463?v=4" width="100px;" alt="fossabot" /><br /><sub><b>fossabot</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=fossabot" title="Code">code</a> </td><td align="center" valign="top" width="14.285714285714286%"><a href="https://github.com/piercypixel"><img src="https://avatars.githubusercontent.com/u/13911621?v=4" width="100px;" alt="piercypixel" /><br /><sub><b>piercypixel</b></sub></a><br /><a href="https://github.com/tailwarden/komiser/commits?author=piercypixel" title="Code">code</a> </td></tr>
</tbody>
<tfoot>
<tr>
<td align="center" size="13px" colspan="7">
</img>
</td>
</tr>
</tfoot>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->