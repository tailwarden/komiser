import { Component, OnInit, OnDestroy } from '@angular/core';
import { CivoService } from '../../../services/civo.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../../services/store.service';

@Component({
  selector: 'civo-network',
  templateUrl: './civo.component.html',
  styleUrls: ['./civo.component.css']
})
export class CivoNetworkComponent implements OnInit, OnDestroy {
  public loadBalancers: number;
  public privateNetworks: number;
  public dnsDomains: number;

  public loadingLoadBalancers: boolean;
  public loadingPrivateNetworks: boolean;
  public loadingDNSDomains: boolean;

  private _subscription: Subscription;

  constructor(private civoService: CivoService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }



  ngOnInit() {
  }

  private initState() {
    this.loadBalancers = 0;
    this.privateNetworks = 0;
    this.dnsDomains = 0;

    this.loadingLoadBalancers = true;
    this.loadingPrivateNetworks = true;
    this.loadingDNSDomains = true;

    this.civoService.getLoadBalancers().subscribe(data => {
      this.loadBalancers = data;
      this.loadingLoadBalancers = false;
    }, err => {
      this.loadBalancers = 0;
      this.loadingLoadBalancers = false;
    });

    this.civoService.getPrivateNetworks().subscribe(data => {
      this.privateNetworks = data;
      this.loadingPrivateNetworks = false;
    }, err => {
      this.privateNetworks = 0;
      this.loadingPrivateNetworks = false;
    })

    this.civoService.getDNSDomains().subscribe(data => {
      this.dnsDomains = data;
      this.loadingDNSDomains = false;
    }, err => {
      this.dnsDomains = 0;
      this.loadingDNSDomains = false;
    })

  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

}
