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
}
