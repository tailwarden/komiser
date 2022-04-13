
import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { StoreService } from './store.service';
import { environment } from '../../environments/environment';

@Injectable()
export class SettingsService {
    private BASE_URL = `${environment.apiUrl}`;

    constructor(private http: HttpClient, private storeService: StoreService) { }

    public getIntegrations(): any {
        return this.http
            .get(`${this.BASE_URL}/integrations`).pipe(
                catchError((err) => {
                    let payload = err.error;
                    if (payload && payload.error) this.storeService.add(payload.error);
                    return observableThrowError(err.json().error);
                }));
    }

    public setupSlackIntegration(config): any {
        return this.http
            .post(`${this.BASE_URL}/integrations/slack`, config).pipe(
                catchError((err) => {
                    let payload = err.error;
                    if (payload && payload.error) this.storeService.add(payload.error);
                    return observableThrowError(err.json().error);
                }));
    }

}