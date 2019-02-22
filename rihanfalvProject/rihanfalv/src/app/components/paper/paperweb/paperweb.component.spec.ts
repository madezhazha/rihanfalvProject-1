import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PaperwebComponent } from './paperweb.component';

describe('PaperwebComponent', () => {
  let component: PaperwebComponent;
  let fixture: ComponentFixture<PaperwebComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PaperwebComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PaperwebComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
