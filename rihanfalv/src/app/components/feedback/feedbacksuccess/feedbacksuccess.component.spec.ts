import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FeedbacksuccessComponent } from './feedbacksuccess.component';

describe('FeedbacksuccessComponent', () => {
  let component: FeedbacksuccessComponent;
  let fixture: ComponentFixture<FeedbacksuccessComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FeedbacksuccessComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FeedbacksuccessComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
