import { TestBed, inject } from '@angular/core/testing';

import { AwsService } from './aws.service';

describe('AwsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AwsService]
    });
  });

  it('should be created', inject([AwsService], (service: AwsService) => {
    expect(service).toBeTruthy();
  }));
});
