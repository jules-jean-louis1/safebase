import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleBackupComponent } from './schedule-backup.component';

describe('ScheduleBackupComponent', () => {
  let component: ScheduleBackupComponent;
  let fixture: ComponentFixture<ScheduleBackupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ScheduleBackupComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ScheduleBackupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
