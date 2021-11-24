import { Component, OnInit, OnDestroy } from '@angular/core';
import { AzureService } from '../../azure.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../store.service';

@Component({
  selector: 'azure-storage',
  templateUrl: './azure.component.html',
  styleUrls: ['./azure.component.css']
})
export class AzureStorageComponent implements OnInit, OnDestroy {
  public snapshotsNumber: number;
  public snapshotsSize: string;
  public disksNumber: number;
  public disksSize: string;
  public mysqlInstances: number;

  public loadingSnapshotsNumber: boolean;
  public loadingSnapshotsSize: boolean;
  public loadingDisksNumber: boolean;
  public loadingDisksSize: boolean;
  public loadingMySQLInstances: boolean;

  private _subscription: Subscription;

  constructor(private azureService: AzureService, private storeService: StoreService) {
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
    this.disksNumber = 0;
    this.disksSize = '0 KB';
    this.mysqlInstances = 0;

    this.loadingSnapshotsNumber = true;
    this.loadingSnapshotsSize = true;
    this.loadingDisksNumber = true;
    this.loadingDisksSize = true;
    this.loadingMySQLInstances = true;

    this.azureService.getSnapshots().subscribe(data => {
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

    this.azureService.getDisks().subscribe(data => {
      this.disksNumber = data.length;
      let total = 0;
      data.forEach(disk => {
        total += disk.size;
      });
      this.disksSize = this.bytesToSizeWithUnit(total * 1024 * 1024 * 1024);
      this.loadingDisksNumber = false;
      this.loadingDisksSize = false;
    }, err => {
      this.disksNumber = 0;
      this.disksSize = '0 KB';
      this.loadingDisksNumber = false;
      this.loadingDisksSize = false;
    });

  }

  private bytesToSizeWithUnit(bytes) {
    var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    if (bytes == 0) return '0 Byte';
    var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)).toString());
    return Math.round(bytes / Math.pow(1024, i)) + ' ' + sizes[i];
  };
}
