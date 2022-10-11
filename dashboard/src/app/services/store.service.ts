import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { CloudService } from './cloud.service';
@Injectable()
export class StoreService {
    private notifications: Map<string, Object> = new Map();
    private views = new Array<any>();

    public newNotification: Subject<Map<string, Object>> = new Subject<
        Map<string, Object>
    >();
    public newView: Subject<Array<any>> = new Subject<Array<any>>();

    constructor(private cloudService: CloudService) {
        this.cloudService.getViews().subscribe((views) => {
            this.views = views;
            this.newView.next(this.views);
        });
    }

    public addView(view: any) {
        this.cloudService.saveView(view).subscribe((data) => {
            this.views.push(view);
            this.newView.next(this.views);
        });
    }

    public getView(id) {
        let view = {};
        this.views.forEach((v) => {
            if (v.id == id) {
                view = v;
            }
        });
        return view;
    }

    public getViews() {
        return this.views;
    }

    public add(notification: string) {
        let item = this.notifications[notification];
        if (item) {
            this.notifications[notification] = {
                content: notification,
                timestamp: new Date(),
                total: item.total + 1,
            };
        } else {
            this.notifications[notification] = {
                content: notification,
                timestamp: new Date(),
                total: 1,
            };
        }
        this.newNotification.next(this.notifications);
    }

    public list() {
        return this.notifications;
    }

    public cleanNotifications() {
        this.notifications = new Map();
        this.newNotification.next(this.notifications);
    }
}
