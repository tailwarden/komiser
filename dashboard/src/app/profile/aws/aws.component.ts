import { Component, OnInit, OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import { Subject, Subscription } from 'rxjs';

@Component({
  selector: 'aws-profile',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsProfileComponent implements OnInit, OnDestroy {
  public account : Object = {};
  public organization: Object = {};

  private _subscription: Subscription;

  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();
    
    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.account = {};
      this.organization = {};

      this.initState();
    });
  }

  private initState(){
    this.awsService.getAccountName().subscribe(data => {
      this.account = data;
    }, err => {
      this.account = {};
    });

    this.awsService.getOrganization().subscribe(data => {
      this.organization = data;
    }, err => {
      this.organization = {};
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  ngOnInit() {
  }

}
