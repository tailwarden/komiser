import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { ComputeComponent } from './compute/compute.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { StorageComponent } from './storage/storage.component';
import { NetworkComponent } from './network/network.component';
import { SecurityComponent } from './security/security.component';
import { DataAndAiComponent } from './data-and-ai/data-and-ai.component';

import { AwsService } from './aws.service';
import { StoreService } from './store.service';
import { PaginationModule } from 'ngx-bootstrap/pagination';
import { ProfileComponent } from './profile/profile.component';
import { LimitsComponent } from './limits/limits.component';
import { GoogleAnalyticsService } from './google-analytics.service';



const appRoutes: Routes = [
  { 
    path: 'compute',
    component: ComputeComponent,
    data: { title: 'Compute - Komiser' }
  },
  { 
    path: 'storage',
    component: StorageComponent,
    data: { title: 'Storage - Komiser' }
  },
  { 
    path: 'network',
    component: NetworkComponent,
    data: { title: 'Network - Komiser' }
  },
  { 
    path: 'security',
    component: SecurityComponent,
    data: { title: 'Security - Komiser' }
  },
  { 
    path: 'data-and-ai',
    component: DataAndAiComponent,
    data: { title: 'Data & AI - Komiser' }
  },
  { 
    path: 'profile',
    component: ProfileComponent,
    data: { title: 'Profile - Komiser' }
  },
  { 
    path: 'limits',
    component: LimitsComponent,
    data: { title: 'Service Limits Checks - Komiser' }
  },
  { path: '',
    component: DashboardComponent,
    data: { title: 'Dashboard - Komiser' }
  }
];

@NgModule({
  declarations: [
    AppComponent,
    ComputeComponent,
    DashboardComponent,
    StorageComponent,
    NetworkComponent,
    SecurityComponent,
    DataAndAiComponent,
    ProfileComponent,
    LimitsComponent
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes
    ),
    HttpModule,
    BrowserModule,
    PaginationModule.forRoot()
  ],
  providers: [
    AwsService,
    StoreService,
    GoogleAnalyticsService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(protected _googleAnalyticsService: GoogleAnalyticsService){}
}
