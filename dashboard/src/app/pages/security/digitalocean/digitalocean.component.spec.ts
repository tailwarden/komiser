import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { DigitaloceanComponent } from './digitalocean.component';

describe('DigitaloceanComponent', () => {
    let component: DigitaloceanComponent;
    let fixture: ComponentFixture<DigitaloceanComponent>;

    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            declarations: [DigitaloceanComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(DigitaloceanComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
