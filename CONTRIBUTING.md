We are always thrilled to receive pull requests, and do our best to process them as fast as possible. Not sure if that typo is worth a pull request? Do it! We will appreciate it.

`Note: If your pull request is not accepted on the first try, don’t be discouraged! If there’s a problem with the implementation, hopefully you received feedback on what to improve.`

## Guidelines 

We recommend discussing your plans on our Discord (join our <a href="https://discord.tailwarden.com">community server</a>) head to the `contributors` section before starting to code - especially for more ambitious contributions. This gives other contributors a chance to point you in the right direction, give feedback on your design, and maybe point out if someone else is working on the same thing.

Any significant improvements should be documented as a github issue before anybody starts working on it. Please take a moment to check that an issue doesn’t already exist documenting your bug report or improvement proposal. If it does, it never hurts to add a quick “+1” or “I have this problem too”. This will help prioritize the most common problems and requests

Feel free to communicate through the `#feedback` and `#feature-request` Discord channels. If we identify you as a `contributor`, we will add you to a private `#contributors` channel, to expedite internal communication (Hope to see you there!).

## Conventions

Fork the repo and make changes on your fork in a feature branch based on the master branch:

- If it’s a bugfix branch, name it fix/XXX-something where XXX is the number of the issue
- If it’s a feature branch, create an enhancement issue to announce your intentions, and name it feature/XXX-something where XXX is the number of the issue.
- Submit unit tests for your changes. Go has a great test framework built in; use it! Take a look at existing tests for inspiration. Run the full test suite on your branch before submitting a pull request.
- Make sure you include relevant updates or additions to documentation when creating or modifying features.
- Write clean code. Universally formatted code promotes ease of writing, reading, and maintenance. Always run go fmt before committing your changes. Most editors have plugins that do this automatically.
- While submitting Pull Request, Always remember to change the base branch from <a href="https://github.com/tailwarden/komiser/tree/master">master</a> to <a href="https://github.com/tailwarden/komiser/tree/develop">develop</a>. This will keep your Pull Request away from conflicts. **Master brach always reflects the releases and major fixes, So that it can be used by the end users.** 

## How to add a new cloud provider?

1. Create a `provider.go` under `providers/provider` folder with the following content:

```go
package aws

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

2. Add provider SDK client in `providers/providers.go`
3. Include provider configuration in TOML format under `config.toml`

## How to add a new cloud service/resource?

The process to follow for adding a new cloud service is:

1. Create a new folder under the `providers/providername/service` path called `servicename.go`
2. Inside the new file, add the following:

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

3. Call the function from `providers/providername/providername.go` by adding `MyServiceResources()` to `listOfSupportedServices()` function.

## General tips

A few important things to note when adding functions that call the cloud provider API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch all the resources.
- Make sure the resource has a tags JSON column (if possible). Sometimes this requires additional SDK calls.
- Code is required to be formatted using gofmt, this covers most code style requirements. It is also highly recommended to use goimports to automatically order imports.
- Please try to keep lines length under 80 characters, the exact number of characters is not strict but it generally helps with readability.

## How can I contribute to Komiser dashboard?

* Clone the project
* Install Go dependencies:

```
go mod download
```

* Switch to dashboard folder and install npm dependencies:

```
cd dashboard
npm install
```

* Deploy to a local server:

```
npm run dev
```

* Once you implemented your frontend changes, build the artifact, build it as golang assets:

```
go-bindata-assetfs -o template.go dist/ dist/assets/images/
```

## About Dependency graph feature

A dependency graph is a graph that aims to show the relationship that exists between individual cloud resources. This feature would help users of Komiser get a better view of their running resources.

### Approach we are following and work done by now

We have decided to divide the entire feature into two different parts, i.e., frontend and backend.

- Backend: The role of the backend is to fetch resources along with their list of relationships that exist. For now, the relationships are manually found, i.e., for individual resources, individual functions are written where relationships are found and returned along with resource data. The frontend can get this data by calling URL/resources/relations`.
- Frontend: We managed to prepare the skeleton for the feature in the frontend.

### TODO 

For **Backend**:
Support for more resources. For now, the Komiser backend only returns resources from AWS, specifically Instances and Elastic IPs with relations.
- To do so, a function needs to be created at the individual resource level that populates the `Relation` array field of the return value `resources` with relationships. For example, refer to AWS Instances and ElasticIp code.
- Logic written for fetching relations is currently manual work and can be improved if there is a way to do that automatically.
- Once we have the Relation field populated, the backend will handle the rest of things on its own.

For **Frontend**:
The initial attempt for the graph was to use [ReactFlow](https://reactflow.dev/docs/quickstart/) to represent resources as nodes and their edges as their relationships with other resources. However, it was found that scalability was a concern, and we were also a bit restricted when it came to UI.

Another alternative that came to discussion was libraries like [d3-force](https://github.com/d3/d3-force), [reaflow](https://github.com/reaviz/reaflow), [sigma js](https://github.com/jacomyal/sigma.js) etc. 

[React-Sigma](https://github.com/sim51/react-sigma) seems to fulfil all our requirements. Still, the work is in progress. In order to contribute to the frontend part, trying react-sigma or any other library of choice would be the best way to start. Make sure to consider the following things while working with a library:

- The library supports large number of nodes and edges
- Positioning of nodes are in circular form and edges are directed
- Nodes and edges UI are highly customizable
