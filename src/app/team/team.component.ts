import { Component, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-home',
    templateUrl: './team.component.html',
    styleUrls: ['./team.component.scss'],
    standalone: true,
    imports: [TranslateModule]
})
export class TeamComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

}
