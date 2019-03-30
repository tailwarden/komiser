import { Component, OnInit } from '@angular/core';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
declare var $: any;

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor() { }

  ngOnInit() {
    this.showLastSixMonth();
    this.showEC2InstancesPerRegion();
  }


  public showEC2InstancesPerRegion(){
    $(".mapcontainer").mapael({
      map : {
        name : "world_countries",
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
      plots: {
        'paris': {
          latitude: 48.86,
          longitude: 2.3444,
          value: 500000000,
          tooltip: {content: "Paris<br />Population: 500000000"}
      },
      'newyork': {
          latitude: 40.667,
          longitude: -73.833,
          value: 200001,
          tooltip: {content: "New york<br />Population: 200001"}
      },
      'sydney': {
          latitude: -33.917,
          longitude: 151.167,
          value: 600000,
          tooltip: {content: "Sydney<br />Population: 600000"}
      },
      'brasilia': {
          latitude: -15.781682,
          longitude: -47.924195,
          value: 200000001,
          tooltip: {content: "Brasilia<br />Population: 200000001"}
      },
      'tokyo': {
          latitude: 35.687418,
          longitude: 139.692306,
          value: 200001,
          tooltip: {content: "Tokyo<br />Population: 200001"}
      }
      },
        });
  }

  public showLastSixMonth(){
    var dataSales = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
      series: [
        [
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          },
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          }
        ],
        [
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          },
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          }
        ],
        [
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          },
          {
            meta: 'EC2', value: 1
          },
          {
            meta: 'Lambda', value: 5
          },
          {
            meta: 'RDS', value: 6
          }
        ]
      ]
    }
    
    var optionChartSales = {
      plugins: [
      Chartist.plugins.tooltip()
      ],
      seriesBarDistance: 10,
      axisX: {
        showGrid: false
      },
      height: "245px",
    }
    
    var responsiveChartSales = [
    ['screen and (max-width: 640px)', {
      seriesBarDistance: 5,
      axisX: {
        labelInterpolationFnc: function (value) {
          console.log(value)
          return value[0];
        }
      }
    }]
    ];
    
    new Chartist.Bar('#salesChart', dataSales, optionChartSales);
  }

}
