import { Component, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import {RouterLink, RouterLinkActive} from '@angular/router';

@Component({
    selector: 'app-detail',
    templateUrl: './detail.component.html',
    styleUrls: ['./detail.component.scss'],
    standalone: true,
  imports: [RouterLink, TranslateModule, RouterLinkActive]
})
export class DetailComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
   }

}
