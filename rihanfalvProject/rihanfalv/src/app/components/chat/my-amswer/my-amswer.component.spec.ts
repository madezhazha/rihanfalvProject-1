import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyAmswerComponent } from './my-amswer.component';

describe('MyAmswerComponent', () => {
  let component: MyAmswerComponent;
  let fixture: ComponentFixture<MyAmswerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyAmswerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyAmswerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
