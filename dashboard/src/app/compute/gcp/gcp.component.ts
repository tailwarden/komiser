import { Component, OnInit, AfterViewInit } from '@angular/core';
import { GcpService } from '../../gcp.service';
import { PageChangedEvent } from 'ngx-bootstrap/pagination';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
declare var Chart: any;
declare var Circles: any;
import * as $ from "jquery";
declare var moment: any;

@Component({
  selector: 'gcp-compute',
  templateUrl: './gcp.component.html',
  styleUrls: ['./gcp.component.css']
})
export class GcpComputeComponent implements OnInit, AfterViewInit {
  public runningVMInstances: number = 0;
  public stoppedVMInstances: number = 0;
  public cloudFunctions: any;
  public kubernetesClusters: number;
  public kubernetesNodes: number;
  public images: number;
  public instancesWithCPUUtilization: Array<any> = [];
  public returnedInstancesWithCPUUtilization: Array<any> = [];

  public loadingRunningInstances: boolean = true;
  public loadingStoppedInstances: boolean = true;
  public loadingInstancesPrivacyChart: boolean = true;
  public loadingInstancesFamilyChart: boolean = true;
  public loadingCloudFunctions: boolean = true;
  public loadingKubernetesNodes: boolean = true;
  public loadingKubernetesClusters: boolean = true;
  public loadingImages: boolean = true;
  public loadingGaeBandwidthChart: boolean = true;

  private zones: Map<string, any> = new Map<string, any>([
    ["asia-east1", { "latitude": "23.697809", "longitude": "120.960518" }],
    ["asia-east2", { "latitude": "22.396427", "longitude": "114.109497" }],
    ["asia-northeast1", { "latitude": "35.689487", "longitude": "139.691711" }],
    ["asia-south1", { "latitude": "19.075983", "longitude": "72.877655" }],
    ["asia-southeast1", { "latitude": "1.339637", "longitude": "103.707339" }],
    ["australia-southeast1", { "latitude": "43.498299", "longitude": "2.375200" }],
    ["europe-north1", { "latitude": "60.568890", "longitude": "27.188188" }],
    ["europe-west1", { "latitude": "50.447748", "longitude": "3.819524" }],
    ["europe-west2", { "latitude": "51.507322", "longitude": "-0.127647" }],
    ["europe-west3", { "latitude": "50.110644", "longitude": "8.682092" }],
    ["europe-west4", { "latitude": "53.448402", "longitude": "6.846503" }],
    ["europe-west6", { "latitude": "47.376888", "longitude": "8.541694" }],
    ["northamerica-northeast1", { "latitude": "45.509060", "longitude": "-73.553360" }],
    ["southamerica-east1", { "latitude": "23.550651", "longitude": "-46.633382" }],
    ["us-central1", { "latitude": "41.262128", "longitude": "-95.861391" }],
    ["us-east1", { "latitude": "33.196003", "longitude": "-80.013137" }],
    ["us-east4", { "latitude": "39.029265", "longitude": "-77.467387" }],
    ["us-west1", { "latitude": "45.601506", "longitude": "-121.184159" }],
    ["us-west2", { "latitude": "34.053691", "longitude": "-118.242767" }],
  ])

  constructor(private gcpService: GcpService) {
    this.cloudFunctions = {
      golang: 0,
      python: 0,
      node: 0
    };

    this.gcpService.getComputeInstances().subscribe(data => {
      this.loadingRunningInstances = false;
      this.loadingStoppedInstances = false;
      this.loadingInstancesPrivacyChart = false;
      this.loadingInstancesFamilyChart = false;

      let publicInstances = 0;
      let machineTypes = new Map<string, number>();
      data.forEach(instance => {
        if (instance.status == 'RUNNING') {
          this.runningVMInstances++;
        }
        if (instance.status == 'TERMINATED') {
          this.stoppedVMInstances++;
        }

        if (instance.public) {
          publicInstances++;
        }

        machineTypes[instance.machineType] = (machineTypes[instance.machineType] ? machineTypes[instance.machineType] : 0) + 1;
      })

      let labels = [];
      let series = [];
      let colors = []
      for (var machine in machineTypes) {
        labels.push(machine);
        series.push(machineTypes[machine]);
        colors.push(this.getRandomColor());
      }

      this.showInstanceFamilies(labels, series, colors);
      this.showInstancesPrivacy([publicInstances, data.length - publicInstances]);
    }, err => {
      this.loadingRunningInstances = false;
      this.loadingStoppedInstances = false;
      this.loadingInstancesPrivacyChart = false;
      this.loadingInstancesFamilyChart = false;
    })

    this.gcpService.getCloudFunctions().subscribe(data => {
      Object.keys(data).forEach(runtime => {
        if (runtime.startsWith("nodejs")) {
          this.cloudFunctions.node++;
        }
        if (runtime.startsWith("go")) {
          this.cloudFunctions.golang++;
        }
        if (runtime.startsWith("python")) {
          this.cloudFunctions.python++;
        }
      });
      this.loadingCloudFunctions = false;
    }, err => {
      this.loadingCloudFunctions = false;
      this.cloudFunctions = {
        golang: 0,
        python: 0,
        node: 0
      };
    });

    this.gcpService.getKubernetesClusters().subscribe(data => {
      this.loadingKubernetesClusters = false;
      this.loadingKubernetesNodes = false;

      this.kubernetesClusters = data.length;
      this.kubernetesNodes = 0;
      data.forEach(cluster => {
        this.kubernetesNodes += cluster.nodes;
      })

      let scope = this;
      let _usedRegions = new Map<string, number>();
      let plots = {};

      data.forEach(cluster => {
        let region = cluster.zone.substring(0, cluster.zone.lastIndexOf("-"));
        _usedRegions[region] = (_usedRegions[region] ? _usedRegions[region] : 0) + 1;
      })

      for (var region in _usedRegions) {
        plots[region] = {
          latitude: scope.zones.get(region).latitude,
          longitude: scope.zones.get(region).longitude,
          value: [_usedRegions[region], 1],
          tooltip: { content: `${region}<br />Clusters: ${_usedRegions[region]}` }
        }
      }

      Array.from(this.zones.keys()).forEach(region => {
        let found = false;
        for (let _region in plots) {
          if (_region == region) {
            found = true;
          }
        }
        if (!found) {
          plots[region] = {
            latitude: this.zones.get(region).latitude,
            longitude: this.zones.get(region).longitude,
            value: [_usedRegions[region], 0],
            tooltip: { content: `${region}<br />Clusters: 0` }
          }
        }
      });

      this.showKubernetesClusters(plots);
    }, err => {
      this.kubernetesClusters = 0;
      this.kubernetesNodes = 0;
      this.loadingKubernetesClusters = false;
      this.loadingKubernetesNodes = false;
    });

    this.gcpService.getComputeImages().subscribe(data => {
      this.images = data.length;
      this.loadingImages = false;
    }, err => {
      this.images = 0;
      this.loadingImages = false;
    });

    this.gcpService.getComputeCPUUtilization().subscribe(data => {
      data.forEach(metric => {
        let series = []
        metric.points.forEach(point => {
          series.push(+point.value.doubleValue.toFixed(1))
        })
        if (metric.points.length <= 1) {
          series.push(0)
        }
        this.instancesWithCPUUtilization.push({
          name: metric.metric.labels.instance_name,
          series: series.reverse()
        })
      })

      this.returnedInstancesWithCPUUtilization = this.instancesWithCPUUtilization.slice(0, 3);
      console.log(this.returnedInstancesWithCPUUtilization);
    }, err => {
      this.instancesWithCPUUtilization = [];
      this.returnedInstancesWithCPUUtilization = [];
    });

    this.gcpService.getAppEngineBandwidth().subscribe(data => {
      let availablePeriods = []
      data.forEach(resource => {
        resource.points.forEach(point => {
          let timestamp = new Date(point.interval.endTime).toISOString().split('T')[0]
          if (!availablePeriods.includes(timestamp)) {
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
            if (timestamp == period) {
              serie.push({
                meta: 'Sent Bytes',
                value: point.value.int64Value
              })
              found = true
            }
          })
          if (!found) {
            serie.push({
              meta: 'Sent Bytes',
              value: 0
            })
          }
        })

        total += (serie[serie.length - 1].value / 1024) / 1024;
        series.push(serie)
      });

      console.log(series);
      this.loadingGaeBandwidthChart = false;
      this.showAppEngineBandwidth(availablePeriods, series);
    }, err => {
      this.loadingGaeBandwidthChart = false;
    });
  }

  private showAppEngineBandwidth(labels, series) {
    let scope = this;
    new Chartist.Bar('#gaeBandwidthChart', {
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

  pageChanged(event: PageChangedEvent): void {
    const startItem = (event.page - 1) * event.itemsPerPage;
    const endItem = event.page * event.itemsPerPage;
    this.returnedInstancesWithCPUUtilization = this.instancesWithCPUUtilization.slice(startItem, endItem);
  }

  ngAfterViewInit(): void {
    this.showKubernetesClusters({});
  }

  public showKubernetesClusters(plots) {
    var canvas: any = $(".kubeclustersmap");
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

  private getRandomColor() {
    var letters = '789ABCD'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.round(Math.random() * 6)];
    }
    return color;
  }

  private showInstancesPrivacy(series) {
    var canvas: any = document.getElementById('instancesPrivacyChart');
    var ctx = canvas.getContext('2d');
    new Chart(ctx, {
      type: 'pie',
      data: {
        datasets: [{
          data: series,
          backgroundColor: ['#36A2EB', '#4BC0C0']
        }],
        labels: ['Public Instances', 'Private Instances']
      },
      options: {}
    });
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
    var chart = new Chart.PolarArea(ctx, config);
  }

  ngOnInit() {
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

}
