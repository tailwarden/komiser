import { Component, OnInit } from '@angular/core';
import { AwsService } from '../aws.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public account : Object = {};
  public organization: Object = {};

  constructor(private awsService: AwsService) {
    this.awsService.getAccountName().subscribe(data => {
      this.account = data;
    }, err => {
      this.account = {};
    });

    this.awsService.getOrganization().subscribe(data => {
      this.organization = data;
    }, err => {
      this.organization = {};
    });
  }

  ngOnInit() {
  }

}
