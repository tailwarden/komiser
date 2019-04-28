import { Component, OnDestroy } from '@angular/core';
import { AwsService } from './aws.service';
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
  
  constructor(private awsService: AwsService, private storeService: StoreService){

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

}
