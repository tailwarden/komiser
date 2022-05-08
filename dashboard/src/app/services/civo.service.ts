
import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { StoreService } from './store.service';
import { environment } from '../../environments/environment';

@Injectable()
export class CivoService {
  private BASE_URL = `${environment.apiUrl}/civo`;
  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/instances`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }));
  }
  public getVolumes(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/volumes`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getDisks(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/disks`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getKubernetesClusters(): any {
    return this.http
      .get(`${this.BASE_URL}/kubernetes/clusters`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getFirewallRules(): any {
    return this.http
      .get(`${this.BASE_URL}/security/firewallrules`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getPrivateNetworks(): any {
    return this.http
      .get(`${this.BASE_URL}/network/private`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getLoadBalancers(): any {
    return this.http
      .get(`${this.BASE_URL}/network/loadbalancers`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getSSHKeys(): any {
    return this.http
      .get(`${this.BASE_URL}/security/sshkeys`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getDNSDomains(): any {
    return this.http
      .get(`${this.BASE_URL}/network/dnsdomains`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getRegions(): any {
    return this.http
      .get(`${this.BASE_URL}/resources/regions`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
}

