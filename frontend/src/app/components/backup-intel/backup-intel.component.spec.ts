import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BackupIntelComponent } from './backup-intel.component';

describe('BackupIntelComponent', () => {
  let component: BackupIntelComponent;
  let fixture: ComponentFixture<BackupIntelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BackupIntelComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BackupIntelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
