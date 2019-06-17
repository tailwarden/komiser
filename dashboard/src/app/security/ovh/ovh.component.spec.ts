import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OvhComponent } from './ovh.component';

describe('OvhComponent', () => {
  let component: OvhComponent;
  let fixture: ComponentFixture<OvhComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OvhComponent ]
    })
    .compileComponents();
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
