import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { OvhComponent } from './ovh.component';

describe('OvhComponent', () => {
    let component: OvhComponent;
    let fixture: ComponentFixture<OvhComponent>;

    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            declarations: [OvhComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(OvhComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
