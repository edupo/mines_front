import { TestBed } from '@angular/core/testing';

import { MinesService } from './mines.service';

describe('MinesService', () => {
  let service: MinesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MinesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
