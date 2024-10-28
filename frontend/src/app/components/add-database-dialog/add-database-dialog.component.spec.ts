import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddDatabaseDialogComponent } from './add-database-dialog.component';

describe('AddDatabaseDialogComponent', () => {
  let component: AddDatabaseDialogComponent;
  let fixture: ComponentFixture<AddDatabaseDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddDatabaseDialogComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AddDatabaseDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
