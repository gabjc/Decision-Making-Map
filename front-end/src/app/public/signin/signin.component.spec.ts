import { ComponentFixture, TestBed } from '@angular/core/testing';

import { signinComponent } from './signin.component';

describe('signinComponent', () => {
  let component: signinComponent;
  let fixture: ComponentFixture<signinComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ signinComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(signinComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
