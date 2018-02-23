import { Component, AfterViewInit } from '@angular/core';
import { CostExplorerService } from '@app/services';

import * as Datamap from 'datamaps';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  public billingChartData:Array<any> = []
  public billingChartLabels:Array<any> = []
  public billingChartOptions:any = {
    responsive: true
  };
  public billingChartColors:Array<any> = [
    { // grey
      backgroundColor: 'rgba(148,159,177,0.2)',
      borderColor: 'rgba(148,159,177,1)',
      pointBackgroundColor: 'rgba(148,159,177,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(148,159,177,0.8)'
    }
  ];
  public billingChartLegend:boolean = true;
  public billingChartType:string = 'line';
  public ec2familliesChartLabels:string[] = ['t2.micro', 't2.micro', 'c4.large'];
  public ec2familliesChartData:number[] = [20, 5, 2];
  public ec2familliesChartType:string = 'pie';
  public ec2familliesChartLegend:boolean = false;

  constructor(private costExplorerService: CostExplorerService){
    let res = this.costExplorerService.getBilling()
    let values = []
    let labels = []
    res.forEach(item => {
      values.push(item.amount)
      labels.push(item.end)
    })
    this.billingChartData = [
      {
        data: values,
        label: 'USD'
      }
    ]
    this.billingChartLabels = labels
    
  }

  ngAfterViewInit(){
    var map = new Datamap({
      scope: 'world',
      responsive: true,
      element: document.getElementById('map'),
      geographyConfig: {
        popupOnHover: false,
        highlightOnHover: false
      },
      fills: {
        defaultFill: '#ABDDA4',
        USA: '#A3ACBB',
        RUS: 'red'
      }
    });
    map.bubbles([
      {
        radius: 5,
        centered: 'BRA',
        country: 'USA',
        yeild: 0,
        fillKey: 'USA',
        instances: 5
      },
      {
        radius: 2,
        yeild: 0,
        country: 'USA',
        centered: 'USA',
        fillKey: 'USA',
        instances: 2
      }
    ], {
      popupTemplate: function(geo, data){
        return `<div class="hoverinfo">${data.country}: ${data.instances} EC2`
      }
    })
  }
}
