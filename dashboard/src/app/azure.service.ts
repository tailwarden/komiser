import 'rxjs/add/operator/map';

import { Observable } from 'rxjs/Rx';

import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import { StoreService } from './store.service';

@Injectable()
export class AzureService {
  private BASE_URL = "/azure";
  constructor(private http: Http, private storeService: StoreService) {}
  private getHeaders() {
    let headers = new Headers();
    headers.append("profile", localStorage.getItem("profile"));
    return headers;
  }
  public getProjects() {
    return this.http
      .get(`${this.BASE_URL}/projects`)
      .map((res) => {
        return res.json();
      })
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }
  public getVMs() {
    return this.http
      .get(`${this.BASE_URL}/compute/vms`)
      .map((res) => {
        return res.json();
      })
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      });
  }
  public getSnapshots() {
    return this.http
      .get(`${this.BASE_URL}/compute/snapshots`)
      .map((res) => {
        return res.json();
      })
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getDisks() {
    return this.http
      .get(`${this.BASE_URL}/compute/disks`)
      .map((res) => {
        return res.json();
      })
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getKubernetesClusters() {
    return this.http
      .get(`${this.BASE_URL}/managedclusters/clusters`)
      .map((res) => {
        return res.json();
      })
      .catch((err) => {
        let payload = JSON.parse(err._body);
        if (payload && payload.error) this.storeService.add(payload.error);
        return Observable.throw(err.json().error);
      })
  }
  public getMySQLs() {
    return this.http
       .get(`${this.BASE_URL}/storage/mysqls`)
       .map((res) => {
         return res.json();
       })
       .catch((err) => {
         let payload = JSON.parse(err._body);
         if (payload && payload.error) this.storeService.add(payload.error);
         return Observable.throw(err.json().error);
       })
  }
  public getPostgreSQLs() {
    return this.http
       .get(`${this.BASE_URL}/storage/postgresqls`)
       .map((res) => {
         return res.json();
       })
       .catch((err) => {
         let payload = JSON.parse(err._body);
         if (payload && payload.error) this.storeService.add(payload.error);
         return Observable.throw(err.json().error);
       })
  }
  public getRedisInstances() {
    return this.http
       .get(`${this.BASE_URL}/storage/redis`)
       .map((res) => {
         return res.json();
       })
       .catch((err) => {
         let payload = JSON.parse(err._body);
         if (payload && payload.error) this.storeService.add(payload.error);
         return Observable.throw(err.json().error);
       })
  }
  public getFirewalls() {
    return this.http
        .get(`${this.BASE_URL}/security/firewalls`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getPublicIPs() {
    return this.http
        .get(`${this.BASE_URL}/network/publicips`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getLoadBalancers() {
    return this.http
        .get(`${this.BASE_URL}/network/loadbalancers`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getProfiles() {
    return this.http
        .get(`${this.BASE_URL}/security/profiles`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getSecurityGroups() {
    return this.http
        .get(`${this.BASE_URL}/security/securitygroups`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getSecurityRules() {
    return this.http
        .get(`${this.BASE_URL}/security/securityrules`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getRouteTables() {
    return this.http
        .get(`${this.BASE_URL}/network/routetables`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getVirtualNetworks() {
    return this.http
        .get(`${this.BASE_URL}/network/virtualnetworks`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getSubnets() {
    return this.http
        .get(`${this.BASE_URL}/network/subnets`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getDNSZones() {
    return this.http
        .get(`${this.BASE_URL}/network/dnszones`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getCertificates() {
    return this.http
        .get(`${this.BASE_URL}/acm/certificates`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
  public getExpiredCertificates() {
    return this.http
        .get(`${this.BASE_URL}/acm/expired`)
        .map((res) => {
          return res.json();
        })
        .catch((err) => {
          let payload = JSON.parse(err._body);
          if (payload && payload.error) this.storeService.add(payload.error);
          return Observable.throw(err.json().error);
        })
  }
}

