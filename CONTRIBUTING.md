# Komiser Contributing Guide

## Welcome

🎉 **Want to contribute to Komiser, but not sure where to start?** We've got you covered!

**Welcome to the Komiser Contributing Guide** where you'll find everything you need to get started with contributing to Komiser.

We are thrilled to have you as part of our community and looking forward to your valuable contributions! 🌟

## Before You Get Started!

Before getting started with your first contribution, here are a few things to keep in mind:

### For Major Feature Enhancements

🚀 **Planning to work on adding a major feature enhancement?** That's amazing we always encourage ambitious contributions. **Keep in mind though that we always recommend discussing your plans with the team first.**

This provides an opportunity for fellow contributors to - **guide you in the right direction**, **offer feedback on your design**, and **potentially identify if another contributor is already working on a similar task**.

Here's how to reach out:

- Join the **#contributors** channel on our [Discord Server](https://discord.tailwarden.com) to start the discussion.
- Share your idea(s) in the **#feature-request** channel and get feedback from the community.

### For Bug Reports and Minor Fixes

🐞 **Found a bug and wish to fix it?** Each and every contribution matters regardless of its size!

Here are a few things to keep in mind in this case:

- Before creating a new issue, please make sure to check for an already [existing issue](https://github.com/tailwarden/komiser/issues) for that bug.
- If an issue doesn't exist, [create a new issue](https://github.com/tailwarden/komiser/issues/new/choose) of **\`type: Bug Report\`** and make sure to provide as much detail as possible.
- Feel free to reach out to us on **#contributors** or **#feedback** channel on our Discord Server, if you need any assistance or have questions.

## General Contribution Flow

This section covers the **general contribution flow** when contributing to any area of **Komiser (Engine or the UI)** and some best practices to follow along the way.

Following these steps will ensure that your contributions are well-received, reviewed, and integrated effectively into Komiser's codebase.

### Fork and Pull Request Flow

1️⃣ Head over to the [Komiser GitHub repo](https://github.com/tailwarden/komiser) and "fork it" into your own GitHub account.

2️⃣ Clone your fork to your local machine, using the following command:
```bash
git clone git@github.com:USERNAME/FORKED-PROJECT.git
```

3️⃣ Create a new branch based-off **\`develop\`** branch:
```bash
git checkout develop
git checkout -b fix/XXX-something develop
```
Make sure to follow the following branch naming conventions:
- For feature/enchancements, use: **\`feature/xxxx-name_of_feature\`**
- For bug fixes, use: **\`fix/xxxx-name_of_bug\`**
> Here, **\`xxxx\`** is the issue number associated with the bug/feature!
> For example:
> ```bash
> git checkout -b feature/1022-kubecost-integration develop
> ```

4️⃣ Implement the changes or additions you intend to contribute. Whether it's **bug fixes**, **new features**, or **enhancements**, this is where you put your coding skills to use.

5️⃣ Once your changes are ready, you may then commit and push the changes from your working branch:
```bash
git commit -m "fix(xxxx-name_of_bug): nice commit description"
git push origin feature/1022-kubecost-integration
```

6️⃣ While submitting Pull Request, **make sure to change the base branch from**: [master](https://github.com/tailwarden/komiser/tree/master) to [develop](v). Making sure to avoid any possible merge conflicts

### Keeping your Fork Up-to-Date

If you plan on doing anything more than just a tiny quick fix, you’ll want to **make sure you keep your fork up to date** by tracking the original [“upstream” repo](https://github.com/tailwarden/komiser) that you forked.

Follow the steps given below to do so:

1️⃣ Add the 'upstream' repo to list of remotes:
```bash
git remote add upstream https://github.com/tailwarden/komiser.git
```

2️⃣ Fetch upstream repo’s branches and latest commits:
```bash
git fetch upstream
```

3️⃣ Checkout to the **\`develop\`** branch and merge the upstream:
```bash
git checkout develop
git merge upstream/develop
```

**Now, your local 'develop' branch is up-to-date with everything modified upstream!**


# 🚀 Contributing to Komiser Engine

The core Komiser Engine is written in Go (Golang) and leverages Go Modules. Here are the prerequisites to run Komiser on your local machine:

1. 🌐 **Go Version**:
   - Latest Go version (currently its **`1.19`**) must be installed if you want to build and/or make changes to the existing code. The binary **`go1.19`** should be available in your path.
     > 💡 If you wish to keep multiple versions of Go in your system and don't want to disturb the existing ones, refer [the guide](https://go.dev/doc/manage-install).

2. 🔧 **GOPATH**:
   - Ensure that the **`GOPATH`** environment variable is configured appropriately.


## 🌟 Ways to Contribute to Komiser Engine

Komiser is an open-source cloud-agnostic resource manager. It helps you break down cloud resource costs at the resource level. As a cloud-agnostic cloud management tool, we always have more providers and cloud services to add, update, and cost-calculate.


### ☁️ Adding a new Cloud Provider

#### 1️⃣ Create provider.
Create `provider_name.go` in `providers/provider_name` directory.

#### 2️⃣  Add the following boilerplate:
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

#### 3️⃣ Add SDK client details:
Add your client details to [**`providers/provider.go`**](https://github.com/tailwarden/komiser/blob/develop/providers/providers.go)

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

#### 4️⃣ Add provider configuration:
Add provider configuration in TOML format in **`config.toml`**

```toml
[[gcp]]
name="production"
source="ENVIRONMENT_VARIABLES"
# path="./path/to/credentials/file" specify if 'CREDENTIALS_FILE' is used as source
profile="production"
```

#### 5️⃣ Update Dashboard Utils
Navigate to the `dashboard/utils` folder in your project and locate the file named `servicehelper.ts`. Open the file and follow these steps:

##### a. Add Provider to Union Type
Locate the `Providers` union type at the top of the file. Add the new cloud provider in lowercase to the union.

```typescript
// dashboard/utils/servicehelper.ts

export type Providers =
  | 'aws'
  | 'gcp'
  | 'digitalocean'
  | 'azure'
  | 'civo'
  | 'kubernetes'
  | 'linode'
  | 'tencent'
  | 'oci'
  | 'scaleway'
  | 'mongodbatlas'
  | 'ovh'
  | 'scaleway'
  | 'tencent'
  | 'provider_name'; // Add the new provider here

```
##### b. Add Provider and Service to `allProvidersServices`

Find the `allProvidersServices` object in the `servicehelper.ts` file and add the new provider along with the service in lowercase.

```typescript
// dashboard/utils/servicehelper.ts

export const allProvidersServices: { [key in Providers | string]: string[] } = {
  // ... other services
  new_provider: ['new_service_name'],
};
```

#### 6️⃣ Compile a new Komiser binary:
```bash
go build
```

#### 7️⃣ Start a new Komiser development server:

```bash
./komiser start
```

### 🔋 Adding a new Cloud Service/Resource

Here are the general steps to add a new service/resource for a cloud provider in Komiser:

#### 1️⃣ Create Service
Create a new file `servicename.go` under the path `providers/provider_name/servicename`

#### 2️⃣ Add boilerplate
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

#### 3️⃣ Edit Provider
Call the `MyServiceResources()` function from the above file, by adding it to `providers/providername/provider.go` file's `listOfSupportedServices()` function.

```go
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

#### 4️⃣ Update Dashboard Utils
Navigate to the `dashboard/utils` folder in your project and locate the file named `servicehelper.ts`. Open the file and find the `allProvidersServices` object. Add the new service name in lowercase under the respective provider, adhering to the existing structure.

Here's an example of how to add the new service:

```typescript
// dashboard/utils/servicehelper.ts

export const allProvidersServices: { [key in Providers | string]: string[] } = {
  // ... other services
  your_provider_name: [
    'existing_service_1',
    'existing_service_2',
    // ... other existing services
    'new_service_name', // Add the new service here
  ],
};
```

#### 5️⃣
Do above mentioned steps [4](#4️⃣-add-provider-configuration), [5](#5️⃣-compile-a-new-komiser-binary) and [6](#6️⃣-start-a-new-komiser-development-server). You'll see a new resource/service added to Komiser, in the dashboard!

Additionally, [here](https://youtu.be/Vn5uc2elcVg?feature=shared) is a video tutorial of the entire process for your reference.

> 💡 Tip: you can also start the server via `go run *.go start --config ./config.toml` if you do want to skip the compile step!

### 3️⃣ Enhance existing Cloud service/resource

**So, you wish to improve the code quality of an existing cloud service/resource?** Feel free to discuss your ideas with us on our [Discord Server](https://discord.tailwarden.com) and [open a new issue](https://github.com/tailwarden/komiser/issues).

## 🧪 Testing Your Changes

We leverage the [testing](https://pkg.go.dev/testing) package for tests. Test names follow the `TestXxx(*testing.T)` format where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.
For creating a new test you create a `[name]_test.go` next to the file you'd like to test and replace `[name]` with your filename of the implementation. Look at any of the `*_test.go` files for an example or read the [official docs](https://pkg.go.dev/testing).
You then can run it with `go test /path/to/your/folder/where/the/test/is`. You can run all of our engine tests with `make tests`. You should see something similar to this:

```logtalk
go test ./... | grep -v /dashboard/
...
ok  	github.com/tailwarden/komiser/internal	(cached) [no tests to run]
?   	github.com/tailwarden/komiser/providers/aws/ecr	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/ecs	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/efs	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/eks	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/elasticache	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/elb	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/iam	[no test files]
ok  	github.com/tailwarden/komiser/providers/aws/ec2	(cached)
?   	github.com/tailwarden/komiser/providers/aws/kms	[no test files]
?   	github.com/tailwarden/komiser/providers/aws/lambda	[no test file
...
```

# 🚀 Contributing to Komiser Dashboard UI

Komiser Dashboard utilizes a modern tech stack. Here's a brief about it:

- **Framework**: [**Next.js**](https://nextjs.org/)
- **Language**: [**Typescript**](https://www.typescriptlang.org/)
- **CSS**: [**Tailwind**](https://tailwindcss.com/)
- **Component Library**: [**Storybook**](https://storybook.js.org/docs/react/get-started/why-storybook)
- **Testing**:
  - [**Jest**](https://jestjs.io/)
  - [**React Testing Library**](https://testing-library.com/docs/react-testing-library/intro/)

## 🧩 Prerequisites

1. While working on Komiser UI, you'd typically need the Komiser Server as well. Follow the instructions in the [**Komiser Installation**](#komiser-installation) section above.
2. Ensure you're using the **latest versions** of the tech stack mentioned above.

## 🛠 Setup a Local Development Server

Let's get your hands dirty by setting up the Komiser dashboard:

### 1️⃣ Grab the Go Dependencies
```bash
go mod download
```

### 2️⃣ Configure `config.toml`

If it doesn't exist, create `config.toml` in the root and use this template:

```toml
[[aws]]
	name="sandbox"
	source="CREDENTIALS_FILE"
	path="./path/to/credentials/file"
	profile="default"

[[aws]]
	name="staging"
	source="CREDENTIALS_FILE"
	path="./path/to/credentials/file"
	profile="staging-account"

[[gcp]]
	name="production"
	source="ENVIRONMENT_VARIABLES"
	# path="./path/to/credentials/file" specify if CREDENTIALS_FILE is used
	profile="production"

[sqlite]
  file="komiser.db"
```

> 📘 Dive deeper in our [Quickstart Guide](https://docs.komiser.io/getting-started/quickstart) for more configurations like connecting to PostgreSQL.

Now, craft your credentials. Check out the guides [here](https://docs.komiser.io/getting-started/quickstart#self-hosted) to tailor the config for your needs.

### 3️⃣ Boot Up the Komiser Backend
```bash
go run *.go start --config ./config.toml
```

> 🖥️ This fires up the dashboard at [`http://localhost:3002/`](http://localhost:3002). However, it mirrors the **`master`** branch. Let's make it reflect your development branch!

### 4️⃣ Navigate to the Dashboard Directory
```bash
cd dashboard
```

### 5️⃣ Set up Environment Variables
Create or update the **`.env`** file:
```
NEXT_PUBLIC_API_URL=http://localhost:3000
```

### 6️⃣ Spin up the Dev Server
Install dependencies and fire up the dev server:
```bash
npm install
npm run dev
```

> 🟢 NodeJS: Ensure you're on the `18.x.x` LTS release.

Once done, open up [**`http://localhost:3002/`**](http://localhost:3002)

![Komiser Dashboard](https://hackmd.io/_uploads/ryvOPmFla.png)

📺 For a more detailed walkthrough, check the [video tutorial](https://youtu.be/uwxj11-eRt8?feature=shared).

## 🎨 Understanding the UI Components

Komiser's UI elegance is sculpted using [Storybook](https://storybook.js.org/). Dive deep into the [components section](https://github.com/tailwarden/komiser/tree/develop/dashboard#components) to grasp our conventions.

## 🧪 Testing Your Changes

We leverage **Jest** and **React Testing Library** for unit tests. Navigate to the [testing section](https://github.com/tailwarden/komiser/tree/develop/dashboard#testing) for insights on writing tests.

## 📦 Building a Go Artifact

After refining the frontend, craft a new Go artifact:
```
go-bindata-assetfs -o template.go dist/ dist/assets/images/
```


# 📚 Contributing Best Practices

> 📘 For frontend endeavors, do explore the [Getting Started Guide](https://github.com/tailwarden/komiser/tree/develop/dashboard#getting-started) in the `/dashboard` README.

Embarking on backend development? Here are golden practices to streamline your contributions and facilitate our review process:

### 1. 🧪 Write Comprehensive Unit Tests
- Every code change craves well-thought-out unit tests.
- Leverage [Go's built-in testing framework](https://pkg.go.dev/testing) for this.
- Seek inspiration from tests already enhancing the project.
- 🚀 Before firing that pull request, run the entire test suite on your branch ensuring robust validation of your changes.

### 2. 📖 Keep Documentation Updated
- Newly added or tinkered features? Ensure the documentation mirrors your magic.

### 3. 🎨 Prioritize Clean Code
- Seal your changes with **`go fmt`** for a consistent code style.
- Tools like **VSCode** come handy with plugins automating this formatting quest.

### 4. 🖊 Mindful Code Comments
- Complex logic? Esoteric algorithms? Or just proud of your intricate code snippet? Comment them.
- Thoughtful comments pave way for an insightful review, letting others decipher your masterpiece with ease.

---

# 🌟 Ending Note

We trust this guide lights up your contribution path. Diving into Komiser's codebase should now be thrilling and engaging.

In wrapping up, a **MASSIVE THANK YOU** echoes for your invaluable time and efforts. You're making Komiser even more radiant and welcoming for the community.

> 📞 Need a chat? Have queries buzzing? Buzz us on our [Discord Server](https://discord.tailwarden.com).

---

# 📜 License

Komiser, an emblem of open-source, basks under the [Elastic License 2.0 (ELv2)](https://github.com/tailwarden/komiser/blob/develop/LICENSE).
