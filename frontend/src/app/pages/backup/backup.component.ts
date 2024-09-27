import { Component } from '@angular/core';
import { BackupListComponent } from '../../components/backup-list/backup-list.component';
import { BackupUploadComponent } from '../../components/backup-upload/backup-upload.component';

@Component({
  selector: 'app-backup',
  standalone: true,
  imports: [BackupListComponent, BackupUploadComponent],
  templateUrl: './backup.component.html',
  styleUrl: './backup.component.css',
})
export class BackupComponent {}
