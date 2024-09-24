import {
  ApplicationConfig,
  importProvidersFrom,
  LOCALE_ID,
  provideZoneChangeDetection,
} from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideClientHydration } from '@angular/platform-browser';
import { provideHttpClient, withFetch } from '@angular/common/http';
import { provideAnimations } from '@angular/platform-browser/animations';
import {
  LucideAngularModule,
  Database,
  DatabaseBackup,
  List,
  IterationCw,
  Trash,
  AlarmClockCheck,
  Plus,
  Calendar,
  Save,
  Plug,
  Info,
} from 'lucide-angular';
import { registerLocaleData } from '@angular/common';
import localeFr from '@angular/common/locales/fr';

registerLocaleData(localeFr);

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideClientHydration(),
    provideHttpClient(withFetch()),
    provideAnimations(),
    importProvidersFrom(
      LucideAngularModule.pick({
        Database,
        DatabaseBackup,
        List,
        IterationCw,
        Trash,
        AlarmClockCheck,
        Plus,
        Calendar,
        Save,
        Plug,
        Info
      })
    ),
    { provide: LOCALE_ID, useValue: 'fr' },
  ],
};
