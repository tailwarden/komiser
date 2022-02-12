import 'rxjs/add/operator/map';

import { Observable } from 'rxjs/Rx';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { StoreService } from './store.service';
import { environment } from '../environments/environment';

@Injectable()
export class AzureService {
  private BASE_URL = `${environment.apiUrl}/azure`;
  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getVMs() {
    return this.http
      .get(`${this.BASE_URL}/compute/vms`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }
  public getSnapshots() {
    return this.http
      .get(`${this.BASE_URL}/compute/snapshots`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getDisks() {
    return this.http
      .get(`${this.BASE_URL}/compute/disks`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getKubernetesClusters() {
    return this.http
      .get(`${this.BASE_URL}/managedclusters/clusters`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getMySQLs() {
    return this.http
      .get(`${this.BASE_URL}/storage/mysqls`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getPostgreSQLs() {
    return this.http
      .get(`${this.BASE_URL}/storage/postgresqls`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getRedisInstances() {
    return this.http
      .get(`${this.BASE_URL}/storage/redis`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getFirewalls() {
    return this.http
      .get(`${this.BASE_URL}/security/firewalls`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getPublicIPs() {
    return this.http
      .get(`${this.BASE_URL}/network/publicips`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getLoadBalancers() {
    return this.http
      .get(`${this.BASE_URL}/network/loadbalancers`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getProfiles() {
    return this.http
      .get(`${this.BASE_URL}/security/profiles`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getSecurityGroups() {
    return this.http
      .get(`${this.BASE_URL}/security/securitygroups`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getSecurityRules() {
    return this.http
      .get(`${this.BASE_URL}/security/securityrules`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getRouteTables() {
    return this.http
      .get(`${this.BASE_URL}/network/routetables`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getVirtualNetworks() {
    return this.http
      .get(`${this.BASE_URL}/network/virtualnetworks`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getSubnets() {
    return this.http
      .get(`${this.BASE_URL}/network/subnets`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getDNSZones() {
    return this.http
      .get(`${this.BASE_URL}/network/dnszones`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getCertificates() {
    return this.http
      .get(`${this.BASE_URL}/acm/certificates`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getExpiredCertificates() {
    return this.http
      .get(`${this.BASE_URL}/acm/expired`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getTotalCost() {
    return this.http
      .get(`${this.BASE_URL}/billing/total`)
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
}

