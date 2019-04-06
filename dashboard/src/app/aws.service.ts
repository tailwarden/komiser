import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";

@Injectable()
export class AwsService {

  private BASE_URL = 'http://localhost:3000/aws'

  constructor(private http: Http) { }

  public getCurrentCost(){
    return this.http
     .get(`${this.BASE_URL}/cost/current`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
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
      return Observable.throw(err.json().error)
    })
  }


}
