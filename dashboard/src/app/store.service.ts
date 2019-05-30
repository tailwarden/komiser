import { Injectable } from '@angular/core';
import { Subject } from "rxjs/Subject";
@Injectable()
export class StoreService {

  private provider: string;

  private notifications: Map<string, Object> = new Map();

  public newNotification: Subject<Map<string, Object>> = new Subject<Map<string, Object>>();

  public providerChanged: Subject<string> = new Subject<string>();

  public profileChanged: Subject<string> = new Subject<string>();

  constructor() {
    if(localStorage.getItem('provider')){
      this.provider = localStorage.getItem('provider');
    } else {
      this.provider = 'aws';
      localStorage.setItem('provider', this.provider);
    }
    this.providerChanged.next(this.provider);
  }

  public getProvider(){
    return this.provider;
  }

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

  public cleanNotifications(){
    this.notifications = new Map();
    this.newNotification.next(this.notifications);
  }

  public onProviderChanged(provider: string){
    this.provider = provider;
    localStorage.setItem('provider', this.provider);
    this.providerChanged.next(this.provider);
    this.notifications = new Map();
    this.newNotification.next(this.notifications);
  }

  public onProfileChanged(profile: string){
    this.profileChanged.next(profile);
    this.notifications = new Map();
    this.newNotification.next(this.notifications);
  }

}
