import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { SeriesComponent } from './series.component';
import { TranslateModule } from '@ngx-translate/core';
import { RouterTestingModule } from '@angular/router/testing';


describe('SeriesComponent', () => {
  let component: SeriesComponent;
  let fixture: ComponentFixture<SeriesComponent>;

  beforeEach(waitForAsync(() => {
    void TestBed.configureTestingModule({
    imports: [TranslateModule.forRoot(), RouterTestingModule, SeriesComponent]
}).compileComponents();

    fixture = TestBed.createComponent(SeriesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render title in a h1 tag', waitForAsync(() => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain(
      'PAGES.SERIES.TITLE'
    );
  }));
});
