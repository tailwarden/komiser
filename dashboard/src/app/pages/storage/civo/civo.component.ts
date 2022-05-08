import { Component, OnInit, OnDestroy } from '@angular/core';
import { CivoService } from '../../../services/civo.service';
import { Subscription } from 'rxjs';
import { StoreService } from '../../../services/store.service';

@Component({
  selector: 'civo-storage',
  templateUrl: './civo.component.html',
  styleUrls: ['./civo.component.css']
})
export class CivoStorageComponent implements OnInit, OnDestroy {
  public disksNumber: number;
  public volumes: number;

  public loadingDisksNumber: boolean;
  public loadingVolumes: boolean;

  private _subscription: Subscription;

  constructor(private civoService: CivoService, private storeService: StoreService) {
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
    this.disksNumber = 0;
    this.volumes = 0;

    this.loadingDisksNumber = true;
    this.loadingVolumes = true;

    this.civoService.getDisks().subscribe(data => {
      this.disksNumber = data;
      this.loadingDisksNumber = false;
    }, err => {
      this.disksNumber = 0;
      this.loadingDisksNumber = false;
    });

    this.civoService.getVolumes().subscribe(data => {
      this.volumes = data;
      this.loadingVolumes = false;
    }, err => {
      this.volumes = 0;
      this.loadingVolumes = false;
    });

  }
}
