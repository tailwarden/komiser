import { Injectable } from '@angular/core';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';

@Injectable()
export class AwsService {

  private BASE_URL = '/aws'

  constructor(private http: Http, private storeService: StoreService) { }

  public getProfiles(){
    return this.http
     .get(`${this.BASE_URL}/profiles`)
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

  public getCurrentCost(){
    return this.http
     .get(`${this.BASE_URL}/cost/current`, {headers: this.getHeaders()})
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

  public getCostAndUsage(){
    return this.http
     .get(`${this.BASE_URL}/cost/history`, {headers: this.getHeaders()})
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
      .get(`${this.BASE_URL}/iam/users`, {headers: this.getHeaders()})
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

  public getInstancesPerRegion(){
    return this.http
      .get(`${this.BASE_URL}/ec2/regions`, {headers: this.getHeaders()})
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

  public getUsedRegions(){
    return this.http
      .get(`${this.BASE_URL}/resources/regions`, {headers: this.getHeaders()})
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

  public getCloudwatchAlarms(){
    return this.http
      .get(`${this.BASE_URL}/cloudwatch/alarms`, {headers: this.getHeaders()})
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

  public getLambdaFunctions(){
    return this.http
      .get(`${this.BASE_URL}/lambda/functions`, {headers: this.getHeaders()})
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

  public getLambdaInvocationMetrics(){
    return this.http
      .get(`${this.BASE_URL}/lambda/invocations`, {headers: this.getHeaders()})
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

  public getAccountName(){
    return this.http
      .get(`${this.BASE_URL}/iam/account`, {headers: this.getHeaders()})
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

  public getNumberOfS3Buckets(){
    return this.http
      .get(`${this.BASE_URL}/s3/buckets`, {headers: this.getHeaders()})
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

  public getBucketObjects(){
    return this.http
    .get(`${this.BASE_URL}/s3/objects`, {headers: this.getHeaders()})
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

  public getBucketSize(){
    return this.http
    .get(`${this.BASE_URL}/s3/size`, {headers: this.getHeaders()})
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

  public getEBS(){
    return this.http
    .get(`${this.BASE_URL}/ebs`, {headers: this.getHeaders()})
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

  public getRDSInstances(){
    return this.http
    .get(`${this.BASE_URL}/rds/instances`, {headers: this.getHeaders()})
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

  public getDynamoDBTables(){
    return this.http
    .get(`${this.BASE_URL}/dynamodb/tables`, {headers: this.getHeaders()})
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

  public getElasticacheClusters(){
    return this.http
    .get(`${this.BASE_URL}/elasticache/clusters`, {headers: this.getHeaders()})
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

  public getVirtualPrivateClouds(){
    return this.http
    .get(`${this.BASE_URL}/vpc`, {headers: this.getHeaders()})
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

  public getAccessControlLists(){
    return this.http
    .get(`${this.BASE_URL}/acl`, {headers: this.getHeaders()})
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

  public getRouteTables(){
    return this.http
    .get(`${this.BASE_URL}/route_tables`, {headers: this.getHeaders()})
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

  public getCloudFrontRequests(){
    return this.http
    .get(`${this.BASE_URL}/cloudfront/requests`, {headers: this.getHeaders()})
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

  public getCloudFrontDistributions(){
    return this.http
    .get(`${this.BASE_URL}/cloudfront/distributions`, {headers: this.getHeaders()})
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

  public getApiGatewayRequests(){
    return this.http
    .get(`${this.BASE_URL}/apigateway/requests`, {headers: this.getHeaders()})
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

  public getApiGatewayRestAPIs(){
    return this.http
    .get(`${this.BASE_URL}/apigateway/apis`, {headers: this.getHeaders()})
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

  public getELBRequests(){
    return this.http
    .get(`${this.BASE_URL}/elb/requests`, {headers: this.getHeaders()})
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

  public getELBFamily(){
    return this.http
    .get(`${this.BASE_URL}/elb/family`, {headers: this.getHeaders()})
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

  public getKMSKeys(){
    return this.http
    .get(`${this.BASE_URL}/kms`, {headers: this.getHeaders()})
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

  public getSecurityGroups(){
    return this.http
    .get(`${this.BASE_URL}/security_groups`, {headers: this.getHeaders()})
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

  public getKeyPairs(){
    return this.http
    .get(`${this.BASE_URL}/key_pairs`, {headers: this.getHeaders()})
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

  public getACMListCertificates(){
    return this.http
    .get(`${this.BASE_URL}/acm/certificates`, {headers: this.getHeaders()})
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

  public getACMExpiredCertificates(){
    return this.http
    .get(`${this.BASE_URL}/acm/expired`, {headers: this.getHeaders()})
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

  public getUnrestrictedSecurityGroups(){
    return this.http
    .get(`${this.BASE_URL}/security_groups/unrestricted`, {headers: this.getHeaders()})
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

  public getSQSPublishedMessagesMetrics(){
    return this.http
    .get(`${this.BASE_URL}/sqs/messages`, {headers: this.getHeaders()})
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

  public getSQSQueues(){
    return this.http
    .get(`${this.BASE_URL}/sqs/queues`, {headers: this.getHeaders()})
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

  public getSNSTopics(){
    return this.http
    .get(`${this.BASE_URL}/sns/topics`, {headers: this.getHeaders()})
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

  public getActiveMQBrokers(){
    return this.http
    .get(`${this.BASE_URL}/mq/brokers`, {headers: this.getHeaders()})
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

  public getKinesisStreams(){
    return this.http
    .get(`${this.BASE_URL}/kinesis/streams`, {headers: this.getHeaders()})
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

  public getKinesisShards(){
    return this.http
    .get(`${this.BASE_URL}/kinesis/shards`, {headers: this.getHeaders()})
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

  public getGlueCrawlers(){
    return this.http
    .get(`${this.BASE_URL}/glue/crawlers`, {headers: this.getHeaders()})
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

  public getGlueJobs(){
    return this.http
    .get(`${this.BASE_URL}/glue/jobs`, {headers: this.getHeaders()})
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

  public getDataPipelines(){
    return this.http
    .get(`${this.BASE_URL}/datapipeline/pipelines`, {headers: this.getHeaders()})
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

  public getESDomains(){
    return this.http
    .get(`${this.BASE_URL}/es/domains`, {headers: this.getHeaders()})
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

  public getSWFDomains(){
    return this.http
    .get(`${this.BASE_URL}/swf/domains`, {headers: this.getHeaders()})
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

  public getOpenSupportTickets(){
    return this.http
    .get(`${this.BASE_URL}/support/open`, {headers: this.getHeaders()})
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

  public getSupportTicketsHistory(){
    return this.http
    .get(`${this.BASE_URL}/support/history`, {headers: this.getHeaders()})
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

  public getECS(){
    return this.http
    .get(`${this.BASE_URL}/ecs`, {headers: this.getHeaders()})
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

  public getRoute53Records(){
    return this.http
    .get(`${this.BASE_URL}/route53/records`, {headers: this.getHeaders()})
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

  public getRoute53Zones(){
    return this.http
    .get(`${this.BASE_URL}/route53/zones`, {headers: this.getHeaders()})
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

  public getLogsVolume(){
    return this.http
    .get(`${this.BASE_URL}/logs/volume`, {headers: this.getHeaders()})
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

  public getConsoleLoginEvents(){
    return this.http
    .get(`${this.BASE_URL}/cloudtrail/sign_in_event`, {headers: this.getHeaders()})
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

  public getLambdaErrors(){
    return this.http
    .get(`${this.BASE_URL}/lambda/errors`, {headers: this.getHeaders()})
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

  public getReservedInstances(){
    return this.http
    .get(`${this.BASE_URL}/ec2/reserved`, {headers: this.getHeaders()})
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

  public getScheduledInstances(){
    return this.http
    .get(`${this.BASE_URL}/ec2/scheduled`, {headers: this.getHeaders()})
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

  public getSpotInstances(){
    return this.http
    .get(`${this.BASE_URL}/ec2/spot`, {headers: this.getHeaders()})
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

  public getCostPerInstanceType(){
    return this.http
    .get(`${this.BASE_URL}/cost/instance_type`, {headers: this.getHeaders()})
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

  public getEKSClusters(){
    return this.http
    .get(`${this.BASE_URL}/eks/clusters`, {headers: this.getHeaders()})
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

  public getConsoleLoginSourceIps(){
    return this.http
    .get(`${this.BASE_URL}/cloudtrail/source_ip`, {headers: this.getHeaders()})
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

  public getLogsRetentionPeriod(){
    return this.http
    .get(`${this.BASE_URL}/logs/retention`, {headers: this.getHeaders()})
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

  public getNatGatewayTraffic(){
    return this.http
    .get(`${this.BASE_URL}/nat/traffic`, {headers: this.getHeaders()})
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

  public getOrganization(){
    return this.http
    .get(`${this.BASE_URL}/iam/organization`, {headers: this.getHeaders()})
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

  public getServiceLimits(){
    return this.http
    .get(`${this.BASE_URL}/service/limits`, {headers: this.getHeaders()})
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

  public getEmptyBuckets(){
    return this.http
    .get(`${this.BASE_URL}/s3/empty`, {headers: this.getHeaders()})
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

  public getDetachedElasticIps(){
    return this.http
    .get(`${this.BASE_URL}/eip/detached`, {headers: this.getHeaders()})
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

  public getRedshiftClusters(){
    return this.http
    .get(`${this.BASE_URL}/redshift/clusters`, {headers: this.getHeaders()})
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
    .get(`${this.BASE_URL}/vpc/subnets`, {headers: this.getHeaders()})
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

  public getForecastPrice(){
    return this.http
    .get(`${this.BASE_URL}/cost/forecast`, {headers: this.getHeaders()})
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

  private getHeaders(){
    let headers = new Headers();
    headers.append('profile', localStorage.getItem('profile'));
    return headers;
  }
}
