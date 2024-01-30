import { Component, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-settings',
    templateUrl: './settings.component.html',
    styleUrl: './settings.component.scss',
    standalone: true,
    imports: [TranslateModule]
})
export class SettingsComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

}
