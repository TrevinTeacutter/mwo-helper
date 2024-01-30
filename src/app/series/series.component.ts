import {Component, OnInit} from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-series',
    templateUrl: './series.component.html',
    styleUrl: './series.component.scss',
    standalone: true,
    imports: [TranslateModule]
})
export class SeriesComponent implements OnInit  {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

}
