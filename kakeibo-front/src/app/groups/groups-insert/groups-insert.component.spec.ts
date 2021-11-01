import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupsInsertComponent } from './groups-insert.component';

describe('GroupsInsertComponent', () => {
  let component: GroupsInsertComponent;
  let fixture: ComponentFixture<GroupsInsertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GroupsInsertComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupsInsertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
