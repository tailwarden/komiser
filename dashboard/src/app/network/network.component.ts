import { Component, OnInit, OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  styleUrls: ['./network.component.css']
})
export class NetworkComponent implements OnInit, OnDestroy {
  public provider: string;
  public _subscription: Subscription;

  constructor(private storeService: StoreService) {
    this.provider = this.storeService.getProvider();
    this._subscription = this.storeService.providerChanged.subscribe(provider => {
      console.log(provider);
      this.provider = provider;
    })
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }
}
