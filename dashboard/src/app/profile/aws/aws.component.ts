import { Component, OnInit } from '@angular/core';
import { AwsService } from '../../aws.service';

@Component({
  selector: 'aws-profile',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsProfileComponent implements OnInit {
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
