import { Component, OnInit, AfterViewInit } from '@angular/core';
import { GcpService } from '../../gcp.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';
import * as $ from "jquery";
declare var Chart: any;

@Component({
  selector: 'gcp-dashboard',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpDashboardComponent implements OnInit, AfterViewInit {

  public iamUsers: number = 0;
  public usedRegions: number = 0;
  public forecastBill: string = '0';
  public project: string = 'project id';
  public currentBill: number = 0;
  public openTickets: number = 0;
  public resolvedTickets: number = 0;
  public mostUsedServices: Array<any> = [];

  public loadingProject: boolean = true;
  public loadingCurrentBill: boolean = true;
  public loadingIamUsers: boolean = true;
  public loadingUsedRegions: boolean = true;
  public loadingForecastBill: boolean = true;
  public loadingCostHistoryChart: boolean = true;

  public colors = ['#36A2EB', '#4BBFC0', '#FBAD4B', '#9368E9']

  private zones: Map<string,any> = new Map<string,any>([
    ["asia-east1", {"latitude":"23.697809", "longitude":"120.960518"}],
    ["asia-east2", {"latitude":"22.396427", "longitude":"114.109497"}],
    ["asia-northeast1", {"latitude":"35.689487", "longitude":"139.691711"}],
    ["asia-south1", {"latitude":"19.075983", "longitude":"72.877655"}],
    ["asia-southeast1", {"latitude":"1.339637", "longitude":"103.707339"}],
    ["australia-southeast1", {"latitude":"43.498299", "longitude":"2.375200"}],
    ["europe-north1", {"latitude":"60.568890", "longitude":"27.188188"}],
    ["europe-west1", {"latitude":"50.447748", "longitude":"3.819524"}],
    ["europe-west2", {"latitude":"51.507322", "longitude":"-0.127647"}],
    ["europe-west3", {"latitude":"50.110644", "longitude":"8.682092"}],
    ["europe-west4", {"latitude":"53.448402", "longitude":"6.846503"}],
    ["europe-west6", { "latitude": "47.376888", "longitude": "8.541694" }],
    ["northamerica-northeast1", {"latitude":"45.509060", "longitude":"-73.553360"}],
    ["southamerica-east1", {"latitude":"23.550651", "longitude":"-46.633382"}],
    ["us-central1", {"latitude":"41.262128", "longitude":"-95.861391"}],
    ["us-east1", {"latitude":"33.196003", "longitude":"-80.013137"}],
    ["us-east4", {"latitude":"39.029265", "longitude":"-77.467387"}],
    ["us-west1", {"latitude":"45.601506", "longitude":"-121.184159"}],
    ["us-west2", {"latitude":"34.053691", "longitude":"-118.242767"}],
  ])

  constructor(private gcpService: GcpService) {
    this.gcpService.getProjects().subscribe(data => {
      this.project = data.length > 0 ? data[0].id : this.project;
      this.loadingProject = false;
    }, err => {
      this.project = 'project id';
      this.loadingProject = false;
    });

    let scope = this;

    this.gcpService.getComputeInstances().subscribe(data => {
      let _usedRegions = new Map<string, number>();
      let plots = {};
      
      data.forEach(instance => {
        let region = instance.zone.substring(0, instance.zone.lastIndexOf("-"));
        _usedRegions[region] = (_usedRegions[region] ? _usedRegions[region] : 0) + 1;
      })

      for(var region in _usedRegions){
        this.usedRegions++;
        plots[region] = {
          latitude: scope.zones.get(region).latitude,
          longitude: scope.zones.get(region).longitude,
          value: [_usedRegions[region], 1],
          tooltip: { content: `${region}<br />Instances: ${_usedRegions[region]}` }
        }
      }

      Array.from(this.zones.keys()).forEach(region => {
        let found = false;
        for(let _region in plots){
          if(_region == region){
            found = true;
          }
        }
        if(!found){
          plots[region] = {
            latitude: this.zones.get(region).latitude,
            longitude: this.zones.get(region).longitude,
            value: [_usedRegions[region], 0],
            tooltip: { content: `${region}<br />Instances: 0` }
          }
        }
      });
      
      this.loadingUsedRegions = false;
      this.showComputeInstancesPerRegion(plots);
    }, err => {
      this.loadingUsedRegions = false;
      this.usedRegions = 0;
    });

    this.gcpService.getIAMUsers().subscribe(data => {
      this.loadingIamUsers = false;
      this.iamUsers = data;
    }, err => {
      this.loadingIamUsers = false;
      this.iamUsers = 0;
    });

    this.gcpService.getBillingPerService().subscribe(data => {
      data[data.length - 1].groups.slice(0, 4).forEach(service => {
        this.mostUsedServices.push({
          name: service.service,
          cost: service.cost
        });
      });

      let periods = [];
      let series = []
      data.forEach(period => {
        periods.push(new Date(period.date).toLocaleString('en-us', { month: 'long' }));
      });


      for (let i = 0; i < 6; i++) {
        let serie = []
        for (let j = 0; j < periods.length; j++) {
          let item = data[j].groups[i]
          if(item){
            serie.push({
              meta: item.service, value: item.cost.toFixed(2)
            })
          } else {
            serie.push({
              meta: 'Others', value:0
            })
          }
        }
        series.push(serie)
      }

      this.loadingCostHistoryChart = false;
      this.showLastSixMonth(periods,series);
    }, err => {
      this.mostUsedServices = [];
      this.loadingCostHistoryChart = false;
    });

    this.gcpService.getBillingLastSixMonths().subscribe(data => {
      this.currentBill = data[data.length - 1].cost;
      this.loadingCurrentBill = false;
    }, err => {
      this.loadingCurrentBill = false;
    });
  }

  public showLastSixMonth(labels, series) {
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

    new Chartist.Bar('#costHistoryChart', costHistory, optionChartCostHistory);

  }

  ngOnInit() {
  }

  ngAfterViewInit(): void {
    this.showComputeInstancesPerRegion({});
  }

  public formatNumber(labelValue) {

    // Nine Zeroes for Billions
    return Math.abs(Number(labelValue)) >= 1.0e+9

      ? (Math.abs(Number(labelValue)) / 1.0e+9).toFixed(2) + " B"
      // Six Zeroes for Millions 
      : Math.abs(Number(labelValue)) >= 1.0e+6

        ? (Math.abs(Number(labelValue)) / 1.0e+6).toFixed(2) + " M"
        // Three Zeroes for Thousands
        : Math.abs(Number(labelValue)) >= 1.0e+3

          ? (Math.abs(Number(labelValue)) / 1.0e+3).toFixed(2) + " K"

          : Math.abs(Number(labelValue)).toFixed(2);

  }

  public showComputeInstancesPerRegion(plots) {
    var canvas : any = $(".mapregions");
    canvas.mapael({
      map: {
        name: "world_countries",
        zoom: {
          enabled: true,
          maxLevel: 10
        },
        defaultPlot: {
          attrs: {
            fill: "#004a9b"
            , opacity: 0.6
          }
        },
        defaultArea: {
          attrs: {
            fill: "#e4e4e4"
            , stroke: "#fafafa"
          }
          , attrsHover: {
            fill: "#FBAD4B"
          }
          , text: {
            attrs: {
              fill: "#505444"
            }
            , attrsHover: {
              fill: "#000"
            }
          }
        }
      },
      legend: {
        plot: [
          {
            labelAttrs: {
              fill: "#f4f4e8"
            },
            titleAttrs: {
              fill: "#f4f4e8"
            },
            cssClass: 'density',
            mode: 'horizontal',
            title: "Density",
            marginBottomTitle: 5,
            slices: [{
              label: "< 1",
              max: "0",
              attrs: {
                fill: "#36A2EB"
              },
              legendSpecificAttrs: {
                r: 25
              }
            }, {
              label: "> 1",
              min: "1",
              max: "50000",
              attrs: {
                fill: "#87CB14"
              },
              legendSpecificAttrs: {
                r: 25
              }
            }]
          }
        ]
      },
      plots: plots,
    });
  }

}
