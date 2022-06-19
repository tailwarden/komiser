import { TrendModule } from 'ngx-trend';

import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { AwsService } from './services/aws.service';
import { AzureService } from './services/azure.service';
import { AwsComputeComponent } from './pages/compute/aws/aws.component';
import { AzureComputeComponent } from './pages/compute/azure/azure.component';
import { ComputeComponent } from './pages/compute/compute.component';
import { DigitaloceanComputeComponent } from './pages/compute/digitalocean/digitalocean.component';
import { GcpComputeComponent } from './pages/compute/gcp/gcp.component';
import { OvhComputeComponent } from './pages/compute/ovh/ovh.component';
import { AwsDashboardComponent } from './pages/dashboard/aws/aws.component';
import { AzureDashboardComponent } from './pages/dashboard/azure/azure.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { DigitaloceanDashboardComponent } from './pages/dashboard/digitalocean/digitalocean.component';
import { GcpDashboardComponent } from './pages/dashboard/gcp/gcp.component';
import { OvhDashboardComponent } from './pages/dashboard/ovh/ovh.component';
import { AwsDataAndAIComponent } from './pages/data-and-ai/aws/aws.component';
import { DataAndAiComponent } from './pages/data-and-ai/data-and-ai.component';
import { GcpDataAndAIComponent } from './pages/data-and-ai/gcp/gcp.component';
import { OvhDataAndAIComponent } from './pages/data-and-ai/ovh/ovh.component';
import { DigitaloceanService } from './services/digitalocean.service';
import { GcpService } from './services/gcp.service';
import { GoogleAnalyticsService } from './services/google-analytics.service';
import { AwsNetworkComponent } from './pages/network/aws/aws.component';
import { AzureNetworkComponent } from './pages/network/azure/azure.component';
import { DigitaloceanNetworkComponent } from './pages/network/digitalocean/digitalocean.component';
import { GcpNetworkComponent } from './pages/network/gcp/gcp.component';
import { NetworkComponent } from './pages/network/network.component';
import { OvhNetworkComponent } from './pages/network/ovh/ovh.component';
import { NotificationsComponent } from './pages/notifications/notifications.component';
import { OvhService } from './services/ovh.service';
import { AwsProfileComponent } from './pages/profile/aws/aws.component';
import { GcpProfileComponent } from './pages/profile/gcp/gcp.component';
import { OvhProfileComponent } from './pages/profile/ovh/ovh.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { AwsSecurityComponent } from './pages/security/aws/aws.component';
import { AzureSecurityComponent } from './pages/security/azure/azure.component';
import { DigitaloceanSecurityComponent } from './pages/security/digitalocean/digitalocean.component';
import { GcpSecurityComponent } from './pages/security/gcp/gcp.component';
import { OvhSecurityComponent } from './pages/security/ovh/ovh.component';
import { SecurityComponent } from './pages/security/security.component';
import { AwsStorageComponent } from './pages/storage/aws/aws.component';
import { AzureStorageComponent } from './pages/storage/azure/azure.component';
import { DigitaloceanStorageComponent } from './pages/storage/digitalocean/digitalocean.component';
import { GcpStorageComponent } from './pages/storage/gcp/gcp.component';
import { OvhStorageComponent } from './pages/storage/ovh/ovh.component';
import { StorageComponent } from './pages/storage/storage.component';
import { StoreService } from './services/store.service';
import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap';
import { SettingsService } from './services/settings.service';

const appRoutes: Routes = [
  {
    path: "compute",
    component: ComputeComponent,
    data: { title: "Compute - Komiser" },
  },
  {
    path: "storage",
    component: StorageComponent,
    data: { title: "Storage - Komiser" },
  },
  {
    path: "network",
    component: NetworkComponent,
    data: { title: "Network - Komiser" },
  },
  {
    path: "security",
    component: SecurityComponent,
    data: { title: "Security - Komiser" },
  },
  {
    path: "data-and-ai",
    component: DataAndAiComponent,
    data: { title: "Data & AI - Komiser" },
  },
  {
    path: "profile",
    component: ProfileComponent,
    data: { title: "Profile - Komiser" },
  },
  {
    path: "notifications",
    component: NotificationsComponent,
    data: { title: "Notifications - Komiser" },
  },
  {
    path: "",
    component: DashboardComponent,
    data: { title: "Dashboard - Komiser" },
  },
];

@NgModule({
  declarations: [
    AppComponent,
    AzureComputeComponent,
    ComputeComponent,
    DashboardComponent,
    StorageComponent,
    NetworkComponent,
    SecurityComponent,
    DataAndAiComponent,
    ProfileComponent,
    AwsDashboardComponent,
    GcpDashboardComponent,
    AwsComputeComponent,
    GcpComputeComponent,
    AwsStorageComponent,
    AzureStorageComponent,
    GcpStorageComponent,
    GcpNetworkComponent,
    AwsNetworkComponent,
    AzureNetworkComponent,
    AwsSecurityComponent,
    AzureSecurityComponent,
    GcpSecurityComponent,
    GcpDataAndAIComponent,
    AwsDataAndAIComponent,
    AwsProfileComponent,
    GcpProfileComponent,
    NotificationsComponent,
    OvhDashboardComponent,
    OvhComputeComponent,
    OvhStorageComponent,
    OvhNetworkComponent,
    OvhSecurityComponent,
    OvhDataAndAIComponent,
    OvhProfileComponent,
    DigitaloceanDashboardComponent,
    DigitaloceanComputeComponent,
    DigitaloceanStorageComponent,
    DigitaloceanNetworkComponent,
    DigitaloceanSecurityComponent,
    AzureDashboardComponent,
  ],
  imports: [
    RouterModule.forRoot(appRoutes, { relativeLinkResolution: 'legacy' }),
    HttpClientModule,
    BrowserModule,
    BrowserAnimationsModule,
    TrendModule,
    FormsModule,
    NgbModalModule,
  ],
  providers: [
    AwsService,
    DigitaloceanService,
    StoreService,
    GoogleAnalyticsService,
    GcpService,
    OvhService,
    AzureService,
    SettingsService
  ],
  bootstrap: [AppComponent],
})
export class AppModule {
  constructor(protected _googleAnalyticsService: GoogleAnalyticsService) { }
}
