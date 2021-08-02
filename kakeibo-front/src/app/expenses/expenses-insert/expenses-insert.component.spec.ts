import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ExpensesInsertComponent } from './expenses-insert.component';

describe('ExpensesInsertComponent', () => {
  let component: ExpensesInsertComponent;
  let fixture: ComponentFixture<ExpensesInsertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ExpensesInsertComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ExpensesInsertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
