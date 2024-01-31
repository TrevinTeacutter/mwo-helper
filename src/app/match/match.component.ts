import {Component, OnInit} from '@angular/core';
import {TranslateModule} from '@ngx-translate/core';
import {FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {MatTableModule} from '@angular/material/table';
import {HttpClient, HttpClientModule} from "@angular/common/http";
import {MwoService} from "../mwo/mwo.service";

export interface PilotDetails {
  pilot: string;
  lance: string;
  mech: string;
  kills: number;
  assists: number;
  damage: number;
  teamDamage: number;
  status: string;
  health: number;
  matchScore: number;
}

@Component({
  selector: 'app-match',
  templateUrl: './match.component.html',
  styleUrls: ['./match.component.scss'],
  standalone: true,
  imports: [HttpClientModule, TranslateModule, ReactiveFormsModule, MatTableModule],
})
export class MatchComponent implements OnInit {

  displayedColumns: string[] = ['pilot', 'lance', 'mech', 'kills', 'assists', 'damage', 'teamDamage', 'status', 'health', 'matchScore'];
  team1: PilotDetails[] = [];
  team2: PilotDetails[] = [];

  matchSubmit = new FormGroup({
    apiKey: new FormControl(''),
    matchID: new FormControl(''),
  });

  constructor(private mwo: MwoService) {
  }

  ngOnInit(): void {
    console.log('DetailComponent INIT');
  }

  onSubmit() {
    this.team1 = [];
    this.team2 = [];

    this.mwo.getMatchDetails(
      <string>this.matchSubmit.value.apiKey,
      <string>this.matchSubmit.value.matchID,
      results => {
        console.debug(results);

        results.UserDetails.sort((a, b) => {
          if (a.Lance < b.Lance) {
            return -1;
          }

          if (a.Lance > b.Lance) {
            return 1;
          }

          return 0;
        }).forEach(element => {
          if (element.IsSpectator) {
            return
          }

          const pilotDetails: PilotDetails = {
            pilot: element.Username,
            lance: element.Lance,
            mech: element.MechName,
            kills: element.Kills,
            assists: element.Assists,
            damage: element.Damage,
            teamDamage: element.TeamDamage,
            status: element.HealthPercentage > 0 ? 'Alive' : 'Dead',
            health: element.HealthPercentage,
            matchScore: element.MatchScore,
          }

          switch (element.Team) {
            case "1":
              this.team1.push(pilotDetails)
              break;
            default:
              this.team2.push(pilotDetails)
              break;
          }
        })
      })
  }
}
