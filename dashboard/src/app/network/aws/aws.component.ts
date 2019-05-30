import { Component, OnInit, OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import { Subject, Subscription } from 'rxjs';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
declare var moment: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'aws-network',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsNetworkComponent implements OnInit, OnDestroy {
  private cloudfrontRequests:any;
  private apigatewayRequests:any;
  private elbRequests: any;
  private elbFamilyType:any;
  private natGatewayChartTraffic:any;

  public vpcNumber: number;
  public aclNumber: number;
  public subnetNumbers: number;
  public routeTablesNumber: number;
  public cloudfrontDistributions: number;
  public cdnYesterdayRequests: string;
  public cdnTodayRequests: string;
  public apigatewayYesterdayRequests: string;
  public apigatewayTodayRequests: string;
  public apigatewayApis: number;
  public elbYesterdayRequests: string;
  public elbTodayRequests: string;
  public loadBalancers: number;
  public route53Records: number;
  public route53Zones: number;
  public natGatewayAvailableRegions: Array<String> = [];
  public natGatewayTraffic: Array<any> = [];

  public loadingVPCNumbers: boolean = true;
  public loadingACLNumbers: boolean = true;
  public loadingSubnetNumbers: boolean = true;
  public loadingRouteTablesNumber: boolean = true;
  public loadingCDNNumbers: boolean = true;
  public loadingCDNRequests: boolean = true;
  public loadingAPIGateways: boolean = true;
  public loadingAPIRequests: boolean = true;
  public loadingELBNumber: boolean = true;
  public loadingElbRequests: boolean = true;
  public loadingRoute53Zones: boolean = true;
  public loadingRoute53ARecords: boolean = true;
  public loadingCloudfrontRequestsChart: boolean = true;
  public loadingApigatewayRequestsChart: boolean = true;
  public loadingElbRequestsChart: boolean = true;
  public loadingNatGatewayTrafficChart: boolean = true;
  public loadingElbFamilyType: boolean = true;

  private _subscription: Subscription;

  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();
    
    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.cloudfrontRequests.destroy();
      this.apigatewayRequests.destroy();
      this.elbRequests.destroy();
      this.elbFamilyType.destroy();
      //this.natGatewayChartTraffic.detach();

      let tooltips = document.getElementsByClassName('chartist-tooltip')
      for (let i = 0; i < tooltips.length; i++) {
        tooltips[i].outerHTML = ""
      }
      for (let j = 0; j < 3; j++) {
        let charts = document.getElementsByTagName('svg');
        for (let i = 0; i < charts.length; i++) {
          charts[i].outerHTML = ""
        }
      }

      this.vpcNumber = 0;
      this.aclNumber = 0;
      this.subnetNumbers = 0;
      this.routeTablesNumber = 0;
      this.cloudfrontDistributions = 0;
      this.cdnYesterdayRequests = '0';
      this.cdnTodayRequests = '0';
      this.apigatewayYesterdayRequests = '0';
      this.apigatewayTodayRequests = '0';
      this.apigatewayApis = 0;
      this.elbYesterdayRequests = '0';
      this.elbTodayRequests = '0';
      this.loadBalancers = 0;
      this.route53Records = 0;
      this.route53Zones = 0;
      this.natGatewayAvailableRegions = [];
      this.natGatewayTraffic = [];

      this.loadingVPCNumbers= true;
      this.loadingACLNumbers = true;
      this.loadingSubnetNumbers = true;
      this.loadingRouteTablesNumber = true;
      this.loadingCDNNumbers = true;
      this.loadingCDNRequests = true;
      this.loadingAPIGateways = true;
      this.loadingAPIRequests = true;
      this.loadingELBNumber = true;
      this.loadingElbRequests = true;
      this.loadingRoute53Zones = true;
      this.loadingRoute53ARecords = true;
      this.loadingCloudfrontRequestsChart = true;
      this.loadingApigatewayRequestsChart = true;
      this.loadingElbRequestsChart = true;
      this.loadingNatGatewayTrafficChart = true;
      this.loadingElbFamilyType = true;

      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.awsService.getVirtualPrivateClouds().subscribe(data => {
      this.vpcNumber = data;
      this.loadingVPCNumbers = false;
    }, err => {
      this.vpcNumber = 0;
      this.loadingVPCNumbers = false;
    });

    this.awsService.getAccessControlLists().subscribe(data => {
      this.aclNumber = data;
      this.loadingACLNumbers = false;
    }, err => {
      this.aclNumber = 0;
      this.loadingACLNumbers = false;
    });

    this.awsService.getRouteTables().subscribe(data => {
      this.routeTablesNumber = data;
      this.loadingRouteTablesNumber = false;
    }, err => {
      this.routeTablesNumber = 0;
      this.loadingRouteTablesNumber = false;
    });

    this.awsService.getCloudFrontDistributions().subscribe(data => {
      this.cloudfrontDistributions = data;
      this.loadingCDNNumbers = false;
    }, err => {
      this.cloudfrontDistributions = 0;
      this.loadingCDNNumbers = false;
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

      this.loadingCDNRequests = false;
      this.cdnYesterdayRequests = this.formatNumber(yesterdayRequests).toString();
      this.cdnTodayRequests = this.formatNumber(todayRequests).toString();
      this.loadingCloudfrontRequestsChart = false;
      this.showCloudFrontRequests(datasets);
    }, err => {
      this.cdnTodayRequests = '0';
      this.cdnYesterdayRequests = '0';
      this.loadingCloudfrontRequestsChart = false;
      this.loadingCDNRequests = false;
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
      this.loadingAPIRequests = false;
      this.loadingApigatewayRequestsChart = false;
      this.apigatewayYesterdayRequests = this.formatNumber(yesterdayRequests).toString();
      this.apigatewayTodayRequests = this.formatNumber(todayRequests).toString()
      this.showApiGatewayRequests(datasets);
    }, err => {
      this.apigatewayYesterdayRequests = '0';
      this.apigatewayTodayRequests = '0';
      this.loadingAPIRequests = false;
      this.loadingApigatewayRequestsChart = false;
    });

    this.awsService.getApiGatewayRestAPIs().subscribe(data => {
      this.apigatewayApis = data;
      this.loadingAPIGateways = false;
    }, err => {
      this.apigatewayApis = 0;
      this.loadingAPIGateways = false;
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
      this.loadingElbRequests = false;
      this.loadingElbRequestsChart = false;
      this.elbYesterdayRequests = this.formatNumber(yesterdayRequests).toString();
      this.elbTodayRequests = this.formatNumber(todayRequests).toString();
      this.showELBRequests(datasets);
    }, err => {
      this.elbYesterdayRequests = '0';
      this.elbTodayRequests = '0';
      this.loadingElbRequests = false;
      this.loadingElbRequestsChart = false;
    });


    this.awsService.getELBFamily().subscribe(data => {
      let labels = [];
      let dataset = [];
      let total = 0;
      Object.keys(data).forEach(key => {
        labels.push(key)
        dataset.push(data[key]);
        total += data[key];
      });
      this.loadBalancers = total;

      this.loadingElbFamilyType = false;
      this.loadingELBNumber = false;
      this.showELBFamily(labels, dataset);
    }, err => {
      this.loadingElbFamilyType = false;
      this.loadingELBNumber = false;
      this.loadBalancers = 0;
      console.log(err);
    });

    this.awsService.getRoute53Records().subscribe(data => {
      this.route53Records = data;
      this.loadingRoute53ARecords = false;
    }, err => {
      this.route53Records = 0;
      this.loadingRoute53ARecords = false;
    });

    this.awsService.getRoute53Zones().subscribe(data => {
      this.route53Zones = data;
      this.loadingRoute53Zones = false;
    }, err => {
      this.route53Zones = 0;
      this.loadingRoute53Zones = false;
    });

    this.awsService.getNatGatewayTraffic().subscribe(data => {
      this.natGatewayTraffic = data;

      Object.keys(data).forEach(region => {
        this.natGatewayAvailableRegions.push(region);
      });

      if (this.natGatewayAvailableRegions.length > 0) {
        this.showNatGatewayTrafficInRegion(this.natGatewayAvailableRegions[0]);
      }

      this.loadingNatGatewayTrafficChart = false;
    }, err => {
      this.natGatewayTraffic = [];
      this.loadingNatGatewayTrafficChart = false;
    });

    this.awsService.getVPCSubnets().subscribe(data => {
      this.subnetNumbers = data;
      this.loadingSubnetNumbers = false;
    }, err => {
      this.subnetNumbers = 0;
      this.loadingSubnetNumbers = false;
    })
  }

  private dynamicColors() {
    var r = Math.floor(Math.random() * 255);
    var g = Math.floor(Math.random() * 255);
    var b = Math.floor(Math.random() * 255);
    return "rgba(" + r + "," + g + "," + b + ", 0.5)";
  }

  ngOnInit() {
  }

  private showNatGatewayTraffic(labels, series) {
    let scope = this;
    this.natGatewayChartTraffic = new Chartist.Bar('#natGatewayChartTraffic', {
      labels: labels,
      series: series
    }, {
        plugins: [
          Chartist.plugins.tooltip()
        ],
        stackBars: true,
        axisY: {
          offset: 80,
          labelInterpolationFnc: function (value) {
            return scope.bytesToSizeWithUnit(value);
          }
        }
      }).on('draw', function (data) {
        if (data.type === 'bar') {
          data.element.attr({
            style: 'stroke-width: 30px'
          });
        }
      });
  }

  public showNatGatewayTrafficInRegion(region: any) {

    let labels = [];
    let series = [];
    let serieBytesIn = [];
    let serieBytesOut = [];
    Object.keys(this.natGatewayTraffic[region]).forEach(timestamp => {
      labels.push(timestamp);
      serieBytesIn.push({
        meta: 'BytesInFromDestination', value: this.natGatewayTraffic[region][timestamp].BytesInFromDestination
      })
      serieBytesOut.push({
        meta: 'BytesOutToDestination', value: this.natGatewayTraffic[region][timestamp].BytesOutToDestination
      })
    })
    series.push(serieBytesIn);
    series.push(serieBytesOut);

    this.showNatGatewayTraffic(labels, series);
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };


  private showELBFamily(labels, dataset) {
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

    var canvas: any = document.getElementById('elbFamilyType');
    var ctx = canvas.getContext('2d');
    this.elbFamilyType = new Chart(ctx, {
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

  private formatNumber(labelValue) {

    // Nine Zeroes for Billions
    return Math.abs(Number(labelValue)) >= 1.0e+9

      ? (Math.abs(Number(labelValue)) / 1.0e+9).toFixed(2) + " B"
      // Six Zeroes for Millions 
      : Math.abs(Number(labelValue)) >= 1.0e+6

        ? (Math.abs(Number(labelValue)) / 1.0e+6).toFixed(2) + " M"
        // Three Zeroes for Thousands
        : Math.abs(Number(labelValue)) >= 1.0e+3

          ? (Math.abs(Number(labelValue)) / 1.0e+3).toFixed(2) + " K"

          : Math.abs(Number(labelValue));

  }

  private showApiGatewayRequests(datasets) {
    let scope = this;
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
              beginAtZero: true,
              callback: function (value, index, values) {
                return scope.formatNumber(value);
              }
            }
          }]
        }
      }
    };

    var canvas: any = document.getElementById('apigatewayRequests');
    var ctx = canvas.getContext('2d');
    this.apigatewayRequests = new Chart(ctx, config);
  }

  private showCloudFrontRequests(datasets) {
    let scope = this;
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
              beginAtZero: true,
              callback: function (value, index, values) {
                return scope.formatNumber(value);
              }
            }
          }]
        }
      }
    };

    var canvas: any = document.getElementById('cloudfrontRequests');
    var ctx = canvas.getContext('2d');
    this.cloudfrontRequests = new Chart(ctx, config);
  }

  private showELBRequests(datasets) {
    let scope = this;
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
              beginAtZero: true,
              callback: function (value, index, values) {
                return scope.formatNumber(value);
              }
            }
          }]
        }
      }
    };

    var canvas: any = document.getElementById('elbRequests');
    var ctx = canvas.getContext('2d');
    this.elbRequests = new Chart(ctx, config);
  }

}
