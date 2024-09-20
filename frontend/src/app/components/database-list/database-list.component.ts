import { Component, OnInit } from '@angular/core';
import { TableModule } from 'primeng/table';
import { CommonModule } from '@angular/common';
import { DatabaseService } from '../../services/database.service';
import { NotificationService } from '../../services/notification.service';
import { ButtonModule } from 'primeng/button';
import { MessageService } from 'primeng/api';
import { MessageModule } from 'primeng/message';
import { ToastModule } from 'primeng/toast';
import { BackupService } from '../../services/backup.service';
import { ScheduleBackupComponent } from '../schedule-backup/schedule-backup.component';
import { RestoreDatabaseDialogComponent } from '../restore-database-dialog/restore-database-dialog.component';
import { LucideAngularModule } from 'lucide-angular';
import { TooltipModule } from 'primeng/tooltip';

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
  providers: [
    DatabaseService,
    MessageService,
    NotificationService,
    BackupService,
  ],
})
export class DatabaseListComponent implements OnInit {
  databases: Database[] = [];
  errorMessage: string = '';

  constructor(
    private databaseService: DatabaseService,
    private messageService: MessageService,
    private notificationService: NotificationService,
    private backupService: BackupService
  ) {}

  ngOnInit() {
    this.loadDatabases();

    // S'abonner à la notification pour rafraîchir la liste
    this.notificationService.refreshList$.subscribe(() => {
      console.log('Refresh list notification received in parent'); // Ajout d'un log pour déboguer
      this.loadDatabases();
    });
  }

  loadDatabases() {
    console.log('Loading databases...');
    // Appel de votre service pour charger les bases de données
    this.databaseService.getDatabases().subscribe((databases) => {
      console.log('Databases loaded', databases);
      this.databases = databases;
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
