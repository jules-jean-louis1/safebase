import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import {
  FormGroup,
  Validators,
  ReactiveFormsModule,
  FormControl,
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
import { LucideAngularModule } from 'lucide-angular';

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
    LucideAngularModule,
  ],
  templateUrl: './add-database-dialog.component.html',
  styleUrls: ['./add-database-dialog.component.css'],
  providers: [MessageService],
})
export class AddDatabaseDialogComponent implements OnInit {
  @Output() databaseAdded = new EventEmitter<void>();
  dbTypes: any[] = [];
  databaseForm: FormGroup = new FormGroup({});
  visible: boolean = false;

  constructor(
    private databaseService: DatabaseService,
    private messageService: MessageService,
  ) {}

  ngOnInit() {
    this.dbTypes = [
      { label: 'MySQL 8.0', value: 'mysql' },
      { label: 'PostgreSQL 16.0', value: 'postgres' },
    ];

    this.databaseForm = new FormGroup({
      name: new FormControl('', Validators.required),
      type: new FormControl('', Validators.required),
      host: new FormControl('', Validators.required),
      port: new FormControl('', Validators.required),
      username: new FormControl('', Validators.required),
      password: new FormControl(''),
      database_name: new FormControl('', Validators.required),
    });
  }

  showDialog() {
    this.databaseForm.reset();
    this.visible = true;
  }

  onSubmit() {
    if (this.databaseForm.valid) {
      this.databaseService.addDatabase(this.databaseForm.value).subscribe({
        next: (data) => {
          this.databaseAdded.emit();// Émet l'événement
          this.messageService.add({
            severity: 'success',
            summary: 'Database added',
            detail: 'Database added successfully',
          });
          this.visible = false; // Fermez le dialog après soumission
        },
        error: (error) => {
          console.log('Database add failed', error);
          console.log(error.error);
          this.messageService.add({
            severity: 'error',
            summary: 'Database add failed',
            detail: error.error.error,
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
          this.messageService.add({
            severity: 'success',
            summary: 'Connection successful',
            detail: 'Connection to the database was successful',
          });
        },
        error: (error) => {
          this.messageService.add({
            severity: 'error',
            summary: 'Connection failed',
            detail: error.error.error,
          });
        },
      });
    }
  }

}
