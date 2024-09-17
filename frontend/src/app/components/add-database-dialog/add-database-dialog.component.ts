import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  Validators,
  ReactiveFormsModule,
} from '@angular/forms';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { DropdownModule } from 'primeng/dropdown';
import { CommonModule } from '@angular/common';
import { DatabaseService } from '../../services/database.service';
import { MessageModule } from 'primeng/message';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { NotificationService } from '../../services/notification.service';

@Component({
  selector: 'app-add-database-dialog',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    DialogModule,
    ButtonModule,
    InputTextModule,
    DropdownModule,
    MessageModule,
    ToastModule,
  ],
  templateUrl: './add-database-dialog.component.html',
  styleUrls: ['./add-database-dialog.component.css'],
  providers: [MessageService, NotificationService],
})
export class AddDatabaseDialogComponent implements OnInit {
  dbTypes: any[];
  databaseForm!: FormGroup;
  visible: boolean = false;

  constructor(
    private fb: FormBuilder,
    private databaseService: DatabaseService,
    private messageService: MessageService,
    private notificationService: NotificationService
  ) {
    this.dbTypes = [
      { label: 'MySQL', value: 'mysql' },
      { label: 'PostgreSQL', value: 'postgres' },
    ];
  }

  ngOnInit() {
    this.databaseForm = this.fb.group({
      name: ['', Validators.required],
      type: ['', Validators.required],
      host: ['', Validators.required],
      port: ['', Validators.required],
      username: ['', Validators.required],
      password: ['', Validators.required],
      database_name: ['', Validators.required],
    });
  }

  showDialog() {
    this.databaseForm.reset();
    this.visible = true;
  }

  onSubmit() {
    if (this.databaseForm.valid) {
      console.log(this.databaseForm.value);
      // Traitez les données du formulaire ici
      this.databaseService.addDatabase(this.databaseForm.value).subscribe({
        next: (data) => {
          console.log('Database added', data);
          this.messageService.add({
            severity: 'success',
            summary: 'Database added',
            detail: 'Database added successfully',
          });
          this.notificationService.notifyRefreshList();
          this.visible = false; // Fermez le dialog après soumission
        },
        error: (error) => {
          console.error('Database add failed', error);
          this.messageService.add({
            severity: 'error',
            summary: 'Database add failed',
            detail: 'Failed to add database',
          });
        },
      });
    } else {
      // Marquez tous les champs comme touchés pour afficher les erreurs
      Object.values(this.databaseForm.controls).forEach((control) => {
        control.markAsTouched();
      });
      this.messageService.add({
        severity: 'error',
        summary: 'Form validation failed',
        detail: 'Please fill in all required fields',
      });
    }
  }

  testConnection() {
    if (this.databaseForm.valid) {
      this.databaseService.testConnection(this.databaseForm.value).subscribe({
        next: (data) => {
          console.log('Connection successful', data);
          this.messageService.add({
            severity: 'success',
            summary: 'Connection successful',
            detail: 'Connection to the database was successful',
          });
        },
        error: (error) => {
          console.error('Connection failed', error);
          this.messageService.add({
            severity: 'error',
            summary: 'Connection failed',
            detail: 'Connection to the database failed',
          });
        },
      });
    }
  }
}
