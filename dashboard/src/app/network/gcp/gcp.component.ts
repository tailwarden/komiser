import { Component, OnInit } from '@angular/core';
import { GcpService } from '../../gcp.service';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
declare var moment: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'gcp-network',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpNetworkComponent implements OnInit {

  public dnsZones: number;
  public networks: number;
  public routers: number;
  public loadBalancers: number;
  public subnets: number;
  public externalAddresses: number;
  public aRecords: number;
  public natGateways: number;

  public loadingDnsZones: boolean = true;
  public loadingNetworks: boolean = true;
  public loadingRouters: boolean = true;
  public loadingLBRequestsChart: boolean = true;
  public loadingConsumedRequestsChart: boolean = true;
  public loadingLoadBalancers: boolean = true;
  public loadingSubnets: boolean = true;
  public loadingExternalAddresses: boolean = true;
  public loadingARecords: boolean = true;
  public loadingNatGateways: boolean = true;
  
  constructor(private gcpService: GcpService) {
    this.gcpService.getDNSZones().subscribe(data => {
      this.dnsZones = data;
      this.loadingDnsZones = false;
    }, err => {
      this.dnsZones = 0;
      this.loadingDnsZones = false;
    });

    this.gcpService.getVpcNetworks().subscribe(data => {
      this.networks = data;
      this.loadingNetworks = false;
    }, err => {
      this.networks = 0;
      this.loadingNetworks = false;
    });

    this.gcpService.getVpcRouters().subscribe(data => {
      this.routers = data;
      this.loadingRouters = false;
    }, err => {
      this.routers = 0;
      this.loadingRouters = false;
    });

    this.gcpService.getLBRequests().subscribe(data => {
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
                meta: 'Requests',
                value: point.value.int64Value
              })
              found = true
            }
          })
          if(!found){
            serie.push({
              meta: 'Requests',
              value: 0
            })
          }  
        })   
           
        total +=(serie[serie.length - 1].value / 1024)/1024;
        series.push(serie)
      });

      this.loadingLBRequestsChart = false;
      this.showLBRequests(availablePeriods, series);
    }, err => {
      this.loadingLBRequestsChart = false;
    });

    this.gcpService.getAPIRequests().subscribe(data => {
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
                meta: resource.resource.labels.service,
                value: point.value.int64Value
              })
              found = true
            }
          })
          if(!found){
            serie.push({
              meta: resource.resource.labels.service,
              value: 0
            })
          }  
        })   
           
        total +=(serie[serie.length - 1].value / 1024)/1024;
        series.push(serie)
      });

      this.loadingConsumedRequestsChart = false;
      this.showConsumedRequests(availablePeriods, series);
    }, err => {
      this.loadingConsumedRequestsChart = false;
    });

    this.gcpService.getTotalLoadBalancers().subscribe(data => {
      this.loadBalancers = data;
      this.loadingLoadBalancers = false;
    }, err => {
      this.loadBalancers = 0;
      this.loadingLoadBalancers = false;
    });

    this.gcpService.getVPCSubnets().subscribe(data => {
      this.subnets = data;
      this.loadingSubnets = false;
    }, err => {
      this.subnets = 0;
      this.loadingSubnets = false;
    });

    this.gcpService.getVPCAddresses().subscribe(data => {
      this.externalAddresses = data;
      this.loadingExternalAddresses = false;
    }, err => {
      this.externalAddresses = 0;
      this.loadingExternalAddresses = false;
    });

    this.gcpService.getDnsARecords().subscribe(data => {
      this.aRecords = data;
      this.loadingARecords = false;
    }, err => {
      this.aRecords = 0;
      this.loadingARecords = false;
    });

    this.gcpService.getNatGateways().subscribe(data => {
      this.natGateways = data;
      this.loadingNatGateways = false;
    }, err => {
      this.natGateways = 0;
      this.loadingNatGateways = false;
    });
  }

  private showConsumedRequests(labels, series) {
    let scope = this;
    new Chartist.Bar('#consumedRequestsChart', {
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

  private showLBRequests(labels, series) {
    let scope = this;
    new Chartist.Bar('#lbRequestsChart', {
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


  ngOnInit() {
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


}
