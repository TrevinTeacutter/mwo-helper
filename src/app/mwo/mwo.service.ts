import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {catchError, retry, throwError} from "rxjs";
import {PilotDetails} from "../match/match.component";

interface Response {
  MatchDetails: MatchDetails,
  UserDetails: UserDetails[],
}

interface MatchDetails {
  Map: string;
  ViewMode: string;
  TimeOfDay: string;
  GameMode: string;
  Region: string;
  MatchTimeMinutes: string;
  UseStockLoadout: boolean;
  NoMechQuirks: boolean;
  NoMechEfficiencies: boolean;
  WinningTeam: string;
  Team1Score: number;
  Team2Score: number;
  MatchDuration: string;
  CompleteTime: string;
}

interface UserDetails {
  Username: string;
  IsSpectator: boolean;
  Team: string;
  Lance: string;
  MechItemID: number;
  MechName: string;
  SkillTier: number;
  HealthPercentage: number;
  Kills: number;
  KillsMostDamage: number;
  Assists: number;
  ComponentsDestroyed: number;
  MatchScore: number;
  Damage: number;
  TeamDamage: number;
  UnitTag: string;
}

@Injectable({
  providedIn: 'root'
})
export class MwoService {
  constructor(private http: HttpClient) { }

  public getMatchDetails(apiKey: string, matchID: string, consumerFn: (response: Response) => void) {
    const request = this.http.get<Response>(`/api/v1/matches/${matchID}?api_token=${apiKey}`)
      .pipe(
        retry(3),
        catchError(error => {
          if (error.status === 0) {
            console.error('An error occurred:', error.error);
          } else {
            console.error(
              `Backend returned code ${error.status}, body was: `, error.error);
          }

          return throwError(() => new Error('Something bad happened; please try again later.'));
        })
      );

    request.subscribe(consumerFn).unsubscribe();
  }
}
