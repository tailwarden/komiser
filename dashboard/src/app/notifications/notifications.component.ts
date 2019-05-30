import { Component, OnInit, OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';
import * as moment from 'moment';

@Component({
  selector: 'app-notifications',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.css']
})
export class NotificationsComponent implements OnInit, OnDestroy {
  public notifications: Array<Object> = [];
  public _subscription: Subscription;

  constructor(private storeService: StoreService) {

    let temp = this.storeService.list();
    Object.keys(temp).forEach(key => {
      this.notifications.push(temp[key]);
    })

    this._subscription = this.storeService.newNotification.subscribe(notifications => {
      this.notifications = [];
      Object.keys(notifications).forEach(key => {
        this.notifications.push(notifications[key]);
      })
    })
  }

  public calcMoment(timestamp) {
    return moment(timestamp).fromNow();
  }

  ngOnDestroy(){
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }

}
