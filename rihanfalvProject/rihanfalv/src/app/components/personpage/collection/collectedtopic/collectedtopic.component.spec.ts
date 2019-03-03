import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CollectedtopicComponent } from './collectedtopic.component';

describe('CollectedtopicComponent', () => {
  let component: CollectedtopicComponent;
  let fixture: ComponentFixture<CollectedtopicComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CollectedtopicComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CollectedtopicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
