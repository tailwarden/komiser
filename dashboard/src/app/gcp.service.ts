import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';
import { environment } from '../environments/environment';

@Injectable()
export class GcpService {

  private BASE_URL = `${environment.apiUrl}/gcp`

  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getProjects() {
    return this.http
      .get(`${this.BASE_URL}/resourcemanager/projects`)
      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getComputeInstances() {
    return this.http
      .get(`${this.BASE_URL}/compute/instances`)
      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getIamRoles() {
    return this.http
      .get(`${this.BASE_URL}/iam/roles`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getStorageBuckets() {
    return this.http
      .get(`${this.BASE_URL}/storage/buckets`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getComputeDisks() {
    return this.http
      .get(`${this.BASE_URL}/compute/disks`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDNSZones() {
    return this.http
      .get(`${this.BASE_URL}/dns/zones`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getPubSubTopics() {
    return this.http
      .get(`${this.BASE_URL}/pubsub/topics`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getCloudFunctions() {
    return this.http
      .get(`${this.BASE_URL}/cloudfunctions/functions`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSqlInstances() {
    return this.http
      .get(`${this.BASE_URL}/sql/instances`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVpcNetworks() {
    return this.http
      .get(`${this.BASE_URL}/vpc/networks`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVpcFirewalls() {
    return this.http
      .get(`${this.BASE_URL}/vpc/firewalls`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVpcRouters() {
    return this.http
      .get(`${this.BASE_URL}/vpc/routers`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDiskSnapshots() {
    return this.http
      .get(`${this.BASE_URL}/compute/snapshots`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBucketsSize() {
    return this.http
      .get(`${this.BASE_URL}/storage/size`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBucketsObjects() {
    return this.http
      .get(`${this.BASE_URL}/storage/objects`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getIngestedLoggingBytes() {
    return this.http
      .get(`${this.BASE_URL}/logging/bytes_ingested`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getKubernetesClusters() {
    return this.http
      .get(`${this.BASE_URL}/kubernetes/clusters`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getComputeImages() {
    return this.http
      .get(`${this.BASE_URL}/compute/images`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getRedisInstances() {
    return this.http
      .get(`${this.BASE_URL}/redis/instances`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getComputeCPUUtilization() {
    return this.http
      .get(`${this.BASE_URL}/compute/cpu`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getIAMUsers() {
    return this.http
      .get(`${this.BASE_URL}/iam/users`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBigQueryStatements() {
    return this.http
      .get(`${this.BASE_URL}/bigquery/statements`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBigQueryStorage() {
    return this.http
      .get(`${this.BASE_URL}/bigquery/storage`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBigQueryTables() {
    return this.http
      .get(`${this.BASE_URL}/bigquery/tables`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBigQueryDatasets() {
    return this.http
      .get(`${this.BASE_URL}/bigquery/datasets`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getQuotas() {
    return this.http
      .get(`${this.BASE_URL}/compute/quotas`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getLBRequests() {
    return this.http
      .get(`${this.BASE_URL}/lb/requests`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getAPIRequests() {
    return this.http
      .get(`${this.BASE_URL}/api/requests`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getTotalLoadBalancers() {
    return this.http
      .get(`${this.BASE_URL}/lb/total`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVPCSubnets() {
    return this.http
      .get(`${this.BASE_URL}/vpc/subnets`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVPCAddresses() {
    return this.http
      .get(`${this.BASE_URL}/vpc/addresses`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getVPNTunnels() {
    return this.http
      .get(`${this.BASE_URL}/vpn/tunnels`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSSLPolicies() {
    return this.http
      .get(`${this.BASE_URL}/ssl/policies`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSSLCertificates() {
    return this.http
      .get(`${this.BASE_URL}/ssl/certificates`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getSecurityPolicies() {
    return this.http
      .get(`${this.BASE_URL}/security/policies`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getKMSCryptoKeys() {
    return this.http
      .get(`${this.BASE_URL}/kms/cryptokeys`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getAppEngineBandwidth() {
    return this.http
      .get(`${this.BASE_URL}/gae/bandwidth`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getEnabledAPIs() {
    return this.http
      .get(`${this.BASE_URL}/serviceusage/apis`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDataprocJobs() {
    return this.http
      .get(`${this.BASE_URL}/dataproc/jobs`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDataprocClusters() {
    return this.http
      .get(`${this.BASE_URL}/dataproc/clusters`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBillingLastSixMonths() {
    return this.http
      .get(`${this.BASE_URL}/billing/history`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getBillingPerService() {
    return this.http
      .get(`${this.BASE_URL}/billing/service`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDnsARecords() {
    return this.http
      .get(`${this.BASE_URL}/dns/records`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getServiceAccounts() {
    return this.http
      .get(`${this.BASE_URL}/iam/service_accounts`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getDataflowJobs() {
    return this.http
      .get(`${this.BASE_URL}/dataflow/jobs`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }

  public getNatGateways() {
    return this.http
      .get(`${this.BASE_URL}/nat/gateways`)

      .catch(err => {
        let payload = JSON.parse(err._body)
        if (payload && payload.error)
          this.storeService.add(payload.error);
        return Observable.throw(err.json().error)
      })
  }
}