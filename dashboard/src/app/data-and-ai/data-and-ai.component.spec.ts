import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DataAndAiComponent } from './data-and-ai.component';

describe('DataAndAiComponent', () => {
  let component: DataAndAiComponent;
  let fixture: ComponentFixture<DataAndAiComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DataAndAiComponent ]
    })
    .compileComponents();
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
