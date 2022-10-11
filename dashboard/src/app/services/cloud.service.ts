import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { StoreService } from './store.service';
import { environment } from '../../environments/environment';

@Injectable()
export class CloudService {
    private BASE_URL = `${environment.apiUrl}`;

    constructor(private http: HttpClient, private storeService: StoreService) {}

    public getCloudAccounts(): any {
        return this.http.get(`${this.BASE_URL}/accounts`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public getCloudRegions(): any {
        return this.http.get(`${this.BASE_URL}/regions`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public getCloudCostByCloudProvider(): any {
        return this.http.get(`${this.BASE_URL}/billing/providers`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public getCloudCostByCloudRegion(): any {
        return this.http.get(`${this.BASE_URL}/billing/regions`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public getCloudCostByCloudAccount(): any {
        return this.http.get(`${this.BASE_URL}/billing/accounts`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public getViews(): any {
        return this.http.get(`${this.BASE_URL}/views`).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }

    public saveView(view): any {
        return this.http.post(`${this.BASE_URL}/views`, view).pipe(
            catchError((err) => {
                return observableThrowError(err.error);
            })
        );
    }
}
