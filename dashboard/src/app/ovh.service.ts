import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';

@Injectable()
export class OvhService {

  private BASE_URL = '/ovh'

  constructor(private http: Http, private storeService: StoreService) { }

  public getCloudProjects(){
    return this.http
     .get(`${this.BASE_URL}/cloud/projects`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudInstances(){
    return this.http
     .get(`${this.BASE_URL}/cloud/instances`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getStorageContainers(){
    return this.http
     .get(`${this.BASE_URL}/cloud/storage`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getUsers(){
    return this.http
     .get(`${this.BASE_URL}/cloud/users`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudVolumes(){
    return this.http
     .get(`${this.BASE_URL}/cloud/volumes`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudSnapshots(){
    return this.http
     .get(`${this.BASE_URL}/cloud/snapshots`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudAlerts(){
    return this.http
     .get(`${this.BASE_URL}/cloud/alerts`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentBill(){
    return this.http
     .get(`${this.BASE_URL}/cloud/current`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudImages(){
    return this.http
     .get(`${this.BASE_URL}/cloud/images`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCloudIps(){
    return this.http
     .get(`${this.BASE_URL}/cloud/ip`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getPublicNetworks(){
    return this.http
     .get(`${this.BASE_URL}/cloud/network/public`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getPrivateNetworks(){
    return this.http
     .get(`${this.BASE_URL}/cloud/network/private`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getFailoverIps(){
    return this.http
     .get(`${this.BASE_URL}/cloud/failover/ip`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getVRacks(){
    return this.http
     .get(`${this.BASE_URL}/cloud/vrack`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getKubeClusters(){
    return this.http
     .get(`${this.BASE_URL}/cloud/kube/clusters`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getKubeNodes(){
    return this.http
     .get(`${this.BASE_URL}/cloud/kube/nodes`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getSSHKeys(){
    return this.http
     .get(`${this.BASE_URL}/cloud/sshkeys`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getLimits(){
    return this.http
     .get(`${this.BASE_URL}/cloud/quotas`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getSSLCertificates(){
    return this.http
     .get(`${this.BASE_URL}/cloud/ssl/certificates`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getSSLGateways(){
    return this.http
     .get(`${this.BASE_URL}/cloud/ssl/gateways`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getProfile(){
    return this.http
     .get(`${this.BASE_URL}/cloud/profile`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getTickets(){
    return this.http
     .get(`${this.BASE_URL}/cloud/tickets`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }
}