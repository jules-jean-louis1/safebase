import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { DatabaseListComponent } from '../../components/database-list/database-list.component';
import { AddDatabaseDialogComponent } from '../../components/add-database-dialog/add-database-dialog.component';

@Component({
  selector: 'app-database',
  standalone: true,
  imports: [CommonModule, DatabaseListComponent, AddDatabaseDialogComponent],
  templateUrl: './database.component.html',
  styleUrl: './database.component.css'
})
export class DatabaseComponent {

}
