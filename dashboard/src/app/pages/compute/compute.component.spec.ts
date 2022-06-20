import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ComputeComponent } from './compute.component';

describe('ComputeComponent', () => {
    let component: ComputeComponent;
    let fixture: ComponentFixture<ComputeComponent>;

    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            declarations: [ComputeComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(ComputeComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
