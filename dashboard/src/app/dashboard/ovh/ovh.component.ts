import { Component, OnInit, AfterViewInit } from '@angular/core';
import { OvhService } from '../../ovh.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';
import * as $ from "jquery";
declare var Chart: any;

@Component({
  selector: 'ovh-dashboard',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhDashboardComponent implements OnInit, AfterViewInit {
  public projects: number = 0;
  public usedRegions: number = 0;
  public users: number = 0;
  public alerts: number = 0;
  public currentBill: number = 0;
  public openTickets: number = 0;
  public resolvedTickets: number = 0;
  public mostUsedServices: Array<any> = [];

  public loadingProjects: boolean = true;
  public loadingUsedRegions: boolean = true;
  public loadingUsers: boolean = true;
  public loadingAlerts: boolean = true;
  public loadingOpenTickets: boolean = true;
  public loadingResolvedTickets: boolean = true;

  public colors = ['#36A2EB', '#4BBFC0', '#FBAD4B', '#9368E9', '#FE656D'];

  private regions: Map<string, any> = new Map<string, any>([
    ["BHS5", { "latitude": "45.313978", "longitude": "-73.875834" }],
    ["DE1", { "latitude": "50.110924", "longitude": "8.682127" }],
    ["GRA5", { "latitude": "50.986938", "longitude": "2.125890" }],
    ["SBG5", { "latitude": "48.573406", "longitude": "7.752111" }],
    ["UK1", { "latitude": "51.507351", "longitude": "-0.127758" }],
    ["WAW1", { "latitude": "52.229675", "longitude": "21.012230" }],
    ["SGP1", { "latitude": "1.352083", "longitude": "103.819839" }],
    ["SYD1", { "latitude": "-33.868820", "longitude": "151.209290" }],
  ])

  constructor(private ovhService: OvhService) {
    this.ovhService.getCloudProjects().subscribe(data => {
      this.projects = data.length;
      this.loadingProjects = false;
    }, err => {
      this.projects = 0;
      this.loadingProjects = false;
    });

    this.ovhService.getUsers().subscribe(data => {
      this.loadingUsers = false;
      this.users = data;
    }, err => {
      this.users = 0;
      this.loadingUsers = false;
    });

    this.ovhService.getCloudAlerts().subscribe(data => {
      this.alerts = data;
      this.loadingAlerts = false;
    }, err => {
      this.alerts = 0;
      this.loadingAlerts = false;
    });

    this.ovhService.getCurrentBill().subscribe(data => {
      this.currentBill = data.total;
      data.services.forEach(service => {
        this.mostUsedServices.push({
          name: service.label,
          cost: service.total
        });
      })
    }, err => {
      this.currentBill = 0;
    });

    this.ovhService.getCloudInstances().subscribe(data => {
      let plots = {}
      let _usedRegions = new Map<string, number>();

      data.forEach(instance => {
        _usedRegions[instance.region] = (_usedRegions[instance.region] ? _usedRegions[instance.region] : 0) + 1;
      });

      let scope = this;

      for (var region in _usedRegions) {
        this.usedRegions++;
        plots[region] = {
          latitude: scope.regions.get(region).latitude,
          longitude: scope.regions.get(region).longitude,
          value: [_usedRegions[region], 1],
          tooltip: { content: `${region}<br />Instances: ${_usedRegions[region]}` }
        }
      }

      Array.from(this.regions.keys()).forEach(region => {
        let found = false;
        for (let _region in plots) {
          if (_region == region) {
            found = true;
          }
        }
        if (!found) {
          plots[region] = {
            latitude: this.regions.get(region).latitude,
            longitude: this.regions.get(region).longitude,
            value: [_usedRegions[region], 0],
            tooltip: { content: `${region}<br />Instances: 0` }
          }
        }

        this.loadingUsedRegions = false;
        this.showInstancesPerRegion(plots);
      });
    }, err => {
      this.usedRegions = 0;
      this.loadingUsedRegions = false;
      let plots = {}
      Array.from(this.regions.keys()).forEach(region => {
        plots[region] = {
          latitude: this.regions.get(region).latitude,
          longitude: this.regions.get(region).longitude,
          value: [0, 0],
          tooltip: { content: `${region}<br />Instances: 0` }
        }
      });
      this.showInstancesPerRegion(plots);
    });

    this.ovhService.getTickets().subscribe(data => {
      this.openTickets = data.open;
      this.resolvedTickets = data.close;
      this.loadingOpenTickets = false;
      this.loadingResolvedTickets = false;
    }, err => {
      this.openTickets = 0;
      this.resolvedTickets = 0;
      this.loadingOpenTickets = false;
      this.loadingResolvedTickets = false;
    });
  }

  ngOnInit() {
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

  ngAfterViewInit(): void {
    this.showInstancesPerRegion({});
  }

  public showInstancesPerRegion(plots) {
    var canvas: any = $(".mapregions");
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
