import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BackupUploadComponent } from './backup-upload.component';

describe('BackupUploadComponent', () => {
  let component: BackupUploadComponent;
  let fixture: ComponentFixture<BackupUploadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BackupUploadComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BackupUploadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
