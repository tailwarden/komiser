import { Component, OnInit } from '@angular/core';
import { StoreService } from '../../services/store.service';
import { Subscription } from 'rxjs';
import { AwsService } from '../../services/aws.service';

@Component({
    selector: 'app-inventory',
    templateUrl: './inventory.component.html',
    styleUrls: ['./inventory.component.css'],
})
export class InventoryComponent implements OnInit {
    public provider: string;
    public _subscription: Subscription;
    public services: Array<any> = new Array<any>();

    constructor(
        private storeService: StoreService,
        private awsService: AwsService
    ) {
        this.provider = this.storeService.getProvider();
        this._subscription = this.storeService.providerChanged.subscribe(
            (provider) => {
                console.log(provider);
                this.provider = provider;
            }
        );

        this.awsService.getLambdaFunctions().subscribe((data) => {
            data.forEach((f) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'Lambda',
                    name: f.name,
                    tags: f.tags,
                    region: f.region,
                });
            });
        });

        this.awsService.getNumberOfS3Buckets().subscribe((data) => {
            data.forEach((bucket) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'S3',
                    name: bucket.name,
                    tags: bucket.tags,
                    region: bucket.region,
                });
            });
        });

        this.awsService.getVirtualPrivateClouds().subscribe((data) => {
            data.forEach((vpc) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'VPC',
                    name: vpc.name,
                    tags: vpc.tags,
                    region: vpc.region,
                });
            });
        });

        this.awsService.getRouteTables().subscribe((data) => {
            data.forEach((rt) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'Route Table',
                    name: rt.name,
                    tags: rt.tags,
                    region: rt.region,
                });
            });
        });

        this.awsService.getVPCSubnets().subscribe((data) => {
            data.forEach((subnet) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'Subnet',
                    name: subnet.name,
                    tags: subnet.tags,
                    region: subnet.region,
                });
            });
        });

        this.awsService.getSecurityGroups().subscribe((data) => {
            data.forEach((sg) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'Security Group',
                    name: sg.name,
                    tags: sg.tags,
                    region: sg.region,
                });
            });
        });

        this.awsService.getSQSQueues().subscribe((data) => {
            data.forEach((queue) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'SQS',
                    name: queue.name,
                    tags: queue.tags,
                    region: queue.region,
                });
            });
        });

        this.awsService.getInstancesPerRegion().subscribe((data) => {
            data.forEach((item) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'EC2',
                    name: item.id,
                    tags: item.tags,
                    region: item.region,
                });
            });
        });

        this.awsService.getECS().subscribe((data) => {
            data.services.forEach((service) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'ECS Services',
                    name: service.Name,
                    tags: service.tags,
                    region: service.region,
                });
            });
            data.tasks.forEach((task) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'ECS Tasks',
                    name: task.ARN,
                    tags: task.tags,
                    region: task.region,
                });
            });
            data.clusters.forEach((cluster) => {
                this.services.push({
                    account: 'Sandbox',
                    service: 'ECS Clusters',
                    name: cluster.Name,
                    tags: cluster.tags,
                    region: cluster.region,
                });
            });
        });
    }

    ngOnInit(): void {}
}
