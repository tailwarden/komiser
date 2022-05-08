import { Component, OnInit, OnDestroy } from '@angular/core';
import { CivoService } from '../../../services/civo.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../../services/store.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';
import * as moment from 'moment';

@Component({
  selector: 'civo-security',
  templateUrl: './civo.component.html',
  styleUrls: ['./civo.component.css']
})
export class CivoSecurityComponent implements OnInit, OnDestroy {
  public firewallRules: number;
  public sshKeys: number;
  public expiredCertificates: number;
  public profiles: number;
  public securityGroups: number;
  public securityRules: number;

  public loadingFirewallRules: boolean;
  public loadingSSHKeys: boolean;
  public loadingExpiredCertificates: boolean;
  public loadingProfiles: boolean;
  public loadingSecurityGroups: boolean;
  public loadingSecurityRules: boolean;

  private _subscription: Subscription;

  constructor(private civoService: CivoService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.firewallRules = 0;
    this.sshKeys = 0;

    this.loadingSSHKeys = true;
    this.loadingFirewallRules = true;

    this.civoService.getFirewallRules().subscribe(data => {
      this.firewallRules = data;
      this.loadingFirewallRules = false;
    }, err => {
      this.firewallRules = 0;
      this.loadingFirewallRules = false;
    });

    this.civoService.getSSHKeys().subscribe(data => {
      this.sshKeys = data;
      this.loadingSSHKeys = false;
    }, err => {
      this.sshKeys = 0;
      this.loadingSSHKeys = false;
    });

  }

  ngOnInit() {
  }

}
