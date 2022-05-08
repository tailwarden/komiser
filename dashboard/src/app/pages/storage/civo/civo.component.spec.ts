import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { CivoComponent } from './civo.component';

describe('CivoComponent', () => {
  let component: CivoComponent;
  let fixture: ComponentFixture<CivoComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ CivoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CivoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
