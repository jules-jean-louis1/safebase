import { Component, Input, OnInit } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { DropdownModule } from 'primeng/dropdown';
import { DatabaseService } from '../../services/database.service';
import { InputSwitchModule } from 'primeng/inputswitch';
import { MessageService } from 'primeng/api';
import { MessageModule } from 'primeng/message';
import { ToastModule } from 'primeng/toast';
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
  is_cron_active?: boolean;
  cron_schedule?: string;
  created_at: string;
  updated_at: string;
}

interface CronOption {
  label: string;
  cron: string;
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
    InputSwitchModule,
    MessageModule,
    ToastModule,
    LucideAngularModule,
    TooltipModule,
  ],
  templateUrl: './schedule-backup.component.html',
  styleUrl: './schedule-backup.component.css',
  providers: [DatabaseService, MessageService],
})
export class ScheduleBackupComponent implements OnInit {
  @Input() database!: Database;
  cronOptions: CronOption[] | undefined;
  visible: boolean = false;

  scheduleForm: FormGroup = new FormGroup({});

  constructor(
    private databaseService: DatabaseService,
    private messageService: MessageService
  ) {}

  ngOnInit(): void {
    this.cronOptions = [
      { label: 'Toutes les minutes', cron: '* * * * *' },
      { label: 'Toutes les 5 minutes', cron: '*/5 * * * *' },
      { label: 'Toutes les heures', cron: '0 * * * *' },
      { label: 'Tous les jours à minuit', cron: '0 0 * * *' },
      { label: 'Chaque lundi à 8h00', cron: '0 8 * * 1' },
      { label: 'Le premier jour du mois à minuit', cron: '0 0 1 * *' },
      { label: 'Chaque dimanche à 10h00', cron: '0 10 * * 7' },
      { label: 'Chaque jour de la semaine à 7h30', cron: '30 7 * * 1-5' },
      { label: 'Tous les 15 jours à minuit', cron: '0 0 */15 * *' },
      { label: 'Tous les mois le 15 à 14h', cron: '0 14 15 * *' },
    ];

    this.scheduleForm = new FormGroup({
      selectedCron: new FormControl<CronOption | null>(
        null,
        Validators.required
      ),
      isActiveCron: new FormControl<boolean>(false),
    });
  }

  showDialog() {
    this.scheduleForm.patchValue({
      selectedCron:
        this.cronOptions?.find(
          (option) => option.cron === this.database.cron_schedule
        ) || null,
      isActiveCron: this.database.is_cron_active || false,
    });
    this.visible = true;
  }

  onSubmit() {

    if (!this.scheduleForm.valid) {
      this.messageService.add({
        severity: 'error',
        summary: 'Formulaire invalide',
        detail: 'Veuillez remplir tous les champs',
      });
      console.log('Formulaire invalide');
      return;
    }
    const formValues = this.scheduleForm.value;
    this.database.cron_schedule = formValues.selectedCron?.cron || '';
    this.database.is_cron_active = formValues.isActiveCron;

    this.databaseService.updateDatabase(this.database).subscribe({
      next: (response) => {
        console.log('Database updated successfully', response);
        this.messageService.add({
          severity: 'success',
          summary: 'Succès',
          detail: 'Base de données mise à jour avec succès',
        });
        this.visible = false;
      },
      error: (error) => {
        console.error('Error updating database', error);
        this.messageService.add({
          severity: 'error',
          summary: 'Erreur',
          detail: 'Erreur lors de la mise à jour de la base de données',
        });
      },
    });
  }
}
