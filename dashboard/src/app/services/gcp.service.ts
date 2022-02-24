
import { throwError as observableThrowError, Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { StoreService } from './store.service';
import { environment } from '../../environments/environment';


@Injectable()
export class GcpService {

  private BASE_URL = `${environment.apiUrl}/gcp`

  constructor(private http: HttpClient, private storeService: StoreService) { }

  public getProjects(): any {
    return this.http
      .get(`${this.BASE_URL}/resourcemanager/projects`).pipe(
        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getComputeInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/instances`).pipe(
        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getIamRoles(): any {
    return this.http
      .get(`${this.BASE_URL}/iam/roles`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getStorageBuckets(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/buckets`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getComputeDisks(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/disks`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDNSZones(): any {
    return this.http
      .get(`${this.BASE_URL}/dns/zones`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getPubSubTopics(): any {
    return this.http
      .get(`${this.BASE_URL}/pubsub/topics`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getCloudFunctions(): any {
    return this.http
      .get(`${this.BASE_URL}/cloudfunctions/functions`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSqlInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/sql/instances`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVpcNetworks(): any {
    return this.http
      .get(`${this.BASE_URL}/vpc/networks`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVpcFirewalls(): any {
    return this.http
      .get(`${this.BASE_URL}/vpc/firewalls`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVpcRouters(): any {
    return this.http
      .get(`${this.BASE_URL}/vpc/routers`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDiskSnapshots(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/snapshots`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBucketsSize(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/size`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBucketsObjects(): any {
    return this.http
      .get(`${this.BASE_URL}/storage/objects`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getIngestedLoggingBytes(): any {
    return this.http
      .get(`${this.BASE_URL}/logging/bytes_ingested`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getKubernetesClusters(): any {
    return this.http
      .get(`${this.BASE_URL}/kubernetes/clusters`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getComputeImages(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/images`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getRedisInstances(): any {
    return this.http
      .get(`${this.BASE_URL}/redis/instances`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getComputeCPUUtilization(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/cpu`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getIAMUsers(): any {
    return this.http
      .get(`${this.BASE_URL}/iam/users`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBigQueryStatements(): any {
    return this.http
      .get(`${this.BASE_URL}/bigquery/statements`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBigQueryStorage(): any {
    return this.http
      .get(`${this.BASE_URL}/bigquery/storage`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBigQueryTables(): any {
    return this.http
      .get(`${this.BASE_URL}/bigquery/tables`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBigQueryDatasets(): any {
    return this.http
      .get(`${this.BASE_URL}/bigquery/datasets`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getQuotas(): any {
    return this.http
      .get(`${this.BASE_URL}/compute/quotas`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getLBRequests(): any {
    return this.http
      .get(`${this.BASE_URL}/lb/requests`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getAPIRequests(): any {
    return this.http
      .get(`${this.BASE_URL}/api/requests`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getTotalLoadBalancers(): any {
    return this.http
      .get(`${this.BASE_URL}/lb/total`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVPCSubnets(): any {
    return this.http
      .get(`${this.BASE_URL}/vpc/subnets`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVPCAddresses(): any {
    return this.http
      .get(`${this.BASE_URL}/vpc/addresses`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getVPNTunnels(): any {
    return this.http
      .get(`${this.BASE_URL}/vpn/tunnels`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSSLPolicies(): any {
    return this.http
      .get(`${this.BASE_URL}/ssl/policies`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSSLCertificates(): any {
    return this.http
      .get(`${this.BASE_URL}/ssl/certificates`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getSecurityPolicies(): any {
    return this.http
      .get(`${this.BASE_URL}/security/policies`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getKMSCryptoKeys(): any {
    return this.http
      .get(`${this.BASE_URL}/kms/cryptokeys`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getAppEngineBandwidth(): any {
    return this.http
      .get(`${this.BASE_URL}/gae/bandwidth`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getEnabledAPIs(): any {
    return this.http
      .get(`${this.BASE_URL}/serviceusage/apis`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDataprocJobs(): any {
    return this.http
      .get(`${this.BASE_URL}/dataproc/jobs`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDataprocClusters(): any {
    return this.http
      .get(`${this.BASE_URL}/dataproc/clusters`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBillingLastSixMonths(): any {
    return this.http
      .get(`${this.BASE_URL}/billing/history`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getBillingPerService(): any {
    return this.http
      .get(`${this.BASE_URL}/billing/service`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDnsARecords(): any {
    return this.http
      .get(`${this.BASE_URL}/dns/records`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getServiceAccounts(): any {
    return this.http
      .get(`${this.BASE_URL}/iam/service_accounts`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getDataflowJobs(): any {
    return this.http
      .get(`${this.BASE_URL}/dataflow/jobs`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }

  public getNatGateways(): any {
    return this.http
      .get(`${this.BASE_URL}/nat/gateways`).pipe(

        catchError(err => {
          let payload = err.error;
          if (payload && payload.error)
            this.storeService.add(payload.error);
          return observableThrowError(err.json().error)
        }))
  }
}