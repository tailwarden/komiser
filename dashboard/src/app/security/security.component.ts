import { Component, OnInit } from '@angular/core';
import { AwsService } from '../aws.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';

@Component({
  selector: 'app-security',
  templateUrl: './security.component.html',
  styleUrls: ['./security.component.css']
})
export class SecurityComponent implements OnInit {

  public kmsKeys: number;
  public securityGroups: number;
  public keyPairs: number;
  public routeTables: number;
  public acmCertificates: number;
  public acmExpiredCertificates: number;
  public unrestrictedSecurityGroups: Array<any> = [];
  public returnedUnrestrictedSecurityGroups: Array<any> = [];

  constructor(private awsService: AwsService) {
    this.awsService.getKMSKeys().subscribe(data => {
      this.kmsKeys = data;
    }, err => {
      this.kmsKeys = 0;
    });

    this.awsService.getSecurityGroups().subscribe(data => {
      this.securityGroups = data;
    }, err => {
      this.securityGroups = 0;
    });

    this.awsService.getKeyPairs().subscribe(data => {
      this.keyPairs = data;
    }, err => {
      this.keyPairs = 0;
    });

    this.awsService.getRouteTables().subscribe(data => {
      this.routeTables = data;
    }, err => {
      this.routeTables = 0;
    });

    this.awsService.getACMListCertificates().subscribe(data => {
      this.acmCertificates = data;
    }, err => {
      this.acmCertificates = 0;
    });

    this.awsService.getACMExpiredCertificates().subscribe(data => {
      this.acmExpiredCertificates = data;
    }, err => {
      this.acmExpiredCertificates = 0;
    });

    this.awsService.getUnrestrictedSecurityGroups().subscribe(data => {
      this.unrestrictedSecurityGroups = data;
      this.returnedUnrestrictedSecurityGroups = this.unrestrictedSecurityGroups.slice(0, 20);
    }, err => {
      this.unrestrictedSecurityGroups = [];
      this.returnedUnrestrictedSecurityGroups = [];
    })
  }

  pageChanged(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedUnrestrictedSecurityGroups = this.unrestrictedSecurityGroups.slice(startItem, endItem);
  }

  ngOnInit() {
  }

}
