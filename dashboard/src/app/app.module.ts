import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { ComputeComponent } from './compute/compute.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { StorageComponent } from './storage/storage.component';
import { NetworkComponent } from './network/network.component';
import { SecurityComponent } from './security/security.component';
import { DataAndAiComponent } from './data-and-ai/data-and-ai.component';
import { MonitoringComponent } from './monitoring/monitoring.component';


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
    path: 'monitoring',
    component: MonitoringComponent,
    data: { title: 'Monitoring - Komiser' }
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
    MonitoringComponent
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    ),
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
