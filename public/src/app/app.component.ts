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
    responsive: true,
    cutoutPercentage: 70

  };
  public billingChartLegend:boolean = true;
  public billingChartType:string = 'bar';
  public ec2familliesChartLabels:string[] = ['t2.micro', 't2.micro', 'c4.large'];
  public ec2familliesChartData:number[] = [20, 5, 2];
  public ec2familliesChartType:string = 'doughnut';
  public ec2familliesChartLegend:boolean = false;

   public polarAreaChartLabels:string[] = ['Download Sales', 'In-Store Sales', 'Mail Sales', 'Telesales', 'Corporate Sales'];
  public polarAreaChartData:number[] = [300, 500, 100, 40, 120];
  public polarAreaLegend:boolean = true;
  public polarAreaOptions: any = {
    responsive: true
  }
 
  public polarAreaChartType:string = 'polarArea';

   public pieChartLabels:string[] = ['Download Sales', 'In-Store Sales', 'Mail Sales'];
  public pieChartData:number[] = [300, 500, 100];
  public pieChartType:string = 'pie';

  public costAndUsageChartData:Array<any> = [];
  public costAndUsageChartLabels:Array<any> = [];
  public costAndUsageChartOptions:any = {
    responsive: true,
    scales: {
                    xAxes: [{
                        display: !0,
                        gridLines: {
                            display: !1
                        }
                    }],
                    yAxes: [{
                        ticks: {
                            max: 40,
                            min: 0,
                            stepSize: .5
                        },
                        display: !1,
                        gridLines: {
                            display: !1
                        }
                    }]
                },
                legend: {
                    display: !1
                }
  };
  public costAndUsageChartType:string = 'line';

  public barChartOptions:any = {
    responsive: true,
    scales: {
                xAxes: [{
                    display: !1
                }],
                yAxes: [{
                    display: !1
                }]
            },
            legend: {
                display: !1
            }
  };
  public barChartLabels:string[] = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "November", "December"];
  public barChartType:string = 'bar';
 
  public barChartData:any[] = [
    {
      label: "Data Set 1",
                backgroundColor: ["rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)"],
                borderColor: ["rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)", "rgb(121, 106, 238)"],
                borderWidth: 1,
                data: [35, 49, 55, 68, 81, 95, 85, 40, 30, 27, 22, 15]
    }
  ];

  public barChartOptions2:any = {
    responsive: true,
    scales: {
                    xAxes: [{
                        display: !0,
                        gridLines: {
                            display: !1
                        }
                    }],
                    yAxes: [{
                        display: !0,
                        gridLines: {
                            display: !1
                        }
                    }]
                },
                legend: {
                    display: !0
                }
  };
  public barChartLabels2:string[] = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"];
  public barChartType2:string = 'line';
 
  public barChartData2:any[] = [
    {
                    label: "Page Visitors",
                    fill: !0,
                    lineTension: 0,
                    backgroundColor: "transparent",
                    borderColor: "#f15765",
                    pointBorderColor: "#da4c59",
                    pointHoverBackgroundColor: "#da4c59",
                    borderCapStyle: "butt",
                    borderDash: [],
                    borderDashOffset: 0,
                    borderJoinStyle: "miter",
                    borderWidth: 1,
                    pointBackgroundColor: "#fff",
                    pointBorderWidth: 1,
                    pointHoverRadius: 5,
                    pointHoverBorderColor: "#fff",
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 0,
                    data: [50, 20, 60, 31, 52, 22, 40, 25, 30, 68, 56, 40, 60, 43, 55, 39, 47],
                    spanGaps: !1
                }, {
                    label: "Page Views",
                    fill: !0,
                    lineTension: 0,
                    backgroundColor: "transparent",
                    borderColor: "#54e69d",
                    pointHoverBackgroundColor: "#44c384",
                    borderCapStyle: "butt",
                    borderDash: [],
                    borderDashOffset: 0,
                    borderJoinStyle: "miter",
                    borderWidth: 1,
                    pointBorderColor: "#44c384",
                    pointBackgroundColor: "#fff",
                    pointBorderWidth: 1,
                    pointHoverRadius: 5,
                    pointHoverBorderColor: "#fff",
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    data: [20, 7, 35, 17, 26, 8, 18, 10, 14, 46, 30, 30, 14, 28, 17, 25, 17, 40],
                    spanGaps: !1
                }
  ];

  public doughnutChartLabels2:string[] = ["First", "Second", "Third", "Fourth"];
  public doughnutChartData2:number[] = [300, 50, 100, 60];
  public doughnutChartType2:string = 'doughnut';
  public doughnutChartOptions2: any = {
    cutoutPercentage: 80,
                legend: {
                    display: !1
                }
  }

  public currentBill: number;
  public billUnit : string;
  public currentDate: string;

  constructor(private costExplorerService: CostExplorerService){
    this.getCostAndUsage()
  }

  private getCostAndUsage(): void {
    var values = []
    var labels = []
    this.costExplorerService.getBilling().subscribe(res => {
      this.currentBill = res[res.length - 1].Amount.toFixed(2)
      this.billUnit = res[res.length - 1].Unit;
      this.currentDate = `${(new Date().toLocaleString("en-us", { month: "long" }))} ${(new Date()).getFullYear()}`
      res.forEach(item => {
        values.push(item.Amount.toFixed(3))
        labels.push(new Date(item.End).toLocaleString("en-us", { month: "long" }))
      })
      this.costAndUsageChartData = [
        {
          label: "Bill",
          fill: 0,
          lineTension: 0,
          backgroundColor: "transparent",
          borderColor: "#6ccef0",
          pointBorderColor: "#59c2e6",
          pointHoverBackgroundColor: "#59c2e6",
          borderCapStyle: "butt",
          borderDash: [],
          borderDashOffset: 0,
          borderJoinStyle: "miter",
          borderWidth: 3,
          pointBackgroundColor: "#59c2e6",
          pointBorderWidth: 0,
          pointHoverRadius: 4,
          pointHoverBorderColor: "#fff",
          pointHoverBorderWidth: 0,
          pointRadius: 4,
          pointHitRadius: 0,
          data: values,
          spanGaps: 1
        }
      ]
      this.costAndUsageChartLabels = labels
    }) 
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
