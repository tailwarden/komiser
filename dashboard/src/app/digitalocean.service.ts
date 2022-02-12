import 'rxjs/add/operator/map';

import { Observable } from 'rxjs/Rx';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { StoreService } from './store.service';
import { environment } from '../environments/environment';

@Injectable()
export class DigitaloceanService {
  private BASE_URL = `${environment.apiUrl}/digitalocean`;

  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getProfile() {
    return this.http
      .get(`${this.BASE_URL}/account`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getActionsHistory() {
    return this.http
      .get(`${this.BASE_URL}/actions`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getContentDeliveryNetworks() {
    return this.http
      .get(`${this.BASE_URL}/cdns`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getCertificates() {
    return this.http
      .get(`${this.BASE_URL}/certificates`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getDomains() {
    return this.http
      .get(`${this.BASE_URL}/domains`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getDroplets() {
    return this.http
      .get(`${this.BASE_URL}/droplets`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getListOfFirewalls() {
    return this.http
      .get(`${this.BASE_URL}/firewalls/list`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getUnsecureFirewalls() {
    return this.http
      .get(`${this.BASE_URL}/firewalls/unsecure`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getFloatingIps() {
    return this.http
      .get(`${this.BASE_URL}/floatingips`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getKubernetesClusters() {
    return this.http
      .get(`${this.BASE_URL}/k8s`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getSshKeys() {
    return this.http
      .get(`${this.BASE_URL}/keys`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getLoadBalancers() {
    return this.http
      .get(`${this.BASE_URL}/loadbalancers`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getProjects() {
    return this.http
      .get(`${this.BASE_URL}/projects`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getRecords() {
    return this.http
      .get(`${this.BASE_URL}/records`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getSnapshots() {
    return this.http
      .get(`${this.BASE_URL}/snapshots`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getVolumes() {
    return this.http
      .get(`${this.BASE_URL}/volumes`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }

  public getDatabases() {
    return this.http
      .get(`${this.BASE_URL}/databases`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }
}
