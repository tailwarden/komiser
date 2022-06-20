import { TestBed, inject } from '@angular/core/testing';

import { AzureService } from './azure.service';

describe('AzureService', () => {
    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [AzureService],
        });
    });

    it('should be created', inject([AzureService], (service: AzureService) => {
        expect(service).toBeTruthy();
    }));
});
