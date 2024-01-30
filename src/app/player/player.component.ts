import {Component, OnInit} from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-player',
    templateUrl: './player.component.html',
    styleUrl: './player.component.scss',
    standalone: true,
    imports: [TranslateModule]
})
export class PlayerComponent implements OnInit  {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

}
