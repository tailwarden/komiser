import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { environment } from '@env/environment'; 
import 'rxjs/add/operator/map';

@Injectable()
export class CostExplorerService {

  private apiURL : string = environment.API_URL;
  
  constructor(private http: Http) { }

  public getBilling(){
    return this.http
     .get(`${this.apiURL}/cost`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentVPC(){
    return this.http
     .get(`${this.apiURL}/vpc/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentACL(){
    return this.http
     .get(`${this.apiURL}/acl/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSecurityGroup(){
    return this.http
     .get(`${this.apiURL}/security_group/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentNatGateway(){
    return this.http
     .get(`${this.apiURL}/nat/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentInternetGateway(){
    return this.http
     .get(`${this.apiURL}/internet_gateway/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentElasticIP(){
    return this.http
     .get(`${this.apiURL}/eip/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentKeyPair(){
    return this.http
     .get(`${this.apiURL}/key_pair/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentAutoscalingGroup(){
    return this.http
     .get(`${this.apiURL}/autoscaling_group/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentRouteTable(){
    return this.http
     .get(`${this.apiURL}/route_table/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentDynamoDBTable(){
    return this.http
     .get(`${this.apiURL}/dynamodb/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentDynamoDBThroughput(){
    return this.http
     .get(`${this.apiURL}/dynamodb/throughput`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEBSFamily(){
    return this.http
     .get(`${this.apiURL}/ebs/family`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEBSSize(){
    return this.http
     .get(`${this.apiURL}/ebs/size`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2Family(){
    return this.http
     .get(`${this.apiURL}/ec2/family`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2State(){
    return this.http
     .get(`${this.apiURL}/ec2/state`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentEC2Region(){
    return this.http
     .get(`${this.apiURL}/ec2/region`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSnapshot(){
    return this.http
     .get(`${this.apiURL}/snapshot/total`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentSnapshotSize(){
    return this.http
     .get(`${this.apiURL}/snapshot/size`)
     .map(res => {
       return res.json()
     })
  }

  public getCurrentLambdaRuntime(){
    return this.http
     .get(`${this.apiURL}/lambda/runtime`)
     .map(res => {
       return res.json()
     })
  }

}
