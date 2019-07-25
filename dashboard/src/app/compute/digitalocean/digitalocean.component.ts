import { Component, OnInit, OnDestroy, AfterViewInit } from '@angular/core';
import { DigitaloceanService } from '../../digitalocean.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../store.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';
import * as $ from "jquery";
declare var Chart: any;

@Component({
  selector: 'digitalocean-compute',
  templateUrl: './digitalocean.component.html',
  styleUrls: ['./digitalocean.component.css']
})
export class DigitaloceanComputeComponent implements OnInit, OnDestroy, AfterViewInit {
  public kubernetesClusters: number;
  public kubernetesNodes: number;
  public activeDroplets: number;
  public offDroplets: number;
  public archivedDroplets: number;

  public images: any = {};

  public loadingKubernetesClusters: boolean;
  public loadingKubernetesNodes: boolean;
  public loadingArchivedDroplets: boolean;
  public loadingOffDroplets: boolean;
  public loadingActiveDroplets: boolean;
  public loadingImages: boolean = true;

  private regions: Map<string, any> = new Map<string, any>([
    ["nyc", { "latitude": "40.712776", "longitude": "-74.005974" }],
    ["ams", { "latitude": "52.370216", "longitude": "4.895168" }],
    ["sfo", { "latitude": "37.774929", "longitude": "-122.419418" }],
    ["sgp", { "latitude": "1.352083", "longitude": "103.819839" }],
    ["lon", { "latitude": "51.507351", "longitude": "-0.127758" }],
    ["fra", { "latitude": "50.110924", "longitude": "8.682127" }],
    ["tor", { "latitude": "43.653225", "longitude": "-79.383186" }],
    ["blr", { "latitude": "12.971599", "longitude": "77.594566" }],
  ]);

  private _subscription: Subscription;

  constructor(private digitaloceanService: DigitaloceanService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }

  private initState() {
    this.images = {
      ubuntu: 0,
      bsd: 0,
      debian: 0,
      fedora: 0,
      centos: 0
    }
    this.kubernetesClusters = 0;
    this.kubernetesNodes = 0;
    this.activeDroplets = 0;
    this.offDroplets = 0;
    this.archivedDroplets = 0;
    this.loadingActiveDroplets = true;
    this.loadingArchivedDroplets = true;
    this.loadingOffDroplets = true;
    this.loadingKubernetesClusters = true;
    this.loadingKubernetesNodes = true;
    this.loadingImages = true;

    this.digitaloceanService.getDroplets().subscribe(data => {
      data.forEach(droplet => {
        if (droplet.status == 'active') {
          this.activeDroplets++;
        } else if (droplet.status == 'off') {
          this.offDroplets++;
        } else {
          this.archivedDroplets++;
        }

        switch (droplet.image) {
          case 'Ubuntu':
            this.images.ubuntu++;
            break;
          case 'CentOS':
            this.images.centos++;
            break;
          case 'FreeBSD':
            this.images.bsd++;
            break;
          case 'Debian':
            this.images.debian++;
            break;
          case 'Fedora':
            this.images.fedora++;
            break;
        }
      });
      this.loadingActiveDroplets = false;
      this.loadingArchivedDroplets = false;
      this.loadingOffDroplets = false;
      this.loadingImages = false;
    }, err => {
      this.activeDroplets = 0;
      this.offDroplets = 0;
      this.archivedDroplets = 0;
      this.loadingActiveDroplets = false;
      this.loadingArchivedDroplets = false;
      this.loadingOffDroplets = false;
      this.images = {
        ubuntu: 0,
        bsd: 0,
        debian: 0,
        fedora: 0,
        centos: 0
      };
      this.loadingImages = false;
    });

    this.digitaloceanService.getKubernetesClusters().subscribe(data => {
      this.kubernetesClusters = data.length;
      data.forEach(cluster => {
        this.kubernetesNodes += cluster.nodes;
      });
      this.loadingKubernetesClusters = false;
      this.loadingKubernetesNodes = false;


      let _usedRegions = new Map<string, number>();
      let plots = {};
      let scope = this;

      data.forEach(cluster => {
        let region = cluster.region.substring(0, cluster.region.length - 1);
        _usedRegions[region] = (_usedRegions[region] ? _usedRegions[region] : 0) + 1;
      })


      for (var region in _usedRegions) {
        plots[region] = {
          latitude: scope.regions.get(region).latitude,
          longitude: scope.regions.get(region).longitude,
          value: [_usedRegions[region], 1],
          tooltip: { content: `${region}<br />Clusters: ${_usedRegions[region]}` }
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
            tooltip: { content: `${region}<br />Clusters: 0` }
          }
        }
      });

      this.showKubernetesClustersPerRegion(plots);

    }, err => {
      this.kubernetesClusters = 0;
      this.kubernetesNodes = 0;
      this.loadingKubernetesClusters = false;
      this.loadingKubernetesNodes = false;
    });
  }

  ngAfterViewInit(): void {
    this.showKubernetesClustersPerRegion({});
  }

  private showKubernetesClustersPerRegion(plots) {
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

}
