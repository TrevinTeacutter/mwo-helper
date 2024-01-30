import { Component, OnInit } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import {FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";

@Component({
    selector: 'app-match',
    templateUrl: './match.component.html',
    styleUrls: ['./match.component.scss'],
    standalone: true,
    imports: [TranslateModule, ReactiveFormsModule],
})
export class MatchComponent implements OnInit {

  matchSubmit = new FormGroup({
    apiKey: new FormControl(''),
    matchID: new FormControl(''),
  });

  constructor() { }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

  onSubmit() {
    // TODO: Use EventEmitter with form value
    console.warn(this.matchSubmit.value);
  }

}
