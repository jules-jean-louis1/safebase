import { Component } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { MessageService } from 'primeng/api';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { FileUploadModule } from 'primeng/fileupload';
import { ToastModule } from 'primeng/toast';
import { BackupService } from '../../services/backup.service';

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
  styleUrl: './backup-upload.component.css',
  providers: [MessageService, BackupService],
})
export class BackupUploadComponent {
  visible: boolean = false;

  constructor(
    private messageService: MessageService,
    private backupService: BackupService
  ) {}

  showDialog() {
    this.visible = true;
  }

  onUpload(event: UploadEvent) {
    this.messageService.add({
      severity: 'info',
      summary: 'Success',
      detail: 'File Uploaded with Basic Mode',
    });
  }
}
