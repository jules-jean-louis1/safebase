import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateManualBackupComponent } from './create-manual-backup.component';

describe('CreateManualBackupComponent', () => {
  let component: CreateManualBackupComponent;
  let fixture: ComponentFixture<CreateManualBackupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateManualBackupComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateManualBackupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
