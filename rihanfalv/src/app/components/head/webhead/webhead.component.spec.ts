import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WebheadComponent } from './webhead.component';

describe('WebheadComponent', () => {
  let component: WebheadComponent;
  let fixture: ComponentFixture<WebheadComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WebheadComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WebheadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
