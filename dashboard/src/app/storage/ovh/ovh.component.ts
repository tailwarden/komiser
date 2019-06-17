import { Component, OnInit } from '@angular/core';
import { OvhService } from '../../ovh.service';
declare var Chart: any;
declare var $: any;
declare var window: any;
declare var Circles: any;
import * as Chartist from 'chartist';
import 'chartist-plugin-tooltips';

@Component({
  selector: 'ovh-storage',
  templateUrl: './ovh.component.html',
  styleUrls: ['./ovh.component.css']
})
export class OvhStorageComponent implements OnInit {
  public storageContainers: number = 0;
  public storedObjects: string = '0';
  public storageContainerTotalSize: string = '0 KB';
  public emptyStorageContainers: number = 0;
  public totalVolumes: number = 0;
  public volumesSize: string = '0 KB';
  public unattachedVolumes: number = 0;
  public snapshots: string = '0 KB';

  public loadingStorageContainers: boolean = true;
  public loadingStoredObjects: boolean = true;
  public loadingStorageContainerTotalSize: boolean = true;
  public loadingEmptyStorageContainers: boolean = true;
  public loadingTotalVolumes: boolean = true;
  public loadingVolumesSize: boolean = true;
  public loadingUnattachedVolumes: boolean = true;
  public loadingVolumesTypeChart: boolean = true;
  public loadingSnapshots: boolean = true;

  constructor(private ovhService: OvhService) {
    this.ovhService.getStorageContainers().subscribe(data => {
      let totalObjects = 0;
      let totalSize = 0;
      data.forEach(container => {
        if(container.storedObjects == 0) {
          this.emptyStorageContainers++;
        }
        totalObjects += container.storedObjects;
        totalSize += container.storedBytes;
      });
      this.storedObjects = this.formatNumber(totalObjects).toString();
      this.storageContainerTotalSize = this.bytesToSizeWithUnit(totalSize);
      this.storageContainers = data.length;
      this.loadingStorageContainers = false;
      this.loadingEmptyStorageContainers = false;
      this.loadingStorageContainerTotalSize = false;
      this.loadingStoredObjects = false;
    }, err => {
      this.storageContainers = 0;
      this.emptyStorageContainers = 0;
      this.storageContainers = 0;
      this.storageContainerTotalSize = '0 KB';
      this.loadingStorageContainers = false;
      this.loadingEmptyStorageContainers = false;
      this.loadingStorageContainerTotalSize = false;
    });

    this.ovhService.getCloudVolumes().subscribe(data => {
      this.totalVolumes = data.length;

      let labels = ['classic', 'high-speed']
      let series = [0, 0]
      let totalSize = 0;
      data.forEach(volume => {
        totalSize += volume.size;
        if(volume.attachedTo.length == 0){
          this.unattachedVolumes++;
        }
        if(volume.type == 'classic'){
          series[0]++;
        }else{
          series[1]++;
        }
      });

      this.showVolumesFamily(labels, series);
      this.loadingVolumesTypeChart = false;
      this.volumesSize = this.bytesToSizeWithUnit(totalSize * 1024 * 1024 * 1024);
      this.loadingTotalVolumes = false;
      this.loadingVolumesSize = false;
      this.loadingUnattachedVolumes = false;
    }, err => {
      this.loadingTotalVolumes = false;
      this.loadingVolumesSize = false;
      this.loadingUnattachedVolumes = false;
      this.totalVolumes = 0;
      this.volumesSize = '0 KB';
      this.unattachedVolumes = 0;
      this.loadingVolumesTypeChart = false;
    });

    this.ovhService.getCloudSnapshots().subscribe(data => {
      let totalSize = 0;
      data.forEach(snapshot => {
        totalSize += snapshot.size;
      });
      this.snapshots = this.bytesToSizeWithUnit(totalSize * 1024 * 1024 * 1024);
      this.loadingSnapshots = false;
    }, err => {
      this.snapshots = '0 KB';
      this.loadingSnapshots = false;
    });
  }

  private showVolumesFamily(labels, series) {
    var barChartData = {
      labels: labels,
      datasets: [{
        backgroundColor: [
          "#36A2EB",
          "#4BC0C0",
          "#FFCD56",
          "#FF6385"
        ],
        borderWidth: 1,
        data: series
      }]

    };

    let canvas: any = document.getElementById('volumesTypeChart');
    var ctx = canvas.getContext('2d');
    new Chart(ctx, {
      type: 'pie',
      data: barChartData,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        legend: {
          position: 'top',
        },
      }
    });
  }

  ngOnInit() {
  }

  private formatNumber(labelValue: number) {

    // Nine Zeroes for Billions
    return Math.abs(Number(labelValue)) >= 1.0e+9

      ? (Math.abs(Number(labelValue)) / 1.0e+9).toFixed(2) + " B"
      // Six Zeroes for Millions 
      : Math.abs(Number(labelValue)) >= 1.0e+6

        ? (Math.abs(Number(labelValue)) / 1.0e+6).toFixed(2) + " M"
        // Three Zeroes for Thousands
        : Math.abs(Number(labelValue)) >= 1.0e+3

          ? (Math.abs(Number(labelValue)) / 1.0e+3).toFixed(2) + " K"

          : Math.abs(Number(labelValue));

  }

  private dynamicColors() {
    var r = Math.floor(Math.random() * 255);
    var g = Math.floor(Math.random() * 255);
    var b = Math.floor(Math.random() * 255);
    return "rgba(" + r + "," + g + "," + b + ", 0.5)";
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };

  private bytesToSize(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i))
  };

}
