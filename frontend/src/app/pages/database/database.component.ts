import { CommonModule } from '@angular/common';
import { Component, ViewChild } from '@angular/core';
import { DatabaseListComponent } from '../../components/database-list/database-list.component';
import { AddDatabaseDialogComponent } from '../../components/add-database-dialog/add-database-dialog.component';

@Component({
  selector: 'app-database',
  standalone: true,
  imports: [CommonModule, DatabaseListComponent, AddDatabaseDialogComponent],
  templateUrl: './database.component.html',
  styleUrl: './database.component.css',
})
export class DatabaseComponent {
  @ViewChild(DatabaseListComponent) databaseListComponent!: DatabaseListComponent;

  onDatabaseAdded() {
    console.log('Database added event received in parent');
    this.databaseListComponent.loadDatabases(); // Appelle la méthode loadDatabases() du composant DatabaseListComponent
  }
}
