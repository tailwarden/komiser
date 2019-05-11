import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';

@Component({
  selector: 'gcp-limits',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpLimitsComponent implements OnInit {
  public limits: Array<any> = [];
  public loadingServiceLimits: boolean = true;

  constructor(private gcpService: GcpService) {
    this.gcpService.getQuotas().subscribe(data => {
      this.limits = data;
      this.loadingServiceLimits = false;
    }, err => {
      this.limits = [];
      this.loadingServiceLimits = false;
    });
  }

  public fixLabel(label){
    let value = label.toLowerCase().split('_').join(' ');
    let data = [];
    value.split(' ').forEach(part => {
      data.push(part.charAt(0).toUpperCase() + part.slice(1));
    })
    return data.join(' ');
  }

  public getServiceLogo(name: string){
    if (name.indexOf('Snapshots') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/disk.png';
    }
    else if (name.indexOf('Ssl') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/ssl.png';
    }
    else if (name.indexOf('Firewalls') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/firewalls.png';
    }
    else if (name.indexOf('Networks') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/network.png';
    }
    else if (name.indexOf('Routes') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/routes.png';
    }
    else if (name.indexOf('Buckets') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/compute.png';
    }
    else if (name.indexOf('Services') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/bucket.png';
    }
    else if (name.indexOf('Images') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/os.png';
    }
    else if (name.indexOf('Instances') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/compute.png';
    }
    else if (name.indexOf('Router') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/router.png';
    }
    else if (name.indexOf('Vpn') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/vpn.png';
    }
    else if (name.indexOf('Https') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/https.png';
    }
    else if (name.indexOf('Security') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/security.png';
    }
    else if (name.indexOf('Addresses') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/ip.png';
    }
    else if (name.indexOf('Health Checks') != -1) {
      return 'https://cdn.komiser.io/images/services/gcp/white/healthcheck.png';
    } else {
      return 'https://cdn.komiser.io/images/services/gcp/white/gcp.png';
    }
  }

  public getColor(quota) {
    if(quota.limit == quota.usage) {
      return 'card card-stats card-danger';
    } else if(quota.limit > quota.usage && quota.usage > 0) {
      return 'card card-stats card-warning';
    } else {
      return 'card card-stats card-success';
    }
  }

  ngOnInit() {
  }

}
