import { Component, OnInit, OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import { Subject, Subscription } from 'rxjs';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'aws-storage',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsStorageComponent implements OnInit, OnDestroy {
  private s3BucketsSizeChart: any;
  private s3BucketsObjectsChart: any;
  private ebsFamilyChart: any;
  private logsVolumeChart: any;

  public s3Buckets: number;
  public emptyBuckets: number;
  public s3BucketSize: string;
  public s3BucketObjects: string;
  public ebsTotal: number;
  public ebsTotalSize: string;
  public ebsUsed: number;
  public dynamodbTables: number;
  public rdsInstances: number;
  public docdbInstances: number;
  public memcachedClusters: number;
  public redisClusters: number;
  public logsRetentionPeriod: number;
  public redshiftClusters: number;

  public loadingS3Buckets: boolean = true;
  public loadingS3BucketSize: boolean = true;
  public loadingS3BucketObjects: boolean = true;
  public loadingEmptyBuckets: boolean = true;
  public loadingEbsTotal: boolean = true;
  public loadingEbsTotalSize: boolean = true;
  public loadingEbsUsed: boolean = true;
  public loadingLogsRetentionPeriod: boolean = true;
  public loadingDynamoTables: boolean = true;
  public loadingRdsInstances: boolean = true;
  public loadingDocDbInstances: boolean = true;
  public loadingRedshiftClusters: boolean = true;
  public loadingMemCachedClusters: boolean = true;
  public loadingRedisClusters: boolean = true;
  public loadingS3BucketsSizeChart: boolean = true;
  public loadingS3BucketsObjectsChart: boolean = true;
  public loadingEbsFamilyChart: boolean = true;
  public loadingLogsVolumeChart: boolean = true;

  private _subscription: Subscription;

  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.s3BucketsSizeChart.detach();
      this.s3BucketsObjectsChart.detach();
      this.logsVolumeChart.detach();
      this.ebsFamilyChart.destroy();

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


      this.s3Buckets = 0;
      this.emptyBuckets = 0;
      this.s3BucketSize = '0 KB';
      this.s3BucketObjects = '0';
      this.ebsTotal = 0;
      this.ebsTotalSize = '0 KB';
      this.ebsUsed = 0;
      this.dynamodbTables = 0;
      this.rdsInstances = 0;
      this.docdbInstances = 0;
      this.memcachedClusters = 0;
      this.redisClusters = 0;
      this.logsRetentionPeriod = 0;
      this.redshiftClusters = 0;

      this.loadingS3Buckets = true;
      this.loadingS3BucketSize = true;
      this.loadingS3BucketObjects = true;
      this.loadingEmptyBuckets = true;
      this.loadingEbsTotal = true;
      this.loadingEbsTotalSize = true;
      this.loadingEbsUsed = true;
      this.loadingLogsRetentionPeriod = true;
      this.loadingDynamoTables = true;
      this.loadingRdsInstances = true;
      this.loadingDocDbInstances = true;
      this.loadingRedshiftClusters = true;
      this.loadingMemCachedClusters = true;
      this.loadingRedisClusters = true;
      this.loadingS3BucketsSizeChart = true;
      this.loadingS3BucketsObjectsChart = true;
      this.loadingEbsFamilyChart = true;
      this.loadingLogsVolumeChart = true;

      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.awsService.getNumberOfS3Buckets().subscribe(data => {
      this.s3Buckets = data ? data : 0;
      this.loadingS3Buckets = false;
    }, err => {
      this.s3Buckets = 0;
      this.loadingS3Buckets = false;
    });

    this.awsService.getEBS().subscribe(data => {
      let sum = 0
      let labels = []
      let series = []
      Object.keys(data.family).forEach(key => {
        sum += data.family[key]
        labels.push(key);
        series.push(data.family[key])
      })
      this.ebsTotal = sum;
      this.loadingEbsTotal = false;
      this.ebsTotalSize = this.bytesToSizeWithUnit(data.total * 1024 * 1024);
      this.loadingEbsTotalSize = false;
      this.ebsUsed = data.state['in-use'];
      this.loadingEbsUsed = false;
      this.loadingEbsFamilyChart = false;
      this.showEBSFamily(labels, series);
    }, err => {
      this.ebsTotal = 0
      this.ebsTotalSize = '0 KB'
      this.ebsUsed = 0
      this.loadingEbsTotal = false;
      this.loadingEbsTotalSize = false;
      this.loadingEbsUsed = false;
      this.loadingEbsFamilyChart = false;
    });

    this.awsService.getBucketObjects().subscribe(data => {
      let total = 0;
      Object.keys(data).forEach(region => {
        total += data[region][Object.keys(data[region])[Object.keys(data[region]).length - 1]]
      })

      let labels = [];
      let i = 0;
      let series = [];
      Object.keys(data).forEach(region => {
        let serie = [];
        Object.keys(data[region]).forEach(timestamp => {
          serie.push({
            meta: region, value: data[region][timestamp]
          })
          if (i == 0) {
            labels.push(timestamp)
          }
        })
        series.push(serie)
        i++;
      })

      this.loadingS3BucketObjects = false;
      this.s3BucketObjects = this.formatNumber(total).toString();
      this.loadingS3BucketsObjectsChart = false;

      this.showS3BucketsObjects(labels, series);

    }, err => {
      this.emptyBuckets = 0;
      this.s3BucketSize = '0 KB';
      this.loadingS3BucketObjects = false;
      this.loadingS3BucketsObjectsChart = false;
    });

    this.awsService.getBucketSize().subscribe(data => {
      let labels = [];
      let i = 0;
      let series = [];
      Object.keys(data).forEach(region => {
        let serie = [];
        Object.keys(data[region]).forEach(timestamp => {
          serie.push({
            meta: region, value: data[region][timestamp]
          })
          if (i == 0) {
            labels.push(timestamp)
          }
        })
        series.push(serie)
        i++;
      })

      let total = 0;
      Object.keys(data).forEach(region => {
        total += data[region][Object.keys(data[region])[Object.keys(data[region]).length - 1]]
      });

      this.loadingS3BucketsSizeChart = false;
      this.loadingS3BucketSize = false;
      this.s3BucketSize = this.bytesToSizeWithUnit(total);
      this.showS3BucketsSize(labels, series);
    }, err => {
      this.s3BucketObjects = '0';
      this.loadingS3BucketsSizeChart = false;
      this.loadingS3BucketSize = false;
    });

    this.awsService.getEmptyBuckets().subscribe(data => {
      this.emptyBuckets = data;
      this.loadingEmptyBuckets = false;
    }, err => {
      this.emptyBuckets = 0;
      this.loadingEmptyBuckets = false;
    });


    this.awsService.getDynamoDBTables().subscribe(data => {
      this.dynamodbTables = data.total;
      this.loadingDynamoTables = false;
    }, err => {
      this.dynamodbTables = 0;
      this.loadingDynamoTables = false;
    });

    this.awsService.getRDSInstances().subscribe(data => {
      this.docdbInstances = data.docdb ? data.docdb : 0;
      let total = 0;
      Object.keys(data).forEach(key => {
        if (key != "docdb") {
          total += data[key]
        }
      })
      this.loadingDocDbInstances = false;
      this.loadingRdsInstances = false;
      this.rdsInstances = total;
    }, err => {
      this.loadingRdsInstances = false;
      this.docdbInstances = 0;
      this.rdsInstances = 0;
    });

    this.awsService.getElasticacheClusters().subscribe(data => {
      this.redisClusters = data['redis'] ? data['redis'] : 0;
      this.memcachedClusters = data['memcached'] ? data['memcached'] : 0;
      this.loadingRedisClusters = false;
      this.loadingMemCachedClusters = false;
    }, err => {
      this.redisClusters = 0;
      this.memcachedClusters = 0;
      this.loadingRedisClusters = false;
      this.loadingMemCachedClusters = false;
    });

    this.awsService.getLogsVolume().subscribe(data => {
      let seriesIncomingBytes = [];
      let seriesIncomingLogEvents = [];
      let labels = [];

      Object.keys(data[0].Datapoints).forEach(key => {
        labels.push(key)
        seriesIncomingBytes.push({
          meta: 'IncomingBytes',
          value: data[0].Datapoints[key]
        })
        seriesIncomingLogEvents.push({
          meta: 'IncomingLogEvents',
          value: data[1].Datapoints[key]
        })
      })

      this.loadingLogsVolumeChart = false;
      this.showLogsVolume(labels, [
        seriesIncomingBytes,
        seriesIncomingLogEvents
      ]);
    }, err => {
      this.loadingLogsVolumeChart = false;
    });

    this.awsService.getLogsRetentionPeriod().subscribe(data => {
      this.logsRetentionPeriod = data;
      this.loadingLogsRetentionPeriod = false;
    }, err => {
      this.logsRetentionPeriod = 0;
      this.loadingLogsRetentionPeriod = false;
    });

    this.awsService.getRedshiftClusters().subscribe(data => {
      this.redshiftClusters = data;
      this.loadingRedshiftClusters = false;
    }, err => {
      this.redshiftClusters = 0;
      this.loadingRedshiftClusters = false;
    })
  }


  private showS3BucketsObjects(labels, series) {
    let scope = this;
    this.s3BucketsObjectsChart = new Chartist.Bar('#s3BucketsObjectsChart', {
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
            return scope.formatNumber(value);
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

  private showS3BucketsSize(labels, series) {
    let scope = this;
    this.s3BucketsSizeChart = new Chartist.Bar('#s3BucketsSizeChart', {
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

  private showLogsVolume(labels, series) {
    let scope = this;
    this.logsVolumeChart = new Chartist.Bar('#logsVolumeChart', {
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

  ngOnInit() { }

  private showEBSFamily(labels, series) {
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
        data: series
      }]

    };

    let canvas: any = document.getElementById('ebsFamilyChart');
    var ctx = canvas.getContext('2d');
    this.ebsFamilyChart = new Chart(ctx, {
      type: 'pie',
      data: barChartData,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        legend: {
          position: 'top',
        },
      }
    });
  }

  private formatNumber(labelValue: number) {

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

  private dynamicColors() {
    var r = Math.floor(Math.random() * 255);
    var g = Math.floor(Math.random() * 255);
    var b = Math.floor(Math.random() * 255);
    return "rgba(" + r + "," + g + "," + b + ", 0.5)";
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

  private bytesToSize(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i))
  };


}
