import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';

@Component({
  selector: 'gcp-profile',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpProfileComponent implements OnInit {
  public project: any = {};

  constructor(private gcpService: GcpService) {
    this.gcpService.getProjects().subscribe(data => {
      this.project = data[0];
    }, err => {
      this.project = {};
    })
  }

  ngOnInit() {
  }

}
