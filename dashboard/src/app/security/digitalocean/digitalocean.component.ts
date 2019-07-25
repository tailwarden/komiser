import { Component, OnInit, OnDestroy } from '@angular/core';
import { DigitaloceanService } from '../../digitalocean.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../store.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';
import * as moment from 'moment';

@Component({
  selector: 'digitalocean-security',
  templateUrl: './digitalocean.component.html',
  styleUrls: ['./digitalocean.component.css']
})
export class DigitaloceanSecurityComponent implements OnInit, OnDestroy {
  public sshKeys: number;
  public firewalls: number;
  public customCertificates: number;
  public letsEncryptCertificates: number;
  public unrestrictedFirewalls: Array<any> = [];
  public returnedUnrestrictedFirewalls: Array<any> = [];
  public actions: Array<any> = [];
  public returnedActions: Array<any> = [];

  public loadingSSHKeys: boolean;
  public loadingFirewalls: boolean;
  public loadingCustomCertificates: boolean;
  public loadingLetsEncryptCertificates: boolean;

  private _subscription: Subscription;

  constructor(private digitaloceanService: DigitaloceanService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnDestroy(){
    this._subscription.unsubscribe();
  }

  private initState(){
    this.sshKeys = 0;
    this.firewalls = 0;
    this.customCertificates = 0;
    this.letsEncryptCertificates = 0;
    this.unrestrictedFirewalls = [];
    this.returnedUnrestrictedFirewalls = [];
    this.actions = [];
    this.returnedActions = [];

    this.loadingCustomCertificates = true;
    this.loadingFirewalls = true;
    this.loadingLetsEncryptCertificates = true;
    this.loadingSSHKeys = true;

    this.digitaloceanService.getSshKeys().subscribe(data => {
      this.sshKeys = data;
      this.loadingSSHKeys = false;
    }, err => {
      this.sshKeys = 0;
      this.loadingSSHKeys = false;
    });

    this.digitaloceanService.getListOfFirewalls().subscribe(data => {
      this.firewalls = data;
      this.loadingFirewalls = false;
    }, err => {
      this.firewalls = 0;
      this.loadingFirewalls = false;
    });

    this.digitaloceanService.getCertificates().subscribe(data => {
      this.letsEncryptCertificates = data.letsEncrypt;
      this.customCertificates = data.custom;
      this.loadingCustomCertificates = false;
      this.loadingLetsEncryptCertificates = false;
    }, err => {
      this.letsEncryptCertificates = 0;
      this.customCertificates = 0;
      this.loadingCustomCertificates = false;
      this.loadingLetsEncryptCertificates = false;
    });

    this.digitaloceanService.getUnsecureFirewalls().subscribe(data => {
      this.unrestrictedFirewalls = data;
      this.returnedUnrestrictedFirewalls = this.unrestrictedFirewalls.slice(0, 20);
    }, err => {
      this.unrestrictedFirewalls = [];
      this.returnedUnrestrictedFirewalls = [];
    });

    this.digitaloceanService.getActionsHistory().subscribe(data => {
      this.actions = data;
      this.returnedActions = this.actions.slice(0, 20);
    }, err => {
      this.returnedActions = [];
      this.actions = [];
    });
  }

  public calcMoment(timestamp) {
    return moment(timestamp).fromNow();
  }

  pageChangedActions(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedActions = this.actions.slice(startItem, endItem);
  }

  pageChanged(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedUnrestrictedFirewalls = this.unrestrictedFirewalls.slice(startItem, endItem);
  }

  public getFlagIcon(region){
    switch(region){
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
