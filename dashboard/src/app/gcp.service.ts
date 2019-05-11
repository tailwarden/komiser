import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';

@Injectable()
export class GcpService {

  private BASE_URL = '/gcp'

  constructor(private http: Http, private storeService: StoreService) { }

  public getProjects(){
    return this.http
     .get(`${this.BASE_URL}/resourcemanager/projects`)
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

  public getComputeInstances(){
    return this.http
     .get(`${this.BASE_URL}/compute/instances`)
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

  public getIamRoles(){
    return this.http
     .get(`${this.BASE_URL}/iam/roles`)
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

  public getStorageBuckets(){
    return this.http
     .get(`${this.BASE_URL}/storage/buckets`)
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

  public getComputeDisks(){
    return this.http
     .get(`${this.BASE_URL}/compute/disks`)
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

  public getDNSZones(){
    return this.http
     .get(`${this.BASE_URL}/dns/zones`)
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

  public getPubSubTopics(){
    return this.http
     .get(`${this.BASE_URL}/pubsub/topics`)
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

  public getCloudFunctions(){
    return this.http
     .get(`${this.BASE_URL}/cloudfunctions/functions`)
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

  public getSqlInstances(){
    return this.http
     .get(`${this.BASE_URL}/sql/instances`)
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

  public getVpcNetworks(){
    return this.http
     .get(`${this.BASE_URL}/vpc/networks`)
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

  public getVpcFirewalls(){
    return this.http
     .get(`${this.BASE_URL}/vpc/firewalls`)
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

  public getVpcRouters(){
    return this.http
     .get(`${this.BASE_URL}/vpc/routers`)
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

  public getDiskSnapshots(){
    return this.http
     .get(`${this.BASE_URL}/compute/snapshots`)
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

  public getBucketsSize(){
    return this.http
     .get(`${this.BASE_URL}/storage/size`)
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

  public getBucketsObjects(){
    return this.http
     .get(`${this.BASE_URL}/storage/objects`)
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

  public getIngestedLoggingBytes(){
    return this.http
     .get(`${this.BASE_URL}/logging/bytes_ingested`)
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

  public getKubernetesClusters(){
    return this.http
     .get(`${this.BASE_URL}/kubernetes/clusters`)
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

  public getComputeImages(){
    return this.http
     .get(`${this.BASE_URL}/compute/images`)
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

  public getRedisInstances(){
    return this.http
     .get(`${this.BASE_URL}/redis/instances`)
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

  public getComputeCPUUtilization(){
    return this.http
     .get(`${this.BASE_URL}/compute/cpu`)
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

  public getIAMUsers(){
    return this.http
     .get(`${this.BASE_URL}/iam/users`)
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

  public getBigQueryStatements(){
    return this.http
     .get(`${this.BASE_URL}/bigquery/statements`)
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

  public getBigQueryStorage(){
    return this.http
     .get(`${this.BASE_URL}/bigquery/storage`)
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

  public getBigQueryTables(){
    return this.http
     .get(`${this.BASE_URL}/bigquery/tables`)
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

  public getBigQueryDatasets(){
    return this.http
     .get(`${this.BASE_URL}/bigquery/datasets`)
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

  public getQuotas(){
    return this.http
     .get(`${this.BASE_URL}/compute/quotas`)
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

  public getLBRequests(){
    return this.http
     .get(`${this.BASE_URL}/lb/requests`)
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

  public getAPIRequests(){
    return this.http
     .get(`${this.BASE_URL}/api/requests`)
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

  public getTotalLoadBalancers(){
    return this.http
     .get(`${this.BASE_URL}/lb/total`)
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

  public getVPCSubnets(){
    return this.http
     .get(`${this.BASE_URL}/vpc/subnets`)
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

  public getVPCAddresses(){
    return this.http
     .get(`${this.BASE_URL}/vpc/addresses`)
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

  public getVPNTunnels(){
    return this.http
     .get(`${this.BASE_URL}/vpn/tunnels`)
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

  public getSSLPolicies(){
    return this.http
     .get(`${this.BASE_URL}/ssl/policies`)
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
     .get(`${this.BASE_URL}/ssl/certificates`)
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

  public getSecurityPolicies(){
    return this.http
     .get(`${this.BASE_URL}/security/policies`)
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

  public getKMSCryptoKeys(){
    return this.http
     .get(`${this.BASE_URL}/kms/cryptokeys`)
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

  public getAppEngineBandwidth(){
    return this.http
     .get(`${this.BASE_URL}/gae/bandwidth`)
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

  public getEnabledAPIs(){
    return this.http
     .get(`${this.BASE_URL}/serviceusage/apis`)
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

  public getDataprocJobs(){
    return this.http
     .get(`${this.BASE_URL}/dataproc/jobs`)
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

  public getDataprocClusters(){
    return this.http
     .get(`${this.BASE_URL}/dataproc/clusters`)
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

  public getBillingLastSixMonths(){
    return this.http
     .get(`${this.BASE_URL}/billing/history`)
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

  public getBillingPerService(){
    return this.http
     .get(`${this.BASE_URL}/billing/service`)
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

  public getDnsARecords(){
    return this.http
     .get(`${this.BASE_URL}/dns/records`)
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

  public getServiceAccounts(){
    return this.http
     .get(`${this.BASE_URL}/iam/service_accounts`)
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

  public getDataflowJobs(){
    return this.http
     .get(`${this.BASE_URL}/dataflow/jobs`)
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

  public getNatGateways(){
    return this.http
     .get(`${this.BASE_URL}/nat/gateways`)
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