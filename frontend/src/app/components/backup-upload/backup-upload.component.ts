import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { MessageService } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { FileUploadModule } from 'primeng/fileupload';
import { ToastModule } from 'primeng/toast';
import { BackupService } from '../../services/backup.service';
import { HttpErrorResponse } from '@angular/common/http';

interface UploadEvent {
  originalEvent: any;
  files: File[];
}

@Component({
  selector: 'app-backup-upload',
  standalone: true,
  imports: [
    DialogModule,
    LucideAngularModule,
    ButtonModule,
    FileUploadModule,
    ToastModule,
  ],
  templateUrl: './backup-upload.component.html',
  styleUrls: ['./backup-upload.component.css'],
  providers: [MessageService, BackupService],
})
export class BackupUploadComponent {
  visible: boolean = false;
  uploadedFile: File | null = null;  // Pour stocker le fichier sélectionné

  constructor(
    private messageService: MessageService,
    private backupService: BackupService
  ) {}

  showDialog() {
    this.visible = true;
  }

  onSelect(event: UploadEvent) {
    this.uploadedFile = event.files[0]; // Stocke le fichier sélectionné
  }

  onUpload() {
    if (this.uploadedFile) {
      this.backupService.uploadBackup(this.uploadedFile).subscribe(
        (response) => {
          this.messageService.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Backup uploaded successfully!',
          });
          this.visible = false;
        },
        (error: HttpErrorResponse) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Error',
            detail: 'Backup upload failed: ' + error.message,
          });
        }
      );
    } else {
      this.messageService.add({
        severity: 'warn',
        summary: 'No File',
        detail: 'Please select a file before uploading.',
      });
    }
  }
}
