import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';

declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'gcp-storage',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpStorageComponent implements OnInit {

  public storageBuckets: number;
  public disksTotal: number;
  public disksTotalSize: string;
  public mysqlInstances: number;
  public postgresInstances: number;
  public totalSnapshots: number;
  public snapshotsSize: string;
  public storageBucketSize: string;
  public storageBucketObjects: string;
  public imagesSize: string;
  public redisInstances: number;

  public loadingStorageBuckets: boolean = true;
  public loadingDisksTotal: boolean = true;
  public loadingDisksTotalSize: boolean = true;
  public loadingMySQLInstances: boolean = true;
  public loadingPostgresInstances: boolean = true;
  public loadingTotalSnapshots: boolean = true;
  public loadingSnapshotsSize: boolean = true;
  public loadingStorageBucketsSizeChart: boolean = true;
  public loadingStorageBucketsObjectsChart: boolean = true;
  public loadingStorageBucketObjects: boolean = true;
  public loadingStorageBucketSize: boolean = true;
  public loadingLogsVolumeChart: boolean = true;
  public loadingImagesSize: boolean = true;
  public loadingRedisInstances: boolean = true;

  constructor(private gcpService: GcpService) {
    this.gcpService.getStorageBuckets().subscribe(data => {
      this.storageBuckets = data;
      this.loadingStorageBuckets = false;
    }, err => {
      this.storageBuckets = 0;
      this.loadingStorageBuckets = false;
    });

    this.gcpService.getComputeDisks().subscribe(data => {
      let total = 0;
      this.disksTotal = 0;
      data.forEach(disk => {
        this.disksTotal++;
        total+=disk.size;
      });
      this.disksTotalSize = this.bytesToSizeWithUnit(total*1024*1024*1024);
      this.loadingDisksTotal = false;
      this.loadingDisksTotalSize= false;
    }, err => {
      this.disksTotal = 0;
      this.disksTotalSize = '0 KB';
      this.loadingDisksTotal = false;
      this.loadingDisksTotalSize= false;
    });

    this.gcpService.getSqlInstances().subscribe(data => {
      this.mysqlInstances = 0;
      this.postgresInstances = 0;
      this.loadingMySQLInstances = false;
      this.loadingPostgresInstances = false;

      data.forEach(instance => {
        if(instance.kind.startsWith('POSTGRES')){
          this.mysqlInstances++;
        }
        if(instance.kind.startsWith('MYSQL')){
          this.postgresInstances++;
        }
      })
    }, err => {
      this.mysqlInstances = 0;
      this.postgresInstances = 0;
      this.loadingMySQLInstances = false;
      this.loadingPostgresInstances = false;
    });

    this.gcpService.getDiskSnapshots().subscribe(data => {
      this.loadingTotalSnapshots = false;
      this.loadingSnapshotsSize = false;
      this.totalSnapshots = data.length;
      let total = 0;
      data.forEach(snapshot => {
        total+=snapshot.size;
      });
      this.snapshotsSize = this.bytesToSizeWithUnit(total*1024*1024*1024);
    }, err => {
      this.loadingTotalSnapshots = false;
      this.loadingSnapshotsSize = false;
      this.snapshotsSize = '0 KB';
      this.totalSnapshots = 0;
    });

    this.gcpService.getBucketsSize().subscribe(data => {
      let points = data[0].points;
      let labels = [];
      let series = [];
      points.forEach(point => {
        labels.push(new Date(point.interval.endTime).toISOString().split('T')[0]);
        series.push({
          meta: "Size",
          value: point.value.doubleValue
        })
      })
      
      labels = labels.reverse();
      series = series.reverse();

      this.storageBucketSize = this.bytesToSizeWithUnit(series[series.length - 1].value);
      this.loadingStorageBucketSize = false;
      this.loadingStorageBucketsSizeChart = false;
      this.showStorageBucketsSize(labels, [series]);
    }, err => {
      this.storageBucketSize = 'O KB';
      this.loadingStorageBucketSize = false;
      this.loadingStorageBucketsSizeChart = false;
    });

    this.gcpService.getBucketsObjects().subscribe(data => {
      let points = data[0].points;
      let labels = [];
      let series = [];
      points.forEach(point => {
        labels.push(new Date(point.interval.endTime).toISOString().split('T')[0]);
        series.push({
          meta: "Objects",
          value: point.value.int64Value
        })
      })
      
      labels = labels.reverse();
      series = series.reverse();

      this.storageBucketObjects = this.formatNumber(series[series.length - 1].value).toString();
      this.loadingStorageBucketObjects = false;
      this.loadingStorageBucketsObjectsChart = false;
      this.showStorageBucketsObjects(labels, [series]);
    }, err => {
      this.storageBucketObjects = '0';
      this.loadingStorageBucketObjects = false;
      this.loadingStorageBucketsObjectsChart = false;
    });

    this.gcpService.getIngestedLoggingBytes().subscribe(data => {
      let availablePeriods = []
      data.forEach(resource => {
        resource.points.forEach(point => {
          let timestamp = new Date(point.interval.endTime).toISOString().split('T')[0]
          if(!availablePeriods.includes(timestamp)){
            availablePeriods.push(timestamp)
          }
        })
      });

      availablePeriods.sort((a, b) => {
        return new Date(b).getTime() - new Date(a).getTime();
      })
      availablePeriods = availablePeriods.reverse();

      let series = [];
      data.forEach(resource => {
        let serie = []
       
        availablePeriods.forEach(period => {
          let found = false;
          resource.points.forEach(point => {
            let timestamp = new Date(point.interval.endTime).toISOString().split('T')[0]
            if(timestamp == period){
              serie.push({
                meta: resource.metric.labels.resource_type,
                value: point.value.int64Value
              })
              found = true
            }
          })
          if(!found){
            serie.push({
              meta: resource.metric.labels.resource_type,
              value: 0
            })
          }  
        })   
           
        series.push(serie)
      });

      this.loadingLogsVolumeChart = false;
      this.showIngestedBytes(availablePeriods, series);
    }, err => {
      this.loadingLogsVolumeChart = false;
    });

    this.gcpService.getComputeImages().subscribe(data => {
      let total = 0;
      data.forEach(image => {
        total += image.size;
      });
      this.loadingImagesSize = false;
      this.imagesSize = this.bytesToSizeWithUnit(total*1024*1024*1024);
    }, err => {
      this.loadingImagesSize = false;
      this.imagesSize = '0 KB';
    });

    this.gcpService.getRedisInstances().subscribe(data => {
      this.redisInstances = data;
      this.loadingRedisInstances = false;
    }, err => {
      this.redisInstances = 0;
      this.loadingRedisInstances = false;
    })
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

  private showStorageBucketsSize(labels, series) {
    let scope = this;
    new Chartist.Bar('#storageBucketsSizeChart', {
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

  private showStorageBucketsObjects(labels, series) {
    let scope = this;
    new Chartist.Bar('#storageBucketsObjectsChart', {
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

  private showIngestedBytes(labels, series) {
    let scope = this;
    new Chartist.Bar('#logsVolumeChart', {
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

  ngOnInit() {
  }

}
