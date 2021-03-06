import { Injectable } from '@angular/core';
import { ThemeService } from '../../../sugarloaf/components/smart/theme-switch/theme-switch.service';

@Injectable()
export class InitService {
  constructor(private themeService: ThemeService) {}

  init(): void {
    this.themeService.loadTheme().catch(e => {
      console.error('Unable to load theme:', e);
    });
  }
}
