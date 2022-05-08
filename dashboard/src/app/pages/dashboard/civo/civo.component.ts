import 'chartist-plugin-tooltips';
import 'jquery-mapael';
import 'jquery-mapael/js/maps/world_countries.js';

import * as Chartist from 'chartist';
import * as $ from 'jquery';
import { Subscription } from 'rxjs';

import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { CivoService } from '../../../services/civo.service';
import { StoreService } from '../../../services/store.service';

declare var Chart: any;

@Component({
  selector: "civo-dashboard",
  templateUrl: "./civo.component.html",
  styleUrls: ["./civo.component.css"],
})
export class CivoDashboardComponent
  implements OnInit, AfterViewInit, OnDestroy {
  public usedRegions: number;

  public loadingUsedRegions: boolean = true;

  private _subscription: Subscription;

  constructor(
    private civoService: CivoService,
    private storeService: StoreService,
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

  ngOnInit() {
  }

  private initState() {
    this.usedRegions = 0;

    this.loadingUsedRegions = true;

    this.civoService.getRegions().subscribe(data => {
      this.usedRegions = data;
      this.loadingUsedRegions = false;
    }, err => {
      this.usedRegions = 0;
      this.loadingUsedRegions = false;
    });

  }

  ngAfterViewInit(): void {
  }
}
