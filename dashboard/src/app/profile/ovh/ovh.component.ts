import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';

@Component({
  selector: 'ovh-profile',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhProfileComponent implements OnInit {

  public profile: any;

  constructor(private ovhService: OvhService) {
    this.ovhService.getProfile().subscribe(data => {
      this.profile = data;
    }, err => {
      this.profile = {};
    });
  }

  ngOnInit() {
  }

}
