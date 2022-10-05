import { Component, OnInit } from '@angular/core';
import { StoreService } from '../../services/store.service';
import { Subscription } from 'rxjs';
import { AwsService } from '../../services/aws.service';
import { CloudService } from '../../services/cloud.service';

@Component({
    selector: 'app-inventory',
    templateUrl: './inventory.component.html',
    styleUrls: ['./inventory.component.css'],
})
export class InventoryComponent implements OnInit {
    public provider: string;
    public _subscription: Subscription;
    public services: Array<any> = new Array<any>();
    public accounts: Array<any> = new Array<any>();
    public term: string = '';
    public regions: Set<any> = new Set<any>();

    constructor(
        private storeService: StoreService,
        private awsService: AwsService,
        private cloudService: CloudService
    ) {
        this.cloudService.getCloudAccounts().subscribe(accounts => {
            this.accounts = accounts;
            if(this.accounts) {
                if (this.accounts['AWS']){
                    this.accounts['AWS'].forEach(account => {
                        this.getAWSResources(account);
                    })
                }
            }
        })
    }

    private getAWSResources(account){
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

                this.getRegions();
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

                this.getRegions();
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

                this.getRegions();
            });
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
          var value = (hash >> (i * 8)) & 0xFF;
          colour += ('00' + value.toString(16)).substr(-2);
        }
        return colour;
    }

    public getRegions(){
        let tempRegions = new Set<any>();
        this.services.forEach(service => {
            tempRegions.add(service.region)
        })
        this.regions = tempRegions;
    }

    public getCloudAccounts(){
        let count = 0;
        Object.keys(this.accounts).forEach(provider => {
            count += (this.accounts[provider]?.length)
        })
        return count;
    }

    ngOnInit(): void {}
}
