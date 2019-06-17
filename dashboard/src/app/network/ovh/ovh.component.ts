import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';

@Component({
  selector: 'ovh-network',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhNetworkComponent implements OnInit {
  public cloudIps: number = 0;
  public publicNetworks: number = 0;
  public privateNetworks: number = 0;
  public vRacks: number = 0;
  public failoverIps: number = 0;

  public loadingCloudIps: boolean = true;
  public loadingPublicNetworks: boolean = true;
  public loadingPrivateNetworks: boolean = true;
  public loadingVRacks: boolean = true;
  public loadingFailoverIps: boolean = true;

  constructor(private ovhService: OvhService) {
    this.ovhService.getCloudIps().subscribe(data => {
      this.cloudIps = data;
      this.loadingCloudIps = false;
    }, err => {
      this.cloudIps = 0;
      this.loadingCloudIps = false;
    });

    this.ovhService.getPublicNetworks().subscribe(data => {
      this.publicNetworks = data;
      this.loadingPublicNetworks = false;
    }, err => {
      this.publicNetworks = 0;
      this.loadingPublicNetworks = false;
    });

    this.ovhService.getPrivateNetworks().subscribe(data => {
      this.privateNetworks = data;
      this.loadingPrivateNetworks = false;
    }, err => {
      this.privateNetworks = 0;
      this.loadingPrivateNetworks = false;
    });

    this.ovhService.getFailoverIps().subscribe(data => {
      this.failoverIps = data;
      this.loadingFailoverIps = false;
    }, err => {
      this.failoverIps = 0;
      this.loadingFailoverIps = false;
    });

    this.ovhService.getVRacks().subscribe(data => {
      this.vRacks = data;
      this.loadingVRacks = false;
    }, err => {
      this.vRacks = 0;
      this.loadingVRacks = false;
    });
  }

  ngOnInit() {
  }

}
