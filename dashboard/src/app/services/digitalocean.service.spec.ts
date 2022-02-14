import { TestBed, inject } from '@angular/core/testing';

import { DigitaloceanService } from './digitalocean.service';

describe('DigitaloceanService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [DigitaloceanService]
    });
  });

  it('should be created', inject([DigitaloceanService], (service: DigitaloceanService) => {
    expect(service).toBeTruthy();
  }));
});
