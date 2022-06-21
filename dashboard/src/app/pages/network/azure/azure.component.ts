import { Component, OnInit, OnDestroy } from '@angular/core';
import { AzureService } from '../../../services/azure.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../../services/store.service';

@Component({
    selector: 'azure-network',
    templateUrl: './azure.component.html',
    styleUrls: ['./azure.component.css'],
})
export class AzureNetworkComponent implements OnInit, OnDestroy {
    public loadBalancers: number;
    public publicIPs: number;
    public routeTables: number;
    public virtualNetworks: number;
    public subnets: number;
    public dnsZones: number;

    public loadingLoadBalancers: boolean;
    public loadingPublicIPs: boolean;
    public loadingRouteTables: boolean;
    public loadingVirtualNetworks: boolean;
    public loadingSubnets: boolean;
    public loadingDNSZones: boolean;

    private _subscription: Subscription;

    constructor(
        private azureService: AzureService,
        private storeService: StoreService
    ) {
        this.initState();

        this._subscription = this.storeService.profileChanged.subscribe(
            (account) => {
                this.initState();
            }
        );
    }

    ngOnInit() {}

    private initState() {
        this.loadBalancers = 0;
        this.publicIPs = 0;
        this.routeTables = 0;
        this.virtualNetworks = 0;
        this.subnets = 0;
        this.dnsZones = 0;

        this.loadingLoadBalancers = true;
        this.loadingPublicIPs = true;
        this.loadingRouteTables = true;
        this.loadingVirtualNetworks = true;
        this.loadingSubnets = true;
        this.loadingDNSZones = true;

        this.azureService.getLoadBalancers().subscribe(
            (data) => {
                this.loadBalancers = data;
                this.loadingLoadBalancers = false;
            },
            (err) => {
                this.loadBalancers = 0;
                this.loadingLoadBalancers = false;
            }
        );

        this.azureService.getPublicIPs().subscribe(
            (data) => {
                this.publicIPs = data;
                this.loadingPublicIPs = false;
            },
            (err) => {
                this.publicIPs = 0;
                this.loadingPublicIPs = false;
            }
        );

        this.azureService.getRouteTables().subscribe(
            (data) => {
                this.routeTables = data;
                this.loadingRouteTables = false;
            },
            (err) => {
                this.routeTables = 0;
                this.loadingRouteTables = false;
            }
        );

        this.azureService.getVirtualNetworks().subscribe(
            (data) => {
                this.virtualNetworks = data.length;
                this.loadingVirtualNetworks = false;
            },
            (err) => {
                this.virtualNetworks = 0;
                this.loadingVirtualNetworks = false;
            }
        );

        this.azureService.getSubnets().subscribe(
            (data) => {
                this.subnets = data;
                this.loadingSubnets = false;
            },
            (err) => {
                this.subnets = 0;
                this.loadingSubnets = false;
            }
        );

        this.azureService.getDNSZones().subscribe(
            (data) => {
                this.dnsZones = data;
                this.loadingDNSZones = false;
            },
            (err) => {
                this.dnsZones = 0;
                this.loadingDNSZones = false;
            }
        );
    }

    ngOnDestroy() {
        this._subscription.unsubscribe();
    }
}
