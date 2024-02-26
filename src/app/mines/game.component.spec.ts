import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MinesGameComponent } from './game.component';

describe('MinesGameComponent', () => {
  let component: MinesGameComponent;
  let fixture: ComponentFixture<MinesGameComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MinesGameComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(MinesGameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
