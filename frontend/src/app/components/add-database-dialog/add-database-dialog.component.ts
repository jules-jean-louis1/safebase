import { Component } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';

@Component({
  selector: 'app-add-database-dialog',
  standalone: true,
  imports: [DialogModule, ButtonModule, InputTextModule],
  templateUrl: './add-database-dialog.component.html',
  styleUrl: './add-database-dialog.component.css',
})
export class AddDatabaseDialogComponent {
  visible: boolean = false;

  showDialog() {
    this.visible = true;
  }
}
