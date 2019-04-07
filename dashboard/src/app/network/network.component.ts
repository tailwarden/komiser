import { Component, OnInit } from '@angular/core';
import { AwsService } from '../aws.service';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
declare var moment: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  styleUrls: ['./network.component.css']
})
export class NetworkComponent implements OnInit {
  public vpcNumber: number;
  public aclNumber: number;
  public routeTablesNumber: number;
  public cloudfrontDistributions: number;
  public cdnYesterdayRequests: number;
  public cdnTodayRequests: number;
  public apigatewayYesterdayRequests: number;
  public apigatewayTodayRequests: number;
  public apigatewayApis: number;
  public elbYesterdayRequests: number;
  public elbTodayRequests: number;

  constructor(private awsService: AwsService) {
    this.awsService.getVirtualPrivateClouds().subscribe(data => {
      this.vpcNumber = data;
    }, err => {
      this.vpcNumber = 0;
    });

    this.awsService.getAccessControlLists().subscribe(data => {
      this.aclNumber = data;
    }, err => {
      this.aclNumber = 0;
    });

    this.awsService.getRouteTables().subscribe(data => {
      this.routeTablesNumber = data;
    }, err => {
      this.routeTablesNumber = 0;
    });

    this.awsService.getCloudFrontDistributions().subscribe(data => {
      this.cloudfrontDistributions = data;
    }, err => {
      this.cloudfrontDistributions = 0;
    })

    this.awsService.getCloudFrontRequests().subscribe(data => {
      let datasets = [];
      let total = 0;

      let todayRequests = 0;
      let yesterdayRequests = 0;

      let todayDate = new Date();
      todayDate.setHours(0, 0, 0, 0);
      let yesterdayDate = new Date(new Date().setDate(new Date().getDate() - 1))
      yesterdayDate.setHours(0, 0, 0, 0);

      data.forEach(distribution => {
        if (distribution && distribution.Distribution) {
          if (distribution.Datapoints.length > 0) {
            let color = this.dynamicColors()
            let dataset = {
              label: distribution.Distribution,
              backgroundColor: color,
              borderColor: color,
              fill: false,
              borderWidth: 1,
              pointStyle: 'line',
              data: []
            }
            let data = []
            distribution.Datapoints.forEach(dt => {
              data.push({
                x: new Date(dt.timestamp),
                y: dt.value
              })

              let dtTimestamp = new Date(dt.timestamp);
              dtTimestamp.setHours(0, 0, 0, 0);

              if (moment(dtTimestamp).isSame(todayDate)) {
                todayRequests += dt.value;
              }

              if (moment(dtTimestamp).isSame(yesterdayDate)) {
                yesterdayRequests += dt.value;
              }
            })
            dataset.data = data
            datasets.push(dataset)
          }
        }
      })
      this.cdnYesterdayRequests = yesterdayRequests;
      this.cdnTodayRequests = todayRequests;
      this.showCloudFrontRequests(datasets);
    }, err => {
      this.cdnTodayRequests = 0;
      this.cdnYesterdayRequests = 0;
    })

    this.awsService.getApiGatewayRequests().subscribe(data => {
      let datasets = [];
      let total = 0;

      let todayRequests = 0;
      let yesterdayRequests = 0;

      let todayDate = new Date();
      todayDate.setHours(0, 0, 0, 0);
      let yesterdayDate = new Date(new Date().setDate(new Date().getDate() - 1))
      yesterdayDate.setHours(0, 0, 0, 0);

      data.forEach(region => {
        if (region && region.Region) {
          if (region.Datapoints.length > 0) {
            let color = this.dynamicColors()
            let dataset = {
              label: region.Region,
              backgroundColor: color,
              borderColor: color,
              fill: false,
              borderWidth: 1,
              pointStyle: 'line',
              data: []
            }
            let data = []
            region.Datapoints.forEach(dt => {
              data.push({
                x: new Date(dt.timestamp),
                y: dt.value
              })

              let dtTimestamp = new Date(dt.timestamp);
              dtTimestamp.setHours(0, 0, 0, 0);

              if (moment(dtTimestamp).isSame(todayDate)) {
                todayRequests += dt.value;
              }

              if (moment(dtTimestamp).isSame(yesterdayDate)) {
                yesterdayRequests += dt.value;
              }
            })
            dataset.data = data
            datasets.push(dataset)
          }
        }
      })
      this.apigatewayYesterdayRequests = yesterdayRequests;
      this.apigatewayTodayRequests = todayRequests;
      this.showApiGatewayRequests(datasets);
    }, err => {
      this.apigatewayYesterdayRequests = 0;
      this.apigatewayTodayRequests = 0;
    });

    this.awsService.getApiGatewayRestAPIs().subscribe(data => {
      this.apigatewayApis = data;
    }, err => {
      this.apigatewayApis = 0;
    });

    this.awsService.getELBRequests().subscribe(data => {
      let datasets = [];
      let total = 0;

      let todayRequests = 0;
      let yesterdayRequests = 0;

      let todayDate = new Date();
      todayDate.setHours(0, 0, 0, 0);
      let yesterdayDate = new Date(new Date().setDate(new Date().getDate() - 1))
      yesterdayDate.setHours(0, 0, 0, 0);

      data.forEach(region => {
        if (region && region.Region) {
          if (region.Datapoints.length > 0) {
            let color = this.dynamicColors()
            let dataset = {
              label: region.Region,
              backgroundColor: color,
              borderColor: color,
              fill: false,
              borderWidth: 1,
              pointStyle: 'line',
              data: []
            }
            let data = []
            region.Datapoints.forEach(dt => {
              data.push({
                x: new Date(dt.timestamp),
                y: dt.value
              })

              let dtTimestamp = new Date(dt.timestamp);
              dtTimestamp.setHours(0, 0, 0, 0);

              if (moment(dtTimestamp).isSame(todayDate)) {
                todayRequests += dt.value;
              }

              if (moment(dtTimestamp).isSame(yesterdayDate)) {
                yesterdayRequests += dt.value;
              }
            })
            dataset.data = data
            datasets.push(dataset)
          }
        }
      })
      this.elbYesterdayRequests = yesterdayRequests;
      this.elbTodayRequests = todayRequests;
      this.showELBRequests(datasets);
    }, err => {
      this.elbYesterdayRequests = 0;
      this.elbTodayRequests = 0;
    });


    this.awsService.getELBFamily().subscribe(data => {
      let labels = [];
      let dataset = [];
      Object.keys(data).forEach(key => {
        labels.push(key)
        dataset.push(data[key]);
      });

      this.showELBFamily(labels, dataset);
    }, err => {
      console.log(err);
    });
  }

  private dynamicColors() {
    var r = Math.floor(Math.random() * 255);
    var g = Math.floor(Math.random() * 255);
    var b = Math.floor(Math.random() * 255);
    return "rgba(" + r + "," + g + "," + b + ", 0.5)";
  }

  ngOnInit() {
  }

  private showELBFamily(labels, dataset){
		var barChartData = {
			labels: labels,
			datasets: [{
				backgroundColor: [
          "#36A2EB",
          "#4BC0C0",
          "#FFCD56",
          "#FF6385"
        ],
				borderWidth: 1,
				data: dataset
			}]

    };
    
    var ctx = document.getElementById('elbFamilyType').getContext('2d');
    window.myBar = new Chart(ctx, {
      type: 'pie',
      data: barChartData,
      options: {
        responsive: true,
        legend: {
          position: 'top',
        },
      }
    });

	
  }

  private showApiGatewayRequests(datasets) {
    var config = {
      type: 'line',
      data: {
        datasets: datasets
      },
      options: {
        point: { display: false },
        responsive: true,
        maintainAspectRatio: false,
        title: {
          display: false
        },
        legend: {
          display: false
        },
        tooltips: {
          enabled: true,
          mode: 'index',
          position: 'nearest',
        },
        hover: {
          mode: 'nearest',
          intersect: false
        },
        scales: {
          xAxes: [{
            type: 'time',
            time: {
              parser: 'YYYY-MM-DD HH:mm:ss',
              unit: 'day',
              unitStepSize: 15,
              displayFormats: {
                'day': 'MMM DD'
              }
            }
          }],
          yAxes: [{
            ticks: {
              beginAtZero: true
            }
          }]
        }
      }
    };

    var ctx = document.getElementById('apigatewayRequests').getContext('2d');
    new Chart(ctx, config);
  }

  private showCloudFrontRequests(datasets) {
    var config = {
      type: 'line',
      data: {
        datasets: datasets
      },
      options: {
        point: { display: false },
        responsive: true,
        maintainAspectRatio: false,
        title: {
          display: false
        },
        legend: {
          display: false
        },
        tooltips: {
          enabled: true,
          mode: 'index',
          position: 'nearest',
        },
        hover: {
          mode: 'nearest',
          intersect: false
        },
        scales: {
          xAxes: [{
            type: 'time',
            time: {
              parser: 'YYYY-MM-DD HH:mm:ss',
              unit: 'day',
              unitStepSize: 20,
              displayFormats: {
                'day': 'MMM DD'
              }
            }
          }],
          yAxes: [{
            ticks: {
              beginAtZero: true
            }
          }]
        }
      }
    };

    var ctx = document.getElementById('cloudfrontRequests').getContext('2d');
    new Chart(ctx, config);
  }

  private showELBRequests(datasets) {
    var config = {
      type: 'line',
      data: {
        datasets: datasets
      },
      options: {
        point: { display: false },
        responsive: true,
        maintainAspectRatio: false,
        title: {
          display: false
        },
        legend: {
          display: false
        },
        tooltips: {
          enabled: true,
          mode: 'index',
          position: 'nearest',
        },
        hover: {
          mode: 'nearest',
          intersect: false
        },
        scales: {
          xAxes: [{
            type: 'time',
            time: {
              parser: 'YYYY-MM-DD HH:mm:ss',
              unit: 'day',
              unitStepSize: 20,
              displayFormats: {
                'day': 'MMM DD'
              }
            }
          }],
          yAxes: [{
            ticks: {
              beginAtZero: true
            }
          }]
        }
      }
    };

    var ctx = document.getElementById('elbRequests').getContext('2d');
    new Chart(ctx, config);
  }

}
