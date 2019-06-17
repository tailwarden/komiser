import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';
@Component({
  selector: 'ovh-limits',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhLimitsComponent implements OnInit {

  public limits : Array<any> = new Array();
  public loadingServiceLimits: boolean = true;

  constructor(private ovhService: OvhService) {
    this.ovhService.getLimits().subscribe(data => {
      this.limits = data;
      data.forEach(item => {
        item.volume.usedGigabytes = this.bytesToSizeWithUnit(item.volume.usedGigabytes*1024*1024*1024);
        item.volume.maxGigabytes = this.bytesToSizeWithUnit(item.volume.maxGigabytes*1024*1024*1024);
      });
      this.loadingServiceLimits = false;
    }, err => {
      this.loadingServiceLimits = false;
      this.limits = [];
    })
  }

  public getFlagIcon(region){
    switch(region){
      case 'SBG5':
        return 'https://cdn.komiser.io/images/flags/france.png';
      case 'BHS5':
        return 'https://cdn.komiser.io/images/flags/canada.png';
      case 'GRA5':
        return 'https://cdn.komiser.io/images/flags/france.png';
      case 'WAW1':
        return 'https://cdn.komiser.io/images/flags/poland.png';
      case 'DE1':
        return 'https://cdn.komiser.io/images/flags/germany.png';
      case 'UK1':
        return 'https://cdn.komiser.io/images/flags/uk.png';
      default: 
        return 'https://cdn.komiser.io/images/flags/france.png';
    } 
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

  ngOnInit() {
  }

}
