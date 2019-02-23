import { TestBed } from '@angular/core/testing';

import { LandService } from './land.service';

describe('LandService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: LandService = TestBed.get(LandService);
    expect(service).toBeTruthy();
  });
});
