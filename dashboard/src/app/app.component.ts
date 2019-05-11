import { Component, OnDestroy } from '@angular/core';
import { AwsService } from './aws.service';
import { GcpService } from './gcp.service';
import { StoreService } from './store.service';
import { not } from '@angular/compiler/src/output/output_ast';
import { Subscription } from 'rxjs';
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
  public notifications: Array<Object> = [];
  public _subscription: Subscription;
  public currentProvider: Object;
  public availableProviders : Array<any> = [
    {
      label: 'Amazon Web Services',
      value: 'aws'
    },
    {
      label: 'Google Cloud Platform',
      value: 'gcp'
    }
  ]

  private _storeService: StoreService;

  private providers: Map<String, Object> = new Map<String, Object>();

  constructor(private awsService: AwsService, private gcpService: GcpService, private storeService: StoreService){

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

    this.currentProvider = this.providers[this.storeService.getProvider()];
    this.storeService.onProviderChanged(this.storeService.getProvider());

    this._storeService = storeService;

    if (this.currentProvider == 'aws'){
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
    } else {
      this.redAlarms = 0; 

      this.gcpService.getProjects().subscribe(data => {
        this.accountName = data[0].name;
      }, err => {
        this.accountName = 'Project Name';
      })
    }

    this._subscription = this.storeService.newNotification.subscribe(notifications => {
      this.notifications = [];
      Object.keys(notifications).forEach(key => {
        this.notifications.push(notifications[key]);
      })
    })
  }

  ngOnDestroy() {
     this._subscription.unsubscribe();
   }

   public calcMoment(timestamp){
      return moment(timestamp).fromNow();
   }

   public onCloudProviderSelected(provider){
     this.currentProvider = this.providers[provider];
     this._storeService.onProviderChanged(provider);
   }

}
