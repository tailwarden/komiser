import { Component, OnInit, AfterViewInit, OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
declare var Chart: any;
declare var Circles: any;
import * as $ from "jquery";
import { Subject, Subscription } from 'rxjs';
declare var moment: any;

@Component({
  selector: 'aws-compute',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsComputeComponent implements OnInit, AfterViewInit, OnDestroy {
  private costPerInstanceTypeChart: any;
  private lambdaInvocationsChart: any;
  private lambdaErrorsChart:any;
  private instancesFamilyChart:any;
  private instancesPrivacyChart:any;

  public runningEC2Instances: number = 0;
  public stoppedEC2Instances: number = 0;
  public terminatedEC2Instances: number = 0;
  public lambdaFunctions: any;
  public ecsServices: number = 0;
  public ecsTasks: number = 0;
  public ecsClusters: number = 0;
  public reservedInstances: number = 0;
  public spotInstances: number = 0;
  public scheduledInstances: number = 0;
  public eksClusters: number = 0;
  public detchedElasticIps: number = 0;

  public loadingRunningInstances: boolean = true;
  public loadingStoppedInstances: boolean = true;
  public loadingTerminatedInstances: boolean = true;
  public loadingReservedInstances: boolean = true;
  public loadingSpotInstances: boolean = true;
  public loadingScheduledInstances: boolean = true;
  public loadingDetachedIps: boolean = true;
  public loadingLambdaFunctions: boolean = true;
  public loadingEcsClusters: boolean = true;
  public loadingEcsTasks: boolean = true;
  public loadingEcsServices: boolean = true;
  public loadingEksClusters: boolean = true;
  public loadingInstancesPrivacyChart: boolean = true;
  public loadingInstancesFamilyChart: boolean = true;
  public loadingCostPerInstanceTypeChart: boolean = true;
  public loadingLambdaInvocationsChart: boolean = true;
  public loadingLambdaErrorsChart: boolean = true;

  private _subscription: Subscription;


  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.costPerInstanceTypeChart.detach();
      this.lambdaInvocationsChart.detach();
      this.lambdaErrorsChart.detach();
      this.instancesFamilyChart.destroy();
      this.instancesPrivacyChart.destroy();

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

      this.runningEC2Instances = 0;
      this.stoppedEC2Instances = 0;
      this.terminatedEC2Instances = 0;
      this.lambdaFunctions = {};
      this.ecsServices = 0;
      this.ecsTasks = 0;
      this.ecsClusters = 0;
      this.reservedInstances = 0;
      this.spotInstances = 0;
      this.scheduledInstances = 0;
      this.eksClusters = 0;
      this.detchedElasticIps = 0;

      this.loadingRunningInstances = true;
      this.loadingStoppedInstances = true;
      this.loadingTerminatedInstances = true;
      this.loadingReservedInstances = true;
      this.loadingSpotInstances = true;
      this.loadingScheduledInstances = true;
      this.loadingDetachedIps = true;
      this.loadingLambdaFunctions= true;
      this.loadingEcsClusters= true;
      this.loadingEcsTasks = true;
      this.loadingEcsServices = true;
      this.loadingEksClusters= true;
      this.loadingInstancesPrivacyChart = true;
      this.loadingInstancesFamilyChart= true;
      this.loadingCostPerInstanceTypeChart= true;
      this.loadingLambdaInvocationsChart= true;
      this.loadingLambdaErrorsChart= true;

      this.initState();
    });
  }

  ngOnDestroy() {
     this._subscription.unsubscribe();
   }

  private initState(){
    this.lambdaFunctions = {}

    this.awsService.getDetachedElasticIps().subscribe(data => {
      this.detchedElasticIps = data;
      this.loadingDetachedIps = false;
    }, err => {
      this.loadingDetachedIps = false;
      this.detchedElasticIps = 0;
    });

    this.awsService.getInstancesPerRegion().subscribe(data => {
      this.runningEC2Instances = data.state.running ? data.state.running : 0;
      this.stoppedEC2Instances = data.state.stopped ? data.state.stopped : 0;
      this.terminatedEC2Instances = data.state.terminated ? data.state.terminated : 0;

      this.loadingRunningInstances = false;
      this.loadingStoppedInstances = false;
      this.loadingTerminatedInstances = false;
      this.loadingInstancesPrivacyChart = false;
      this.loadingInstancesFamilyChart = false;

      let labels = [];
      let series = [];
      let colors = []
      Object.keys(data.family).forEach(key => {
        labels.push(key);
        series.push(data.family[key]);
        colors.push(this.getRandomColor());
      })

      this.showInstancesPrivacy([data.public, data.private]);

      this.showInstanceFamilies(labels, series, colors);
    }, err => {
      this.loadingRunningInstances = false;
      this.loadingStoppedInstances = false;
      this.loadingInstancesPrivacyChart = false;
      this.loadingTerminatedInstances = false;
      this.loadingInstancesFamilyChart = false;
      this.runningEC2Instances = 0;
      this.stoppedEC2Instances = 0;
      this.terminatedEC2Instances = 0;
    });

    this.awsService.getLambdaFunctions().subscribe(data => {
      this.lambdaFunctions.golang = data.golang ? data.golang : 0;
      this.lambdaFunctions.ruby = data.ruby ? data.ruby : 0;
      this.lambdaFunctions.java = data.java ? data.java : 0;
      this.lambdaFunctions.csharp = data.csharp ? data.csharp : 0;
      this.lambdaFunctions.python = data.python ? data.python : 0;
      this.lambdaFunctions.node = data.node ? data.node : 0;
      this.lambdaFunctions.custom = data.custom ? data.custom : 0;
      this.loadingLambdaFunctions = false;
    }, err => {
      this.lambdaFunctions = {
        golang: 0,
        ruby: 0,
        java: 0,
        csharp: 0,
        python: 0,
        node: 0,
        custom: 0
      };
      this.loadingLambdaFunctions = false;
    });

    this.awsService.getLambdaInvocationMetrics().subscribe(data => {
      let labels = [];
      data.forEach(period => {
        labels.push(new Date(period.timestamp).toLocaleString('en-us', { month: 'long' }))
      })

      let series = []
      for (let i = 0; i < labels.length; i++) {
        let serie = []
        for (let j = 0; j < labels.length; j++) {
          let item = data[j].metrics[i]
          if(item){
            serie.push({
              meta: item.label, value: item.value
            })
          } else {
            serie.push({
              meta: 'others', value: 0
            })
          }
        }
        series.push(serie)
      }

      this.loadingLambdaInvocationsChart = false;
      this.showLambdaInvocations(labels, series);
    }, err => {
      this.loadingLambdaInvocationsChart = false;
    });

    this.awsService.getECS().subscribe(data => {
      this.ecsServices = data.services;
      this.ecsClusters = data.clusters;
      this.ecsTasks = data.tasks;
      this.loadingEcsClusters = false;
      this.loadingEcsServices = false;
      this.loadingEcsTasks = false;
    }, err => {
      this.ecsServices = 0;
      this.ecsClusters = 0;
      this.ecsTasks = 0;
      this.loadingEcsClusters = false;
      this.loadingEcsServices = false;
      this.loadingEcsTasks = false;
    });

    this.awsService.getLambdaErrors().subscribe(data => {
      let labels = [];
      data.forEach(period => {
        labels.push(new Date(period.timestamp).toISOString().split('T')[0])
      })

      let series = []
      for (let i = 0; i < labels.length; i++) {
        let serie = []
        for (let j = 0; j < labels.length; j++) {
          let item = data[j].metrics[i]
          if(item){
            serie.push({
              meta: item.label, value: item.value
            })
          } else {
            serie.push({
              meta: 'others', value: 0
            })
          }
        }
        series.push(serie)
      }

      this.loadingLambdaErrorsChart = false;
      this.showLambdaErrors(labels, series);
    }, err => {
      this.loadingLambdaErrorsChart = false;
    });

    this.awsService.getReservedInstances().subscribe(data => {
      this.reservedInstances = data;
      this.loadingReservedInstances = false;
    }, err => {
      this.reservedInstances = 0;
      this.loadingReservedInstances = false;
    });

    this.awsService.getScheduledInstances().subscribe(data => {
      this.scheduledInstances = data;
      this.loadingScheduledInstances = false;
    }, err => {
      this.scheduledInstances = 0;
      this.loadingScheduledInstances = false;
    });

    this.awsService.getSpotInstances().subscribe(data => {
      this.spotInstances = data;
      this.loadingSpotInstances= false;
    }, err => {
      this.spotInstances = 0;
      this.loadingSpotInstances= false;
    });

    this.awsService.getEKSClusters().subscribe(data => {
      this.eksClusters = data;
      this.loadingEksClusters = false;
    }, err => {
      this.eksClusters = 0;
      this.loadingEksClusters = false;
    });

    this.awsService.getCostPerInstanceType().subscribe(data => {
      let periods = [];
      let series = []
      data.history.forEach(period => {
        periods.push(new Date(period.start).toLocaleString('en-us', { month: 'long' }));
      });

      for (let i = 0; i < periods.length; i++) {
        let serie = []
        for (let j = 0; j < periods.length; j++) {
          let item = data.history[j].groups[i]
          if(item){
            serie.push({
              meta: item.key, value: item.amount.toFixed(2)
            })
          }else{
            serie.push({
              meta: 'others', value: 0
            })
          }
        }
        series.push(serie)
      }

      this.loadingCostPerInstanceTypeChart = false;

      this.showCostPerInstanceType(periods, series);
    }, err => {
      this.loadingCostPerInstanceTypeChart = false;
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

  ngOnInit() {}

  ngAfterViewInit(): void {}

  private showCostPerInstanceType(labels, series) {
    let scope = this;
    var costHistory = {
      labels: labels,
      series: series
    }

    var optionChartCostHistory = {
      plugins: [
        Chartist.plugins.tooltip()
      ],
      seriesBarDistance: 10,
      axisX: {
        showGrid: false
      },
      axisY: {
        offset: 80,
        labelInterpolationFnc: function(value) {
          return scope.formatNumber(value)
        },
      },
      height: "245px",
    }

    this.costPerInstanceTypeChart = new Chartist.Bar('.costPerInstanceTypeChart', costHistory, optionChartCostHistory);

  }

  private showInstancesPrivacy(series){
    var canvas : any = document.getElementById('instancesPrivacyChart');
    var ctx = canvas.getContext('2d');
    this.instancesPrivacyChart = new Chart(ctx, {
        type: 'pie',
        data: {
          datasets: [{
            data: series,
            backgroundColor: ['#36A2EB','#4BC0C0']
          }],
          labels: ['Public Instances', 'Private Instances']
        },
        options: {}
    });
  }

  private showLambdaErrors(labels, series){
    let scope = this;
    this.lambdaErrorsChart = new Chartist.Line('.lambdaErrorsChart', {
      labels: labels,
      series: series
    }, {
      plugins: [
        Chartist.plugins.tooltip()
      ],
      axisY: {
        offset: 80,
        labelInterpolationFnc: function(value) {
          return scope.formatNumber(value)
        }
      }
    }).on('draw', function(data) {
      if(data.type === 'line') {
        data.element.attr({
          style: 'stroke-width: 1px'
        });
      }
    });
  }

  private showLambdaInvocations(labels, series){
    let scope = this;
    this.lambdaInvocationsChart = new Chartist.Bar('.lambdaInvocationsChart', {
      labels: labels,
      series: series
    }, {
      plugins: [
        Chartist.plugins.tooltip()
      ],
      stackBars: true,
      axisY: {
        offset: 80,
        labelInterpolationFnc: function(value) {
          return scope.formatNumber(value)
        }
      }
    }).on('draw', function(data) {
      if(data.type === 'bar') {
        data.element.attr({
          style: 'stroke-width: 30px'
        });
      }
    });
  }

  private getRandomColor() {
    var letters = '789ABCD'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.round(Math.random() * 6)];
    }
    return color;
  }

  private showInstanceFamilies(labels, series, colors) {
    var color = Chart.helpers.color;
    var config = {
      data: {
        datasets: [{
          data: series,
          backgroundColor: colors,
          label: 'My dataset' // for legend
        }],
        labels: labels,
      },
      options: {
        maintainAspectRatio: false,
        responsive: true,
        legend: {
          position: 'bottom'
        },
        title: {
          display: false
        },
        scale: {
          ticks: {
            beginAtZero: true
          },
          reverse: false
        },
        animation: {
          animateRotate: false,
          animateScale: true
        }
      }
    };

    var ctx = document.getElementById('instancesFamilyChart');
    this.instancesFamilyChart = new Chart.PolarArea(ctx, config);
  }

}
