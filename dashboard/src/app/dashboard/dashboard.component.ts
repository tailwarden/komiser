import { Component, OnInit, AfterViewInit, OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit, AfterViewInit, OnDestroy {
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

  ngOnInit() { }

  ngAfterViewInit(): void {

  }
}
