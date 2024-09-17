import { Component, OnInit } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { DropdownModule } from 'primeng/dropdown';
import { DatabaseService } from '../../services/database.service';

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

@Component({
  selector: 'app-schedule-backup',
  standalone: true,
  imports: [
    DialogModule,
    ButtonModule,
    InputTextModule,
    ReactiveFormsModule,
    DropdownModule,
  ],
  templateUrl: './schedule-backup.component.html',
  styleUrl: './schedule-backup.component.css',
  providers: [DatabaseService],
})
export class ScheduleBackupComponent implements OnInit {
  databases: Database[] = [];
  visible: boolean = false;

  formGroup: FormGroup = new FormGroup({});
  showDialog() {
    this.visible = true;
  }
  constructor(private databaseService: DatabaseService) {}

  loadDatabases() {
    this.databaseService.getDatabases().subscribe(
      (data) => {
        this.databases = data;
      },
      (error) => {
        console.error(error);
      }
    );
  }

  // scheduleForm = new FormGroup({
  //   name: new FormControl<string>('')
  // });

  ngOnInit(): void {
    this.loadDatabases();

    this.formGroup = new FormGroup({
      selectedDatabase: new FormControl<Database | null>(null),
    });
  }

  // onSubmit() {
  //   console.warn(this.scheduleForm.value);
  // }
}
