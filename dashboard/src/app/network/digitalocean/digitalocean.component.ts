import { Component, OnInit, OnDestroy } from '@angular/core';
import { DigitaloceanService } from '../../digitalocean.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../store.service';

@Component({
  selector: 'digitalocean-network',
  templateUrl: './digitalocean.component.html',
  styleUrls: ['./digitalocean.component.css']
})
export class DigitaloceanNetworkComponent implements OnInit, OnDestroy {
  public floatingIps: number;
  public domains: number;
  public loadBalancers: number;
  public cdnEndpoints: number;
  public aRecords: number;
  public cnameRecords: number;

  public loadingFloatingIps: boolean; 
  public loadingDomains: boolean;
  public loadingLoadBalancers: boolean;
  public loadingCDNEndpoints: boolean;
  public loadingARecords: boolean;
  public loadingCNAMERecords: boolean;

  private _subscription: Subscription;

  constructor(private digitaloceanService: DigitaloceanService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  

  ngOnInit() {
  }

  private initState(){
    this.floatingIps = 0;
    this.domains = 0;
    this.loadBalancers = 0;
    this.cdnEndpoints = 0;
    this.aRecords = 0;
    this.cnameRecords = 0;

    this.loadingCDNEndpoints = true;
    this.loadingCNAMERecords = true;
    this.loadingARecords = true;
    this.loadingDomains = true;
    this.loadingFloatingIps = true;
    this.loadingLoadBalancers = true;

    this.digitaloceanService.getFloatingIps().subscribe(data => {
      this.floatingIps = data;
      this.loadingFloatingIps = false;
    }, err => {
      this.floatingIps = 0;
      this.loadingFloatingIps = false;
    });

    this.digitaloceanService.getDomains().subscribe(data => {
      this.domains = data;
      this.loadingDomains = false;
    }, err => {
      this.domains = 0;
      this.loadingDomains = false;
    });

    this.digitaloceanService.getLoadBalancers().subscribe(data => {
      this.loadBalancers = data;
      this.loadingLoadBalancers = false;
    }, err => {
      this.loadBalancers = 0;
      this.loadingLoadBalancers = false;
    });

    this.digitaloceanService.getContentDeliveryNetworks().subscribe(data => {
      this.cdnEndpoints = data;
      this.loadingCDNEndpoints = false;
    }, err => {
      this.cdnEndpoints = 0;
      this.loadingCDNEndpoints = false;
    });

    this.digitaloceanService.getRecords().subscribe(data => {
      this.aRecords = data.a;
      this.cnameRecords = data.cname;
      this.loadingARecords = false;
      this.loadingCNAMERecords = false;
    }, err => {
      this.aRecords = 0;
      this.cnameRecords = 0;
      this.loadingARecords = false;
      this.loadingCNAMERecords = false;
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

}
