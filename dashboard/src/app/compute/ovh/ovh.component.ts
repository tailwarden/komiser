import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';
declare var Chart: any;
declare var Circles: any;
import * as $ from "jquery";
declare var moment: any;

@Component({
  selector: 'ovh-compute',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhComputeComponent implements OnInit {
  public activeInstances: number = 0;
  public stoppedInstances: number = 0;
  public pausedInstances: number = 0;
  public windowsImages: number = 0;
  public linuxImages: number = 0;
  public kubernetesClusters: number = 0;
  public kubernetesNodes: number = 0;

  public loadingActiveInstances: boolean = true;
  public loadingPausedInstances: boolean = true;
  public loadingStoppedInstances: boolean = true;
  public loadingInstancesPrivacyChart: boolean = true;
  public loadingInstancesFamilyChart: boolean = true;
  public loadingWindowsImages: boolean = true;
  public loadingLinuxImages: boolean = true;
  public loadingKubernetesClusters: boolean = true;
  public loadingKubernetesNodes: boolean = true;

  constructor(private ovhService: OvhService) {
    this.ovhService.getCloudInstances().subscribe(data => {
      let machineTypes = new Map<string, number>();

      let privateInstances = 0;

      data.forEach(instance => {
        if(instance.status == 'ACTIVE'){
          this.activeInstances++;
        }
        if(instance.status == 'STOPPED'){
          this.stoppedInstances++;
        }
        if(instance.status == 'PAUSED'){
          this.pausedInstances++;
        }
        instance.ipAddresses.forEach(ip => {
          if(ip.type == 'private'){
            privateInstances++
          }
        });
        machineTypes[instance.planCode] = (machineTypes[instance.planCode] ? machineTypes[instance.planCode] : 0) + 1;
      })
      
      let labels = [];
      let series = [];
      let colors = []
      for (var machine in machineTypes) {
        labels.push(machine);
        series.push(machineTypes[machine]);
        colors.push(this.getRandomColor());
      }

      this.loadingInstancesPrivacyChart = false;
      this.showInstancesPrivacy([data.length - privateInstances, privateInstances]);

      this.loadingInstancesFamilyChart = false;
      this.showInstanceFamilies(labels, series, colors);

      this.loadingActiveInstances = false;
      this.loadingPausedInstances = false;
      this.loadingStoppedInstances = false;
    }, err => {
      this.activeInstances = 0;
      this.stoppedInstances = 0;
      this.pausedInstances = 0;
      this.loadingInstancesPrivacyChart = false;
      this.loadingInstancesFamilyChart = false;
      this.loadingActiveInstances = false;
      this.loadingPausedInstances = false;
      this.loadingStoppedInstances = false;
    });

    this.ovhService.getCloudImages().subscribe(data => {
      this.linuxImages = data.linux;
      this.windowsImages = data.windows;
      this.loadingLinuxImages = false;
      this.loadingWindowsImages = false;
    }, err => {
      this.linuxImages = 0;
      this.windowsImages = 0;
      this.loadingLinuxImages = false;
      this.loadingWindowsImages = false;
    });

    this.ovhService.getKubeClusters().subscribe(data => {
      this.kubernetesClusters = data;
      this.loadingKubernetesClusters = false;
    }, err => {
      this.kubernetesClusters = 0;
      this.loadingKubernetesClusters = false;
    });

    this.ovhService.getKubeNodes().subscribe(data => {
      this.kubernetesNodes = data;
      this.loadingKubernetesNodes = false;
    }, err => {
      this.kubernetesNodes = 0;
      this.loadingKubernetesNodes = false;
    });
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
          label: 'My dataset'
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

  private getRandomColor() {
    var letters = '789ABCD'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.round(Math.random() * 6)];
    }
    return color;
  }

  ngOnInit() {
  }

}
