import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GcpComponent } from './gcp.component';

describe('GcpComponent', () => {
  let component: GcpComponent;
  let fixture: ComponentFixture<GcpComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GcpComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GcpComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
