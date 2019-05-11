import { Component, OnInit,OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-data-and-ai',
  templateUrl: './data-and-ai.component.html',
  styleUrls: ['./data-and-ai.component.css']
})
export class DataAndAiComponent implements OnInit, OnDestroy {
  public provider: string;
  public _subscription: Subscription;

  constructor(private storeService: StoreService) {
    this.provider = this.storeService.getProvider();
    this._subscription = this.storeService.providerChanged.subscribe(provider => {
      this.provider = provider;
    })
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }
}
