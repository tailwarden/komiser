import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';
import { environment } from '../environments/environment';

@Injectable()
export class OvhService {

  private BASE_URL = `${environment.apiUrl}/ovh`

  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getCloudProjects() {
    return this.http
      .get(`${this.BASE_URL}/cloud/projects`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudInstances() {
    return this.http
      .get(`${this.BASE_URL}/cloud/instances`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getStorageContainers() {
    return this.http
      .get(`${this.BASE_URL}/cloud/storage`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getUsers() {
    return this.http
      .get(`${this.BASE_URL}/cloud/users`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudVolumes() {
    return this.http
      .get(`${this.BASE_URL}/cloud/volumes`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudSnapshots() {
    return this.http
      .get(`${this.BASE_URL}/cloud/snapshots`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudAlerts() {
    return this.http
      .get(`${this.BASE_URL}/cloud/alerts`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCurrentBill() {
    return this.http
      .get(`${this.BASE_URL}/cloud/current`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudImages() {
    return this.http
      .get(`${this.BASE_URL}/cloud/images`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudIps() {
    return this.http
      .get(`${this.BASE_URL}/cloud/ip`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getPublicNetworks() {
    return this.http
      .get(`${this.BASE_URL}/cloud/network/public`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getPrivateNetworks() {
    return this.http
      .get(`${this.BASE_URL}/cloud/network/private`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getFailoverIps() {
    return this.http
      .get(`${this.BASE_URL}/cloud/failover/ip`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVRacks() {
    return this.http
      .get(`${this.BASE_URL}/cloud/vrack`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getKubeClusters() {
    return this.http
      .get(`${this.BASE_URL}/cloud/kube/clusters`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getKubeNodes() {
    return this.http
      .get(`${this.BASE_URL}/cloud/kube/nodes`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSSHKeys() {
    return this.http
      .get(`${this.BASE_URL}/cloud/sshkeys`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getLimits() {
    return this.http
      .get(`${this.BASE_URL}/cloud/quotas`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSSLCertificates() {
    return this.http
      .get(`${this.BASE_URL}/cloud/ssl/certificates`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSSLGateways() {
    return this.http
      .get(`${this.BASE_URL}/cloud/ssl/gateways`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getProfile() {
    return this.http
      .get(`${this.BASE_URL}/cloud/profile`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getTickets() {
    return this.http
      .get(`${this.BASE_URL}/cloud/tickets`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }
}