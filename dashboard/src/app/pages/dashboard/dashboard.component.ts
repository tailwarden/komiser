import { Component, OnInit, AfterViewInit, OnDestroy } from '@angular/core';
import { StoreService } from '../../services/store.service';
import { AwsService } from '../../services/aws.service';
import { SettingsService } from '../../services/settings.service';
import { Subject, Subscription } from 'rxjs';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';
import * as $ from 'jquery';
declare var Chart: any;
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { CloudService } from '../../services/cloud.service';

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
    styleUrls: ['./dashboard.component.css'],
})
export class DashboardComponent implements OnInit, AfterViewInit, OnDestroy {
    public accounts: Array<any> = new Array<any>();
    public regions: Array<any> = new Array<any>();
    public plots: Array<any> = new Array<any>();
    public cost: number = 0;
    public groupBy: string = 'account';
    public monthlyCostBreakdownChart: any;

    constructor(private cloudService: CloudService) {
        this.getMonthlyCostBreakdownByCloudAccount();
        this.cloudService.getCloudAccounts().subscribe((data) => {
            this.accounts = data;
        });

        this.getCloudRegions();
    }

    public getCloudAccounts() {
        let count = 0;
        Object.keys(this.accounts).forEach((provider) => {
            count += this.accounts[provider]?.length;
        });
        return count;
    }

    public onGroupByChanged(groupBy) {
        console.log(groupBy);
        switch (groupBy) {
            case 'account':
                this.getMonthlyCostBreakdownByCloudAccount();
                break;
            case 'region':
                this.getMonthlyCostBreakdownByCloudRegion();
                break;
            case 'provider':
                this.getMonthlyCostBreakdownByCloudProvider();
                break;
        }
    }

    public getCloudRegions() {
        this.cloudService.getCloudRegions().subscribe((data) => {
            this.regions = data;
            this.plots = data;
            this.plots.forEach((plot) => {
                plot.value = [2, 1];
                plot.tooltip = {
                    content: `Region: <b>${plot.label}</b><br/>
                    City: <b>${plot.name}</b><br/>`,
                };
            });
            this.getCloudServiceMap();
        });
    }

    private getMonthlyCostBreakdownByCloudProvider() {
        this.cloudService.getCloudCostByCloudProvider().subscribe((data) => {
            if (data) {
                let labels = [];
                let values = [];
                let colors = [];
                Object.keys(data).forEach((key) => {
                    labels.push(key);
                    values.push(data[key]);
                    colors.push(this.stringToColour(key));
                });
                this.showCostBreakdown(labels, values, colors);
            }
        });
    }

    private getMonthlyCostBreakdownByCloudRegion() {
        this.cloudService.getCloudCostByCloudRegion().subscribe((data) => {
            if (data) {
                let labels = [];
                let values = [];
                let colors = [];
                Object.keys(data).forEach((key) => {
                    labels.push(key);
                    values.push(data[key]);
                    colors.push(this.stringToColour(key));
                });
                this.showCostBreakdown(labels, values, colors);
            }
        });
    }

    private getMonthlyCostBreakdownByCloudAccount() {
        this.cloudService.getCloudCostByCloudAccount().subscribe((data) => {
            if (data) {
                let labels = [];
                let values = [];
                let colors = [];
                this.cost = 0;
                Object.keys(data).forEach((key) => {
                    labels.push(key);
                    values.push(data[key]);
                    colors.push(this.stringToColour(key));
                    if (data[key] > 0) {
                        this.cost += data[key];
                    }
                });
                this.showCostBreakdown(labels, values, colors);
            }
        });
    }

    private showCostBreakdown(labels, values, colors) {
        this.monthlyCostBreakdownChart?.destroy();
        const canvas: any = document.getElementById(
            'monthlyCostBreakdownChart'
        );
        const ctx = canvas.getContext('2d');
        this.monthlyCostBreakdownChart = new Chart(ctx, {
            type: 'doughnut',
            data: {
                labels: labels,
                datasets: [
                    {
                        data: values,
                        backgroundColor: colors,
                        borderColor: '#FFFFFF',
                        hoverOffset: 15,
                    },
                ],
            },
            options: {
                aspectRatio: 2,
                layout: {
                    padding: 5,
                },
                plugins: {
                    legend: {
                        position: 'right',
                        labels: {
                            font: {
                                family: 'Noto Sans',
                                color: '#091126',
                            },
                            usePointStyle: true,
                            padding: 16,
                            generateLabels: (chart) => {
                                const datasets = chart.data.datasets;
                                return datasets[0].data.map((data, i) => ({
                                    text: `${chart.data.labels[i]} $${data}`,
                                    fillStyle: datasets[0].backgroundColor[i],
                                    strokeStyle: '#fff',
                                    hidden: !chart.getDataVisibility(i),
                                    index: i,
                                }));
                            },
                        },
                    },
                    tooltip: {
                        backgroundColor: 'rgba(0,0,0,.75)',
                        multiKeyBackground: '#282828',
                        boxPadding: 8,
                        padding: 12,
                        usePointStyle: true,
                        bodyFont: {
                            family: 'Noto Sans',
                        },
                        callbacks: {
                            label(label) {
                                return `${label.label}: $${label.formattedValue}`;
                            },
                        },
                    },
                },
            },
        });
    }

    private getCloudServiceMap() {
        const canvas: any = $('.mapregions');
        canvas.mapael({
            map: {
                name: 'world_countries',
                zoom: {
                    enabled: true,
                    maxLevel: 10,
                },
                defaultPlot: {
                    attrs: {
                        fill: '#387beb',
                        opacity: 1,
                    },
                },
                defaultArea: {
                    attrs: {
                        fill: '#F4F2F8',
                        stroke: '#D9D9D9',
                    },
                    attrsHover: {
                        fill: '#e2e3e9',
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
                                    fill: '#999999',
                                },
                                legendSpecificAttrs: {
                                    r: 25,
                                },
                            },
                            {
                                label: '> 1',
                                min: '1',
                                attrs: {
                                    fill: '#387beb',
                                },
                                legendSpecificAttrs: {
                                    r: 25,
                                },
                            },
                        ],
                    },
                ],
            },
            plots: this.plots,
        });
    }

    public stringToColour(str) {
        var hash = 0;
        for (var i = 0; i < str.length; i++) {
            hash = str.charCodeAt(i) + ((hash << 5) - hash);
        }
        var colour = '#';
        for (var i = 0; i < 3; i++) {
            var value = (hash >> (i * 8)) & 0xff;
            colour += ('00' + value.toString(16)).substr(-2);
        }
        return colour;
    }

    ngOnDestroy() {}

    ngOnInit() {}

    ngAfterViewInit(): void {}
}
