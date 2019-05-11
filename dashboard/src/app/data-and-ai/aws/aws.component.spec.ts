import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AwsComponent } from './aws.component';

describe('AwsComponent', () => {
  let component: AwsComponent;
  let fixture: ComponentFixture<AwsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AwsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AwsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
