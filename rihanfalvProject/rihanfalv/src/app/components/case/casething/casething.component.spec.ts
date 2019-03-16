import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CasethingComponent } from './casething.component';

describe('CasethingComponent', () => {
  let component: CasethingComponent;
  let fixture: ComponentFixture<CasethingComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CasethingComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CasethingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
