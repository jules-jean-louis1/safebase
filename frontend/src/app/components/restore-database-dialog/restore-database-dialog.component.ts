import { Component, Input, OnInit } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TooltipModule } from 'primeng/tooltip';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { DropdownModule } from 'primeng/dropdown';
import { BackupService } from '../../services/backup.service';
import { RestoreService } from '../../services/restore.service';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { RippleModule } from 'primeng/ripple';
import { ButtonModule } from 'primeng/button';

interface Database {
  id: string;
  name: string;
  type: string;
  host: string;
  port: string;
  username: string;
  password: string;
  database_name: string;
  is_cron_active?: boolean;
  cron_schedule?: string;
  created_at: string;
  updated_at: string;
}

interface Backup {
  id: string;
  database_id?: string;
  status: string;
  backup_type: string;
  filename: string;
  size?: string;
  error_msg?: string;
  log?: string;
  created_at: string;
  updated_at: string;
}

@Component({
  selector: 'app-restore-database-dialog',
  standalone: true,
  imports: [
    DialogModule,
    InputTextModule,
    LucideAngularModule,
    TooltipModule,
    DropdownModule,
    RippleModule,
    ToastModule,
    ReactiveFormsModule,
    ButtonModule,
  ],
  templateUrl: './restore-database-dialog.component.html',
  styleUrl: './restore-database-dialog.component.css',
  providers: [BackupService, RestoreService, MessageService],
})
export class RestoreDatabaseDialogComponent implements OnInit {
  @Input() database!: Database;
  backups: Backup[] = [];
  visible: boolean = false;

  restoreForm: FormGroup = new FormGroup({
    selectedBackup: new FormControl<Backup | null>(null),
  });

  constructor(
    private backupService: BackupService,
    private restoreService: RestoreService,
    private messageService: MessageService
  ) {}
  ngOnInit(): void {}

  showDialog() {
    this.visible = true;
    this.loadBackups();
  }

  loadBackups() {
    const params = {
      dbType: this.database.type,
    };
    this.backupService.getBackupsBy(params).subscribe((data: Backup[]) => {
      this.backups = data;
    });
  }

  onSubmit() {
    console.log(this.restoreForm.value);
    if (!this.restoreForm.valid) {
      this.messageService.add({
        severity: 'error',
        summary: 'Error',
        detail: 'Please select a backup',
      });
      return;
    }
    const formValues = this.restoreForm.value;
    const data = {
      database_id: this.database.id,
      backup_id: formValues.selectedBackup.ID,
    };
    this.restoreService.insertRestore(data).subscribe({
      next: (data) => {
        console.log(data);
        this.messageService.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Restore operation has been initiated',
        });
        this.visible = false;
      },
      error: (error) => {
        const errorMessage =
          error?.Message || 'Failed to initiate restore operation';
        this.messageService.add({
          severity: 'error',
          summary: 'Error',
          detail: errorMessage,
        });
        console.error(error);
      },
    });
  }
}
