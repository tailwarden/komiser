import { TestBed, inject } from '@angular/core/testing';

import { GcpService } from './gcp.service';

describe('GcpService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [GcpService]
    });
  });

  it('should be created', inject([GcpService], (service: GcpService) => {
    expect(service).toBeTruthy();
  }));
});
