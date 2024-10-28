import { Component, OnInit } from '@angular/core';
import { TableModule } from 'primeng/table';
import { CommonModule } from '@angular/common';
import { DatabaseService } from '../../services/database.service';
import { ButtonModule } from 'primeng/button';
import { MessageService } from 'primeng/api';
import { MessageModule } from 'primeng/message';
import { ToastModule } from 'primeng/toast';
import { BackupService } from '../../services/backup.service';
import { ScheduleBackupComponent } from '../schedule-backup/schedule-backup.component';
import { RestoreDatabaseDialogComponent } from '../restore-database-dialog/restore-database-dialog.component';
import { LucideAngularModule } from 'lucide-angular';
import { TooltipModule } from 'primeng/tooltip';
import { NotificationService } from '../../services/notification.service';
import { BackendService } from '../../services/backend.service';

interface Database {
  id: string;
  name: string;
  type: string;
  host: string;
  port: string;
  username: string;
  password: string;
  database_name: string;
  is_cron_active: boolean;
  cron_schedule: string;
  created_at: string;
  updated_at: string;
}

@Component({
  selector: 'app-database-list',
  standalone: true,
  imports: [
    TableModule,
    CommonModule,
    ButtonModule,
    MessageModule,
    ToastModule,
    TooltipModule,
    LucideAngularModule,
    ScheduleBackupComponent,
    RestoreDatabaseDialogComponent,
  ],
  templateUrl: './database-list.component.html',
  styleUrls: ['./database-list.component.css'],
  providers: [DatabaseService, MessageService, BackupService, NotificationService, BackendService],
})
export class DatabaseListComponent implements OnInit {
  databases: Database[] = [];
  errorMessage: string = '';

  constructor(
    private databaseService: DatabaseService,
    private messageService: MessageService,
    private backupService: BackupService,
    private backendService: BackendService,
    private notificationService: NotificationService
  ) {}

  ngOnInit() {
    this.loadDatabases();

    this.notificationService.getRefreshListObservable().subscribe(() => {
      this.loadDatabases();
    });
  }

  loadDatabases() {
    this.backendService.isBackendReachable().subscribe((isReachable) => {
      if (isReachable) {
        this.databaseService.getDatabases().subscribe({
          next: (databases) => {
            this.databases = databases;
          },
          error: (error) => {
            console.error('Error loading databases', error);
            this.errorMessage = 'Failed to load databases';
          },
        });
      } else {
        this.errorMessage = 'Backend server is not reachable';
      }
    });
  }

  createBackup(databaseId: string) {
    console.log('Creating backup for database with ID:', databaseId);
    if (!databaseId) {
      this.messageService.add({
        severity: 'error',
        summary: 'Erreur',
        detail: 'Veuillez sélectionner une base de données.',
      });
      return;
    }
    this.backupService.createBackup(databaseId).subscribe({
      next: (data) => {
        console.log('Backup created', data);
        this.messageService.add({
          severity: 'success',
          summary: 'Succès',
          detail: 'Backup créé avec succès.',
        });
      },
      error: (error) => {
        console.error(error);
        this.messageService.add({
          severity: 'error',
          summary: 'Erreur',
          detail: 'Erreur lors de la création du backup.',
        });
      },
    });
  }

  deleteDatabase(databaseId: string) {
    if (!databaseId) {
      this.messageService.add({
        severity: 'error',
        summary: 'Erreur',
        detail: 'Veuillez sélectionner une base de données.',
      });
      return;
    }
    this.databaseService.deleteDatabase(databaseId).subscribe({
      next: (data) => {
        console.log('Database deleted', data);
        this.messageService.add({
          severity: 'success',
          summary: 'Succès',
          detail: 'Base de données supprimée avec succès.',
        });
        this.loadDatabases();
      },
      error: (error) => {
        console.error(error);
        this.messageService.add({
          severity: 'error',
          summary: 'Erreur',
          detail: 'Erreur lors de la suppression de la base de données.',
        });
      },
    });
    console.log('Deleting database with ID:', databaseId);
  }
}
