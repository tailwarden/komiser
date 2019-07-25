import { Component, OnDestroy } from '@angular/core';
import { AwsService } from './aws.service';
import { GcpService } from './gcp.service';
import { OvhService } from './ovh.service';
import { DigitaloceanService } from './digitalocean.service';
import { StoreService } from './store.service';
import { not } from '@angular/compiler/src/output/output_ast';
import { Subscription } from 'rxjs';
import { Subject } from "rxjs/Subject";
import * as moment from 'moment';

declare var ga: Function;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnDestroy {

  public accountName: string = 'Username';
  public redAlarms: number;
  public profiles: Array<string> = [];
  public currentProfile: string;
  public notifications: Array<Object> = [];
  public _subscription: Subscription;
  public currentProvider: any;
  public availableProviders: Array<any> = [
    {
      label: 'Amazon Web Services',
      value: 'aws'
    },
    {
      label: 'Google Cloud Platform',
      value: 'gcp'
    },
    {
      label: 'OVH',
      value: 'ovh'
    },
    {
      label: 'DigitalOcean',
      value: 'digitalocean'
    }
  ];

  private _storeService: StoreService;

  private providers: Map<String, Object> = new Map<String, Object>();

  constructor(private awsService: AwsService, private gcpService: GcpService, private storeService: StoreService, private digitaloceanService: DigitaloceanService, private ovhService: OvhService) {

    this.providers['aws'] = {
      label: 'Amazon Web Services',
      value: 'aws',
      logo: 'https://cdn.komiser.io/images/aws.png'
    };

    this.providers['gcp'] = {
      label: 'Google Cloud Platform',
      value: 'gcp',
      logo: 'https://cdn.komiser.io/images/gcp.png'
    };

    this.providers['ovh'] = {
      label: 'OVH',
      value: 'ovh',
      logo: 'https://cdn.komiser.io/images/ovh.jpg'
    };

    this.providers['digitalocean'] = {
      label: 'DigitalOcean',
      value: 'digitalocean',
      logo: 'https://cdn.komiser.io/images/digitalocean.png'
    };

    //if (this.storeService.getProvider() == 'aws') {
      if (localStorage.getItem('profile')) {
        this.currentProfile = localStorage.getItem('profile');
      } else {
        this.currentProfile = 'default';
        localStorage.setItem('profile', this.currentProfile);
      }

      this.awsService.getProfiles().subscribe(profiles => {
        this.profiles = profiles;
        if (this.profiles.length > 0 && this.profiles.indexOf(this.currentProfile) == -1) {
          this.currentProfile = this.profiles[0];
          localStorage.setItem('profile', this.currentProfile);
        }
      }, err => {
        this.profiles = [];
      })
   // }


    this.currentProvider = this.providers[this.storeService.getProvider()];
    this.storeService.onProviderChanged(this.storeService.getProvider());

    this._storeService = storeService;

    this.getAccountName();

    this._subscription = this.storeService.newNotification.subscribe(notifications => {
      this.notifications = [];
      Object.keys(notifications).forEach(key => {
        this.notifications.push(notifications[key]);
      })
    })
  }

  private getAccountName() {
    if (this.currentProvider.value == 'aws') {
      this.awsService.getAccountName().subscribe(data => {
        this.accountName = data.username;
      }, err => {
        this.accountName = 'Username';
      });

      this.awsService.getCloudwatchAlarms().subscribe(data => {
        this.redAlarms = data.ALARM;
      }, err => {
        this.redAlarms = 0;
      });
    } else if(this.currentProvider.value == 'ovh'){
      this.ovhService.getCloudAlerts().subscribe(data => {
        this.redAlarms = data;
      }, err => {
        this.redAlarms = 0;
      });

      this.ovhService.getProfile().subscribe(data => {
        this.accountName = data.nichandle;
      }, err => {
        this.accountName = 'Username';
      });
    }else if(this.currentProvider.value == 'digitalocean'){
      this.redAlarms = 0;
      this.digitaloceanService.getProfile().subscribe(data => {
        this.accountName = data.email.substring(0, data.email.indexOf('@'))
      }, err => {
        this.accountName = 'Username';
      });
    } else {
      this.redAlarms = 0;

      this.gcpService.getProjects().subscribe(data => {
        this.accountName = data[0].name;
      }, err => {
        this.accountName = 'Project Name';
      })
    }
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  public calcMoment(timestamp) {
    return moment(timestamp).fromNow();
  }

  public onCloudProviderSelected(provider) {
    this.currentProvider = this.providers[provider];
    this._storeService.onProviderChanged(provider);
    this.getAccountName();
  }

  public onProfileSelected(profile) {
    this.currentProfile = profile;
    localStorage.setItem('profile', this.currentProfile);
    this._storeService.onProfileChanged(profile);
    this.getAccountName();
  }

}
