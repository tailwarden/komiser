import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { AwsService } from '../../services/aws.service';
import { CloudService } from '../../services/cloud.service';
import { DigitaloceanService } from '../../services/digitalocean.service';
import { GcpService } from '../../services/gcp.service';
import { StoreService } from '../../services/store.service';

@Component({
    selector: 'app-views',
    templateUrl: './views.component.html',
    styleUrls: ['./views.component.css'],
})
export class ViewsComponent implements OnInit {
    public view: any = {};
    public services: Array<any> = new Array<any>();
    public filteredServices: Array<any> = new Array<any>();

    constructor(
        private activatedRoute: ActivatedRoute,
        private storeService: StoreService,
        private cloudService: CloudService,
        private awsService: AwsService,
        private gcpService: GcpService,
        private digitalOceanService: DigitaloceanService
    ) {
        this.activatedRoute.paramMap.subscribe((params: Params) => {
            this.view = this.storeService.getView(params.get('id'));
            this.getFilteredServices();
        });
    }

    private getFilteredServices() {
        this.cloudService.getCloudAccounts().subscribe((accounts) => {
            if (accounts) {
                if (accounts['AWS']) {
                    accounts['AWS'].forEach((account) => {
                        this.getAWSResources(account);
                    });
                }
                if (accounts['DIGITALOCEAN']) {
                    accounts['DIGITALOCEAN'].forEach((account) => {
                        this.getDigitalOceanResources(account);
                    });
                }
                if (accounts['GCP']) {
                    accounts['GCP'].forEach((account) => {
                        this.getGCPResources(account);
                    });
                }
            }
        });
    }

    private getGCPResources(account) {
        this.gcpService.getComputeInstances().subscribe((data) => {
            data.forEach((instance) => {
                this.services.push({
                    provider: 'GCP',
                    account: account,
                    service: 'Compute Engine',
                    name: instance.name,
                    tags: instance.tags,
                    region: instance.region,
                });
            });
            this.onNewServices();
        });
    }

    private getDigitalOceanResources(account) {
        this.digitalOceanService.getDroplets().subscribe((data) => {
            data.forEach((droplet) => {
                this.services.push({
                    provider: 'DIGITALOCEAN',
                    account: account,
                    service: 'Droplet',
                    name: droplet.name,
                    tags: droplet.tags,
                    region: droplet.region,
                });
            });
            this.onNewServices();
        });

        this.digitalOceanService.getSnapshots().subscribe((data) => {
            data.forEach((snapshot) => {
                this.services.push({
                    provider: 'DIGITALOCEAN',
                    account: account,
                    service: 'Snapshot',
                    name: snapshot.name,
                    tags: snapshot.tags,
                    region: snapshot.region,
                });
            });
            this.onNewServices();
        });

        this.digitalOceanService.getVolumes().subscribe((data) => {
            data.forEach((volume) => {
                this.services.push({
                    provider: 'DIGITALOCEAN',
                    account: account,
                    service: 'Volume',
                    name: volume.name,
                    tags: volume.tags,
                    region: volume.region,
                });
            });
            this.onNewServices();
        });

        this.digitalOceanService.getDatabases().subscribe((data) => {
            data.forEach((database) => {
                this.services.push({
                    provider: 'DIGITALOCEAN',
                    account: account,
                    service: 'Database',
                    name: database.name,
                    tags: database.tags,
                    region: database.region,
                });
            });
            this.onNewServices();
        });
    }

    private getAWSResources(account) {
        this.awsService.getLambdaFunctions(account).subscribe((data) => {
            data.forEach((f) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'Lambda',
                    name: f.name,
                    tags: f.tags,
                    region: f.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getNumberOfS3Buckets(account).subscribe((data) => {
            data.forEach((bucket) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'S3',
                    name: bucket.name,
                    tags: bucket.tags,
                    region: bucket.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getVirtualPrivateClouds(account).subscribe((data) => {
            data.forEach((vpc) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'VPC',
                    name: vpc.name,
                    tags: vpc.tags,
                    region: vpc.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getRouteTables(account).subscribe((data) => {
            data.forEach((rt) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'Route Table',
                    name: rt.name,
                    tags: rt.tags,
                    region: rt.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getVPCSubnets(account).subscribe((data) => {
            data.forEach((subnet) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'Subnet',
                    name: subnet.name,
                    tags: subnet.tags,
                    region: subnet.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getSecurityGroups(account).subscribe((data) => {
            data.forEach((sg) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'Security Group',
                    name: sg.name,
                    tags: sg.tags,
                    region: sg.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getSQSQueues(account).subscribe((data) => {
            data.forEach((queue) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'SQS',
                    name: queue.name,
                    tags: queue.tags,
                    region: queue.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getDynamoDBTables(account).subscribe((data) => {
            data.forEach((table) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'DynamoDB',
                    name: table.name,
                    tags: table.tags,
                    region: table.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getInstancesPerRegion(account).subscribe((data) => {
            data.forEach((item) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'EC2',
                    name: item.id,
                    tags: item.tags,
                    region: item.region,
                });
            });

            this.onNewServices();
        });

        this.awsService.getECS(account).subscribe((data) => {
            data.services.forEach((service) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'ECS Services',
                    name: service.Name,
                    tags: service.tags,
                    region: service.region,
                });
            });
            data.tasks.forEach((task) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'ECS Tasks',
                    name: task.ARN,
                    tags: task.tags,
                    region: task.region,
                });
            });
            data.clusters.forEach((cluster) => {
                this.services.push({
                    provider: 'AWS',
                    account: account,
                    service: 'ECS Clusters',
                    name: cluster.Name,
                    tags: cluster.tags,
                    region: cluster.region,
                });
            });

            this.onNewServices();
        });
    }

    public onNewServices() {
        this.filteredServices = [];
        this.services.forEach((service) => {
            let found = false;
            service.tags?.forEach((serviceTag) => {
                this.view.tags?.forEach((tag) => {
                    if (
                        serviceTag.includes(tag.key) ||
                        serviceTag.includes(tag.value)
                    ) {
                        found = true;
                    }
                });
            });
            if (found) {
                this.filteredServices.push(service);
            }
        });
    }

    ngOnInit(): void {}

    public stringToColour(str) {
        var hash = 0;
        for (var i = 0; i < str.length; i++) {
            hash = str.charCodeAt(i) + ((hash << 5) - hash);
        }
        var colour = '#';
        for (var i = 0; i < 3; i++) {
            var value = (hash >> (i * 8)) & 0xff;
            colour += ('00' + value.toString(16)).substr(-2);
        }
        return colour;
    }
}
