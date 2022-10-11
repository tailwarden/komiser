import * as moment from 'moment';
import { Subscription, Subject } from 'rxjs';

import { not } from '@angular/compiler/src/output/output_ast';
import { Component, OnDestroy } from '@angular/core';

import { AwsService } from './services/aws.service';
import { AzureService } from './services/azure.service';
import { DigitaloceanService } from './services/digitalocean.service';
import { GcpService } from './services/gcp.service';
import { OvhService } from './services/ovh.service';
import { StoreService } from './services/store.service';

declare var ga: Function;

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnDestroy {
    public currentProfile: string;
    public notifications: Array<Object> = [];
    public views: Array<Object> = [];
    public _subscriptionNotifications: Subscription;
    public _subscriptionViews: Subscription;

    constructor(private storeService: StoreService) {
        this._subscriptionNotifications =
            this.storeService.newNotification.subscribe((notifications) => {
                this.notifications = [];
                Object.keys(notifications).forEach((key) => {
                    this.notifications.push(notifications[key]);
                });
            });

        this._subscriptionViews = this.storeService.newView.subscribe(
            (views) => {
                this.views = views;
            }
        );
    }

    ngOnDestroy() {
        this._subscriptionNotifications.unsubscribe();
        this._subscriptionViews.unsubscribe();
    }

    public calcMoment(timestamp) {
        return moment(timestamp).fromNow();
    }
}
