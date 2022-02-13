import { TestBed, inject } from '@angular/core/testing';

import { OvhService } from './ovh.service';

describe('OvhService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [OvhService]
    });
  });

  it('should be created', inject([OvhService], (service: OvhService) => {
    expect(service).toBeTruthy();
  }));
});
