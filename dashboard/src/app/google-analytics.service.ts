import { Injectable } from '@angular/core';
import {Router, NavigationEnd} from '@angular/router';
declare var ga:Function; // <-- Here we declare GA variable

@Injectable()
export class GoogleAnalyticsService {

  constructor(router: Router) {
    router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        ga('set', 'page', event.url);
        ga('send', 'pageview');
      }
    })
    
  }

}
