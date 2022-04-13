
import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { StoreService } from './store.service';
import { environment } from '../../environments/environment';


@Injectable()
export class OvhService {

  private BASE_URL = `${environment.apiUrl}/ovh`

  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getCloudProjects(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/projects`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/instances`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getStorageContainers(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/storage`).pipe(
        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getUsers(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/users`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudVolumes(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/volumes`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudSnapshots(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/snapshots`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudAlerts(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/alerts`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCurrentBill(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/current`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudImages(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/images`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudIps(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/ip`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getPublicNetworks(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/network/public`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getPrivateNetworks(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/network/private`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getFailoverIps(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/failover/ip`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVRacks(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/vrack`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getKubeClusters(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/kube/clusters`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getKubeNodes(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/kube/nodes`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSSHKeys(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/sshkeys`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getLimits(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/quotas`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSSLCertificates(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/ssl/certificates`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSSLGateways(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/ssl/gateways`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getProfile(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/profile`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getTickets(): any {
    return this.http
      .get(`${this.BASE_URL}/cloud/tickets`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }
}