import { Component, OnInit, AfterViewInit } from '@angular/core';
import { AwsService } from '../../aws.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';

import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';
import * as $ from "jquery";
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'aws-security',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsSecurityComponent implements OnInit, AfterViewInit {

  public kmsKeys: number;
  public securityGroups: number;
  public keyPairs: number;
  public routeTables: number;
  public acmCertificates: number;
  public acmExpiredCertificates: number;
  public unrestrictedSecurityGroups: Array<any> = [];
  public returnedUnrestrictedSecurityGroups: Array<any> = [];
  public consoleLoginSourceIps: Array<any> = [];

  public loadingKMSKeys: boolean = true;
  public loadingSecurityGroups: boolean = true;
  public loadingKeyPairs: boolean = true;
  public loadingRouteTables: boolean = true;
  public loadingACMCertificates: boolean = true;
  public loadingACMExpiredCertificates: boolean = true;
  public loadingSignInEventsChart: boolean = true;

  constructor(private awsService: AwsService) {
    this.awsService.getKMSKeys().subscribe(data => {
      this.kmsKeys = data;
      this.loadingKMSKeys = false;
    }, err => {
      this.kmsKeys = 0;
      this.loadingKMSKeys = false;
    });

    this.awsService.getSecurityGroups().subscribe(data => {
      this.securityGroups = data;
      this.loadingSecurityGroups = false;
    }, err => {
      this.securityGroups = 0;
      this.loadingSecurityGroups = false;
    });

    this.awsService.getKeyPairs().subscribe(data => {
      this.keyPairs = data;
      this.loadingKeyPairs = false;
    }, err => {
      this.keyPairs = 0;
      this.loadingKeyPairs = false;
    });

    this.awsService.getRouteTables().subscribe(data => {
      this.routeTables = data;
      this.loadingRouteTables = false;
    }, err => {
      this.routeTables = 0;
      this.loadingRouteTables = false;
    });

    this.awsService.getACMListCertificates().subscribe(data => {
      this.acmCertificates = data;
      this.loadingACMCertificates = false;
    }, err => {
      this.acmCertificates = 0;
      this.loadingACMCertificates = false;
    });

    this.awsService.getACMExpiredCertificates().subscribe(data => {
      this.acmExpiredCertificates = data;
      this.loadingACMExpiredCertificates = false;
    }, err => {
      this.acmExpiredCertificates = 0;
      this.loadingACMExpiredCertificates = false;
    });

    this.awsService.getUnrestrictedSecurityGroups().subscribe(data => {
      this.unrestrictedSecurityGroups = data;
      this.returnedUnrestrictedSecurityGroups = this.unrestrictedSecurityGroups.slice(0, 20);
    }, err => {
      this.unrestrictedSecurityGroups = [];
      this.returnedUnrestrictedSecurityGroups = [];
    });

    this.awsService.getConsoleLoginEvents().subscribe(data => {
      let labels = [];
      let series = [];

      Object.keys(data).forEach(period => {
        labels.push(period)

      })

      for (let i = 0; i < labels.length; i++) {
        let serie = []
        for (let j = 0; j < labels.length; j++) {
          let username = Object.keys(data[labels[j]])[i]
          if (username) {
            serie.push({
              meta: username, value: data[labels[j]][username]
            })
          } else {
            serie.push({
              meta: 'others', value: 0
            })
          }
        }
        series.push(serie)
      }

      this.loadingSignInEventsChart = false;
      this.showSignInEvents(labels, series);
    }, err => {
      this.loadingSignInEventsChart = false;
      console.log(err);
    });

    this.awsService.getConsoleLoginSourceIps().subscribe(data => {
      let plots = {}
      Object.keys(data).forEach(ip => {
        this.consoleLoginSourceIps.push({
          ip: ip,
          total: data[ip].total
        });

        plots[ip] = {
          latitude: data[ip].coordinate.lat,
          longitude: data[ip].coordinate.lon,
          value: [data[ip].total, 1],
          tooltip: { content: `${ip}<br />Total: ${data[ip].total}` }
        }
      })
      this.showSourceIpLogin(plots);
    }, err => {
      this.consoleLoginSourceIps = [];
    });
  }

  pageChanged(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedUnrestrictedSecurityGroups = this.unrestrictedSecurityGroups.slice(startItem, endItem);
  }


  ngAfterViewInit(): void {
    this.showSourceIpLogin({});
  }

  ngOnInit() { }

  private showSourceIpLogin(plots) {
    var canvas : any = $("#sourceIpsChart");
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
            fill: "#59d05d"
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
                fill: "#4B5F91"
              },
              legendSpecificAttrs: {
                r: 25
              }
            }, {
              label: "> 1",
              min: "1",
              max: "50000",
              attrs: {
                fill: "#59D05D"
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

  private showSignInEvents(labels, series) {
    new Chartist.Bar('#signInEventsChart', {
      labels: labels,
      series: series
    }, {
        plugins: [
          Chartist.plugins.tooltip()
        ],
        stackBars: true,
        axisY: {
          labelInterpolationFnc: function (value) {
            return value
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

}
