import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';

@Injectable()
export class AwsService {

  private BASE_URL = 'http://localhost:3000/aws'

  constructor(private http: Http, private storeService: StoreService) { }

  public getCurrentCost(){
    return this.http
     .get(`${this.BASE_URL}/cost/current`)
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
     .get(`${this.BASE_URL}/cost/history`)
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

  public getInstancesPerRegion(){
    return this.http
      .get(`${this.BASE_URL}/ec2/regions`)
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
      .get(`${this.BASE_URL}/resources/regions`)
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
      .get(`${this.BASE_URL}/cloudwatch/alarms`)
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
      .get(`${this.BASE_URL}/lambda/functions`)
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
      .get(`${this.BASE_URL}/lambda/invocations`)
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
      .get(`${this.BASE_URL}/iam/account`)
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
      .get(`${this.BASE_URL}/s3/buckets`)
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
    .get(`${this.BASE_URL}/s3/objects`)
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
    .get(`${this.BASE_URL}/s3/size`)
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
    .get(`${this.BASE_URL}/ebs`)
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
    .get(`${this.BASE_URL}/rds/instances`)
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
    .get(`${this.BASE_URL}/dynamodb/tables`)
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
    .get(`${this.BASE_URL}/elasticache/clusters`)
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
    .get(`${this.BASE_URL}/vpc`)
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
    .get(`${this.BASE_URL}/acl`)
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
    .get(`${this.BASE_URL}/route_tables`)
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
    .get(`${this.BASE_URL}/cloudfront/requests`)
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
    .get(`${this.BASE_URL}/cloudfront/distributions`)
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
    .get(`${this.BASE_URL}/apigateway/requests`)
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
    .get(`${this.BASE_URL}/apigateway/apis`)
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
    .get(`${this.BASE_URL}/elb/requests`)
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
    .get(`${this.BASE_URL}/elb/family`)
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
    .get(`${this.BASE_URL}/kms`)
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
    .get(`${this.BASE_URL}/security_groups`)
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
    .get(`${this.BASE_URL}/key_pairs`)
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
    .get(`${this.BASE_URL}/acm/certificates`)
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
    .get(`${this.BASE_URL}/acm/expired`)
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
    .get(`${this.BASE_URL}/security_groups/unrestricted`)
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
    .get(`${this.BASE_URL}/sqs/messages`)
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
    .get(`${this.BASE_URL}/sqs/queues`)
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
    .get(`${this.BASE_URL}/sns/topics`)
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
    .get(`${this.BASE_URL}/mq/brokers`)
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
    .get(`${this.BASE_URL}/kinesis/streams`)
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
    .get(`${this.BASE_URL}/kinesis/shards`)
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
    .get(`${this.BASE_URL}/glue/crawlers`)
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
    .get(`${this.BASE_URL}/glue/jobs`)
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
    .get(`${this.BASE_URL}/datapipeline/pipelines`)
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
    .get(`${this.BASE_URL}/es/domains`)
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
    .get(`${this.BASE_URL}/swf/domains`)
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
    .get(`${this.BASE_URL}/support/open`)
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
    .get(`${this.BASE_URL}/support/history`)
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
    .get(`${this.BASE_URL}/ecs`)
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
    .get(`${this.BASE_URL}/route53/records`)
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
    .get(`${this.BASE_URL}/route53/zones`)
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
    .get(`${this.BASE_URL}/logs/volume`)
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
    .get(`${this.BASE_URL}/cloudtrail/sign_in_event`)
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
    .get(`${this.BASE_URL}/lambda/errors`)
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
    .get(`${this.BASE_URL}/ec2/reserved`)
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
    .get(`${this.BASE_URL}/ec2/scheduled`)
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
    .get(`${this.BASE_URL}/ec2/spot`)
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
    .get(`${this.BASE_URL}/cost/instance_type`)
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
    .get(`${this.BASE_URL}/eks/clusters`)
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
    .get(`${this.BASE_URL}/cloudtrail/source_ip`)
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
    .get(`${this.BASE_URL}/logs/retention`)
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
    .get(`${this.BASE_URL}/nat/traffic`)
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
    .get(`${this.BASE_URL}/iam/organization`)
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
    .get(`${this.BASE_URL}/service/limits`)
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
    .get(`${this.BASE_URL}/s3/empty`)
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
    .get(`${this.BASE_URL}/eip/detached`)
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
    .get(`${this.BASE_URL}/redshift/clusters`)
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

  public getForecastPrice(){
    return this.http
    .get(`${this.BASE_URL}/cost/forecast`)
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
