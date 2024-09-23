import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { LucideAngularModule } from 'lucide-angular';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';

@Component({
  selector: 'app-backup-intel',
  standalone: true,
  imports: [DialogModule, ButtonModule, LucideAngularModule, CommonModule],
  templateUrl: './backup-intel.component.html',
  styleUrl: './backup-intel.component.css'
})
export class BackupIntelComponent {
  @Input() backup!: any;
    visible: boolean = false;

    showDialog() {
        this.visible = true;
    }
}
