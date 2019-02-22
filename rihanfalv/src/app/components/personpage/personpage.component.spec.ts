import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonpageComponent } from './personpage.component';

describe('PersonpageComponent', () => {
  let component: PersonpageComponent;
  let fixture: ComponentFixture<PersonpageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PersonpageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PersonpageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
