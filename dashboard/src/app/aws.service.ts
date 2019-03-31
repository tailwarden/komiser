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

  public getIAMRoles(){
    return this.http
      .get(`${this.BASE_URL}/iam/roles`)
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

}
