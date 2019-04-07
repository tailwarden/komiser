import { Component, OnInit } from '@angular/core';
import { AwsService } from '../aws.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
declare var Chart: any;
declare var Circles: any;
declare var $: any;
declare var moment: any;

@Component({
  selector: 'app-compute',
  templateUrl: './compute.component.html',
  styleUrls: ['./compute.component.css']
})
export class ComputeComponent implements OnInit {

  public runningEC2Instances: number;
  public stoppedEC2Instances: number;
  public terminatedEC2Instances: number;
  
  public lambdaFunctions: Object;

  constructor(private awsService: AwsService) {

    this.lambdaFunctions = {}

    this.awsService.getInstancesPerRegion().subscribe(data => {
      this.runningEC2Instances = data.state.running ? data.state.running : 0;
      this.stoppedEC2Instances = data.state.stopped ? data.state.stopped : 0;
      this.terminatedEC2Instances = data.state.terminated ? data.state.terminated : 0;

      let labels = [];
      let series = [];
      let colors = []
      Object.keys(data.family).forEach(key => {
        labels.push(key);
        series.push(data.family[key]);
        colors.push(this.getRandomColor());
      })

      this.showInstanceFamilies(labels, series, colors);
    }, err => {
      this.runningEC2Instances = 0;
      this.stoppedEC2Instances = 0;
      this.terminatedEC2Instances = 0;
    });

    this.awsService.getLambdaFunctions().subscribe(data => {
      this.lambdaFunctions.golang = data.golang ? data.golang : 0;
      this.lambdaFunctions.ruby = data.ruby ? data.ruby : 0;
      this.lambdaFunctions.java = data.java ? data.java : 0;
      this.lambdaFunctions.csharp = data.csharp ? data.csharp : 0;
      this.lambdaFunctions.python = data.python ? data.python : 0;
      this.lambdaFunctions.node = data.node ? data.node : 0;
      this.lambdaFunctions.custom = data.custom ? data.custom : 0;
    }, err => {
      this.lambdaFunctions = {
        golang: 0,
        ruby: 0,
        java: 0,
        csharp: 0,
        python: 0,
        node: 0,
        custom: 0
      };
    });

    this.awsService.getLambdaInvocationMetrics().subscribe(data => {
      let labels = [];
      data.forEach(period => {
        labels.push(new Date(period.timestamp).toLocaleString('en-us', { month: 'long' }))
      })

      let series = []
      for (let i = 0; i < labels.length; i++) {
        let serie = []
        for (let j = 0; j < labels.length; j++) {
          let item = data[j].metrics[i]
          if(item){
            serie.push({
              meta: item.label, value: item.value
            })
          } else {
            serie.push({
              meta: 'others', value: 0
            })
          }
        }
        series.push(serie)
      }
      this.showLambdaInvocations(labels, series);
    }, err => {
      console.log(err)
    })
  }

  ngOnInit() {
    this.showInstanceFamilies([], [], []);
    this.showLambdaInvocations([], []);
    this.showInstancesPrivacy();
  }

  private showInstancesPrivacy(){
    var data = {
      series: [139, 10]
    };
    
    var sum = function(a, b) { return a + b };
    
    new Chartist.Pie('#instances-privacy', data, {
      labelInterpolationFnc: function(value) {
        return Math.round(value / data.series.reduce(sum) * 100) + '%';
      }
    });
  }

  private showLambdaInvocations(labels, series){
    new Chartist.Bar('.lambdaInvocationsChart', {
      labels: labels,
      series: series
    }, {
      plugins: [
        Chartist.plugins.tooltip()
      ],
      stackBars: true,
      axisY: {
        labelInterpolationFnc: function(value) {
          return (value/1000000) + 'M'
        }
      }
    }).on('draw', function(data) {
      if(data.type === 'bar') {
        data.element.attr({
          style: 'stroke-width: 30px'
        });
      }
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

    var ctx = document.getElementById('task-complete2');
    var chart = new Chart.PolarArea(ctx, config);
  }

}
