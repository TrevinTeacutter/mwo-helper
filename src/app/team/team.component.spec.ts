import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { TeamComponent } from './team.component';
import { TranslateModule } from '@ngx-translate/core';
import { RouterTestingModule } from '@angular/router/testing';

describe('TeamComponent', () => {
  let component: TeamComponent;
  let fixture: ComponentFixture<TeamComponent>;

  beforeEach(waitForAsync(() => {
    void TestBed.configureTestingModule({
    imports: [TranslateModule.forRoot(), RouterTestingModule, TeamComponent]
}).compileComponents();

    fixture = TestBed.createComponent(TeamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render title in a h1 tag', waitForAsync(() => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain(
      'PAGES.TEAM.TITLE'
    );
  }));
});
