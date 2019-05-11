import { Component, OnInit,OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-limits',
  templateUrl: './limits.component.html',
  styleUrls: ['./limits.component.css']
})
export class LimitsComponent implements OnInit {
  public provider: string;
  public _subscription: Subscription;

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  constructor(private storeService: StoreService) {
    this.provider = this.storeService.getProvider();
    this._subscription = this.storeService.providerChanged.subscribe(provider => {
      this.provider = provider;
    });
  }

  ngOnInit() {
  }

}

