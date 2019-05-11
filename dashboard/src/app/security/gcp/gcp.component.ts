import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';

@Component({
  selector: 'gcp-security',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpSecurityComponent implements OnInit {

  public iamRoles: number = 0;
  public firewalls: number = 0;
  public sslPolicies: number = 0;
  public sslCertificates: number = 0;
  public securityPolicies: number = 0;
  public vpnTunnels: number = 0;
  public cryptoKeys: number = 0;
  public enabledAPIs: Array<any> = [];
  public returnedEnabledAPIs: Array<any> = [];
  public serviceAccounts : number = 0;

  public loadingIamRoles: boolean = true;
  public loadingFirewalls: boolean = true;
  public loadingSSLCertificates: boolean = true;
  public loadingSSLPolicies: boolean = true;
  public loadingSecurityPolicies: boolean = true;
  public loadingVPNTunnels: boolean = true;
  public loadingCryptoKeys: boolean = true;
  public loadingServiceAccounts: boolean = true;

  constructor(private gcpService: GcpService) {
    this.gcpService.getIamRoles().subscribe(data => {
      this.iamRoles = data;
      this.loadingIamRoles = false;
    }, err => {
      this.loadingIamRoles = false;
    });

    this.gcpService.getVPNTunnels().subscribe(data => {
      this.vpnTunnels = data;
      this.loadingVPNTunnels = false;
    }, err => {
      this.vpnTunnels = 0;
      this.loadingVPNTunnels = false;
    });

    this.gcpService.getVpcFirewalls().subscribe(data => {
      this.firewalls = data;
      this.loadingFirewalls = false;
    }, err => {
      this.firewalls = 0;
      this.loadingFirewalls = false;
    });

    this.gcpService.getSSLCertificates().subscribe(data => {
      this.sslCertificates = data;
      this.loadingSSLCertificates = false;
    }, err => {
      this.sslCertificates = 0;
      this.loadingSSLCertificates = false;
    });

    this.gcpService.getSSLPolicies().subscribe(data => {
      this.sslPolicies = data;
      this.loadingSSLPolicies = false;
    }, err => {
      this.sslPolicies = 0;
      this.loadingSSLPolicies = false;
    });

    this.gcpService.getSecurityPolicies().subscribe(data => {
      this.securityPolicies = data;
      this.loadingSecurityPolicies = false;
    }, err => {
      this.securityPolicies = 0;
      this.loadingSecurityPolicies = false;
    });

    this.gcpService.getKMSCryptoKeys().subscribe(data => {
      this.cryptoKeys = data;
      this.loadingCryptoKeys = false;
    }, err => {
      this.cryptoKeys = 0;
      this.loadingCryptoKeys = false;
    });

    this.gcpService.getEnabledAPIs().subscribe(data => {
      this.enabledAPIs = data;
      this.returnedEnabledAPIs = this.enabledAPIs.slice(0, 10);
    }, err => {
      this.enabledAPIs = [];
      this.returnedEnabledAPIs = [];
    });

    this.gcpService.getServiceAccounts().subscribe(data => {
      this.serviceAccounts = data;
      this.loadingServiceAccounts = false;
    }, err => {
      this.serviceAccounts = 0;
      this.loadingServiceAccounts = false;
    });
  }

  pageChanged(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedEnabledAPIs = this.enabledAPIs.slice(startItem, endItem);
  }

  getState(state){
    if(state){
      return 'badge badge-success';
    } else {
      return 'badge badge-danger';
    }
  }

  getStateLabel(state){
    if(state){
      return 'Enabled';
    } else {
      return 'Disabled';
    }
  }

  ngOnInit() {
  }

}
