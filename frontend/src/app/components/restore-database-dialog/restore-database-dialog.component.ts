import { Component, Input, OnInit } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { TooltipModule } from 'primeng/tooltip';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { DropdownModule } from 'primeng/dropdown';
import { BackupService } from '../../services/backup.service';

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
  ],
  templateUrl: './restore-database-dialog.component.html',
  styleUrl: './restore-database-dialog.component.css',
  providers: [BackupService],
})
export class RestoreDatabaseDialogComponent implements OnInit {
  @Input() database!: Database;
  backups: Backup[] = [];
  visible: boolean = false;

  restoreForm: FormGroup = new FormGroup({});

  constructor(private backupService: BackupService) {}
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
}
