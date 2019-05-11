import { Component, OnInit, AfterViewInit, OnDestroy } from '@angular/core';
import { StoreService } from '../store.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-security',
  templateUrl: './security.component.html',
  styleUrls: ['./security.component.css']
})
export class SecurityComponent implements OnInit, AfterViewInit, OnDestroy {
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

  ngAfterViewInit(): void {
  }

  ngOnInit() { }


}
