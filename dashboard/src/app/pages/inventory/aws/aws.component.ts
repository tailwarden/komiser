import { Component, OnInit } from '@angular/core';
import { AwsService } from '../../../services/aws.service';

@Component({
  selector: 'aws-inventory',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsComponent implements OnInit {

  public services : Array<any> = new Array<any>();

  constructor(private awsService: AwsService) {
    this.awsService.getInstancesPerRegion().subscribe(data => {
      data.forEach(item => {
        this.services.push({
          account: 'Sandbox',
          service: 'EC2',
          name: item.id,
          tags: item.tags,
        })
      })
    })
  }

  ngOnInit(): void {
  }

}
