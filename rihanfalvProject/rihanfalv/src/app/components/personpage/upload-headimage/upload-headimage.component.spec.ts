import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UploadHeadimageComponent } from './upload-headimage.component';

describe('UploadHeadimageComponent', () => {
  let component: UploadHeadimageComponent;
  let fixture: ComponentFixture<UploadHeadimageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UploadHeadimageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UploadHeadimageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
