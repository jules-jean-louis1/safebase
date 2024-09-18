import { Component, OnInit } from '@angular/core';
import { BackupService } from '../../services/backup.service';
import { TableModule } from 'primeng/table';
import { CommonModule } from '@angular/common';
import { TooltipModule } from 'primeng/tooltip';
import { LucideAngularModule } from 'lucide-angular';

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
  database: {
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
  };
}

@Component({
  selector: 'app-backup-list',
  standalone: true,
  imports: [TableModule, CommonModule, TooltipModule, LucideAngularModule],
  templateUrl: './backup-list.component.html',
  styleUrl: './backup-list.component.css',
  providers: [BackupService],
})
export class BackupListComponent implements OnInit {
  backups: Backup[] = [];
  constructor(private backupService: BackupService) {}

  loadBackups() {
    this.backupService.getBackupFull().subscribe((data: Backup[]) => {
      this.backups = data;
    });
  }

  ngOnInit(): void {
    this.loadBackups();
  }

  deleteBackup(backupId: string) {
    if (!backupId) return;
    this.backupService.deleteBackup(backupId).subscribe({
      next: (data) => {
        console.log(data);
        this.loadBackups();
      },
      error: (error) => {
        console.error(error);
      },
    });
  }
}
