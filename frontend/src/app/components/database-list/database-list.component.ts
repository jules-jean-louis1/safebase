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

  ngOnInit(): void {
    this.loadDatabases();

    this.notificationService.refreshList$.subscribe(() => {
      this.loadDatabases();
    });
  }

  loadDatabases() {
    this.databaseService.getDatabases().subscribe({
      next: (data: Database[]) => {
        this.databases = data;
      },
      error: (error) => {
        this.errorMessage =
          'Erreur lors de la récupération des bases de données.';
        console.error(error);
      },
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
    // Ajoutez ici la logique pour créer un backup en utilisant l'ID de la base de données
  }
}
