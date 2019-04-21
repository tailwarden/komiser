import { Injectable } from '@angular/core';
import { Subject } from "rxjs/Subject";

@Injectable()
export class StoreService {

  private notifications: Map<String, Object> = new Map();

  newNotification: Subject<Map<String, Object>> = new Subject<Map<String, Object>>();

  constructor() {}

  public add(notification: string){
    let item = this.notifications[notification];
    if (item) {
      this.notifications[notification] = {
        content: notification,
        timestamp: new Date(),
        total: item.total + 1
      }
    } else {
      this.notifications[notification] = {
        content: notification,
        timestamp: new Date(),
        total: 1
      }
    }
    this.newNotification.next(this.notifications);
  }

  public list(){
    return this.notifications;
  }

}
