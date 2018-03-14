import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';

@Injectable()
export class AWSService {
  
  constructor(private http: Http) { }

  public getBilling(){
    return this.http
     .get(`/cost`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentVPC(){
    return this.http
     .get(`/vpc/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentACL(){
    return this.http
     .get(`/acl/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSecurityGroup(){
    return this.http
     .get(`/security_group/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentNatGateway(){
    return this.http
     .get(`/nat/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentInternetGateway(){
    return this.http
     .get(`/internet_gateway/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentElasticIP(){
    return this.http
     .get(`/eip/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentKeyPair(){
    return this.http
     .get(`/key_pair/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentAutoscalingGroup(){
    return this.http
     .get(`/autoscaling_group/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentRouteTable(){
    return this.http
     .get(`/route_table/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentDynamoDBTable(){
    return this.http
     .get(`/dynamodb/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentDynamoDBThroughput(){
    return this.http
     .get(`/dynamodb/throughput`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEBSFamily(){
    return this.http
     .get(`/ebs/family`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEBSSize(){
    return this.http
     .get(`/ebs/size`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2Family(){
    return this.http
     .get(`/ec2/family`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2State(){
    return this.http
     .get(`/ec2/state`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2Region(){
    return this.http
     .get(`/ec2/region`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSnapshot(){
    return this.http
     .get(`/snapshot/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSnapshotSize(){
    return this.http
     .get(`/snapshot/size`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentLambdaRuntime(){
    return this.http
     .get(`/lambda/runtime`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentELBFamily(){
     return this.http
     .get(`/elb/family`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentS3Buckets(){
     return this.http
     .get(`/s3/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSQSQueues(){
     return this.http
     .get(`/sqs/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSNSTopics(){
     return this.http
     .get(`/sns/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentHostedZones(){
     return this.http
     .get(`/hosted_zone/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentIAMRoles(){
     return this.http
     .get(`/role/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentIAMPolicies(){
     return this.http
     .get(`/policy/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentIAMGroups(){
     return this.http
     .get(`/group/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentIAMUsers(){
     return this.http
     .get(`/user/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentCloudwatchAlarmsState(){
     return this.http
     .get(`/cloudwatch/state`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentCloudFrontDistributions(){
     return this.http
     .get(`/cloudfront/total`)
     .map(res => {
       return res.json()
     })
  }
}
