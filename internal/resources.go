package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/aws"
	"github.com/tailwarden/komiser/providers/azure"
	"github.com/tailwarden/komiser/providers/civo"
	"github.com/tailwarden/komiser/providers/digitalocean"
	"github.com/tailwarden/komiser/providers/gcp"
	"github.com/tailwarden/komiser/providers/k8s"
	"github.com/tailwarden/komiser/providers/linode"
	"github.com/tailwarden/komiser/providers/mongodbatlas"
	"github.com/tailwarden/komiser/providers/oci"
	"github.com/tailwarden/komiser/providers/ovh"
	"github.com/tailwarden/komiser/providers/scaleway"
	"github.com/tailwarden/komiser/providers/tencent"
)

func fetchResources(ctx context.Context, clients []providers.ProviderClient, regions []string, telemetry bool) {
	numWorkers := 64
	wp := providers.NewWorkerPool(numWorkers)
	wp.Start()

	var wwg sync.WaitGroup
	workflowTrigger := func(client providers.ProviderClient, provider string) {
		wwg.Add(1)
		go func() {
			defer wwg.Done()
			triggerFetchingWorkflow(ctx, client, provider, telemetry, regions, wp)
		}()
	}

	for _, client := range clients {
		if client.AWSClient != nil {
			workflowTrigger(client, "AWS")
		} else if client.DigitalOceanClient != nil {
			workflowTrigger(client, "DigitalOcean")
		} else if client.OciClient != nil {
			workflowTrigger(client, "OCI")
		} else if client.CivoClient != nil {
			workflowTrigger(client, "Civo")
		} else if client.K8sClient != nil {
			workflowTrigger(client, "Kubernetes")
		} else if client.LinodeClient != nil {
			workflowTrigger(client, "Linode")
		} else if client.TencentClient != nil {
			workflowTrigger(client, "Tencent")
		} else if client.AzureClient != nil {
			workflowTrigger(client, "Azure")
		} else if client.ScalewayClient != nil {
			workflowTrigger(client, "Scaleway")
		} else if client.MongoDBAtlasClient != nil {
			workflowTrigger(client, "MongoDBAtlas")
		} else if client.GCPClient != nil {
			workflowTrigger(client, "GCP")
		} else if client.OVHClient != nil {
			workflowTrigger(client, "OVH")
		}
		log.Println("Workflow triggered for client:", client.Name)
	}

	wwg.Wait()
	wp.Wait()
}

func triggerFetchingWorkflow(ctx context.Context, client providers.ProviderClient, provider string, telemetry bool, regions []string, wp *providers.WorkerPool) {
	localHub := sentry.CurrentHub().Clone()

	defer func() {
		err := recover()
		if err != nil {
			log.WithField("err", err).Error(fmt.Sprintf("error fetching %s resources", provider))
			localHub.CaptureException(err.(error))
			localHub.Flush(2 * time.Second)
		}
	}()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("provider", provider)
	})

	if telemetry {
		analytics.TrackEvent("fetching_resources", map[string]interface{}{
			"provider": provider,
		})
	}

	switch provider {
	case "AWS":
		aws.FetchResources(ctx, client, regions, db, telemetry, analytics, wp)
	case "DigitalOcean":
		digitalocean.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "OCI":
		oci.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Civo":
		civo.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Kubernetes":
		k8s.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Linode":
		linode.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Tencent":
		tencent.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Azure":
		azure.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Scaleway":
		scaleway.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "MongoDBAtlas":
		mongodbatlas.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "GCP":
		gcp.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "OVH":
		ovh.FetchResources(ctx, client, db, telemetry, analytics, wp)
	}
}
