import { TestBed } from '@angular/core/testing';

import { GolangService } from './golang.service';

describe('GolangService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: GolangService = TestBed.get(GolangService);
    expect(service).toBeTruthy();
  });
});
