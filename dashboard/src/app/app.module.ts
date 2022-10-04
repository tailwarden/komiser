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
import { AwsDashboardComponent } from './pages/dashboard/aws/aws.component';
import { AzureDashboardComponent } from './pages/dashboard/azure/azure.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { DigitaloceanDashboardComponent } from './pages/dashboard/digitalocean/digitalocean.component';
import { GcpDashboardComponent } from './pages/dashboard/gcp/gcp.component';
import { OvhDashboardComponent } from './pages/dashboard/ovh/ovh.component';
import { DigitaloceanService } from './services/digitalocean.service';
import { GcpService } from './services/gcp.service';
import { GoogleAnalyticsService } from './services/google-analytics.service';
import { NotificationsComponent } from './pages/notifications/notifications.component';
import { OvhService } from './services/ovh.service';
import { StoreService } from './services/store.service';
import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap';
import { SettingsService } from './services/settings.service';
import { InventoryComponent } from './pages/inventory/inventory.component';

const appRoutes: Routes = [
    {
        path: 'inventory',
        component: InventoryComponent,
        data: { title: 'Inventory - Komiser' },
    },
    {
        path: 'notifications',
        component: NotificationsComponent,
        data: { title: 'Notifications - Komiser' },
    },
    {
        path: '',
        component: DashboardComponent,
        data: { title: 'Dashboard - Komiser' },
    },
];

@NgModule({
    declarations: [
        AppComponent,
        DashboardComponent,
        AwsDashboardComponent,
        GcpDashboardComponent,
        NotificationsComponent,
        OvhDashboardComponent,
        DigitaloceanDashboardComponent,
        AzureDashboardComponent,
        InventoryComponent,
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
        SettingsService,
    ],
    bootstrap: [AppComponent],
})
export class AppModule {
    constructor(protected _googleAnalyticsService: GoogleAnalyticsService) {}
}
