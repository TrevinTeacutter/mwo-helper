import {Routes} from "@angular/router";
import {HomeComponent} from "./home/home.component";
import {DetailComponent} from "./detail/detail.component";
import {TeamComponent} from "./team/team.component";
import {PlayerComponent} from "./player/player.component";
import {SeriesComponent} from "./series/series.component";
import {MatchComponent} from "./match/match.component";
import {SettingsComponent} from "./settings/settings.component";
import {PageNotFoundComponent} from "./shared/components";

export const routes: Routes = [
  {path: 'home', component: HomeComponent},
  {path: 'detail', component: DetailComponent},
  {path: 'team', component: TeamComponent},
  {path: 'player', component: PlayerComponent},
  {path: 'series', component: SeriesComponent},
  {path: 'match', component: MatchComponent},
  {path: 'settings', component: SettingsComponent},
  {
    path: '',
    redirectTo: '/home',
    pathMatch: 'full'
  },
  {
    path: '**',
    component: PageNotFoundComponent
  }
];
