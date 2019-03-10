import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ToQuestionComponent } from './to-question.component';

describe('ToQuestionComponent', () => {
  let component: ToQuestionComponent;
  let fixture: ComponentFixture<ToQuestionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ToQuestionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ToQuestionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
