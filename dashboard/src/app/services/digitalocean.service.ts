import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { StoreService } from './store.service';
import { environment } from '../../environments/environment';

@Injectable()
export class DigitaloceanService {
    private BASE_URL = `${environment.apiUrl}/digitalocean`;

    constructor(private http: HttpClient, private storeService: StoreService) {}

    public getProfile(): any {
        return this.http.get(`${this.BASE_URL}/account`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getActionsHistory(): any {
        return this.http.get(`${this.BASE_URL}/actions`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getContentDeliveryNetworks(): any {
        return this.http.get(`${this.BASE_URL}/cdns`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getCertificates(): any {
        return this.http.get(`${this.BASE_URL}/certificates`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getDomains(): any {
        return this.http.get(`${this.BASE_URL}/domains`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getDroplets(): any {
        return this.http.get(`${this.BASE_URL}/droplets`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getListOfFirewalls(): any {
        return this.http.get(`${this.BASE_URL}/firewalls/list`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getUnsecureFirewalls(): any {
        return this.http.get(`${this.BASE_URL}/firewalls/unsecure`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getFloatingIps(): any {
        return this.http.get(`${this.BASE_URL}/floatingips`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getKubernetesClusters(): any {
        return this.http.get(`${this.BASE_URL}/k8s`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getSshKeys(): any {
        return this.http.get(`${this.BASE_URL}/keys`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getLoadBalancers(): any {
        return this.http.get(`${this.BASE_URL}/loadbalancers`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getProjects(): any {
        return this.http.get(`${this.BASE_URL}/projects`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getRecords(): any {
        return this.http.get(`${this.BASE_URL}/records`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getSnapshots(): any {
        return this.http.get(`${this.BASE_URL}/snapshots`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getVolumes(): any {
        return this.http.get(`${this.BASE_URL}/volumes`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }

    public getDatabases(): any {
        return this.http.get(`${this.BASE_URL}/databases`).pipe(
            catchError((err) => {
                let payload = err.error;
                if (payload && payload.error)
                    this.storeService.add(payload.error);
                return observableThrowError(err.json().error);
            })
        );
    }
}
