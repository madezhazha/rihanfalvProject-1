import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonalfeedbackComponent } from './personalfeedback.component';

describe('PersonalfeedbackComponent', () => {
  let component: PersonalfeedbackComponent;
  let fixture: ComponentFixture<PersonalfeedbackComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PersonalfeedbackComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PersonalfeedbackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
