import { BrowserModule } from '@angular/platform-browser';
import { NgModule, enableProdMode } from '@angular/core';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';

import { AWSService } from '@app/services';
import { ChartsModule } from 'ng2-charts';

enableProdMode()

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpModule,
    ChartsModule
  ],
  providers: [
    AWSService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
