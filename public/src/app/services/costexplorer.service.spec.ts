import { TestBed, inject } from '@angular/core/testing';

import { CostExplorerService } from './costexplorer.service';

describe('CostexplorerService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CostExplorerService]
    });
  });

  it('should be created', inject([CostExplorerService], (service: CostExplorerService) => {
    expect(service).toBeTruthy();
  }));
});
