import { Component, OnInit, OnDestroy } from '@angular/core';
import { DigitaloceanService } from '../../digitalocean.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../store.service';

@Component({
  selector: 'digitalocean-storage',
  templateUrl: './digitalocean.component.html',
  styleUrls: ['./digitalocean.component.css']
})
export class DigitaloceanStorageComponent implements OnInit, OnDestroy {
  public snapshotsNumber: number;
  public snapshotsSize: string;
  public volumesNumber: number;
  public volumesSize: string;
  public mysqlInstances: number;

  public loadingSnapshotsNumber: boolean;
  public loadingSnapshotsSize: boolean;
  public loadingVolumesNumber: boolean;
  public loadingVolumesSize: boolean;
  public loadingMySQLInstances: boolean;

  private _subscription: Subscription;

  constructor(private digitaloceanService: DigitaloceanService, private storeService: StoreService) {
    this.initState();
    this._subscription = this.storeService.profileChanged.subscribe(account => {
      this.initState();
    });
  }

  ngOnInit() {
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  private initState() {
    this.snapshotsNumber = 0;
    this.snapshotsSize = '0 KB';
    this.volumesNumber = 0;
    this.volumesSize = '0 KB';
    this.mysqlInstances = 0;

    this.loadingSnapshotsNumber = true;
    this.loadingSnapshotsSize = true;
    this.loadingVolumesNumber = true;
    this.loadingVolumesSize = true;
    this.loadingMySQLInstances = true;

    this.digitaloceanService.getSnapshots().subscribe(data => {
      this.snapshotsNumber = data.length;
      let total = 0;
      data.forEach(volume => {
        total += volume.size;
      });
      this.snapshotsSize = this.bytesToSizeWithUnit(total * 1024 * 1024 * 1024);
      this.loadingSnapshotsNumber = false;
      this.loadingSnapshotsSize = false;
    }, err => {
      this.loadingSnapshotsNumber = false;
      this.loadingSnapshotsSize = false;
      this.snapshotsNumber = 0;
      this.snapshotsSize = '0 KB';
    });

    this.digitaloceanService.getVolumes().subscribe(data => {
      this.volumesNumber = data.length;
      let total = 0;
      data.forEach(volume => {
        total += volume.size;
      });
      this.volumesSize = this.bytesToSizeWithUnit(total * 1024 * 1024 * 1024);
      this.loadingVolumesNumber = false;
      this.loadingVolumesSize = false;
    }, err => {
      this.volumesNumber = 0;
      this.volumesSize = '0 KB';
      this.loadingVolumesNumber = false;
      this.loadingVolumesSize = false;
    });

    this.digitaloceanService.getDatabases().subscribe(data => {
      this.mysqlInstances = data;
      this.loadingMySQLInstances = false;
    }, err => {
      this.mysqlInstances = 0;
      this.loadingMySQLInstances = false;
    });
  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };
}
