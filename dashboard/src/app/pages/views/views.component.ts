import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { StoreService } from '../../services/store.service';

@Component({
    selector: 'app-views',
    templateUrl: './views.component.html',
    styleUrls: ['./views.component.css'],
})
export class ViewsComponent implements OnInit {
    public view: any = {};

    constructor(
        private activatedRoute: ActivatedRoute,
        private storeService: StoreService
    ) {
        this.activatedRoute.paramMap.subscribe((params: Params) => {
            console.log(params.get('id'));
            this.view = this.storeService.getView(params.get('id'));
        });
    }

    ngOnInit(): void {}
}
