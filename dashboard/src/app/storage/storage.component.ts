import { Component, OnInit } from '@angular/core';
import { AwsService } from '../aws.service';

declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'app-storage',
  templateUrl: './storage.component.html',
  styleUrls: ['./storage.component.css']
})
export class StorageComponent implements OnInit {
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

  constructor(private awsService: AwsService) {
    this.awsService.getNumberOfS3Buckets().subscribe(data => {
      this.s3Buckets = data ? data : 0;
    }, err => {
      this.s3Buckets = 0;
    });

    this.awsService.getEBS().subscribe(data => {
      let sum = 0
      Object.keys(data.family).forEach(key => {
        sum +=data.family[key]
      })
      this.ebsTotal = sum;
      this.ebsTotalSize = data.total;
      this.ebsUsed = data.state['in-use'];
    }, err => {
      this.ebsTotal = 0
      this.ebsTotalSize = '0 KB'
      this.ebsUsed = 0
    });

    this.awsService.getBucketObjects().subscribe(data => {
      let emptyBucket = 0;
      let datasets = []
      let total = 0;
      data.forEach(bucket => {
        if (bucket && bucket.Bucket) {
          if (bucket.Datapoints.length == 0) {
            emptyBucket++;
          } else {
            let color = this.dynamicColors()
            let dataset = {
              label: bucket.Bucket,
              backgroundColor: color,
              borderColor: color,
              fill: false,
              borderWidth: 1,
              pointStyle: 'line',
              data: []
            }
            let data = []
            bucket.Datapoints.forEach(dt => {
              data.push({
                x: new Date(dt.timestamp),
                y: dt.value
              })
            })
            dataset.data = data
            datasets.push(dataset)

            total += bucket.Datapoints[bucket.Datapoints.length - 1].value
          }
        }
      })

      this.emptyBuckets = emptyBucket;

      this.showS3BucketsObjects(datasets);

      this.s3BucketSize = this.bytesToSizeWithUnit(total);

    }, err => {
      this.emptyBuckets = 0;
      this.s3BucketSize = '0 KB';
    });

    this.awsService.getBucketSize().subscribe(data => {
      let datasets = []
      let total = 0;
      data.forEach(bucket => {
        if (bucket && bucket.Bucket) {
          if (bucket.Datapoints.length > 0) {
            let color = this.dynamicColors()
            let dataset = {
              label: bucket.Bucket,
              backgroundColor: color,
              borderColor: color,
              fill: false,
              borderWidth: 1,
              pointStyle: 'line',
              data: []
            }
            let data = []
            bucket.Datapoints.forEach(dt => {
              data.push({
                x: new Date(dt.timestamp),
                y: this.bytesToSize(dt.value)
              })
            })
            dataset.data = data
            datasets.push(dataset)
            total += bucket.Datapoints[bucket.Datapoints.length - 1].value
          }
        }
      })

      this.showS3BucketsSize(datasets);

      this.s3BucketObjects = this.formatNumber(total);
    }, err => {
      this.s3BucketObjects = '0';
    });


    this.awsService.getDynamoDBTables().subscribe(data => {
      this.dynamodbTables = data.total;
    }, err => {
      this.dynamodbTables = 0;
    });

    this.awsService.getRDSInstances().subscribe(data => {
      this.docdbInstances = data.docdb;
      let total = 0;
      Object.keys(data).forEach(key => {
        if(key != "docdb"){
          total += data[key]
        }
      })
      this.rdsInstances = total;
    }, err => {
      this.docdbInstances = 0;
      this.rdsInstances = 0;
    });

    this.awsService.getElasticacheClusters().subscribe(data => {
      this.redisClusters = data['redis'] ? data['redis'] : 0;
      this.memcachedClusters = data['memcached'] ? data['memcached'] : 0;
    }, err => {
      this.redisClusters = 0;
      this.memcachedClusters = 0;
    });
  }

  ngOnInit() {
    this.showS3BucketsSize([]);
    this.showS3BucketsObjects([]);
    this.showEBSFamily();
  }

  private showEBSFamily(){
    var data = {
      series: [139, 10]
    };
    
    var sum = function(a, b) { return a + b };
    
    new Chartist.Pie('#ebs-family', data, {
      labelInterpolationFnc: function(value) {
        return Math.round(value / data.series.reduce(sum) * 100) + '%';
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

  private dynamicColors() {
    var r = Math.floor(Math.random() * 255);
    var g = Math.floor(Math.random() * 255);
    var b = Math.floor(Math.random() * 255);
    return "rgba(" + r + "," + g + "," + b + ", 0.5)";
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes','KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
    return Math.round(bytes / Math.pow(1024, i), 2) + ' ' + sizes[i];
 };

  private bytesToSize(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
    return Math.round(bytes / Math.pow(1024, i), 2)
 };

 private showS3BucketsObjects(datasets){
    var config = {
      type: 'line',
      data: {
        datasets: datasets
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        title: {
          display: false
        },
        legend: {
          display: false
        },
        tooltips: {
          mode: 'index',
          intersect: true
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
            },
            ticks: {
              autoSkip: true,
              maxTicksLimit: 10
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

    var ctx = document.getElementById('s3BucketObjects').getContext('2d');
    new Chart(ctx, config);
 }  

  private showS3BucketsSize(datasets) {
    var config = {
      type: 'line',
      data: {
        datasets: datasets
      },
      options: {
        point:{display:false},
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

    var ctx = document.getElementById('s3BucketsSize').getContext('2d');
    new Chart(ctx, config);
  }

}
