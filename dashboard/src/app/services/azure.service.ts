
import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { StoreService } from './store.service';
import { environment } from '../../environments/environment';

@Injectable()
export class AzureService {
  private BASE_URL = `${environment.apiUrl}/azure`;
  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getVMs(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/vms`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }));
  }
  public getSnapshots(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/snapshots`).pipe(
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
      .get(`${this.BASE_URL}/managedclusters/clusters`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getMySQLs(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/mysqls`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getPostgreSQLs(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/postgresqls`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getRedisInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/redis`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getFirewalls(): any {
    return this.http
      .get(`${this.BASE_URL}/security/firewalls`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getPublicIPs(): any {
    return this.http
      .get(`${this.BASE_URL}/network/publicips`).pipe(
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
  public getProfiles(): any {
    return this.http
      .get(`${this.BASE_URL}/security/profiles`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getSecurityGroups(): any {
    return this.http
      .get(`${this.BASE_URL}/security/securitygroups`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getSecurityRules(): any {
    return this.http
      .get(`${this.BASE_URL}/security/securityrules`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getRouteTables(): any {
    return this.http
      .get(`${this.BASE_URL}/network/routetables`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getVirtualNetworks(): any {
    return this.http
      .get(`${this.BASE_URL}/network/virtualnetworks`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getSubnets(): any {
    return this.http
      .get(`${this.BASE_URL}/network/subnets`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getDNSZones(): any {
    return this.http
      .get(`${this.BASE_URL}/network/dnszones`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getCertificates(): any {
    return this.http
      .get(`${this.BASE_URL}/acm/certificates`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getExpiredCertificates(): any {
    return this.http
      .get(`${this.BASE_URL}/acm/expired`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
  public getTotalCost(): any {
    return this.http
      .get(`${this.BASE_URL}/billing/total`).pipe(
        catchError((err) => {
          let payload = err.error;
          if (payload && payload.error) this.storeService.add(payload.error);
          return observableThrowError(err.json().error);
        }))
  }
}

