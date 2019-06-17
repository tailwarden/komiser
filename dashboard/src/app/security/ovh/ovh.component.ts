import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';

@Component({
  selector: 'ovh-security',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhSecurityComponent implements OnInit {

  public sshKeys: number = 0;
  public sslCertificates: number = 0;
  public sslGateways: number = 0;

  public loadingSSHKeys: boolean = true;
  public loadingSSLCertificates: boolean = true;
  public loadingSSLGateways: boolean = true;

  constructor(private ovhService: OvhService) {
    this.ovhService.getSSHKeys().subscribe(data => {
      this.sshKeys = data;
      this.loadingSSHKeys = false;
    }, err => {
      this.sshKeys = 0;
      this.loadingSSHKeys = false;
    });

    this.ovhService.getSSLCertificates().subscribe(data => {
      this.sslCertificates = data;
      this.loadingSSLCertificates = false;
    }, err => {
      this.sslCertificates = 0;
      this.loadingSSLCertificates = false;
    });

    this.ovhService.getSSLGateways().subscribe(data => {
      this.sslGateways = data;
      this.loadingSSLGateways = false;
    }, err => {
      this.sslGateways = 0;
      this.loadingSSLGateways = false;
    });
    
  }

  ngOnInit() {
  }

}
