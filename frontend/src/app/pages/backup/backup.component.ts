import { Component } from '@angular/core';
import { BackupListComponent } from '../../components/backup-list/backup-list.component';

@Component({
  selector: 'app-backup',
  standalone: true,
  imports: [BackupListComponent],
  templateUrl: './backup.component.html',
  styleUrl: './backup.component.css'
})
export class BackupComponent {

}
