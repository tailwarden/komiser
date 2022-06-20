import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { AzureComponent } from './azure.component';

describe('DigitaloceanComponent', () => {
    let component: AzureComponent;
    let fixture: ComponentFixture<AzureComponent>;

    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            declarations: [AzureComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(AzureComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
