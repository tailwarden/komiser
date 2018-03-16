import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";

@Injectable()
export class AWSService {
  
  constructor(private http: Http) { }

  public getBilling(){
    return this.http
     .get(`/cost`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentVPCs(){
    return this.http
     .get(`/vpc`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentACLs(){
    return this.http
     .get(`/acl`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentSecurityGroups(){
    return this.http
     .get(`/security_group`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentNatGateways(){
    return this.http
     .get(`/nat`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentInternetGateways(){
    return this.http
     .get(`/internet_gateway`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentElasticIPs(){
    return this.http
     .get(`/eip`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentKeyPairs(){
    return this.http
     .get(`/key_pair`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentAutoscalingGroups(){
    return this.http
     .get(`/autoscaling_group`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentRouteTables(){
    return this.http
     .get(`/route_table`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentDynamoDBTables(){
    return this.http
     .get(`/dynamodb`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentEBSVolumes(){
    return this.http
     .get(`/ebs`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentEC2Instances(){
    return this.http
     .get(`/ec2`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentSnapshots(){
    return this.http
     .get(`/snapshot`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentLambdaFunctions(){
    return this.http
     .get(`/lambda`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentElasticLoadBalancers(){
     return this.http
     .get(`/elb`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentS3Buckets(){
     return this.http
     .get(`/s3`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentSQSQueues(){
     return this.http
     .get(`/sqs`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentSNSTopics(){
     return this.http
     .get(`/sns`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentHostedZones(){
     return this.http
     .get(`/hosted_zone`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentIAMRoles(){
     return this.http
     .get(`/iam/role`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentIAMPolicies(){
     return this.http
     .get(`/iam/policy`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentIAMGroups(){
     return this.http
     .get(`/iam/group`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentIAMUsers(){
     return this.http
     .get(`/iam/user`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentCloudwatchAlarms(){
     return this.http
     .get(`/cloudwatch`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }

  public getCurrentCloudFrontDistributions(){
     return this.http
     .get(`/cloudfront`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
       return Observable.throw(err.json().error)
     })
  }
}
