import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { LucideAngularModule } from 'lucide-angular';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterModule, LucideAngularModule],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.css',
})
export class SidebarComponent {}
