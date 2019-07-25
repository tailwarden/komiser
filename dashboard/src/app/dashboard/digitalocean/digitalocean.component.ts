import { Component, OnInit, AfterViewInit, OnDestroy} from '@angular/core';
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
  selector: 'digitalocean-dashboard',
  templateUrl: './digitalocean.component.html',
  styleUrls: ['./digitalocean.component.css']
})
export class DigitaloceanDashboardComponent implements OnInit, AfterViewInit, OnDestroy {

  public projects: number;
  public usedRegions: number;


  public loadingProjects: boolean = true;
  public loadingUsedRegions: boolean = true;

  private regions: Map<string,any> = new Map<string,any>([
    ["nyc", {"latitude":"40.712776", "longitude":"-74.005974"}],
    ["ams", {"latitude":"52.370216", "longitude":"4.895168"}],
    ["sfo", {"latitude":"37.774929", "longitude":"-122.419418"}],
    ["sgp", {"latitude":"1.352083", "longitude":"103.819839"}],
    ["lon", {"latitude":"51.507351", "longitude":"-0.127758"}],
    ["fra", {"latitude":"50.110924", "longitude":"8.682127"}],
    ["tor", {"latitude":"43.653225", "longitude":"-79.383186"}],
    ["blr", {"latitude":"12.971599", "longitude":"77.594566"}],
  ]);

  private _subscription: Subscription;

  constructor(private digitaloceanService: DigitaloceanService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnDestroy(){
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }

  private initState(){
    this.projects = 0;
    this.usedRegions = 0;

    this.loadingProjects = true;
    this.loadingUsedRegions = true;

    this.digitaloceanService.getProjects().subscribe(data => {
      this.projects = data;
      this.loadingProjects = false;
    }, err => {
      this.loadingProjects = false;
      this.projects = 0;
    });

    this.digitaloceanService.getDroplets().subscribe(data => {
      let _usedRegions = new Map<string, number>();
      let plots = {};
      let scope = this;
      
      data.forEach(droplet => {
        let region = droplet.region.substring(0, droplet.region.length-1);
        _usedRegions[region] = (_usedRegions[region] ? _usedRegions[region] : 0) + 1;
      })

      
      for(var region in _usedRegions){
        this.usedRegions++;
        plots[region] = {
          latitude: scope.regions.get(region).latitude,
          longitude: scope.regions.get(region).longitude,
          value: [_usedRegions[region], 1],
          tooltip: { content: `${region}<br />Droplets: ${_usedRegions[region]}` }
        }
      }

      
      Array.from(this.regions.keys()).forEach(region => {
        let found = false;
        for(let _region in plots){
          if(_region == region){
            found = true;
          }
        }
        if(!found){
          plots[region] = {
            latitude: this.regions.get(region).latitude,
            longitude: this.regions.get(region).longitude,
            value: [_usedRegions[region], 0],
            tooltip: { content: `${region}<br />Droplets: 0` }
          }
        }
      });
      
      this.loadingUsedRegions = false;
      this.showDropletsPerRegion(plots);
    }, err => {
      this.loadingUsedRegions = false;
      this.usedRegions = 0;
    });
  }

  ngAfterViewInit(): void {
    this.showDropletsPerRegion({});
  }

  private showDropletsPerRegion(plots){
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
