import { Component, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-match',
    templateUrl: './match.component.html',
    styleUrls: ['./match.component.scss'],
    standalone: true,
    imports: [TranslateModule]
})
export class MatchComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

}
