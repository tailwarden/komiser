import { Component, OnInit } from '@angular/core';
import { StoreService } from '../../services/store.service';
import { Subscription } from 'rxjs';
import { AwsService } from '../../services/aws.service';
import { DigitaloceanService } from '../../services/digitalocean.service';
import { CloudService } from '../../services/cloud.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';
import { GcpService } from '../../services/gcp.service';

@Component({
    selector: 'app-inventory',
    templateUrl: './inventory.component.html',
    styleUrls: ['./inventory.component.css'],
})
export class InventoryComponent implements OnInit {
    public provider: string;
    public _subscription: Subscription;
    public services: Array<any> = new Array<any>();
    public selectedResources: Array<any> = new Array<any>();
    public accounts: Array<any> = new Array<any>();
    public term: string = '';
    public regions: Set<any> = new Set<any>();
    public currentPage: number = 1;
    public itemsPerPage: number = 10;

    constructor(
        private storeService: StoreService,
        private awsService: AwsService,
        private cloudService: CloudService,
        private digitalOceanService: DigitaloceanService,
        private gcpService: GcpService,
    ) {
        this.cloudService.getCloudAccounts().subscribe((accounts) => {
            this.accounts = accounts;
            if (this.accounts) {
                if (this.accounts['AWS']) {
                    this.accounts['AWS'].forEach((account) => {
                        this.getAWSResources(account);
                    });
                }
                if (this.accounts['DIGITALOCEAN']) {
                    this.accounts['DIGITALOCEAN'].forEach((account) => {
                        this.getDigitalOceanResources(account);
                    });
                }
                if (this.accounts['GCP']) {
                    this.accounts['GCP'].forEach((account) => {
                        this.getGCPResources(account);
                    });
                }
            }
        });
    }

    private getGCPResources(account) {
        this.gcpService.getComputeInstances().subscribe(data => {
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
        })
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();

            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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

            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
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
            this.getRegions();
            this.selectedResources = this.services.slice(0, 10);
        });
    }

    public cleanSelection() {
        this.term = '';
    }

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

    public getRegions() {
        let tempRegions = new Set<any>();
        this.services.forEach((service) => {
            tempRegions.add(service.region);
        });
        this.regions = tempRegions;
    }

    public getCloudAccounts() {
        let count = 0;
        Object.keys(this.accounts).forEach((provider) => {
            count += this.accounts[provider]?.length;
        });
        return count;
    }

    public pageChanged(event: PageChangedEvent): void {
        this.cleanSelection();
        const startItem = (event.page - 1) * event.itemsPerPage;
        const endItem = event.page * event.itemsPerPage;
        this.selectedResources = this.services.slice(startItem, endItem);
    }

    public changeSearchFilter(term) {
        this.selectedResources = this.services.filter((service) => {
            return (
                service.region.toLowerCase().includes(term.toLowerCase()) ||
                service.account.toLowerCase().includes(term.toLowerCase()) ||
                service.provider.toLowerCase().includes(term.toLowerCase()) ||
                service.service.toLowerCase().includes(term.toLowerCase()) ||
                service.name.toLowerCase().includes(term.toLowerCase()) ||
                service.tags?.filter((tag) => {
                    tag.toLowerCase().includes(term.toLowerCase());
                }).length > 0
            );
        });
        this.selectedResources = this.selectedResources.slice(0, 10);
    }

    ngOnInit(): void {}
}
