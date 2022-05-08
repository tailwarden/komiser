import { TestBed, inject } from '@angular/core/testing';

import { CivoService } from './civo.service';

describe('CivoService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CivoService]
    });
  });

  it('should be created', inject([CivoService], (service: CivoService) => {
    expect(service).toBeTruthy();
  }));
});
