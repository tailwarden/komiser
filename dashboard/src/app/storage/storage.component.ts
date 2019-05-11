import { Component, OnInit, OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-storage',
  templateUrl: './storage.component.html',
  styleUrls: ['./storage.component.css']
})
export class StorageComponent implements OnInit, OnDestroy {

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

  ngOnInit() {}

}
