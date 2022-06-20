import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { DataAndAiComponent } from './data-and-ai.component';

describe('DataAndAiComponent', () => {
    let component: DataAndAiComponent;
    let fixture: ComponentFixture<DataAndAiComponent>;

    beforeEach(waitForAsync(() => {
        TestBed.configureTestingModule({
            declarations: [DataAndAiComponent],
        }).compileComponents();
    }));

    beforeEach(() => {
        fixture = TestBed.createComponent(DataAndAiComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
