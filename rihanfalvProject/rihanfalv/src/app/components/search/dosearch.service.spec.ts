import { TestBed } from '@angular/core/testing';

import { DosearchService } from './dosearch.service';

describe('DosearchService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: DosearchService = TestBed.get(DosearchService);
    expect(service).toBeTruthy();
  });
});
