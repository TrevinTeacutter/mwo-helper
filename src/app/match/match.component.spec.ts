import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { MatchComponent } from './match.component';
import { TranslateModule } from '@ngx-translate/core';
import { RouterTestingModule } from '@angular/router/testing';


describe('MatchComponent', () => {
  let component: MatchComponent;
  let fixture: ComponentFixture<MatchComponent>;

  beforeEach(waitForAsync(() => {
    void TestBed.configureTestingModule({
    imports: [TranslateModule.forRoot(), RouterTestingModule, MatchComponent]
}).compileComponents();

    fixture = TestBed.createComponent(MatchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render title in a h1 tag', waitForAsync(() => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain(
      'PAGES.MATCH.TITLE'
    );
  }));
});
