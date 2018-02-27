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
  
  public currentVPC: number;
  public currentACL: number;
  public currentSecurityGroup: number;
  public currentNatGateway: number;
  public currentInternetGateway: number;
  public currentElasticIP: number;
  public currentKeyPair: number;
  public currentAutoscalingGroup: number;
  public currentRouteTable: number;
  public currentDynamoDBReadCapacity: number;
  public currentDynamoDBWriteCapacity: number;
  public currentDynamoDBTable: number;

  public ebsFamilliesChartLabels:string[] = [];
  public ebsFamilliesChartData:number[] = [];
  public ebsFamilliesChartType:string = 'doughnut';
  public ebsFamilliesChartLegend:boolean = false;

  public ec2FamilliesChartLabels:string[] = [];
  public ec2FamilliesChartData:number[] = [];
  public ec2FamilliesChartLegend:boolean = true;
  public ec2FamilliesOptions: any = {
    responsive: true
  }
 
  public ec2FamilliesChartType:string = 'polarArea';

  public doughnutChartOptions: any = {
    cutoutPercentage: 80
  }

  constructor(private costExplorerService: CostExplorerService){
    /*this.getCostAndUsage()
    this.getCurrentVPC()
    this.getCurrentACL()
    this.getCurrentSecurityGroup()
    this.getCurrentNatGateway()
    this.getCurrentInternetGateway()
    this.getCurrentElasticIP()
    this.getCurrentKeyPair()
    this.getCurrentAutoscalingGroup()
    this.getCurrentRouteTable()
    this.getCurrentDynamoDBTable()
    this.getCurrentDynamoDBThroughput()*/
    this.getCurrentEBSFamily()
    this.getCurrentEC2Family()
  }

  private getCurrentVPC(): void {
    this.costExplorerService.getCurrentVPC().subscribe(res => {
      this.currentVPC = res;
    })
  }

  private getCurrentACL(): void {
    this.costExplorerService.getCurrentACL().subscribe(res => {
      this.currentACL = res;
    })
  }

  private getCurrentSecurityGroup(): void {
    this.costExplorerService.getCurrentSecurityGroup().subscribe(res => {
      this.currentSecurityGroup = res;
    })
  }

  private getCurrentNatGateway(): void {
    this.costExplorerService.getCurrentNatGateway().subscribe(res => {
      this.currentNatGateway = res;
    })
  }

  private getCurrentInternetGateway(): void {
    this.costExplorerService.getCurrentInternetGateway().subscribe(res => {
      this.currentInternetGateway = res;
    })
  }

  private getCurrentElasticIP(): void {
    this.costExplorerService.getCurrentElasticIP().subscribe(res => {
      this.currentElasticIP = res;
    })
  }

  private getCurrentKeyPair(): void {
    this.costExplorerService.getCurrentKeyPair().subscribe(res => {
      this.currentKeyPair = res;
    })
  }

  private getCurrentAutoscalingGroup(): void {
    this.costExplorerService.getCurrentAutoscalingGroup().subscribe(res => {
      this.currentAutoscalingGroup = res;
    })
  }

  private getCurrentRouteTable(): void {
    this.costExplorerService.getCurrentRouteTable().subscribe(res => {
      this.currentRouteTable = res;
    })
  }

  private getCurrentDynamoDBTable(): void {
    this.costExplorerService.getCurrentDynamoDBTable().subscribe(res => {
      this.currentDynamoDBTable = res;
    })
  }

  private getCurrentDynamoDBThroughput(): void {
    this.costExplorerService.getCurrentDynamoDBThroughput().subscribe(res => {
      this.currentDynamoDBReadCapacity = res.readCapacity;
      this.currentDynamoDBWriteCapacity = res.writeCapacity;
    })
  }

  private getCurrentEBSFamily(): void {
    this.costExplorerService.getCurrentEBSFamily().subscribe(res => {
      for(var i in res){
        this.ebsFamilliesChartLabels.push(i)
        this.ebsFamilliesChartData.push(res[i])
      }
    })
  }

  private getCurrentEC2Family(): void {
    this.costExplorerService.getCurrentEC2Family().subscribe(res => {
      for(var i in res){
        this.ec2FamilliesChartLabels.push(i)
        this.ec2FamilliesChartData.push(res[i])
      }
    })
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
