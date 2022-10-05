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
import { CloudService } from './services/cloud.service';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { DigitaloceanService } from './services/digitalocean.service';
import { GcpService } from './services/gcp.service';
import { GoogleAnalyticsService } from './services/google-analytics.service';
import { NotificationsComponent } from './pages/notifications/notifications.component';
import { OvhService } from './services/ovh.service';
import { StoreService } from './services/store.service';
//import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap';
import { SettingsService } from './services/settings.service';
import { InventoryComponent } from './pages/inventory/inventory.component';
import { SearchFilterPipe } from './filters/search.pipe';
import { PaginationModule } from 'ngx-bootstrap/pagination';

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
        NotificationsComponent,
        InventoryComponent,
        SearchFilterPipe,
    ],
    imports: [
        RouterModule.forRoot(appRoutes, { relativeLinkResolution: 'legacy' }),
        HttpClientModule,
        BrowserModule,
        BrowserAnimationsModule,
        TrendModule,
        FormsModule,
        //  NgbModalModule,
        PaginationModule.forRoot(),
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
        CloudService,
    ],
    bootstrap: [AppComponent],
})
export class AppModule {
    constructor(protected _googleAnalyticsService: GoogleAnalyticsService) {}
}
