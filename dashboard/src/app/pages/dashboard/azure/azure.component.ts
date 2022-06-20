import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';

import * as Chartist from 'chartist';
import * as $ from 'jquery';
import { Subscription } from 'rxjs';

import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { AzureService } from '../../../services/azure.service';
import { StoreService } from '../../../services/store.service';

declare var Chart: any;

@Component({
    selector: 'azure-dashboard',
    templateUrl: './azure.component.html',
    styleUrls: ['./azure.component.css'],
})
export class AzureDashboardComponent
    implements OnInit, AfterViewInit, OnDestroy
{
    public projects: number;
    public usedRegions: number;
    public totalBill: number = 0;

    public loadingProjects: boolean = true;
    public loadingUsedRegions: boolean = true;
    public loadingTotalBill: boolean = true;

    private regions: Map<string, any> = new Map<string, any>([
        ['eastus', { latitude: '37.3719', longitude: '-79.8164' }],
        ['eastus2', { latitude: '36.6681', longitude: '-78.3889' }],
        ['southcentralus', { latitude: '29.4167', longitude: '-98.5' }],
        ['westus2', { latitude: '47.233', longitude: '-119.852' }],
        ['westus3', { latitude: '33.448376', longitude: '-112.074036' }],
        ['australiaeast', { latitude: '-33.86', longitude: '151.2094' }],
        ['southeastasia', { latitude: '1.283', longitude: '103.833' }],
        ['australiaeas', { latitude: '-31.84', longitude: '145.61' }],
    ]);

    private _subscription: Subscription;

    constructor(
        private azureService: AzureService,
        private storeService: StoreService
    ) {
        this.initState();

        this._subscription = this.storeService.profileChanged.subscribe(
            (account) => {
                this.initState();
            }
        );
    }

    ngOnDestroy() {
        this._subscription.unsubscribe();
    }

    ngOnInit() {}

    private initState() {
        this.projects = 0;
        this.usedRegions = 0;
        this.totalBill = 0;

        this.loadingProjects = true;
        this.loadingUsedRegions = true;
        this.loadingTotalBill = true;

        this.azureService.getTotalCost().subscribe(
            (data) => {
                this.totalBill = data.amount;
                this.loadingTotalBill = false;
            },
            (err) => {
                this.totalBill = 0;
                this.loadingTotalBill = false;
            }
        );

        this.azureService.getVMs().subscribe(
            (data) => {
                let _usedRegions = new Map<string, number>();
                let plots = {};
                let scope = this;

                data.forEach((vm) => {
                    let region = vm.region.substring(0, vm.region.length - 1);
                    _usedRegions[region] =
                        (_usedRegions[region] ? _usedRegions[region] : 0) + 1;
                });

                for (var region in _usedRegions) {
                    this.usedRegions++;
                    plots[region] = {
                        latitude: scope.regions.get(region).latitude,
                        longitude: scope.regions.get(region).longitude,
                        value: [_usedRegions[region], 1],
                        tooltip: {
                            content: `${region}<br />VMs: ${_usedRegions[region]}`,
                        },
                    };
                }

                Array.from(this.regions.keys()).forEach((region) => {
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
                            tooltip: { content: `${region}<br />VMs: 0` },
                        };
                    }
                });

                this.loadingUsedRegions = false;
                this.showVMsPerRegion(plots);
            },
            (err) => {
                this.loadingUsedRegions = false;
                this.usedRegions = 0;
            }
        );
    }

    ngAfterViewInit(): void {
        this.showVMsPerRegion({});
    }

    private showVMsPerRegion(plots) {
        var canvas: any = $('.mapregions');
        canvas.mapael({
            map: {
                name: 'world_countries',
                zoom: {
                    enabled: true,
                    maxLevel: 10,
                },
                defaultPlot: {
                    attrs: {
                        fill: '#004a9b',
                        opacity: 0.6,
                    },
                },
                defaultArea: {
                    attrs: {
                        fill: '#e4e4e4',
                        stroke: '#fafafa',
                    },
                    attrsHover: {
                        fill: '#FBAD4B',
                    },
                    text: {
                        attrs: {
                            fill: '#505444',
                        },
                        attrsHover: {
                            fill: '#000',
                        },
                    },
                },
            },
            legend: {
                plot: [
                    {
                        labelAttrs: {
                            fill: '#f4f4e8',
                        },
                        titleAttrs: {
                            fill: '#f4f4e8',
                        },
                        cssClass: 'density',
                        mode: 'horizontal',
                        title: 'Density',
                        marginBottomTitle: 5,
                        slices: [
                            {
                                label: '< 1',
                                max: '0',
                                attrs: {
                                    fill: '#36A2EB',
                                },
                                legendSpecificAttrs: {
                                    r: 25,
                                },
                            },
                            {
                                label: '> 1',
                                min: '1',
                                max: '50000',
                                attrs: {
                                    fill: '#87CB14',
                                },
                                legendSpecificAttrs: {
                                    r: 25,
                                },
                            },
                        ],
                    },
                ],
            },
            plots: plots,
        });
    }
}
