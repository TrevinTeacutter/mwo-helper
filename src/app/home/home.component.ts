import { Component, OnInit } from '@angular/core';
import {Router, RouterLink, RouterLinkActive, RouterOutlet} from '@angular/router';
import { TranslateModule } from '@ngx-translate/core';

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    standalone: true,
  imports: [RouterLink, TranslateModule, RouterLinkActive, RouterOutlet]
})
export class HomeComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit(): void {
    console.log('HomeComponent INIT');
  }

}
