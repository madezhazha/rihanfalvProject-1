import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CollectedcaseComponent } from './collectedcase.component';

describe('CollectedcaseComponent', () => {
  let component: CollectedcaseComponent;
  let fixture: ComponentFixture<CollectedcaseComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CollectedcaseComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CollectedcaseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
