import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DatabaseListComponent } from './database-list.component';

describe('DatabaseListComponent', () => {
  let component: DatabaseListComponent;
  let fixture: ComponentFixture<DatabaseListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DatabaseListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DatabaseListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
