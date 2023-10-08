# Komiser Contributing Guide

## Welcome

**Want to contribute to Komiser, but not sure where to start?** We've got you covered!

**Welcome to the Komiser Contributing Guide** where you'll find everything you need to get started with contributing to Komiser.

We are thrilled to have you as part of our community and looking forward to your valuable contributions!

## Before You Get Started!

Before getting started with your first contribution, here are a few things to keep in mind:

### For Major Feature Enhancements

**Planning to work on adding a major feature enhancement?** That's amazing we always encourage ambitious contributions. **Keep in mind though that we always recommend discussing your plans with the team first.**

This provides an opportunity for fellow contributors to - **guide you in the right direction**, **offer feedback on your design**, and **potentially identify if another contributor is already working on a similar task**. 

Here's how to reach out:

- Join the **#contributors** channel on our [Discord Server](https://discord.tailwarden.com) to start the discussion.
- Share your idea(s) in the **#feature-request** channel and get feedback from the community.

### For Bug Reports and Minor Fixes

**Found a bug and wish to fix it?** Each and every contribution matters regardless of its size!

Here are a few things to keep in mind in this case:

- Before creating a new issue, please make sure to check for an already [existing issue](https://github.com/tailwarden/komiser/issues) for that bug.
- If an issue doesn't exist, [create a new issue](https://github.com/tailwarden/komiser/issues/new/choose) of **`type: Bug Report`** and make sure to provide as much detail as possible.
- Feel free to reach out to us on **#contributors** or **#feedback** channel on our Discord Server, if you need any assistance or have questions.

## General Contribution Flow

This section covers the **general contribution flow** when contributing to any area of **Komiser (Engine or the UI)** and some best practices to follow along the way.

Following these steps will ensure that your contributions are well-received, reviewed, and integrated effectively into Komiser's codebase. 

### Fork and Pull Request Flow

1. Head over to the [Komiser GitHub repo](https://github.com/tailwarden/komiser) and "fork it" into your own GitHub account.
2. Clone your fork to your local machine, using the following command:
    ```bash
    git clone git@github.com:USERNAME/FORKED-PROJECT.git 
    ```
3. Create a new branch based-off **`develop`** branch:
    ```bash
    git checkout develop
    git checkout -b fix/XXX-something develop
    ```
    Make sure to follow the following branch naming conventions:
    - For feature/enchancements, use: **`feature/xxxx-name_of_feature`**
    - For bug fixes, use: **`fix/xxxx-name_of_bug`** <br><br>

    > Here, **`xxxx`** is the issue number associated with the bug/feature!

    For example:
    ```bash
    git checkout -b feature/1022-kubecost-integration develop
    ```
4. Implement the changes or additions you intend to contribute. Whether it's **bug fixes**, **new features**, or **enhancements**, this is where you put your coding skills to use.
5. Once your changes are ready, you may then commit and push the changes from your working branch:
    ```bash
    git commit -m "nice_commit_description"
    git push origin feature/1022-kubecost-integration
    ```
6. While submitting Pull Request, **make sure to change the base branch from**: [master](https://github.com/tailwarden/komiser/tree/master) to [develop](v). Making sure to avoid any possible merge conflicts

> ### Keeping your Fork Up-to-Date
> 
> If you plan on doing anything more than just a tiny quick fix, you’ll want to **make sure you keep your fork up to date** by tracking the original [“upstream” repo](https://github.com/tailwarden/komiser) that you forked. 
>
> Follow the steps given below to do so: 
> 
> 1. Add the 'upstream' repo to list of remotes:
> 
> ```bash
> git remote add upstream https://github.com/tailwarden/komiser.git
> ```
> 
> 2. Fetch upstream repo’s branches and latest commits:
> 
> ```bash
> git fetch upstream
> ```
> 
> 3. Checkout to the **`develop`** branch and merge the upstream:
> 
> ```bash
> git checkout develop
> git merge upstream/develop
> ```
> 
> **Now, your local 'develop' branch is up-to-date with everything modified upstream!**

## Contributing to Komiser Engine

The core Komiser Engine is written in Go (Golang) and leverages Go Modules. Following are the pre-requisistes needed to run Komiser on your local machine:
1. Latest Go version (currently its **`1.19`**) must be installed if you want to build and/or make changes to the existing code. The binary **`go1.19`** should be available in your path. <br><br>
    > If you wish to keep multiple versions of Go in your system and don't want to disturb the existing ones, refer [the guide](https://go.dev/doc/manage-install).
2. Make sure that the **`GOPATH`** environment variable is configured appropriately. 

### Komiser Installation

**Step 1: Installing Komiser CLI**

Follow the instructions given in the [documentation](https://docs.komiser.io/getting-started/installation) to install the **Komiser CLI**, according to your operating system.

**Step 2: Connect to a Cloud Account** 

In order to deploy a **self-hosted (local) instance** of Komiser, the next step would be to connect your Komiser CLI to a cloud account of your choice. You may refer the documentation of the [supported cloud providers](https://docs.komiser.io/configuration/cloud-providers/aws) and follow the instructions using any one (Let's say AWS).

**Step 3: Accessing the Komiser UI**

Once the local Komiser instance is running, you can access the dashboard UI on **`http://localhost:3000`**

![komiser-dashboard](https://hackmd.io/_uploads/Syo0bMtgT.png)

### Ways to Contribute in Komiser Engine


Komiser is an open source cloud-agnostic resource manager which helps you break down your cloud resources cost at the resource level.

Due to the nature of Komiser, a cloud-agnostic cloud management tool, our work is never really done! There are always more providers and cloud services that can be added, updated and cost calculated.

Therefore, there are mainly three ways you can contribute to the Komiser Engine:

### 1. Adding a new Cloud Provider

Here are the general steps to integrate a new cloud provider in Komiser:

**Step 1:** 
Create a new **`provider_name.go`** file under **`providers/provider_name`** directory.

**Step 2:** 
Add the following boilerplate code, which defines the structure of any new provider to be added:

```go
package PROVIDER_NAME

import (
	"context"
	"log"

	. "github.com/tailwarden/komiser/providers"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []FetchDataFunction {
	return []FetchDataFunction{}
}

func FetchProviderData(ctx context.Context, client ProviderClient, db *bun.DB) {
	for _, function := range listOfSupportedServices() {
		resources, err := function(ctx, client)
		if err != nil {
			log.Printf("[%s][PROVIDER] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}
}
```

Then, the main task is writing the code to fetch all the resources/services using the provider's Go client SDK. You may refer any [existing examples](https://github.com/tailwarden/komiser/tree/develop/providers) to understand better.

**Step 3:**
Add the information about the appropriate provider's SDK client in [**`providers/provider.go`**](https://github.com/tailwarden/komiser/blob/develop/providers/providers.go) file:

```go

type ProviderClient struct {
	AWSClient          *aws.Config
	DigitalOceanClient *godo.Client
	OciClient          common.ConfigurationProvider
	CivoClient         *civogo.Client
	K8sClient          *K8sClient
	LinodeClient       *linodego.Client
	TencentClient      *tccvm.Client
	AzureClient        *AzureClient
	ScalewayClient     *scw.Client
	MongoDBAtlasClient *mongodbatlas.Client
	GCPClient          *GCPClient
	Name               string
}

type AzureClient struct {
	Credentials    *azidentity.ClientSecretCredential
	SubscriptionId string
}

```

**Step 4:**
Add the provider configuration in TOML format in your **`config.toml`** file, which will be used by Komiser to configure your account with the CLI.

An example configuration entry for configuring a Google Cloud account in the **`config.toml`** file would look like this:
```
[[gcp]]
name="production"
source="ENVIRONMENT_VARIABLES"
# path="./path/to/credentials/file" specify if 'CREDENTIALS_FILE' is used as source
profile="production"
```

**Step 5:**
Build a new Komiser binary with the latest code changes by running:

```
go build
```

**Step 6:**
Start a new Komiser development server using this new binary:

```
./komiser start
```

**If everything goes well, you'll see a new cloud provider added in the Komiser Dashboard!**

### 2. Adding a new Cloud Service/Resource

Here are the general steps to add a new service/resource for a cloud provider in Komiser:

**Step 1:**
Create a new file **`servicename.go`** under the path **`providers/provider_name/servicename`**

**Step 2:**
Add the following boilerplate code, which defines the structure of any new service/resource to be added for a cloud provider:

```go
package service

import (
	"context"
	"log"
	"time"

	. "github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/providers"
)

func MyServiceResources(ctx context.Context, client ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)
	
    // Logic goes here

	log.Printf("[%s] Fetched %d CLOUD PROVIDER NAME resources\n", client.Name, len(resources))
	return resources, nil
}
```

To understand how to write the required logic, you may refer any [existing examples](https://github.com/tailwarden/komiser/tree/develop/providers/aws) for inspiration!

**Step 3:**
Call the **`MyServiceResources()`** function from the above file, by adding it to **`providers/providername/provider.go`** file's **`listOfSupportedServices()`** function. 

```
func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		ec2.Instances,
		ec2.ElasticIps,
		lambda.Functions,
		ec2.Acls,
		ec2.Subnets,
		ec2.SecurityGroups,
		ec2.AutoScalingGroups,
		ec2.InternetGateways,
		iam.Roles,
		iam.InstanceProfiles,
		iam.OIDCProviders,
		iam.SamlProviders,
		iam.Groups,
.
.
.
```

**Step 4:**
Repeat steps **`4,5,6`** accordingly and you'll see a new resource/service added to Komiser, in the dashboard!

Additionally, [here](https://youtu.be/Vn5uc2elcVg?feature=shared) is a video tutorial of the entire process for your reference.

#### 3. Enhance existing Cloud service/resource

**So, you wish to improve the code quality of an existing cloud service/resource?** Feel free to discuss your ideas with us on our [Discord Server](https://discord.tailwarden.com) and [open a new issue](https://github.com/tailwarden/komiser/issues).

## Contributing to Komiser Dashboard UI

Komiser Dashboard is built on **Typescript** and **Next.js**. The entire frontend stack used is as follows:
- **Next.js**
- **Typescript**
- **Tailwind**
- **Storybook**
- **Jest** 
- **React Testing Library**

Following are the pre-requisites needed to setup a Dev environment of Komiser dashboard:
- In nearly all cases, while contributing to Komiser UI, you will need to build and run the Komiser Server as well, using the CLI. Make sure to follow the steps mentioned in the **"Komiser Installation"** section above.
- Make sure to have all the **latest versions** of the frontend stack listed above.

### Setup a local Developement Server

Here are the steps to setup and access the Komiser dashboard:

**Step 0:**

Install the necessary Go dependencies using the following command:
```
go mod download
```

**Step 1:**
From the root folder, start the Komiser backend server using the following command:
```
go run *.go start --config /path/to/config.toml
```

> As soon as you run this, you'll be able to access the dashboard at `http://localhost:3000`. 
>
> An important point to note here is, this dashboard only reflects the changes from the **`master`** branch.
> 
> For our purpose, we certainly need changes to be reflected from our development branch! 
> Follow the steps given below to do so 👇
>

**Step 2:**
Head over to the **`dashboard`** directory:

```
cd dashboard
```

**Step 3:**
Create a new environment variable in the **`.env`** file:
```
EXT_PUBLIC_API_URL=http://localhost:3000
```

**Step 4:**
Install the necessary **`npm`** dependencies and start the dev server using the following command:
```
npm install
npm run dev
```

You'll be able to access the dashboard at **`http://localhost:3002/`**

![](https://hackmd.io/_uploads/ryvOPmFla.png)

To understand the installation process in a bit more detail, you may refer the [video walkthrough](https://youtu.be/uwxj11-eRt8?feature=shared).

### Understanding the UI components

The Komiser UI components are being handled and organised using [Storybook](https://storybook.js.org/).
Refer the [components section](https://github.com/tailwarden/komiser/tree/develop/dashboard#components) of the dashboard README to understand the component conventions used.

### Testing Your Changes

The Komiser dashboard uses **Jest** and **React Testing Library** for unit tests. Refer the [testing section](https://github.com/tailwarden/komiser/tree/develop/dashboard#testing) of the dashboard README to understand how you can write simple unit tests, to validate your changes.

### Building a Go Artifact

Once you have implemented the necessary frontend changes, make sure to build a new Go artifact using the following command:
```
go-bindata-assetfs -o template.go dist/ dist/assets/images/
```

## Contributing Best Practices

Here are some best practices to follow during the development process to make your changes more structured and making it easier for us to review:

1. **Write Comprehensive Unit Tests:**   
        
    - When making code changes, be sure to include well-structured unit tests. 
    - Utilize [Go's built-in testing framework](https://pkg.go.dev/testing) for this purpose. 
    - Take inspiration from existing tests in the project. 
    - Before submitting your pull request, run the full test suite on your development branch to ensure your changes are thoroughly validated.

2. **Keep Documentation Updated:**

    - Ensure relevant documentation is updated or added, according to new features being added or modified. 

3. **Prioritize Clean Code:**

    - Make sure to use **`go fmt`** to ensure uniform code formatting before committing your changes. 
    - Using IDEs and code editors like **VSCode** makes this easy, as they offer plugins that automate the formatting process.

4. **Mindful Code Comments:**

    - Use comments to explain complex logic, algorithms, or any non-obvious parts of your code. 
    - Well-placed comments will make your code more accessible to others and will ultimately help in a smoother review process of your changes.


## Ending Note

We hope this guide proves to be helpful and makes contributing to Komiser an exciting and fun process for you all.

At the end, we wanna give you a **HUGE THANK YOU** for taking out your time in contributing and making Komiser better and more accessible to the community!

Feel free to reach out to us on our [Discord Server](https://discord.tailwarden.com) if you need any assistance or have any questions.

## License

Komiser is an open-source software released under the [Elastic License 2.0 (ELv2)](https://github.com/tailwarden/komiser/blob/develop/LICENSE). 