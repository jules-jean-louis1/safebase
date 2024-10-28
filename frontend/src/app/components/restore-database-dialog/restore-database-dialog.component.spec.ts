import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RestoreDatabaseDialogComponent } from './restore-database-dialog.component';

describe('RestoreDatabaseDialogComponent', () => {
  let component: RestoreDatabaseDialogComponent;
  let fixture: ComponentFixture<RestoreDatabaseDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RestoreDatabaseDialogComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(RestoreDatabaseDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
