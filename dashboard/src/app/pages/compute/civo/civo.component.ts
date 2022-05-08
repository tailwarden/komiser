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
  selector: "civo-compute",
  templateUrl: "./civo.component.html",
  styleUrls: ["./civo.component.css"],
})
export class CivoComputeComponent implements OnInit, OnDestroy, AfterViewInit {
  public kubernetesClusters: number;
  public instances: number;

  public loadingKubernetesClusters: boolean;
  public loadingInstances: boolean;

  private _subscription: Subscription;

  constructor(
    private civoService: CivoService,
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

  ngOnInit() { }

  private initState() {
    this.kubernetesClusters = 0;
    this.loadingKubernetesClusters = true;

    this.civoService.getInstances().subscribe(data => {
      this.instances = data;
      this.loadingInstances = false;
    }, err => {
      this.instances = 0;
      this.loadingInstances = false;
    });

    this.civoService.getKubernetesClusters().subscribe(data => {
      this.kubernetesClusters = data;
      this.loadingKubernetesClusters = false;
    }, err => {
      this.kubernetesClusters = 0;
      this.loadingKubernetesClusters = false;
    })

  }

  ngAfterViewInit(): void {
  }
  
}
