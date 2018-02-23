import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { environment } from '@env/environment'; 
import 'rxjs/add/operator/map';

@Injectable()
export class CostExplorerService {

  private mockData = [
    {
      "start" : "2017-11-01",
      "end" : "2017-11-30",
      "amount" : 1.52,
      "unit" : "USD"
    },
    {
      "start" : "2017-12-01",
      "end" : "2017-12-30",
      "amount" : 2.52,
      "unit" : "USD"
    },
    {
      "start" : "2018-01-01",
      "end" : "2018-01-30",
      "amount" : 2.52,
      "unit" : "USD"
    }
  ]

  private apiURL : string = environment.API_URL;
  
  constructor(private http: Http) { }

  public getBilling(){
    return this.mockData
    /*return this.http
     .get(`${this.apiURL}`)
     .map(res => {
       return res.json()
     })*/
  }
}
