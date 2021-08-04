import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IncomeInsertComponent } from './income-insert.component';

describe('IncomeInsertComponent', () => {
  let component: IncomeInsertComponent;
  let fixture: ComponentFixture<IncomeInsertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IncomeInsertComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IncomeInsertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
