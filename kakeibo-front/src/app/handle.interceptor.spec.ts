import { TestBed } from '@angular/core/testing';

import { HandleInterceptor } from './handle.interceptor';

describe('HandleInterceptor', () => {
  beforeEach(() => TestBed.configureTestingModule({
    providers: [
      HandleInterceptor
      ]
  }));

  it('should be created', () => {
    const interceptor: HandleInterceptor = TestBed.inject(HandleInterceptor);
    expect(interceptor).toBeTruthy();
  });
});
