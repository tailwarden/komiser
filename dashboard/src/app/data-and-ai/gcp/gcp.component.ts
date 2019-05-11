import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';

declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'gcp-data-and-ai',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpDataAndAIComponent implements OnInit {

  public pubSubTopics: number = 0;
  public bigquerySize: string;
  public bigqueryDatasets: number;
  public bigqueryTables: number;
  public dataprocClusters: number;
  public dataprocJobs: number;
  public dataflowJobs: number;

  public loadingPubSubTopics: boolean = true;
  public loadingBigqueryStatementsChart: boolean = true;
  public loadingBigqueryStorageChart: boolean = true;
  public loadingBigQuerySize: boolean = true;
  public loadingBigQueryTables: boolean = true;
  public loadingBigQueryDatasets: boolean = true;
  public loadingDataprocClusters: boolean = true;
  public loadingDataprocJobs: boolean = true;
  public loadingDataflowJobs: boolean = true;

  constructor(private gcpService:GcpService) {
    this.gcpService.getPubSubTopics().subscribe(data => {
      this.pubSubTopics = data;
      this.loadingPubSubTopics = false;
    }, err => {
      this.pubSubTopics = 0;
      this.loadingPubSubTopics = false;
    });

    this.gcpService.getBigQueryStatements().subscribe(data => {
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
                meta: resource.metric.labels.statement_type,
                value: point.value.doubleValue
              })
              found = true
            }
          })
          if(!found){
            serie.push({
              meta: resource.metric.labels.statement_type,
              value: 0
            })
          }  
        })   
           
        series.push(serie)
      });

      this.loadingBigqueryStatementsChart = false;
      this.showBigQueryBilledStatements(availablePeriods, series);
    }, err => {
      this.loadingBigqueryStatementsChart = false;
    });

    this.gcpService.getBigQueryStorage().subscribe(data => {
      this.loadingBigqueryStorageChart = false;

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

      let total = 0;
      let series = [];
      data.forEach(resource => {
        let serie = []
       
        availablePeriods.forEach(period => {
          let found = false;
          resource.points.forEach(point => {
            let timestamp = new Date(point.interval.endTime).toISOString().split('T')[0]
            if(timestamp == period){
              serie.push({
                meta: resource.resource.labels.dataset_id,
                value: point.value.int64Value
              })
              found = true
            }
          })
          if(!found){
            serie.push({
              meta: resource.resource.labels.dataset_id,
              value: 0
            })
          }  
        })   
           
        total +=(serie[serie.length - 1].value / 1024)/1024;
        series.push(serie)
      });

      this.bigquerySize = this.bytesToSizeWithUnit(total*1024*1024);

      this.loadingBigQuerySize = false;
      this.showBigQueryStorage(availablePeriods, series);
    }, err => {
      this.loadingBigqueryStorageChart = false;
      this.loadingBigQuerySize = false;
      this.bigquerySize = '0 KB';
    });

    this.gcpService.getBigQueryDatasets().subscribe(data => {
      this.bigqueryDatasets = data;
      this.loadingBigQueryDatasets = false;
    }, err => {
      this.bigqueryDatasets = 0;
      this.loadingBigQueryDatasets = false;
    });

    this.gcpService.getBigQueryTables().subscribe(data => {
      this.bigqueryTables = data;
      this.loadingBigQueryTables = false;
    }, err => {
      this.bigqueryTables = 0;
      this.loadingBigQueryTables = false;
    });

    this.gcpService.getDataprocClusters().subscribe(data => {
      this.dataprocClusters = data;
      this.loadingDataprocClusters = false;
    }, err => {
      this.dataprocClusters = 0;
      this.loadingDataprocClusters = false;
    });

    this.gcpService.getDataprocJobs().subscribe(data => {
      this.dataprocJobs = data;
      this.loadingDataprocJobs = false;
    }, err => {
      this.dataprocJobs = 0;
      this.loadingDataprocJobs = false;
    });

    this.gcpService.getDataflowJobs().subscribe(data => {
      this.dataflowJobs = data;
      this.loadingDataflowJobs = false;
    }, err => {
      this.dataflowJobs = 0;
      this.loadingDataflowJobs = false;
    });
  }

  private showBigQueryBilledStatements(labels, series) {
    let scope = this;
    new Chartist.Bar('#bigqueryStatementsChart', {
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
            return scope.bytesToSizeWithUnit(value)+'/s';
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

  private showBigQueryStorage(labels, series) {
    let scope = this;
    new Chartist.Bar('#bigqueryStorageChart', {
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

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

  ngOnInit() {
  }

}
