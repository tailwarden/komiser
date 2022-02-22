import { Component, OnInit, OnDestroy } from '@angular/core';
import { AzureService } from '../../../services/azure.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../../services/store.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';
import * as moment from 'moment';

@Component({
  selector: 'azure-security',
  templateUrl: './azure.component.html',
  styleUrls: ['./azure.component.css']
})
export class AzureSecurityComponent implements OnInit, OnDestroy {
  public firewalls: number;
  public certificates: number;
  public expiredCertificates: number;
  public profiles: number;
  public securityGroups: number;
  public securityRules: number;

  public loadingFirewalls: boolean;
  public loadingCertificates: boolean;
  public loadingExpiredCertificates: boolean;
  public loadingProfiles: boolean;
  public loadingSecurityGroups: boolean;
  public loadingSecurityRules: boolean;

  private _subscription: Subscription;

  constructor(private azureService: AzureService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.firewalls = 0;
    this.certificates = 0;
    this.expiredCertificates = 0;
    this.profiles = 0;
    this.securityRules = 0;

    this.loadingCertificates = true;
    this.loadingFirewalls = true;
    this.loadingProfiles = true;
    this.loadingSecurityRules = true;

    this.azureService.getFirewalls().subscribe(data => {
      this.firewalls = data;
      this.loadingFirewalls = false;
    }, err => {
      this.firewalls = 0;
      this.loadingFirewalls = false;
    });

    this.azureService.getCertificates().subscribe(data => {
      this.certificates = data;
      this.loadingCertificates = false;
    }, err => {
      this.certificates = 0;
      this.loadingCertificates = false;
    });

    this.azureService.getExpiredCertificates().subscribe(data => {
      this.expiredCertificates = data;
      this.loadingExpiredCertificates = false;
    }, err => {
      this.expiredCertificates = 0;
      this.loadingExpiredCertificates = false;
    })

    this.azureService.getProfiles().subscribe(data => {
      this.profiles = data;
      this.loadingProfiles = false;
    }, err => {
      this.profiles = 0;
      this.loadingProfiles = false;
    })

    this.azureService.getSecurityGroups().subscribe(data => {
      this.securityGroups = data.length;
      this.loadingSecurityGroups = false;
    }, err => {
      this.securityGroups = 0;
      this.loadingSecurityGroups = false;
    })

    this.azureService.getSecurityRules().subscribe(data => {
      this.securityRules = data;
      this.loadingSecurityRules = false;
    }, err => {
      this.securityRules = 0;
      this.loadingSecurityRules = false;
    })

  }

  public calcMoment(timestamp) {
    return moment(timestamp).fromNow();
  }

  public getFlagIcon(region) {
    switch (region) {
      case 'nyc':
        return 'https://cdn.komiser.io/images/flags/usa.png';
      case 'ams':
        return 'https://cdn.komiser.io/images/flags/netherlands.png';
      case 'sfo':
        return 'https://cdn.komiser.io/images/flags/usa.png';
      case 'lon':
        return 'https://cdn.komiser.io/images/flags/uk.png';
      case 'fra':
        return 'https://cdn.komiser.io/images/flags/france.png';
      case 'tor':
        return 'https://cdn.komiser.io/images/flags/canada.png';
      case 'blr':
        return 'https://cdn.komiser.io/images/flags/india.png';
      default:
        return 'https://cdn.komiser.io/images/flags/usa.png';
    }
  }

  ngOnInit() {
  }

}
